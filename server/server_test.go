package server_test

import (
	"context"
	"testing"

	"blog/db"
	"blog/pb"
	"blog/server"

	"github.com/stretchr/testify/assert"
)






func Test_Server(t *testing.T) {
	mockDB := db.NewDBMock()

	ctx := context.TODO()

	post := &pb.CreatePostRequest{
		Content: "content",
		AuthorId: 1234,
	}


	mockCall:= mockDB.On("CreatePost", ctx, post).Return("123456", nil)

	 srv := server.Server{
		DBConn: mockDB,
	 }

	 res, err := srv.CreatePost(ctx, post)
	 assert.NoError(t, err)

	 assert.Equal(t, res.PostId, "123456")

	 mockDB.AssertExpectations(t)

	 mockCall.Unset()
}