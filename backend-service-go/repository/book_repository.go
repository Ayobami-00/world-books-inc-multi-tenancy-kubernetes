package repository

import (
	"context"
	"time"

	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/bootstrap"
	"github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/backend-service-go/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type bookRepository struct {
	Db        bootstrap.Database
	TableName string
	Timeout   time.Duration
}

type BookRepository interface {
	FetchAll(ctx context.Context) ([]domain.Book, error)
	Create(ctx context.Context, book *domain.Book) error
	Fetch(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (domain.Book, error)
	Update(ctx context.Context, book *domain.Book) error
}

func NewBookRepository(db bootstrap.Database, timeout time.Duration, tableName string) BookRepository {
	return &bookRepository{
		Db:        db,
		TableName: tableName,
		Timeout:   timeout,
	}
}

func (br *bookRepository) FetchAll(ctx context.Context) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, br.Timeout)
	defer cancel()

	input := &dynamodb.ScanInput{
		TableName: aws.String(br.TableName),
	}

	res, err := br.Db.Client.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	books := []domain.Book{}
	if err := dynamodbattribute.UnmarshalListOfMaps(res.Items, &books); err != nil {
		return books, err
	}

	return books, nil
}

func (br *bookRepository) Create(ctx context.Context, book *domain.Book) error {

	ctx, cancel := context.WithTimeout(ctx, br.Timeout)
	defer cancel()

	item, err := dynamodbattribute.MarshalMap(book)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(br.TableName),
		Item:      item,
		ExpressionAttributeNames: map[string]*string{
			"#id": aws.String("id"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#id)"),
	}

	if _, err := br.Db.Client.PutItemWithContext(ctx, input); err != nil {

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			return err
		}

		return err
	}

	return nil

}

func (br *bookRepository) Fetch(ctx context.Context) ([]domain.Book, error) {

	ctx, cancel := context.WithTimeout(ctx, br.Timeout)
	defer cancel()

	input := &dynamodb.ScanInput{
		TableName: aws.String(br.TableName),
	}

	res, err := br.Db.Client.ScanWithContext(ctx, input)

	if err != nil {

		return nil, err
	}

	users := []domain.Book{}

	if err := dynamodbattribute.UnmarshalListOfMaps(res.Items, &users); err != nil {

		return users, err
	}

	return users, nil

}

func (br *bookRepository) GetByID(ctx context.Context, id string) (domain.Book, error) {

	ctx, cancel := context.WithTimeout(ctx, br.Timeout)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String(br.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	}

	var book domain.Book

	res, err := br.Db.Client.GetItemWithContext(ctx, input)

	if err != nil {
		return book, err
	}

	if res.Item == nil {
		return book, err
	}

	if err := dynamodbattribute.UnmarshalMap(res.Item, &book); err != nil {

		return book, err
	}

	return book, nil

}

func (br *bookRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, br.Timeout)
	defer cancel()

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(br.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	}

	_, err := br.Db.Client.DeleteItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (br *bookRepository) Update(ctx context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(ctx, br.Timeout)
	defer cancel()

	item, err := dynamodbattribute.MarshalMap(book)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(br.TableName),
		Item:      item,
	}

	if _, err := br.Db.Client.PutItemWithContext(ctx, input); err != nil {

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			return err
		}

		return err
	}

	return err
}
