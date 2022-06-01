package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/alexcogojocaru/collector/config"
	collector "github.com/alexcogojocaru/collector/proto-gen/btrace_proxy"
	"google.golang.org/grpc"
)

// CollectorServiceImpl holds the description for the CollectorService from the proto generated files
type CollectorServiceImpl struct {
	// UnimplementedCollectorServiceServer embedded to have forward compatible implementations
	collector.UnimplementedAgentServer
}

// Creates and returns a pointer to CollectorServiceImpl
func NewCollectorServiceImpl() *CollectorServiceImpl {
	collectorImpl := &CollectorServiceImpl{}
	return collectorImpl
}

// Implementation of the StreamSpan method from the proto file
// Receives a SpanRequest object that holds a slice of Spans and returns (SpanResponse, error)
func (agentServiceImpl *CollectorServiceImpl) Send(ctx context.Context, span *collector.Span) (*collector.Response, error) {
	log.Print(span)

	return &collector.Response{}, nil
}

func main() {
	// Set the microseconds flag to have precise timestamp measurements
	log.SetFlags(log.Lmicroseconds | log.Ldate)

	// TODO in future: use the config file to deploy the collector as a cluster using RAFT as a leader election and log replication
	conf, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal("Error while parsing the config file")
	}

	// Create a TCP socket on the port specified in the config file
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Deploy.Port))
	if err != nil {
		log.Fatal("Failed to listen on port")
	} else {
		log.Printf("Started listening on port %d\n", conf.Deploy.Port)
	}

	grpcServer := grpc.NewServer()
	collectorService := NewCollectorServiceImpl()

	collector.RegisterAgentServer(grpcServer, collectorService)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve")
	}
}
