package domain

import (
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/methods"
)

type Book struct {
	ID          string `json:"id"`
	Name        string `json:"username"`
	CreatedBy   string `json:"created_by"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func GenerateBookID() string {

	return methods.GenerateRandomID()
}

func DomainBooksToPBBooks(books []Book) []*pb.Book {

	pbBooks := make([]*pb.Book, len(books))

	if len(books) == 0 {
		return pbBooks
	}

	for i, book := range books {
		pbBooks[i] = &pb.Book{
			// Perform the necessary field mapping here
			Id:          book.ID,
			CreatedBy:   book.CreatedBy,
			Name:        book.Name,
			Description: book.Description,
			Quantity:    book.Quantity,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		}
	}

	return pbBooks
}

func DomainBookToPBBook(book Book) *pb.Book {

	return &pb.Book{
		// Perform the necessary field mapping here
		Id:          book.ID,
		CreatedBy:   book.CreatedBy,
		Name:        book.Name,
		Description: book.Description,
		Quantity:    book.Quantity,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}
}
