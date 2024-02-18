package controller

import (
	"log"
	"microservicesGRPC/library/internal/handler"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Create(string) error
	Get(string) bool
	Delete(string) error
}

func NewUsecase(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Create(username, authUsername string) error {

	if authUsername != "Not Found" {
		if err := u.repo.Create(username); err != nil {
			return err
		}
	}

	return status.Error(codes.NotFound, "not found")
}

func (u *usecase) Get(username string) bool {
	return u.repo.Get(username)
}

func (u *usecase) Delete(username string) error {
	if err := u.repo.Delete(username); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
