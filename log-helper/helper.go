package log_helper

import (
	"context"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Initialize(debug bool, useJsonLogger bool) {
	var cfg zap.Config
	if useJsonLogger {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	if debug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	zap.S().Debug("Initialize Zap Logger with Debug Level")
}

func GetUnaryClientInterceptors() []grpc.UnaryClientInterceptor {
	list := []grpc.UnaryClientInterceptor{
		grpc_zap.UnaryClientInterceptor(zap.L()),
	}

	if ce := zap.L().Check(zap.DebugLevel, "Lorem Ipsum"); ce != nil {
		list = append(list, grpc_zap.PayloadUnaryClientInterceptor(zap.L(), func(ctx context.Context, fullMethodName string) bool {
			return true
		}))
	}
	return list
}

func GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	list := []grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(zap.L()),
	}

	if ce := zap.L().Check(zap.DebugLevel, "Lorem Ipsum"); ce != nil {
		list = append(list, grpc_zap.PayloadUnaryServerInterceptor(zap.L(), func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return true
		}))
	}
	return list
}

func GetStreamServerInterceptors() []grpc.StreamServerInterceptor {
	list := []grpc.StreamServerInterceptor{
		grpc_zap.StreamServerInterceptor(zap.L()),
	}

	if ce := zap.L().Check(zap.DebugLevel, "Lorem Ipsum"); ce != nil {
		list = append(list, grpc_zap.PayloadStreamServerInterceptor(zap.L(), func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return true
		}))
	}
	return list
}
