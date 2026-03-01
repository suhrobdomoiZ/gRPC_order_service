package middleware

import (
	"context"
	"log/slog"
	"os"
	"time"

	"google.golang.org/grpc"
)

var baseLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		baseLogger.Info("Started to perform request")
		start := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(start)

		if err != nil {
			baseLogger.Error("Request failed", "error", err.Error(), "duration", duration.String())
		} else {
			baseLogger.Info("Request completed", "duration", duration.String())
		}
		return resp, err
	}
}
