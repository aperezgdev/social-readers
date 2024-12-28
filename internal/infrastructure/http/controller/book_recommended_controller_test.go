package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/create"
	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/finder"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testSuiteBookRecommended struct {
	bookRecommendedController *BookRecommendedController
	userRepository *repository.MockUserRepository
	bookRecommendedRepository *repository.MockBookRecommendedRepository
}

func setupSuiteBookRecommended() testSuiteBookRecommended {
	bookRecommendedRepository := &repository.MockBookRecommendedRepository{}
	userRepository := &repository.MockUserRepository{}
	bookRecommendedCreator := create.NewBookRecommendedCreator(slog.Default(), bookRecommendedRepository, userRepository)
	bookRecommendedFinderByUser := finder.NewBookRecommendedFinderByUser(slog.Default(), bookRecommendedRepository)
	bookRecommendedController := NewBookRecommendedController(bookRecommendedFinderByUser, bookRecommendedCreator)

	return testSuiteBookRecommended{
		bookRecommendedController,
		userRepository,
		bookRecommendedRepository,
	}
}

func TestGetBookRecommended(t *testing.T) {
	t.Parallel()

	t.Run("should get book recommended", func(t *testing.T) {
		suite := setupSuiteBookRecommended()
		suite.bookRecommendedRepository.On("FindByUser", mock.Anything, mock.Anything).Return(
			[]models.BookRecommended{
				{
					Id: "1",
				},
			}, nil)
		suite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, nil)

		r := httptest.NewRequest(http.MethodGet, "/user/{userId}/books-recommended", nil)
		r.SetPathValue("userId", "1")
		w := httptest.NewRecorder()

		suite.bookRecommendedController.GetBookRecommendedByUser(w, r)

		var response []map[string]string
		errUnmarshal := json.Unmarshal(w.Body.Bytes(), &response)
		if errUnmarshal != nil {
			t.Fatal(errUnmarshal)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "1", response[0]["id"])
	})

	t.Run("should return internal server error", func(t *testing.T) {
		suite := setupSuiteBookRecommended()
		suite.bookRecommendedRepository.On("FindByUser", mock.Anything, mock.Anything).Return(
			[]models.BookRecommended{}, assert.AnError)
		suite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, nil)

		r := httptest.NewRequest(http.MethodGet, "/user/{userId}/books-recommended", nil)
		r.SetPathValue("userId", "1")
		w := httptest.NewRecorder()

		suite.bookRecommendedController.GetBookRecommendedByUser(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestPostBookRecommended(t *testing.T) {
	t.Parallel()

	t.Run("should create book recommended", func(t *testing.T) {
		suite := setupSuiteBookRecommended()
		suite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, nil)
		suite.bookRecommendedRepository.On("Save", mock.Anything, mock.Anything).Return(nil)

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		json := fmt.Sprintf(`{"isbn":"978-6-6795-0881-8","title":"Hobbit","description":"Description","picture":"picture","userId":"%s"}`, uuid.String()) 
		r := httptest.NewRequest(http.MethodPost, "/user/{userId}/books-recommended", bytes.NewBuffer([]byte(json)))
		r.SetPathValue("userId", uuid.String())
		w := httptest.NewRecorder()

		suite.bookRecommendedController.PostBookRecommended(w, r)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return not found user on create book recommended", func(t *testing.T) { 
		suite := setupSuiteBookRecommended()
		suite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, errors.ErrNotExistUser)

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		json := fmt.Sprintf(`{"isbn":"978-6-6795-0881-8","title":"Hobbit","description":"Description","picture":"picture","userId":"%s"}`, uuid.String()) 
		r := httptest.NewRequest(http.MethodPost, "/user/{userId}/books-recommended", bytes.NewBuffer([]byte(json)))
		r.SetPathValue("userId", uuid.String())
		w := httptest.NewRecorder()

		suite.bookRecommendedController.PostBookRecommended(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should return bad request on invalid book recommended", func(t *testing.T) {
		suite := setupSuiteBookRecommended()

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		json := fmt.Sprintf(`{"isbn":"978-6-675-0881-8","title":"Hobbit","description":"Description","picture":"picture","userId":"%s"}`, uuid.String()) 
		r := httptest.NewRequest(http.MethodPost, "/user/{userId}/books-recommended", bytes.NewBuffer([]byte(json)))
		r.SetPathValue("userId", uuid.String())
		w := httptest.NewRecorder()

		suite.bookRecommendedController.PostBookRecommended(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}