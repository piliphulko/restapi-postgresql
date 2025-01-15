package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/piliphulko/marketplace-example/internal/pkg/donkeyhealth"
	"github.com/piliphulko/marketplace-example/internal/pkg/donkeypacking"
	"github.com/piliphulko/marketplace-example/internal/pkg/f16"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-acct-auth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	grpc_health_v1.UnimplementedHealthServer
}

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("APP_NETWORK", "tcp")
	viper.SetDefault("URL_PROMETHEUS", "/metrics")
	viper.SetDefault("URL_CONN_POSTGRESQL", "postgres://postgres:5432@service-outboard-postgresql.outboard-services.svc.cluster.local:5432/test_db")
	viper.SetDefault("JWT_SECRET", "asdfghjk")

	pb.InitJWTSecret(viper.GetString("JWT_SECRET"))
}

func main() {
	lis, err := net.Listen(
		viper.GetString("APP_NETWORK"),
		fmt.Sprintf(":%s", viper.GetString("APP_PORT")),
	)
	if err != nil {
		log.Fatal(err)
	}

	//grpclog.SetLoggerV2(pb.LogGRPC)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			f16.InterceptorCheckCtx,
			f16.IntrceptorHandlerErrors,
			f16.WarpInteceptorRequestTimer(f16.DurationReq_LableMethodCodeInSummaryVec),
			f16.WarpInteceptorRequestCounter(f16.Counter_LableMethodCodeInCounterVec),
		),
	)

	pgxPool, closePgxPool, err := donkeypacking.GetConnPostrgresql(viper.GetString("URL_CONN_POSTGRESQL"))

	defer closePgxPool()
	if err != nil {
		log.Fatal(err)
	}

	server := pb.StartServer()

	server.InsertPostgresql(pgxPool)

	healthFragmentServer := donkeyhealth.CreateFragmentForCheckingHealthServer(
		donkeyhealth.ServiceAsFollows{
			Postgresql: pgxPool,
		},
	)

	grpc_health_v1.RegisterHealthServer(grpcServer, healthFragmentServer)

	pb.RegisterServer(grpcServer, server)

	go func() {
		http.Handle(viper.GetString("URL_PROMETHEUS"), promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("PROMETHEUS_PORT")), nil)
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
