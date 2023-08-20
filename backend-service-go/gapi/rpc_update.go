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

func (s *Server) Update(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {

	err := validator.ValidateUpdateRequest(req)
	if err != nil {
		return api_response.ErrorBadRequest(&pb.UpdateBookRequest{}, err.Error()).(*pb.UpdateBookResponse), nil
	}

	br := repository.NewBookRepository(s.Db, s.Timeout, s.Db.BooksTableName)
	uu := usecase.NewUpdateUsecase(br)

	book := domain.ParseDomainFromReq(req).(domain.Book)

	err = uu.Update(ctx, &book)
	if err != nil {
		return api_response.ErrorInternal(&pb.UpdateBookResponse{}, err.Error()).(*pb.UpdateBookResponse), nil
	}

	return api_response.Success(&pb.UpdateBookResponse{}, "Book successfully updated", map[string]interface{}{}).(*pb.UpdateBookResponse), nil
}
