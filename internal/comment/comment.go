package comment

import (
	"context"
	"errors"
	"fmt"
)
var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented = errors.New("Not implemented")
)

// Comment - a representation of the comment structure for service
type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

//Store interface
type Store interface{
	GetComment(context.Context, string)(Comment, error)
}

//Service
type Service struct{
	Store Store
}

//NewService - returns pointers to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store, 
	}
}

//Return comments
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error){
	fmt.Println("Retrieving a comment")
	// ctx = context.WithValue(ctx, "request_id", "unique-string") //to:DO
	// fmt.Println(ctx.Value("Request_id"))
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil{
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment (ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment (ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment (ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
