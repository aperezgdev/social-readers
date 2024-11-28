package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aperezgdev/social-readers-api/internal/application/post/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/post/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/pkg"
)

type postRequest struct {
	Comment  string `json:"comment"`
	PostedBy string `json:"postedBy"`
}

type postResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Comment   string    `json:"comment"`
	PostedBy  string    `json:"postedBy"`
}

type PostController struct {
	postCreator      create.PostCreator
	postRecentFinder finder.PostRecentFinder
}

func NewPostController(postCreator create.PostCreator, postRecentFinder finder.PostRecentFinder) *PostController {
	return &PostController{postCreator: postCreator, postRecentFinder: postRecentFinder}
}

func (pc PostController) PostPost(w http.ResponseWriter, r *http.Request) {
	post := postRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	errCreator := pc.postCreator.Run(r.Context(), post.Comment, post.PostedBy)
	if errCreator != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (pc PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	posts, err := pc.postRecentFinder.Run(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	postsResponse := pkg.Map(posts, func(p models.Post) postResponse {
		return postResponse{
			time.Time(p.CreatedAt),
			string(p.Id),
			string(p.Comment),
			string(p.PostedBy),
		}
	})

	if err := json.NewEncoder(w).Encode(postsResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
