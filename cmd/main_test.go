package main_test

import (
	"fibo/cmd/migrate"
	"fibo/internal/api/comment"
	"fibo/internal/api/post"
	"fibo/internal/api/user"
	ca "fibo/internal/comment/adapter"
	cd "fibo/internal/comment/domain"
	cuc "fibo/internal/comment/usecases"
	pa "fibo/internal/post/adapter"
	pd "fibo/internal/post/domain"
	puc "fibo/internal/post/usecases"
	ua "fibo/internal/user/adapter"
	ud "fibo/internal/user/domain"
	uuc "fibo/internal/user/usecases"
	"testing"
)

var TestCommentMySQL cd.CommentRepository
var TestPostMySQL pd.PostRepository
var TestUserMySQL ud.UserRepository

var TestCommentUseCases cuc.CommentUseCases
var TestPostUseCases puc.PostUseCases
var TestUserUseCases uuc.UserUseCases

var TestCommentHandler comment.Handler
var TestPostHandler post.Handler
var TestUserHandler user.Handler

func TestMain(m *testing.M) {
	testRepo := migrate.NewMySQLRepository()
	TestCommentMySQL = ca.NewMySQLCommentRepository(testRepo)
	TestPostMySQL = pa.NewMySQLPostRepository(testRepo)
	TestUserMySQL = ua.NewMySQLUserRepository(testRepo)

	TestCommentUseCases = cuc.NewCommentRepository(TestCommentMySQL)
	TestPostUseCases = puc.NewPostUseCases(TestPostMySQL)
	TestUserUseCases = uuc.NewUserUseCases(TestUserMySQL)

	TestCommentHandler = comment.NewHandler(TestCommentUseCases)
	TestPostHandler = post.NewHandler(TestPostUseCases)
	TestUserHandler = user.NewHandler(TestUserUseCases)

	//app := fiber.New()
	//
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("hello world!")
	//})
	//
	//comment.Router(app.Group("/comments"), commentHandler)
	//post.Rounter(app.Group("/posts"), postHandler)
	//user.Router(app.Group("/users"), userHandler)
	//
	//_ = app.Listen(":3000")

}
