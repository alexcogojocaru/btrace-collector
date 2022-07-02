package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/alexcogojocaru/collector/config"
	"github.com/alexcogojocaru/collector/extensions"
	proxy_grpc "github.com/alexcogojocaru/collector/proto-gen/btrace_proxy"
	"google.golang.org/grpc"
)

// CollectorServiceImpl holds the description for the CollectorService from the proto generated files
type CollectorServiceImpl struct {
	// UnimplementedCollectorServiceServer embedded to have forward compatible implementations
	proxy_grpc.UnimplementedAgentServer
	Extensions []extensions.Pluggable
}

// Creates and returns a pointer to CollectorServiceImpl
func NewCollectorServiceImpl() *CollectorServiceImpl {
	collectorImpl := &CollectorServiceImpl{}

	return collectorImpl
}

func (collectorService *CollectorServiceImpl) AddExtensions(exts ...extensions.Pluggable) {
	for _, extension := range exts {
		collectorService.Extensions = append(collectorService.Extensions, extension)
	}
}

func (cs *CollectorServiceImpl) NotifyExtensions(ctx context.Context, span *proxy_grpc.Span) {
	var wg sync.WaitGroup
	for _, extension := range cs.Extensions {
		wg.Add(1)

		go func(extension extensions.Pluggable) {
			defer wg.Done()
			extension.Send(ctx, span)
		}(extension)
	}

	wg.Wait()
}

func (cs *CollectorServiceImpl) Send(ctx context.Context, span *proxy_grpc.Span) (*proxy_grpc.Response, error) {
	cs.NotifyExtensions(ctx, span)
	return &proxy_grpc.Response{}, nil
}

func (cs *CollectorServiceImpl) Stream(stream proxy_grpc.Agent_StreamServer) error {
	for {
		span, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proxy_grpc.Response{})
		}

		if err != nil {
			return err
		}

		cs.NotifyExtensions(stream.Context(), span)
	}
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)

	conf, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal("Error while parsing the config file")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Deploy.Port))
	if err != nil {
		log.Fatal("Failed to listen on port")
	} else {
		log.Printf("Started listening on port %d\n", conf.Deploy.Port)
	}

	grpcServer := grpc.NewServer()
	collectorService := NewCollectorServiceImpl()
	collectorService.AddExtensions(
		extensions.NewStorageExtension(conf.Connections.Storage.Hostname, int(conf.Connections.Storage.Port)),
		extensions.NewNeo4jExtension(),
	)

	proxy_grpc.RegisterAgentServer(grpcServer, collectorService)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve")
	}
}
