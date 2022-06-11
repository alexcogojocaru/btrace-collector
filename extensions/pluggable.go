package extensions

import (
	"context"

	proxy_grpc "github.com/alexcogojocaru/collector/proto-gen/btrace_proxy"
)

type Pluggable interface {
	Send(ctx context.Context, span *proxy_grpc.Span) error
	Close(ctx context.Context)
}
