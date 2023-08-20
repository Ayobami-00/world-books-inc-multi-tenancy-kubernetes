package gapi

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/usecase"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/api_response"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/validator"
)

func (s *Server) FetchById(ctx context.Context, req *pb.FetchBookByIdRequest) (*pb.FetchBookByIdResponse, error) {

	err := validator.ValidateFetchBookByIdRequest(req)
	if err != nil {
		return api_response.ErrorBadRequest(&pb.FetchBookByIdResponse{}, err.Error()).(*pb.FetchBookByIdResponse), nil
	}

	br := repository.NewBookRepository(s.Db, s.Timeout, s.Db.BooksTableName)
	fbiu := usecase.NewFetchByIdUsecase(br)

	bookId := req.GetId()

	book, err := fbiu.FetchById(ctx, bookId)
	if err != nil {
		return api_response.ErrorInternal(&pb.FetchBookByIdResponse{}, err.Error()).(*pb.FetchBookByIdResponse), nil
	}

	return api_response.Success(&pb.FetchBookByIdResponse{}, "Book successfully fetched", map[string]interface{}{
		"book": book,
	}).(*pb.FetchBookByIdResponse), nil

}
