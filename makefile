run:
	go run main.go

gen:
	protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/service.proto
	grpc_tools_node_protoc --grpc_out=grpc_js:pb proto/*.proto

clean:
	rm -rf pb/*