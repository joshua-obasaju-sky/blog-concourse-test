package db

import (
	"blog/pb"
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)





type DB struct {
	*pg.DB
}

func NewDB() *DB{
  db := pg.Connect(&pg.Options{
    Addr:     ":5432",
    User:     "postgres",
    Password: "1234",
    Database: "blog-pract",
})


	return &DB{db}
}

func (d *DB)CreatePost(ctx context.Context, post *pb.CreatePostRequest) (string, error) {


	p := &pb.Post{
		Content: post.Content,
		 AuthorId: post.AuthorId,
		 Id: uuid.NewString(),
		 Date: timestamppb.New(time.Now()),
	}

	_, err := d.Model(p).Insert()
    if err != nil {
     return "", err
    }

  return p.Id, nil
}

func (d *DB) GetPosts(ctx context.Context) ([]*pb.Post, error) {

	var posts []*pb.Post
	
    err := d.Model(&posts).Select()
    if err != nil {
        return nil, err
    }

	return posts, nil
}

func (d *DB) GetPost(ctx context.Context, ID string) (*pb.Post, error) {

	post := &pb.Post{
		Id: ID,
	}
    err := d.Model(post).WherePK().Select()
    if err != nil {
        return nil, err
    }

	return post, nil
}


type DBInterface interface {
	CreatePost(ctx context.Context, post *pb.CreatePostRequest) (string, error)
	GetPosts(ctx context.Context) ([]*pb.Post, error)
	GetPost(ctx context.Context, ID string) (*pb.Post, error)
}

