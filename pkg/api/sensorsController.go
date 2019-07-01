package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andrleite/relayr-app/pkg/api/models"
	"github.com/andrleite/relayr-app/pkg/api/utils"
	"github.com/gorilla/mux"
)

// PostSensor create new sensor
func (a *Api) PostSensor(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var s models.Sensor
	err := json.Unmarshal(body, &s)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	value, err := s.NewSensor(a.DB)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, value, http.StatusCreated)
}

// GetSensors retrive all sensorss from database
func (a *Api) GetSensors(w http.ResponseWriter, r *http.Request) {
	sensors := models.GetAll(a.DB)
	utils.ToJSON(w, sensors, http.StatusOK)
}

// GetSensor by id
func (a *Api) GetSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensor := models.GetByID(vars["id"], a.DB)
	if sensor == nil {
		w.Write([]byte(""))
		utils.ToJSON(w, sensor, http.StatusNoContent)
	} else {
		utils.ToJSON(w, sensor, http.StatusOK)
	}
}

// PutSensor update sensor info
func (a *Api) PutSensor(w http.ResponseWriter, r *http.Request) {
	var s models.Sensor
	fmt.Println(r)
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 10)
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	s.ID = uint(id)
	value, err := s.UpdateSensor(a.DB)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, value, http.StatusOK)
}

// DeleteSensor from database
func (a *Api) DeleteSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := models.Delete(vars["id"], a.DB)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
