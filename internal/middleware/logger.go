package middleware

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func LoggerInterceptor(logger slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		logger.Info("Started to perform request")
		start := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(start)

		if err != nil {
			logger.Error("Request failed", "error", err.Error(), "duration", duration.String())
		} else {
			logger.Info("Request completed", "duration", duration.String())
		}
		return resp, err
	}
}
