run:
	go run main.go

gen:
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/service.proto
	grpc_tools_node_protoc --js_out=import_style=commonjs,binary:pb --grpc_out=grpc_js:pb proto/service.proto

clean:
	rm -rf pb/*