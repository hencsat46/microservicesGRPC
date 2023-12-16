package controller

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Create(string) int
	Get(int) string
	Delete(string) error
}
