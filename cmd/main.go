package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/config"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/di"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/util"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfigs()
	if util.HasError(context.Background(), err) {
		log.Fatalln("failed to load configs", err)
	}

	service, err := di.InitialiazeDeps(config)
	if util.HasError(context.Background(), err) {
		log.Fatalln("failed to intialize deps", err)
	}

	listener, err := net.Listen("tcp", config.Port)
	if util.HasError(context.Background(), err) {
		log.Fatalln("failed to create listener ", err)
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, service)

	fmt.Println("service raising up ...")
	fmt.Println("serving on port:", config.Port)

	log.Fatalln(server.Serve(listener))
}
