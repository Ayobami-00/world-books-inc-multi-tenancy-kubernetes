package gapi

import (
	"time"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/bootstrap"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
)

type Server struct {
	pb.UnimplementedBooksServiceServer
	Db      bootstrap.Database
	Env     *bootstrap.Env
	Timeout time.Duration
}

// NewServer creates a new gRPC server
func NewServer(env *bootstrap.Env, db bootstrap.Database, timeout time.Duration) (*Server, error) {

	server := &Server{
		Db:      db,
		Env:     env,
		Timeout: timeout,
	}

	return server, nil
}
