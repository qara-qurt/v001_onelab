package http

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
	"v001_onelab/internal/model"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res, err := h.UserService.GetAll()

	if err != nil {
		log.Println("GetAll() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	paramID := params.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	res, err := h.UserService.GetByID(id)
	if err != nil {
		if errors.Is(model.ErrorNotFound, err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("GetByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.UserService.Create(user)
	if err != nil {
		log.Println("Create() error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	paramID := params.ByName("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.UserService.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user.ID = uint(id)
	res, err := h.UserService.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
