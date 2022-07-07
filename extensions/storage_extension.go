package extensions

import (
	"context"
	"fmt"

	proxy_grpc "github.com/alexcogojocaru/collector/proto-gen/btrace_proxy"
	storage_grpc "github.com/alexcogojocaru/collector/proto-gen/btrace_storage"
	"google.golang.org/grpc"
)

type StorageExtension struct {
	Client storage_grpc.StorageClient
}

func ConvertToStorageKeyValue(logs []*proxy_grpc.KeyValue) []*storage_grpc.StorageKeyValue {
	var storageLogs []*storage_grpc.StorageKeyValue

	for _, log := range logs {
		storageLogs = append(storageLogs, &storage_grpc.StorageKeyValue{
			Type:  log.Type,
			Value: log.Value,
		})
	}

	return storageLogs
}

func (st *StorageExtension) Send(ctx context.Context, span *proxy_grpc.Span) error {
	storageSpan := storage_grpc.StorageSpan{
		ServiceName:  span.ServiceName,
		SpanID:       span.SpanID,
		TraceID:      span.TraceID,
		ParentSpanID: span.ParentSpanID,
		SpanName:     span.Name,
		Timestamp: &storage_grpc.StorageTimestamp{
			Started:  span.Timestamp.Started,
			Ended:    span.Timestamp.Ended,
			Duration: span.Timestamp.Duration,
		},
		Logs: ConvertToStorageKeyValue(span.Logs),
	}

	st.Client.Store(ctx, &storageSpan)
	return nil
}

func (st *StorageExtension) Close(ctx context.Context) {

}

func NewStorageExtension(storageServerHost string, storageServerPort int) Pluggable {
	conn, _ := grpc.Dial(fmt.Sprintf("%s:%d", storageServerHost, storageServerPort), grpc.WithInsecure())

	return &StorageExtension{
		Client: storage_grpc.NewStorageClient(conn),
	}
}
