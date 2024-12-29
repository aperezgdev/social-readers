package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/social-readers-api/internal/application/book_to_read/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/book_to_read/finder"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testSuiteBookToRead struct {
	bookToReadController *BookToReadController
	bookToReadRepository *repository.MockBookToReadRepository
	userRepository       *repository.MockUserRepository
}

func setupSuiteBookToRead() testSuiteBookToRead {
	bookToReadRepository := &repository.MockBookToReadRepository{}
	userRepository := &repository.MockUserRepository{}
	bookToReadCreator := create.NewBookToReadCreator(slog.Default(), bookToReadRepository, userRepository)
	bookToReadFinderByUser := finder.NewBookToReadFinderByUser(slog.Default(), bookToReadRepository)
	bookToReadController := NewBookToReadsController(bookToReadFinderByUser, bookToReadCreator)

	return testSuiteBookToRead{
		bookToReadController,
		bookToReadRepository,
		userRepository,
	}
}

func TestGetBookToRead(t *testing.T) {
	t.Parallel()

	t.Run("should get book to read", func(t *testing.T) {
		suite := setupSuiteBookToRead()
		suite.bookToReadRepository.On("FindByUser", mock.Anything, mock.Anything).Return(
			[]models.BookToRead{
				{
					Id: "1",
				},
			}, nil)
		suite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, nil)

		r := httptest.NewRequest(http.MethodGet, "/user/{userId}/books-to-read", nil)
		r.SetPathValue("userId", "1")
		w := httptest.NewRecorder()


		suite.bookToReadController.GetBooksToReadByUser(w, r)

		var response []map[string]string
		errUnmarshal := json.Unmarshal(w.Body.Bytes(), &response)
		if errUnmarshal != nil {
			t.Fatal(errUnmarshal)
		}
		suite.bookToReadRepository.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "1", response[0]["id"])
	})

	t.Run("should not get book to read", func(t *testing.T) {
		suite := setupSuiteBookToRead()
		suite.bookToReadRepository.On("FindByUser", mock.Anything, mock.Anything).Return(
			[]models.BookToRead{}, nil)
		suite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, nil)

		r := httptest.NewRequest(http.MethodGet, "/user/{userId}/books-to-read", nil)
		r.SetPathValue("userId", "1")
		w := httptest.NewRecorder()

		suite.bookToReadController.GetBooksToReadByUser(w, r)

		suite.bookToReadRepository.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestPostBookToRead(t *testing.T) {
	t.Parallel()

	t.Run("should create book to read", func(t *testing.T) {
		suite := setupSuiteBookToRead()
		suite.userRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, nil).
			Once()
		suite.bookToReadRepository.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
		
		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}

		json := []byte(fmt.Sprintf(`{"isbn":"978-2-3571-9106-8","title":"title","description":"description","picture":"picture","userId":"%s"}`, uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/user/{userId}/books-to-read", bytes.NewBuffer(json))
		r.SetPathValue("userId", uuid.String())
		w := httptest.NewRecorder()

		suite.bookToReadController.PostBookToRead(w, r)

		suite.userRepository.AssertExpectations(t)
		suite.bookToReadRepository.AssertExpectations(t)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should fail when user not exist", func(t *testing.T) {
		suite := setupSuiteBookToRead()
		suite.userRepository.On("Find", mock.Anything, mock.Anything).
			Return(models.User{}, errors.ErrNotExistUser)

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}

		json := []byte(fmt.Sprintf(`{"isbn":"978-2-3571-9106-8","title":"title","description":"description","picture":"picture","userId":"%s"}`, uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/user/{userId}/books-to-read", bytes.NewBuffer(json))
		r.SetPathValue("userId", "1")
		w := httptest.NewRecorder()

		suite.bookToReadController.PostBookToRead(w, r)
		suite.userRepository.AssertExpectations(t)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should fail when book to read has invalid isbn", func(t *testing.T) {
		suite := setupSuiteBookToRead()

		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}

		json := []byte(fmt.Sprintf(`{"isbn":"978-2-571-9106-9","title":"title","description":"description","picture":"picture","userId":"%s"}`, uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/user/{userId}/books-to-read", bytes.NewBuffer(json))
		r.SetPathValue("userId", uuid.String())
		w := httptest.NewRecorder()

		suite.bookToReadController.PostBookToRead(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}