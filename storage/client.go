package storage

import (
	"log"

	storage_grpc "github.com/alexcogojocaru/collector/proto-gen/btrace_storage"
	"google.golang.org/grpc"
)

type StorageClient struct {
	Client storage_grpc.StorageClient
}

func NewStorageClient() *StorageClient {
	storageServerHost := "192.168.100.2:50051"
	conn, err := grpc.Dial(storageServerHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial %s", storageServerHost)
	}

	return &StorageClient{
		Client: storage_grpc.NewStorageClient(conn),
	}
}
