package controller

import "microservicesGRPC/library/internal/handler"

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Create(string) int
	// Get(int) string
	// Delete(string) error
}

func NewUsecase(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Create(username string) (int, error) {
	id := u.repo.Create(username)
	return id, nil
}
