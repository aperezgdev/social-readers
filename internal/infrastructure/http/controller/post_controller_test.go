package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/social-readers-api/internal/application/post/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/post/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testSuite struct {
	postRepository *repository.MockPostRepository
	userRepository *repository.MockUserRepository
	postController *PostController
}

func setupSuite () *testSuite {
	postRepository := repository.MockPostRepository{}
	userRepository := repository.MockUserRepository{}
	logger := slog.Default()
	postCreator := create.NewPostCreator(logger, &postRepository, &userRepository)
	postFinder := finder.NewPostRecentFinder(logger, &postRepository)
	postController := NewPostController(postCreator, postFinder)

	return &testSuite{
		&postRepository,
		&userRepository,
		postController,
	}
}

func TestGetPost(t *testing.T) {
	t.Parallel()

	t.Run("should get post", func(t *testing.T) {
		setupSuite := setupSuite()
		r := httptest.NewRequest(http.MethodGet, "/post", nil)
		w := httptest.NewRecorder()
		setupSuite.postRepository.On("FindRecent", mock.Anything).Once().Return(
			[]models.Post{
				{
					Id: "1",
				},
			}, nil)

		setupSuite.postController.GetPost(w, r)
		var response []map[string]string
   		errUnmarshal := json.Unmarshal(w.Body.Bytes(), &response)
		if errUnmarshal != nil {
			t.Fatal(errUnmarshal)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "1", response[0]["id"])
		setupSuite.postRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		setupSuite := setupSuite()
		r := httptest.NewRequest(http.MethodGet, "/post", nil)
		w := httptest.NewRecorder()
		setupSuite.postRepository.On("FindRecent", mock.Anything).Once().Return([]models.Post{}, assert.AnError)

		setupSuite.postController.GetPost(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		setupSuite.postRepository.AssertExpectations(t)
	})
}

func TestPostPost(t *testing.T) {
	t.Parallel()

	t.Run("should create post", func(t *testing.T) {
		setupSuite := setupSuite()
		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		json := []byte(fmt.Sprintf(`{"comment":"comment","postedBy":"%s"}`, uuid.String()))

		r := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(json))
		w := httptest.NewRecorder()
		
		setupSuite.userRepository.On("Find", mock.Anything, mock.Anything).Once().Return(models.User{}, nil)
		setupSuite.postRepository.On("Save", mock.Anything, mock.Anything).Once().Return(nil)

		setupSuite.postController.PostPost(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
		setupSuite.postRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		setupSuite := setupSuite()
		uuid, errUuid := uuid.NewV7()
		if errUuid != nil {
			t.Fatal(errUuid)
		}
		json := []byte(fmt.Sprintf(`{"comment":"comment","postedBy":"%s"}`, uuid.String()))

		r := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(json))
		w := httptest.NewRecorder()
		setupSuite.postRepository.On("Save", mock.Anything, mock.Anything).Once().Return(assert.AnError)
		setupSuite.userRepository.On("Find", mock.Anything, mock.Anything).Once().Return(models.User{}, nil)
		
		setupSuite.postController.PostPost(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		setupSuite.postRepository.AssertExpectations(t)
	})

	t.Run("should return validation error cause invalid post", func(t *testing.T) {
		setupSuite := setupSuite()
		r := httptest.NewRequest(http.MethodPost, "/post", nil)
		w := httptest.NewRecorder()

		setupSuite.postController.PostPost(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}