package repository

import (
	"log"
	"microservicesGRPC/authentication/internal/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repo struct {
	data []models.User
}

func New() *repo {
	return &repo{data: make([]models.User, 0, 20)}
}

func (r *repo) Create(user *models.User) (int, error) {
	r.data = append(r.data, *user)

	id := r.getId(user)
	return id, nil
}

func (r *repo) getId(user *models.User) int {
	for key, value := range r.data {
		if *user == value {
			return key
		}
	}
	return -1
}

func (r *repo) Read(username string) (*models.User, error) {
	for i := 0; i < len(r.data); i++ {
		if username == r.data[i].Username {
			return &r.data[i], nil
		}
	}
	return nil, status.Error(codes.NotFound, "not found")
}

func (r *repo) Update(user *models.User) error {
	for key, value := range r.data {
		log.Println(key, value)
		if value.Username == user.Username {
			r.data[key] = *user
			return nil
		}
	}
	return nil
}

func (r *repo) Delete(user *models.User) error {
	for key, value := range r.data {
		if value.Username == user.Username {
			r.data = deleteElement(r.data, key)
		}
	}
	return nil
}

func deleteElement(slice []models.User, s int) []models.User {
	return append(slice[:s], slice[s+1:]...)
}
