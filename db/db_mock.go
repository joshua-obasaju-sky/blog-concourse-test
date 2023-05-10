package db

import (
	"blog/pb"
	"context"

	"github.com/stretchr/testify/mock"
)


type DBGoodMock struct {
	mock.Mock
}

func NewDBMock() *DBGoodMock {
	return new(DBGoodMock)
}


type DBBadMock struct {
	mock.Mock
}


var posts =  []*pb.Post{
		{
			Content: "content1",
			AuthorId: 123,
			Id: "12345",
		},
			{
			Content: "content2",
			AuthorId: 1234,
			Id: "12345",
		},
			{
			Content: "content3",
			AuthorId: 1235,
			Id: "12345",
		},
	}


func (d *DBBadMock) CreatePost(ctx context.Context, post *pb.CreatePostRequest) (string, error) {
    d.Called(ctx, post)

	return "1234", nil
}

func (d *DBBadMock) GetPosts(ctx context.Context) ([]*pb.Post, error) {
	args := d.Called(ctx)

	return args.Get(0).([]*pb.Post), args.Error(1)
}

func (d *DBBadMock) GetPost(ctx context.Context, ID string) (*pb.Post, error) {
	args := d.Called(ctx, ID)
  return args.Get(0).(*pb.Post), args.Error(1)
}



func (d *DBGoodMock) CreatePost(ctx context.Context, post *pb.CreatePostRequest) (string, error) {
	args:= d.Called(ctx, post)
	return args.String(0), args.Error(1)
}

func (d *DBGoodMock) GetPosts(ctx context.Context) ([]*pb.Post, error) {
	return posts, nil
}

func (d *DBGoodMock) GetPost(ctx context.Context, ID string) (*pb.Post, error) {
	return posts[0], nil
}

