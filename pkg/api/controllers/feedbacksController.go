package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/andrleite/relayr-app/pkg/api/models"
	"github.com/andrleite/relayr-app/pkg/api/utils"
)

//PostFeedback record user feedback to database
func PostFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &feedback)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	rows, err := models.NewFeedback(feedback)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusCreated)
}

// GetFeedbacks retrieve all feedbacks
func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks := models.GetAll(models.FEEDBACKS)
	utils.ToJSON(w, feedbacks, http.StatusOK)
}
