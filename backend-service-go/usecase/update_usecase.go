package usecase

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
)

type updateUsecase struct {
	bookRepository repository.BookRepository
}

type UpdateUsecase interface {
	Update(ctx context.Context, book *domain.Book) error
}

func NewUpdateUsecase(bookRepository repository.BookRepository) UpdateUsecase {
	return &updateUsecase{
		bookRepository: bookRepository,
	}
}

func (cu *updateUsecase) Update(ctx context.Context, book *domain.Book) error {
	return cu.bookRepository.Update(ctx, book)
}
