package usecase

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
)

type deleteByIdUsecase struct {
	bookRepository repository.BookRepository
}

type DeleteByIdUsecase interface {
	DeleteById(ctx context.Context, id string) error
}

func NewDeleteByIdUsecase(bookRepository repository.BookRepository) DeleteByIdUsecase {
	return &deleteByIdUsecase{
		bookRepository: bookRepository,
	}
}

func (dbiu *deleteByIdUsecase) DeleteById(ctx context.Context, id string) error {
	return dbiu.bookRepository.Delete(ctx, id)
}
