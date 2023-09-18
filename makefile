run:
	go run .cmd/
proto:
	protoc internal/pb/*.proto --go_out=. --go-grpc_out=.