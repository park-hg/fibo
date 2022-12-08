package api

import (
	_interface "fibo/external/api/interface"
	cusecases "fibo/internal/comment/usecases"
	pusecases "fibo/internal/post/usecases"
	uusecases "fibo/internal/user/usecases"
)

type commentHandler struct {
	commentUseCases cusecases.CommentUseCases
}

type postHandler struct {
	postUseCases pusecases.PostUseCases
}

type userHandler struct {
	userUseCases uusecases.UserUseCases
}

func NewHandler(
	cuc cusecases.CommentUseCases,
	puc pusecases.PostUseCases,
	uuc uusecases.UserUseCases)(
	_interface.CommentHandler,
	_interface.PostHandler,
	_interface.UserHandler) {

	return (
		&commentHandler{commentUseCases: cuc},
		&postHandler{postUseCases: puc},
		&userHandler{userUseCases: uuc}
		)
}

