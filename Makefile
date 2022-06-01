run:
	@collector.exe

build:
	@go install

generate-proto:
	@protoc -I=btrace-idl/proto --go_out=proto-gen ./btrace-idl/proto/v2/proxy.proto
	@protoc -I=btrace-idl/proto --go-grpc_out=proto-gen ./btrace-idl/proto/v2/proxy.proto

proto_storage:
	@protoc -I=btrace-idl/proto --go_out=proto-gen ./btrace-idl/proto/v2/storage.proto
	@protoc -I=btrace-idl/proto --go-grpc_out=proto-gen ./btrace-idl/proto/v2/storage.proto
