package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/social-readers-api/internal/application/comment/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/comment/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/errors"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testSuiteComment struct {
	commentRepository *repository.MockCommentRepository
	postRepository *repository.MockPostRepository
	userRepository *repository.MockUserRepository
	commentController *CommentController
}

func setupSuiteComment() *testSuiteComment {
	commentRepository := repository.MockCommentRepository{}
	postRepository := repository.MockPostRepository{}
	userRepository := repository.MockUserRepository{}
	commentCreator := create.NewCommentCreator(slog.Default(), &commentRepository, &userRepository, &postRepository)
	commentFindByPost := finder.NewCommentFinderByPost(slog.Default(), &commentRepository, &postRepository)
	commentController := NewCommentController(commentCreator, commentFindByPost)

	return &testSuiteComment{
		commentRepository: &commentRepository,
		postRepository: &postRepository,
		userRepository: &userRepository,
		commentController: commentController,
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	t.Run("should get comments", func(t *testing.T) {
		setupSuite := setupSuiteComment()
		r := httptest.NewRequest(http.MethodGet, "/posts/{postId}/comments", nil)
		r.SetPathValue("postId", "1")
		w := httptest.NewRecorder()
		setupSuite.commentRepository.On("FindByPost", mock.Anything, mock.Anything).Return(
			[]models.Comment{
				{
					Id: "1",
				},
			}, nil)
		setupSuite.postRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.Post{
				Id: "1",
			}, nil)

		setupSuite.commentController.GetCommentByPost(w, r)
		var response []map[string]string
   		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, "1", response[0]["id"])
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		setupSuite := setupSuiteComment()
		r := httptest.NewRequest(http.MethodGet, "/posts/{postId}/comments", nil)
		r.SetPathValue("postId", "1")
		w := httptest.NewRecorder()

		setupSuite.postRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.Post{}, assert.AnError)

		setupSuite.commentController.GetCommentByPost(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestPost(t *testing.T) {
	t.Parallel()

	t.Run("should create comment", func(t *testing.T) {
		setupSuite := setupSuiteComment()
		uuid, err := uuid.NewV7()
		if err != nil {
			t.Error("Error generating uuid")
		}
		json := []byte(fmt.Sprintf(`{"content":"comment","postId":"%s", "commentBy":"%s"}`, uuid.String(), uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(json))
		w := httptest.NewRecorder()
		setupSuite.commentRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		setupSuite.postRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.Post{
				Id: "1",
			}, nil)
		setupSuite.userRepository.On("Find", mock.Anything, mock.Anything).Return(models.User{}, nil)

		setupSuite.commentController.PostComment(w, r)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		setupSuite := setupSuiteComment()
		uuid, err := uuid.NewV7()
		if err != nil {
			t.Error("Error generating uuid")
		}	
		json := []byte(fmt.Sprintf(`{"content":"comment","postId":"%s", "commentBy":"%s"}`, uuid.String(), uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(json))
		w := httptest.NewRecorder()

		setupSuite.postRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.Post{}, assert.AnError)
		setupSuite.userRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.User{}, nil)

		setupSuite.commentController.PostComment(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		setupSuite := setupSuiteComment()	
		uuid, err := uuid.NewV7()
		if err != nil {
			t.Error("Error generating uuid")
		}	
		json := []byte(fmt.Sprintf(`{"content":"comment","postId":"%s", "commentBy":"%s"}`, uuid.String(), uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(json))
		w := httptest.NewRecorder()

		setupSuite.postRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.Post{
				Id: "1",
			}, nil)
		setupSuite.userRepository.On("Find", mock.Anything, mock.Anything).Return(models.User{}, errors.ErrNotExistUser)

		setupSuite.commentController.PostComment(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should return error when comment is invalid", func(t *testing.T) {
		setupSuite := setupSuiteComment()
		uuid, err := uuid.NewV7()
		if err != nil {
			t.Error("Error generating uuid")
		}	
		json := []byte(fmt.Sprintf(`{"content":"","postId":"%s", "commentBy":"%s"}`, uuid.String(), uuid.String()))
		r := httptest.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(json))
		w := httptest.NewRecorder()

		setupSuite.postRepository.On("Find", mock.Anything, mock.Anything).Return(
			models.Post{
				Id: "1",
			}, nil)
		setupSuite.userRepository.On("Find", mock.Anything, mock.Anything).Return(models.User{}, nil)

		setupSuite.commentController.PostComment(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}