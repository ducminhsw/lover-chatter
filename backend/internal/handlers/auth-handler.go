package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.ducminhsw.prepare-project/internal/model"
	"github.ducminhsw.prepare-project/internal/repository"
	db "github.ducminhsw.prepare-project/internal/repository/postgres"
	"github.ducminhsw.prepare-project/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerInterface interface {
	HandlerRegister() http.HandlerFunc
	HandlerLogin() http.HandlerFunc
}

type AuthHandler struct {
	Repository repository.UserRepositoryInterface
}

func NewAuthHandler(conn *sql.DB) AuthHandlerInterface {
	return AuthHandler{
		Repository: db.NewUserInterface(conn),
	}
}

func (ah AuthHandler) HandlerRegister() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req *request = &request{}
		var u *model.User = db.NewUser()

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusBadRequest, nil)
			return
		}

		_, err = ah.Repository.Retrieve(r.Context(), req.Email)
		if err == nil {
			utils.ParseResponseModel(w, http.StatusForbidden, nil)
			return
		}
		(*u).Email = (*req).Email
		(*u).Username = (*req).Name

		bytes, err := bcrypt.GenerateFromPassword([]byte((*req).Password), 12)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusBadGateway, nil)
			return
		}
		(*u).HashPassword = string(bytes)

		err = ah.Repository.Create(r.Context(), *u)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusInternalServerError, nil)
			return
		}

		utils.ParseResponseModel(w, http.StatusOK, nil)
	}
}

func (ah AuthHandler) HandlerLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req *request = &request{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusBadRequest, nil)
			return
		}

		u, err := ah.Repository.Retrieve(r.Context(), req.Email)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusNotFound, nil)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.HashPassword), []byte(req.Password))
		if err != nil {
			utils.ParseResponseModel(w, http.StatusForbidden, nil)
			return
		}

		utils.ParseResponseModel(w, http.StatusOK, nil)
	}
}
