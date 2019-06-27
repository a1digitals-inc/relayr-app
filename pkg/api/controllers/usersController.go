package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrleite/relayr-app/pkg/api/models"
	"github.com/andrleite/relayr-app/pkg/api/utils"
	"github.com/gorilla/mux"
)

// PostUser create new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	rows, err := models.NewUser(user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusCreated)
}

// GetUsers retrive all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAll(models.USERS)
	utils.ToJSON(w, users, http.StatusOK)
}

// GetUser by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := models.GetByID(models.USERS, vars["id"])
	utils.ToJSON(w, user, http.StatusOK)
}

// PutUser update user info
func PutUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	user.ID = uint32(id)
	rows, err := models.UpdateUser(user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusOK)
}

// DeleteUser from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := models.Delete(models.USERS, vars["id"])
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
