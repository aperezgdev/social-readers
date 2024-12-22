package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aperezgdev/social-readers-api/internal/application/book_to_read/create"
	finder "github.com/aperezgdev/social-readers-api/internal/application/book_to_read/find"
	"github.com/aperezgdev/social-readers-api/internal/domain/models"
	"github.com/aperezgdev/social-readers-api/pkg"
)

type BookToReadController struct {
	bookToReadFinderByUser finder.BookToReadFinderByUser
	bookToReadCreator create.BookToReadCreator
}

type bookToReadRequest struct {
	Isbn 	  	string `json:"isbn"`
	Title      	string `json:"title"`
	Description string `json:"description"`
	Picture   	string `json:"picture"`
	UserId    	string `json:"userId"`
}

type bookToReadResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	Id 	   string `json:"id"`
	Isbn 	   string `json:"isbn"`
	Title      string `json:"title"`
	Description string `json:"description"`
	Picture    string `json:"picture"`
	UserId     string `json:"userId"`
}

func NewBookToReadsController(bookToReadFinderByUser finder.BookToReadFinderByUser, bookToReadCreator create.BookToReadCreator) *BookToReadController {
	return &BookToReadController{bookToReadFinderByUser, bookToReadCreator}
}

func (controller *BookToReadController) GetBooksToReadByUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	booksToRead, err := controller.bookToReadFinderByUser.Run(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response := pkg.Map(booksToRead, func(b models.BookToRead) bookToReadResponse {
		return bookToReadResponse{
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