package usecase

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
)

type fetchAllUsecase struct {
	bookRepository repository.BookRepository
}

type FetchAllUsecase interface {
	FetchAll(ctx context.Context) ([]domain.Book, error)
}

func NewFetchAllUsecase(bookRepository repository.BookRepository) FetchAllUsecase {
	return &fetchAllUsecase{
		bookRepository: bookRepository,
	}
}

func (cu *fetchAllUsecase) FetchAll(ctx context.Context) ([]domain.Book, error) {
	return cu.bookRepository.FetchAll(ctx)
}
