package controller

import (
	"log"
	"microservicesGRPC/authentication/internal/models"
)

type usecase struct {
	repo RepositoryIntefaces
}

type RepositoryIntefaces interface {
	Create(*models.User) (int, error)
	Read(string) (*models.User, error)
	Update(*models.User) error
	Delete(*models.User) error
}

func NewUsecase(repo RepositoryIntefaces) *usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(user *models.User) (int, error) {

	id, err := u.repo.Create(user)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return id, nil

}

func (u *usecase) Read(username string) (*models.User, error) {
	user, err := u.repo.Read(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (u *usecase) Update(user *models.User) error {

	if err := u.repo.Update(user); err != nil {
		log.Println(err)
		return nil
	}

	return nil
}

func (u *usecase) Delete(user *models.User) error {

	if err := u.repo.Delete(user); err != nil {
		return nil
	}

	return nil
}

func New(repo RepositoryIntefaces) *usecase {
	return &usecase{repo: repo}
}
