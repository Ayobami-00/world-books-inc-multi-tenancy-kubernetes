proto:
	rm -f pb/*.pb.go
	protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    pb/*.proto

server:
	go run cmd/main.go

.PHONY: proto server