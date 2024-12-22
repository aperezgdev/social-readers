package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/create"
	"github.com/aperezgdev/social-readers-api/internal/application/book_recommended/finder"
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

	if err := json.NewEncoder(w).Encode(bookRecommendeds); err != nil {
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