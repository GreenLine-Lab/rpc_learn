package trace

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const TraceId = "Trace-Id"

func GetIdFromCtx(ctx context.Context) string {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		traceIdHeader := md.Get(TraceId)
		if traceIdHeader != nil {
			return traceIdHeader[0]
		}
	}

	return ""
}
