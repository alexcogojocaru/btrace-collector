package extensions

import (
	"context"
	"log"
	"os"

	proxy_grpc "github.com/alexcogojocaru/collector/proto-gen/btrace_proxy"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jExtension struct {
	Driver  neo4j.Driver
	Session neo4j.Session
}

func (neo *Neo4jExtension) Send(ctx context.Context, span *proxy_grpc.Span) error {
	log.Printf("%s %s %s", span.TraceID, span.SpanID, span.ParentSpanID)

	var query string
	if span.ParentSpanID == "0000000000000000" {
		query = `
			MERGE (s:Span {id: $span_id}) 
			MERGE (t:Trace {id: $trace_id}) 
			CREATE (t)-[:PARENT]->(s)`
	} else {
		query = `
			MERGE (ps:Span {id: $parent_span_id}) 
			MERGE (s:Span {id: $span_id}) 
			CREATE (ps)-[:PARENT]->(s)`
	}

	_, err := neo.Session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			query,
			map[string]interface{}{
				"parent_span_id": span.ParentSpanID,
				"span_id":        span.SpanID,
				"trace_id":       span.TraceID,
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	return err
}

func (neo *Neo4jExtension) Close(ctx context.Context) {
	neo.Session.Close()
	neo.Driver.Close()
}

func NewNeo4jExtension() Pluggable {
	neo := &Neo4jExtension{}

	neo4j_uri := os.Getenv("NEO4J_URI")
	neo4j_user := os.Getenv("NEO4J_USER")
	neo4j_pass := os.Getenv("NEO4J_PASSWORD")

	neo.Driver, _ = neo4j.NewDriver(neo4j_uri, neo4j.BasicAuth(neo4j_user, neo4j_pass, ""))

	neo.Session = neo.Driver.NewSession(
		neo4j.SessionConfig{
			AccessMode: neo4j.AccessModeWrite,
		},
	)

	return neo
}
