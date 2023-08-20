package domain

import (
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/utils/methods"
)

func ParseDomainFromReq(request interface{}) interface{} {

	switch r := request.(type) {

	case *pb.CreateBookRequest:

		return Book{
			ID:          GenerateBookID(),
			Name:        r.GetName(),
			CreatedBy:   r.GetCreatedBy(),
			Description: r.GetDescription(),
			Quantity:    r.GetQuantity(),
			CreatedAt:   methods.GetCurrentTimestamp(),
			UpdatedAt:   methods.GetCurrentTimestamp(),
		}

	case *pb.UpdateBookRequest:

		return Book{
			ID:          r.GetBook().Id,
			Name:        r.GetBook().Name,
			Description: r.GetBook().Description,
			Quantity:    r.GetBook().Quantity,
		}

	default:

		return nil

	}
}
