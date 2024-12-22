package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aperezgdev/social-readers-api/internal/application/book_to_read/create"
)

type BookToReadController struct {
	bookToReadCreator create.BookToReadCreator
}

type bookToReadRequest struct {
	Isbn 	  	string `json:"isbn"`
	Title      	string `json:"title"`
	Description string `json:"description"`
	Picture   	string `json:"picture"`
	UserId    	string `json:"userId"`
}

func NewBookToReadsController(bookToReadCreator create.BookToReadCreator) *BookToReadController {
	return &BookToReadController{bookToReadCreator: bookToReadCreator}
}

func (controller *BookToReadController) PostBookToRead(w http.ResponseWriter, r *http.Request) {
	var bookToReadRequest bookToReadRequest
	err := json.NewDecoder(r.Body).Decode(&bookToReadRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = controller.bookToReadCreator.Run(r.Context(), bookToReadRequest.Isbn, bookToReadRequest.Title, bookToReadRequest.Description, bookToReadRequest.Picture, bookToReadRequest.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}