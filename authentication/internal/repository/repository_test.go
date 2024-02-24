package repository

import (
	"microservicesGRPC/authentication/internal/models"
	"testing"

	assert "github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	testcases := []struct {
		user           models.User
		expectedResult int
		expectedError  error
	}{
		{user: models.User{Username: "lskdf", Password: "1234", FirstName: "hello", SecondName: "shit"}, expectedResult: 0, expectedError: nil},
		{user: models.User{Username: "asdfas", Password: "1sdf234", FirstName: "asdfad", SecondName: "sdfqer"}, expectedResult: 1, expectedError: nil},
		{user: models.User{Username: "2312", Password: "asdgsfgn", FirstName: "gtw", SecondName: "taer"}, expectedResult: 2, expectedError: nil},
		{user: models.User{Username: "asdf", Password: "34rfr", FirstName: "tgrfsd", SecondName: "sdgts"}, expectedResult: 3, expectedError: nil},
		{user: models.User{Username: "dsa", Password: "4re435f", FirstName: "afsdvxs", SecondName: "gtfrgs"}, expectedResult: 4, expectedError: nil},
		{user: models.User{Username: "gbhty", Password: "atgfas", FirstName: "34rgfw", SecondName: "grwrdfgt"}, expectedResult: 5, expectedError: nil},
	}

	repo := &repo{make([]models.User, 0, 20)}

	for _, pair := range testcases {
		got, err := repo.Create(&pair.user)

		assert.Equal(t, pair.expectedResult, got)
		assert.Equal(t, pair.expectedError, err)
	}
}

func TestRead(t *testing.T) {

	data := []models.User{
		{Username: "lskdf", Password: "1234", FirstName: "hello", SecondName: "shit"},
		{Username: "asdfas", Password: "1sdf234", FirstName: "asdfad", SecondName: "sdfqer"},
		{Username: "2312", Password: "asdgsfgn", FirstName: "gtw", SecondName: "taer"},
		{Username: "asdf", Password: "34rfr", FirstName: "tgrfsd", SecondName: "sdgts"},
		{Username: "dsa", Password: "4re435f", FirstName: "afsdvxs", SecondName: "gtfrgs"},
		{Username: "gbhty", Password: "atgfas", FirstName: "34rgfw", SecondName: "grwrdfgt"},
	}

	repo := &repo{data: data}

	testcases := []struct {
		username       string
		expectedResult *models.User
		expectedError  error
	}{
		{username: "lskdf", expectedResult: &models.User{Username: "lskdf", Password: "1234", FirstName: "hello", SecondName: "shit"}, expectedError: nil},
		{username: "asdfas", expectedResult: &models.User{Username: "asdfas", Password: "1sdf234", FirstName: "asdfad", SecondName: "sdfqer"}, expectedError: nil},
		{username: "2312", expectedResult: &models.User{Username: "2312", Password: "asdgsfgn", FirstName: "gtw", SecondName: "taer"}, expectedError: nil},
		{username: "asdf", expectedResult: &models.User{Username: "asdf", Password: "34rfr", FirstName: "tgrfsd", SecondName: "sdgts"}, expectedError: nil},
		{username: "dsa", expectedResult: &models.User{Username: "dsa", Password: "4re435f", FirstName: "afsdvxs", SecondName: "gtfrgs"}, expectedError: nil},
		{username: "gbhty", expectedResult: &models.User{Username: "gbhty", Password: "atgfas", FirstName: "34rgfw", SecondName: "grwrdfgt"}, expectedError: nil},
		{username: "111", expectedResult: nil, expectedError: status.Error(codes.NotFound, "not found")},
	}

	for _, pair := range testcases {
		got, err := repo.Read(pair.username)

		assert.Equal(t, pair.expectedResult, got)
		assert.Equal(t, pair.expectedError, err)
	}
}

func TestUpdate(t *testing.T) {

	repo := &repo{make([]models.User, 0, 20)}

	repo.Create(&models.User{Username: "lskdf", Password: "1234", FirstName: "hello", SecondName: "shit"})
	repo.Create(&models.User{Username: "asdfas", Password: "1sdf234", FirstName: "asdfad", SecondName: "sdfqer"})
	repo.Create(&models.User{Username: "2312", Password: "asdgsfgn", FirstName: "gtw", SecondName: "taer"})
	repo.Create(&models.User{Username: "asdf", Password: "34rfr", FirstName: "tgrfsd", SecondName: "sdgts"})
	repo.Create(&models.User{Username: "dsa", Password: "4re435f", FirstName: "afsdvxs", SecondName: "gtfrgs"})
	repo.Create(&models.User{Username: "gbhty", Password: "atgfas", FirstName: "34rgfw", SecondName: "grwrdfgt"})

	testcases := []struct {
		user          *models.User
		expectedError error
	}{
		{user: &models.User{Username: "lskdf", Password: "rtgsd", FirstName: "hello", SecondName: "shit"}, expectedError: nil},
		{user: &models.User{Username: "asdfas", Password: "sgtr", FirstName: "asdfad", SecondName: "sdfqer"}, expectedError: nil},
		{user: &models.User{Username: "2312", Password: "strg", FirstName: "gtw", SecondName: "taer"}, expectedError: nil},
		{user: &models.User{Username: "asdf", Password: "srtgs", FirstName: "tgrfsd", SecondName: "sdgts"}, expectedError: nil},
		{user: &models.User{Username: "dsa", Password: "srtgsr", FirstName: "afsdvxs", SecondName: "gtfrgs"}, expectedError: nil},
		{user: &models.User{Username: "gbhty", Password: "atgfsgsggas", FirstName: "34rgfw", SecondName: "grwrdfgt"}, expectedError: nil},
		{user: &models.User{Username: "1111"}, expectedError: status.Error(codes.NotFound, "not found")},
	}

	for _, pair := range testcases {
		err := repo.Update(pair.user)

		assert.Equal(t, pair.expectedError, err)
	}

}

func TestDelete(t *testing.T) {
	testcases := []struct {
		user          *models.User
		expectedError error
	}{
		{user: &models.User{Username: "lskdf", Password: "rtgsd", FirstName: "hello", SecondName: "shit"}, expectedError: nil},
		{user: &models.User{Username: "asdfas", Password: "sgtr", FirstName: "asdfad", SecondName: "sdfqer"}, expectedError: nil},
		{user: &models.User{Username: "2312", Password: "strg", FirstName: "gtw", SecondName: "taer"}, expectedError: nil},
		{user: &models.User{Username: "asdf", Password: "srtgs", FirstName: "tgrfsd", SecondName: "sdgts"}, expectedError: nil},
		{user: &models.User{Username: "dsa", Password: "srtgsr", FirstName: "afsdvxs", SecondName: "gtfrgs"}, expectedError: nil},
		{user: &models.User{Username: "gbhty", Password: "atgfsgsggas", FirstName: "34rgfw", SecondName: "grwrdfgt"}, expectedError: nil},
		{user: &models.User{Username: "1111"}, expectedError: status.Error(codes.NotFound, "not found")},
	}

	repo := &repo{make([]models.User, 0, 20)}

	repo.Create(&models.User{Username: "lskdf", Password: "1234", FirstName: "hello", SecondName: "shit"})
	repo.Create(&models.User{Username: "asdfas", Password: "1sdf234", FirstName: "asdfad", SecondName: "sdfqer"})
	repo.Create(&models.User{Username: "2312", Password: "asdgsfgn", FirstName: "gtw", SecondName: "taer"})
	repo.Create(&models.User{Username: "asdf", Password: "34rfr", FirstName: "tgrfsd", SecondName: "sdgts"})
	repo.Create(&models.User{Username: "dsa", Password: "4re435f", FirstName: "afsdvxs", SecondName: "gtfrgs"})
	repo.Create(&models.User{Username: "gbhty", Password: "atgfas", FirstName: "34rgfw", SecondName: "grwrdfgt"})

	for _, pair := range testcases {
		err := repo.Delete(pair.user)

		assert.Equal(t, pair.expectedError, err)
	}
}
