package main

import (
	"fibo/cmd/migrate"
	"fibo/internal/api/comment"
	"fibo/internal/api/post"
	"fibo/internal/api/user"
	ca "fibo/internal/comment/adapter"
	cuc "fibo/internal/comment/usecases"
	pa "fibo/internal/post/adapter"
	puc "fibo/internal/post/usecases"
	ua "fibo/internal/user/adapter"
	uuc "fibo/internal/user/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {

	repo := migrate.NewMySQLRepository()

	commentUseCases := cuc.NewCommentRepository(ca.NewMySQLCommentRepository(repo))
	postUseCases := puc.NewPostUseCases(pa.NewMySQLPostRepository(repo))
	userUseCases := uuc.NewUserUseCases(ua.NewMySQLUserRepository(repo))

	commentHandler := comment.NewHandler(commentUseCases)
	postHandler := post.NewHandler(postUseCases)
	userHandler := user.NewHandler(userUseCases)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})

	comment.Router(app.Group("/comments"), commentHandler)
	post.Rounter(app.Group("/posts"), postHandler)
	user.Router(app.Group("/users"), userHandler)

	_ = app.Listen(":3000")
}
