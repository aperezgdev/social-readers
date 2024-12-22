package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/create"
	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/finder"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/pkg"
)

type BookRecommendedController struct {
	bookRecommendedFinderByUser finder.BookRecommendedFinderByUser
	bookRecommendedCreator create.BookRecommendedCreator
}

type bookRecommendedRequest struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	UserId      string `json:"userId"`
}

type bookRecommendedResponse struct {
	CreatedAt 	time.Time `json:"createdAt"`
	Id 			string `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	UserId      string `json:"userId"`
}

func NewBookRecommendedController(
	bookRecommendedFinderByUser finder.BookRecommendedFinderByUser,
	bookRecommendedCreator create.BookRecommendedCreator,
) *BookRecommendedController {
	return &BookRecommendedController{
		bookRecommendedFinderByUser: bookRecommendedFinderByUser,
		bookRecommendedCreator: bookRecommendedCreator,
	}
}

func (controller *BookRecommendedController) GetBookRecommendedByUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	bookRecommendeds, err := controller.bookRecommendedFinderByUser.Run(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := pkg.Map(bookRecommendeds, func(b models.BookRecommended) bookRecommendedResponse {
		return bookRecommendedResponse{
			CreatedAt: time.Time(b.CreatedAt),
			Id: string(b.Id),
			Isbn: string(b.Isbn),
			Title: string(b.Title),
			Description: string(b.Description),
			Picture: string(b.Picture),
			UserId: string(b.UserId),
		}
	})

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (controller *BookRecommendedController) PostBookRecommended(w http.ResponseWriter, r *http.Request) {
	var bookRecommendedRequest bookRecommendedRequest
	err := json.NewDecoder(r.Body).Decode(&bookRecommendedRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.bookRecommendedCreator.Run(r.Context(), bookRecommendedRequest.Isbn, bookRecommendedRequest.Title, bookRecommendedRequest.Description, bookRecommendedRequest.Picture, bookRecommendedRequest.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}