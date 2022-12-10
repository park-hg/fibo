package post

import (
	"fibo/internal/post/usecases"
)

type handler struct {
	puc usecases.PostUseCases
}

func NewHandler(puc usecases.PostUseCases) Handler {
	return &handler{puc: puc}
}
