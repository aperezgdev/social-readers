package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aperezgdev/social-readers-api/internal/application/comment/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/comment/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/pkg"
)

type CommentController struct {
	commentCreator      create.CommentCreator
	commentFinderByPost finder.CommentFinderByPost
}

func NewCommentController(commentCreator create.CommentCreator, commentFinderByPost finder.CommentFinderByPost) *CommentController {
	return &CommentController{commentCreator: commentCreator, commentFinderByPost: commentFinderByPost}
}

type commentRequest struct {
	Content   string `json:"content"`
	PostId    string `json:"postId"`
	CommentBy string `json:"commentBy"`
}

type commentResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	CommentBy string    `json:"commentBy"`
	PostId    string    `json:"postId"`
}

func (uc CommentController) PostComment(w http.ResponseWriter, r *http.Request) {
	var commentRequest commentRequest
	err := json.NewDecoder(r.Body).Decode(&commentRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = uc.commentCreator.Run(r.Context(), commentRequest.Content, commentRequest.PostId, commentRequest.CommentBy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc CommentController) GetCommentByPost(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("postId")
	comments, err := uc.commentFinderByPost.Run(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	commentsResponse := pkg.Map(comments, func(c models.Comment) commentResponse {
		return commentResponse{
			CreatedAt: time.Time(c.CreatedAt),
			Id:        string(c.Id),
			Content:   string(c.Content),
			CommentBy: string(c.CommentedBy),
			PostId:    string(c.PostId),
		}
	})

	err = json.NewEncoder(w).Encode(commentsResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
