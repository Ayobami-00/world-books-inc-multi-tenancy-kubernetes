package usecase

import (
	"context"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/repository"
)

type createUsecase struct {
	bookRepository repository.BookRepository
}

type CreateUsecase interface {
	Create(ctx context.Context, user *domain.Book) error
}

func NewCreateUsecase(bookRepository repository.BookRepository) CreateUsecase {
	return &createUsecase{
		bookRepository: bookRepository,
	}
}

func (cu *createUsecase) Create(ctx context.Context, user *domain.Book) error {
	return cu.bookRepository.Create(ctx, user)
}
