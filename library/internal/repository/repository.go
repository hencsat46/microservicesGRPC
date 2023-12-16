package repository

import "microservicesGRPC/library/internal/controller"

type repository struct {
	data []string
}

func New() controller.RepositoryInterfaces {
	return &repository{data: make([]string, 0, 20)}
}

func (r *repository) Create(username string) int {
	r.data = append(r.data, username)
	id := r.getId(username)
	return id
}

func (r *repository) getId(user string) int {
	for key, value := range r.data {
		if user == value {
			return key
		}
	}
	return -1
}

func (r *repository) Get(userId int) string {
	if len(r.data)-1 < userId {
		return "no user :("
	}
	return r.data[userId]
}

func (r *repository) Delete(username string) error {
	for key, value := range r.data {
		if value == username {
			r.data = deleteElement(r.data, key)
			return nil
		}
	}
	return nil
}

func deleteElement(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
