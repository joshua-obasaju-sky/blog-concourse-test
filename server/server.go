package server

import (
	"blog/db"
	"blog/pb"
	context "context"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)




type Server struct {
 pb.UnimplementedBlogServer
 DBConn db.DBInterface
}



func (s *Server) CreatePost(ctx context.Context, body *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	
	errs := isValidRequest(body)

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	 uid, err := s.DBConn.CreatePost(ctx, body)
	 if err != nil {
		return nil, err
	 }

	return &pb.CreatePostResponse{
		PostId: uid,
	}, nil

}


func (s *Server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	 
	if req.ID == "" {
		return nil, errors.New("ID is not defined")
	}

	_, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}


	post, err := s.DBConn.GetPost(ctx, req.GetID())
	if err != nil {
		return nil, err
	}


	return post, nil
}


func (s *Server) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	 

	posts, err := s.DBConn.GetPosts(ctx)
	if err != nil {
		return nil, err
	}


	return &pb.GetPostsResponse{
		Posts: posts,
	}, nil
}


func isValidRequest(body *pb.CreatePostRequest) (errs []error) {
	if body.Content == "" {
		errs = append(errs, errors.New("Content is required"))
	}

	if body.AuthorId == 0 {
		errs = append(errs, errors.New("AuthorId is required"))
	}

	return 
}


func StartServer() error {


	db := db.NewDB()

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	log.Println("Postgres DB Connected at port 5432")

	lis, err := net.Listen("tcp", ":8080")

   if err != nil {
      return err
   }

   s := grpc.NewServer()
   server := &Server{
	DBConn: db,
   }

   pb.RegisterBlogServer(s, server)

   log.Printf("Server listening at %v", lis.Addr())
   if err := s.Serve(lis); err != nil {
      return err
   }


   return nil
}