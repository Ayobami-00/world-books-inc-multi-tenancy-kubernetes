package validator

import (
	"fmt"
	"net/mail"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
)

const (
	minStringLength = 3
	maxStringLength = 100
)

func ValidateString(fieldName string, value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("%s must contain from %d-%d characters", fieldName, minLength, maxLength)
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString("password", value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString("email", value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateCreateRequest(req *pb.CreateBookRequest) error {
	if err := ValidateString("name", req.GetName(), minStringLength, maxStringLength); err != nil {
		return err
	}

	if err := ValidateString("created_by", req.GetCreatedBy(), minStringLength, maxStringLength); err != nil {
		return err
	}

	if err := ValidateString("description", req.GetCreatedBy(), minStringLength, maxStringLength); err != nil {
		return err
	}

	return nil
}

func ValidateFetchBookByIdRequest(req *pb.FetchBookByIdRequest) error {

	if err := ValidateString("id", req.GetId(), minStringLength, maxStringLength); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateRequest(req *pb.UpdateBookRequest) error {
	if err := ValidateString("name", req.GetBook().Name, minStringLength, maxStringLength); err != nil {
		return err
	}

	if err := ValidateString("name", req.GetBook().Description, minStringLength, maxStringLength); err != nil {
		return err
	}

	return nil
}

func ValidateDeleteBookRequest(req *pb.DeleteBookRequest) error {

	if err := ValidateString("id", req.GetId(), minStringLength, maxStringLength); err != nil {
		return err
	}

	return nil
}
