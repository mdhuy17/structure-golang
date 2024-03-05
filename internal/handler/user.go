package handler

import (
	"PresentationProject/internal/model"
	"PresentationProject/internal/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo}
}

func (u *UserHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	fmt.Fprintf(w, "Email: %s", email)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Unmarshal
	user := &model.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created with ID: %s", createdUser.Email)
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Unmarshal
	user := &model.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User updated with Email: %s", updatedUser.Email)
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	err := u.userRepo.Delete(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User deleted with Email: %s", email)
}
