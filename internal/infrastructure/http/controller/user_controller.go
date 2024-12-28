package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aperezgdev/social-readers-api/internal/application/user/create"
	finder "github.com/aperezgdev/social-readers-api/internal/domain/use_case/user"
)

type UserController struct {
	userCreator create.UserCreator
	userFinder  finder.UserFinder
}

func NewUserController(userCreator create.UserCreator, userFinder finder.UserFinder) *UserController {
	return &UserController{userCreator: userCreator, userFinder: userFinder}
}

type userRequest struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Mail    string `json:"mail"`
}

type userResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Mail        string `json:"mail"`
}

func (uc UserController) PostUser(w http.ResponseWriter, r *http.Request) {
	var user userRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uc.userCreator.Run(r.Context(), user.Name, user.Mail, user.Picture)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := uc.userFinder.Run(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(userResponse{
		Id:          string(user.Id),
		Name:        string(user.Name),
		Description: string(user.Description),
		Picture:     string(user.Picture),
		Mail:        string(user.Mail),
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
