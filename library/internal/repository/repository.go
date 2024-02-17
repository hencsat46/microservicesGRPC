package repository

import (
	"microservicesGRPC/library/internal/controller"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repository struct {
	data []string
}

func New() controller.RepositoryInterfaces {
	return &repository{data: make([]string, 0, 20)}
}

func (r *repository) Create(username string) error {
	r.data = append(r.data, username)
	return nil
}

func (r *repository) Get(username string) bool {
	for i := 0; i < len(r.data); i++ {
		if username == r.data[i] {
			return true
		}
	}
	return false
}

func (r *repository) Delete(username string) error {
	for key, value := range r.data {
		if value == username {
			r.data = deleteElement(r.data, key)
			return nil
		}
	}
	return status.Error(codes.NotFound, "Not Found")
}

func deleteElement(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
