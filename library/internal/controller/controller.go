package controller

import (
	"microservicesGRPC/library/internal/handler"
)

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Create(string) error
	Get(string) bool
	// Delete(string) error
}

func NewUsecase(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Create(username string) error {

	if err := u.repo.Create(username); err != nil {
		return err
	}
	return nil
}

func (u *usecase) Get(username string) bool {
	return u.repo.Get(username)
}
