package main

import (
	"log"
	"net"
	"time"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/bootstrap"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/gapi"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	server, err := gapi.NewServer(env, *app.Db, timeout)
	if err != nil {
		log.Fatal("Cannot create server :", err)
	}

	gprcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)

	grpcServer := grpc.NewServer(gprcLogger)

	pb.RegisterBooksServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", env.BackendServiceUrl)
	if err != nil {
		log.Fatal("Cannot create listener :", err)
	}

	log.Printf("Starting gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start gRPC server :", err)
	}

}
