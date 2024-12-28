package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/social-readers-api/internal/application/user/create"
	domain_errors "github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
	user_vo "github.com/aperezgdev/social-readers-api/internal/domain/value_object/user"
	"github.com/aperezgdev/social-readers-api/internal/infrastructure/http/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.MockUserRepository{}
var userCreator = create.NewUserCreator(slog.Default(), userRepository)
var userFinder = finder.NewUserFinder(slog.Default(), userRepository)
var userController = controller.NewUserController(userCreator, userFinder)

func TestGetUser(t *testing.T) {
	t.Parallel()

	t.Run("should retrieve an existing user", func(t *testing.T) {
		userRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(models.User{Id: user_vo.UserId("123")}, nil)

		r := httptest.NewRequest(http.MethodGet, "/user/123", &bytes.Buffer{})
		w := httptest.NewRecorder()
		userController.GetUser(w, r)
		
		var response map[string]string
   		json.Unmarshal(w.Body.Bytes(), &response)

		assert.NotEmpty(t, response)
		assert.Equal(t, "123", response["id"])
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should not retrieve an not existing user", func(t *testing.T) {
		userRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(models.User{}, domain_errors.ErrNotExistUser)

		r := httptest.NewRequest(http.MethodGet, "/user/123", &bytes.Buffer{})
		w := httptest.NewRecorder()
		userController.GetUser(w, r)
		
		var response map[string]string
   		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Empty(t, response)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should return a internal server error when error ocurred on save", func(t *testing.T) {
		userRepository.On("Find", mock.Anything, mock.Anything).
		Once().
		Return(models.User{Id: user_vo.UserId("123")}, errors.New(""))

		r := httptest.NewRequest(http.MethodGet, "/user/123", &bytes.Buffer{})
		w := httptest.NewRecorder()
		userController.GetUser(w, r)

		var response map[string]string
   		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Empty(t, response)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestPostUser(t *testing.T) {
	t.Parallel()

	t.Run("should create a user and return created status", func(t *testing.T) {
		userRepository.On("Save", mock.Anything, mock.Anything).Once().Return(nil)
		json := []byte(`{"name": "alex", "mail": "john.doe@example.com", "picture": "picture"}`)
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(json))
		w := httptest.NewRecorder()

		userController.PostUser(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return bad request on an invalid body", func(t *testing.T) {
		userRepository.On("Save", mock.Anything, mock.Anything).Once().Return(nil)
		json := []byte(`{"a": "alex", "mail": "john.doe@example.com", "picture": "picture"}`)
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(json))
		w := httptest.NewRecorder()

		userController.PostUser(w, r)
		
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}