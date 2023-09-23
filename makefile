run:
	go run cmd/*.go
proto:
	protoc internal/pb/*.proto --go_out=. --go-grpc_out=.

