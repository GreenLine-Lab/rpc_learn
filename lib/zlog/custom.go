package zlog

import (
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"rpc-learn/lib/trace"
	"strings"
)

type customLogger struct {
	logger  zerolog.Logger
	server  string
	method  string
	traceId string
}

func ApiLogger(logger zerolog.Logger, ctx context.Context) zerolog.Logger {
	apiLogger := &customLogger{
		logger: logger,
	}

	method, ok := grpc.Method(ctx)
	if ok {
		serverAndMethodName := strings.Split(method, "/")
		apiLogger.method = serverAndMethodName[2]

		serverName := strings.Split(serverAndMethodName[1], ".")
		apiLogger.server = serverName[1]
	} else {
		apiLogger.server = "UNKNOWN"
		apiLogger.method = "Unknown"
	}

	apiLogger.traceId = trace.GetIdFromCtx(ctx)

	return apiLogger.logger.With().Timestamp().
		Str("server", apiLogger.server).
		Str("method", apiLogger.method).
		Str("traceId", apiLogger.traceId).Logger()
}

func SrvLogger(logger zerolog.Logger, srvName string) zerolog.Logger {
	return logger.With().Timestamp().
		Str("server", srvName).Logger()
}
