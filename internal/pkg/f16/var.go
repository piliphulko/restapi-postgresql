package f16

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Counter_LableMethodCodeInCounterVec = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: os.Getenv("NAMESPACE"),
			Name:      "count_use_method",
			Help:      "Total number of calls to the gRPC method",
		}, []string{"method", "code"})

	DurationReq_LableMethodCodeInSummaryVec = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  os.Getenv("NAMESPACE"),
		Name:       "grpc_method_duration_seconds",
		Help:       "Time taken to complete gRPC requests",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"method", "code"})
)
