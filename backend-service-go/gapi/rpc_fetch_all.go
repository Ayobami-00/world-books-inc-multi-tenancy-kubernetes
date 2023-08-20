package gapi

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/usecase"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/api_response"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
)

func (s *Server) FetchAll(ctx context.Context, req *pb.FetchAllBooksRequest) (*pb.FetchAllBooksResponse, error) {

	br := repository.NewBookRepository(s.Db, s.Timeout, s.Db.BooksTableName)
	fau := usecase.NewFetchAllUsecase(br)

	books, err := fau.FetchAll(ctx)
	if err != nil {
		return api_response.ErrorInternal(&pb.FetchAllBooksResponse{}, err.Error()).(*pb.FetchAllBooksResponse), nil
	}

	return api_response.Success(&pb.FetchAllBooksResponse{}, "Books successfully fetched", map[string]interface{}{
		"books": books,
	}).(*pb.FetchAllBooksResponse), nil

}
