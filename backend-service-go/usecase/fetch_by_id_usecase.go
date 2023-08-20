package usecase

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
)

type fetchByIdlUsecase struct {
	bookRepository repository.BookRepository
}

type FetchByIdUsecase interface {
	FetchById(ctx context.Context, id string) (domain.Book, error)
}

func NewFetchByIdUsecase(bookRepository repository.BookRepository) FetchByIdUsecase {
	return &fetchByIdlUsecase{
		bookRepository: bookRepository,
	}
}

func (cu *fetchByIdlUsecase) FetchById(ctx context.Context, id string) (domain.Book, error) {
	return cu.bookRepository.GetByID(ctx, id)
}
