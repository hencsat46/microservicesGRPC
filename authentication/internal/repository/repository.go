package repository

import "microservicesGRPC/authentication/internal/models"

type repo struct {
	data []models.User
}

func New() *repo {
	return &repo{data: make([]models.User, 1, 20)}
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

func (r *repo) Read(id int) (*models.User, error) {
	return &r.data[id], nil
}

func (r *repo) Update(user *models.User) error {
	for key, value := range r.data {
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
			deleteElement(r.data, key)
		}
	}
	return nil
}

func deleteElement(slice []models.User, s int) []models.User {
	return append(slice[:s], slice[s+1:]...)
}
