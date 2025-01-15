package f16

import (
	"context"
	"fmt"
	"time"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var mapClientCodes = map[error]codes.Code{
	apierror.ErrTokenFake:          codes.InvalidArgument,
	apierror.ErrTokenExpired:       codes.InvalidArgument,
	apierror.ErrEmpty:              codes.InvalidArgument,
	apierror.ErrMissingMetadata:    codes.InvalidArgument,
	apierror.ErrMissingToken:       codes.Unauthenticated,
	apierror.ErrOrderStatNotSelect: codes.InvalidArgument,
	apierror.ErrPassLen:            codes.InvalidArgument,
	apierror.ErrIncorrectPass:      codes.InvalidArgument,
	apierror.ErrIncorrectLogin:     codes.Unauthenticated,
	apierror.ErrDataLoss:           codes.DataLoss,
	apierror.ErrLoginBusy:          codes.AlreadyExists,
	apierror.ErrIncorrectCountry:   codes.InvalidArgument,
	apierror.ErrInvalidRequest:     codes.InvalidArgument,
	apierror.ErrNotEnoughGoods:     codes.InvalidArgument,
	apierror.ErrDeliveryLocation:   codes.InvalidArgument,
	apierror.ErrNotCanceled:        codes.InvalidArgument,
	apierror.ErrNotEnoughMoney:     codes.InvalidArgument,
	apierror.ErrInvalidTokenFormat: codes.InvalidArgument,
}

func InterceptorCheckCtx(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return nil, status.New(codes.DeadlineExceeded, "").Err()
	case context.Canceled:
		return nil, status.New(codes.Canceled, "").Err()
	default:
		return handler(ctx, req)
	}
}

func IntrceptorHandlerErrors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, err
	}
	fmt.Println(err)
	codesAnswer, ok := mapClientCodes[err]
	if ok {
		return resp, status.New(codesAnswer, err.Error()).Err()
	} else {
		return resp, status.New(codes.Internal, "").Err()
	}
}

func WarpInteceptorRequestTimer(histogramVec *prometheus.SummaryVec) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		var (
			requestTimer = time.Now()
			code         string
		)
		resp, err = handler(ctx, req)

		methodName, ok := ignoringHealthRequest(info.FullMethod)
		if ok {
			return resp, err
		}

		status, ok := status.FromError(err)
		switch ok {
		case true:
			code = status.Code().String()
		case false:
			code = codes.OK.String()
		}

		defer histogramVec.WithLabelValues(methodName, code).Observe(time.Since(requestTimer).Seconds())

		return resp, err
	}
}

func WarpInteceptorRequestCounter(counterVec *prometheus.CounterVec) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)

		methodName, ok := ignoringHealthRequest(info.FullMethod)
		if ok {
			return resp, err
		}

		var code string
		status, ok := status.FromError(err)
		switch ok {
		case true:
			code = status.Code().String()
		case false:
			code = codes.OK.String()
		}

		counterVec.WithLabelValues(methodName, code).Inc()
		return resp, err
	}
}
