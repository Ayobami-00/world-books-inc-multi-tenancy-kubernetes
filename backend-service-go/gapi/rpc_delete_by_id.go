package gapi

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/usecase"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/api_response"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/validator"
)

func (s *Server) DeleteById(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {

	err := validator.ValidateDeleteBookRequest(req)
	if err != nil {
		return api_response.ErrorBadRequest(&pb.DeleteBookResponse{}, err.Error()).(*pb.DeleteBookResponse), nil
	}

	br := repository.NewBookRepository(s.Db, s.Timeout, s.Db.BooksTableName)
	dbiu := usecase.NewDeleteByIdUsecase(br)

	bookId := req.GetId()

	err = dbiu.DeleteById(ctx, bookId)
	if err != nil {
		return api_response.ErrorInternal(&pb.DeleteBookResponse{}, err.Error()).(*pb.DeleteBookResponse), nil
	}

	return api_response.Success(&pb.DeleteBookResponse{}, "Book successfully fetched", map[string]interface{}{}).(*pb.DeleteBookResponse), nil

}
