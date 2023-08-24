run:
	go run main.go

gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=:pb
	grpc_tools_node_protoc --grpc_out=grpc_js:pb proto/*.proto

clean:
	rm -rf pb/*