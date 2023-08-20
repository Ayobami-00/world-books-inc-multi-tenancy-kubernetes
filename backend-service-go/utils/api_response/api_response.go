package api_response

import (
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/pb"
	"google.golang.org/grpc/codes"
)

type ApiResponse interface {
	Success(message string, args map[string]interface{}) interface{}
	ErrorInternal(message string) interface{}
	ErrorBadRequest(message string) interface{}
	ErrorNotFound(message string) interface{}
	ErrorAlreadyExists(message string) interface{}
}

func Success(response interface{}, message string, args map[string]interface{}) interface{} {

	switch r := response.(type) {

	case *pb.FetchAllBooksResponse:

		rr := r
		rr.Code = int64(codes.OK)
		rr.Message = message
		rr.Data = &pb.FetchAllBooksResponseData{
			Books: domain.DomainBooksToPBBooks(args["books"].([]domain.Book)),
		}
		return rr

	case *pb.CreateBookResponse:

		rr := r
		rr.Code = int64(codes.OK)
		rr.Message = message
		return rr

	case *pb.FetchBookByIdResponse:

		rr := r
		rr.Code = int64(codes.OK)
		rr.Message = message
		rr.Book = domain.DomainBookToPBBook(args["book"].(domain.Book))
		return rr

	case *pb.UpdateBookResponse:

		rr := r
		rr.Code = int64(codes.OK)
		rr.Message = message
		return rr

	case *pb.DeleteBookResponse:

		rr := r
		rr.Code = int64(codes.OK)
		rr.Message = message
		return rr

	default:

		return nil

	}
}

func ErrorInternal(response interface{}, message string) interface{} {

	switch r := response.(type) {

	case *pb.FetchAllBooksResponse:

		rr := r
		rr.Code = int64(codes.Internal)
		rr.Message = message
		return rr

	case *pb.CreateBookResponse:

		rr := r
		rr.Code = int64(codes.Internal)
		rr.Message = message
		return rr

	case *pb.FetchBookByIdResponse:

		rr := r
		rr.Code = int64(codes.Internal)
		rr.Message = message
		return rr

	case *pb.UpdateBookResponse:

		rr := r
		rr.Code = int64(codes.Internal)
		rr.Message = message
		return rr

	case *pb.DeleteBookResponse:

		rr := r
		rr.Code = int64(codes.Internal)
		rr.Message = message
		return rr

	default:

		return nil

	}
}

func ErrorBadRequest(response interface{}, message string) interface{} {

	switch r := response.(type) {

	case *pb.FetchAllBooksResponse:

		rr := r
		rr.Code = int64(codes.InvalidArgument)
		rr.Message = message
		return rr

	case *pb.CreateBookResponse:

		rr := r
		rr.Code = int64(codes.InvalidArgument)
		rr.Message = message
		return rr

	case *pb.FetchBookByIdResponse:

		rr := r
		rr.Code = int64(codes.InvalidArgument)
		rr.Message = message
		return rr

	case *pb.UpdateBookResponse:

		rr := r
		rr.Code = int64(codes.InvalidArgument)
		rr.Message = message
		return rr

	case *pb.DeleteBookResponse:

		rr := r
		rr.Code = int64(codes.InvalidArgument)
		rr.Message = message
		return rr

	default:

		return nil

	}
}

func ErrorNotFound(response interface{}, message string) interface{} {

	switch r := response.(type) {

	case *pb.FetchAllBooksResponse:

		rr := r
		rr.Code = int64(codes.NotFound)
		rr.Message = message
		return rr

	case *pb.CreateBookResponse:

		rr := r
		rr.Code = int64(codes.NotFound)
		rr.Message = message
		return rr

	case *pb.FetchBookByIdResponse:

		rr := r
		rr.Code = int64(codes.NotFound)
		rr.Message = message
		return rr

	case *pb.UpdateBookResponse:

		rr := r
		rr.Code = int64(codes.NotFound)
		rr.Message = message
		return rr

	case *pb.DeleteBookResponse:

		rr := r
		rr.Code = int64(codes.NotFound)
		rr.Message = message
		return rr

	default:

		return nil

	}
}

func ErrorAlreadyExists(response interface{}, message string) interface{} {

	switch r := response.(type) {

	case *pb.FetchAllBooksResponse:

		rr := r
		rr.Code = int64(codes.AlreadyExists)
		rr.Message = message
		return rr

	case *pb.CreateBookResponse:

		rr := r
		rr.Code = int64(codes.AlreadyExists)
		rr.Message = message
		return rr

	case *pb.FetchBookByIdResponse:

		rr := r
		rr.Code = int64(codes.AlreadyExists)
		rr.Message = message
		return rr

	case *pb.UpdateBookResponse:

		rr := r
		rr.Code = int64(codes.AlreadyExists)
		rr.Message = message
		return rr

	case *pb.DeleteBookResponse:

		rr := r
		rr.Code = int64(codes.AlreadyExists)
		rr.Message = message
		return rr

	default:

		return nil

	}
}

func ErrorUnauthorized(response interface{}, message string) interface{} {

	switch r := response.(type) {

	case *pb.FetchAllBooksResponse:

		rr := r
		rr.Code = int64(codes.PermissionDenied)
		rr.Message = message
		return rr

	case *pb.CreateBookResponse:

		rr := r
		rr.Code = int64(codes.PermissionDenied)
		rr.Message = message
		return rr

	case *pb.FetchBookByIdResponse:

		rr := r
		rr.Code = int64(codes.PermissionDenied)
		rr.Message = message
		return rr

	case *pb.UpdateBookResponse:

		rr := r
		rr.Code = int64(codes.PermissionDenied)
		rr.Message = message
		return rr

	case *pb.DeleteBookResponse:

		rr := r
		rr.Code = int64(codes.PermissionDenied)
		rr.Message = message
		return rr

	default:

		return nil

	}
}
