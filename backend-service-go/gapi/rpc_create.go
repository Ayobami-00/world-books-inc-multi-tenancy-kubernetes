package gapi

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/usecase"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/api_response"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/validator"
)

func (s *Server) Create(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {

	err := validator.ValidateCreateRequest(req)
	if err != nil {
		return api_response.ErrorBadRequest(&pb.CreateBookResponse{}, err.Error()).(*pb.CreateBookResponse), nil
	}

	br := repository.NewBookRepository(s.Db, s.Timeout, s.Db.BooksTableName)
	cu := usecase.NewCreateUsecase(br)

	book := domain.ParseDomainFromReq(req).(domain.Book)

	err = cu.Create(ctx, &book)
	if err != nil {
		return api_response.ErrorInternal(&pb.CreateBookResponse{}, err.Error()).(*pb.CreateBookResponse), nil
	}

	return api_response.Success(&pb.CreateBookResponse{}, "Book successfully created", map[string]interface{}{}).(*pb.CreateBookResponse), nil

}
