package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.ducminhsw.prepare-project/internal/repository"
	db "github.ducminhsw.prepare-project/internal/repository/postgres"
	"github.ducminhsw.prepare-project/internal/utils"
)

type UserHandlerInterface interface {
	HandlePairLover() http.HandlerFunc
	HandleSettingUser() http.HandlerFunc
	HandleGetConversation() http.HandlerFunc
	HandleGetAllMemo() http.HandlerFunc
	HandleGetPartOfMemo() http.HandlerFunc
	HandleGetDetailMemo() http.HandlerFunc
}

type UserHandler struct {
	Repository repository.UserRepositoryInterface
}

func NewUserHandler(conn *sql.DB) UserHandlerInterface {
	return UserHandler{
		Repository: db.NewUserInterface(conn),
	}
}

func (uh UserHandler) HandlePairLover() http.HandlerFunc {
	type request struct {
		Email string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		target := r.PathValue("target")
		heartKey := r.PathValue("heart-key")

		if len(target) == 0 {
			utils.ParseResponseModel(w, http.StatusBadRequest, nil)
			return
		}

		req := request{}
		err := json.NewDecoder(r.Body).Decode(&r)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusNotFound, nil)
			return
		}
		sen, err := uh.Repository.Retrieve(r.Context(), req.Email)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusNotFound, nil)
			return
		}

		rec, err := uh.Repository.Retrieve(r.Context(), target)
		if err != nil {
			utils.ParseResponseModel(w, http.StatusNotFound, nil)
			return
		}

		if rec.HeartKey != heartKey {
			utils.ParseResponseModel(w, http.StatusForbidden, nil)
			return
		}

		sen.LoverName = rec.Username
		rec.LoverName = sen.Username

		utils.ParseResponseModel(w, http.StatusOK, nil)
	}
}

func (uh UserHandler) HandleSettingUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) HandleGetConversation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) HandleGetAllMemo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) HandleGetPartOfMemo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) HandleGetDetailMemo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
