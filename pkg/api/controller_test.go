package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andrleite/relayr-app/pkg/api/models"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

var a Api

func TestMain(m *testing.M) {
	a = Api{}
	a.Initialize("root", "", "127.0.0.1", "3306", "relayrTest")

	ensureTableExists()

	code := m.Run()

	os.Exit(code)
}

func ensureTableExists() {
	a.DB.DropTableIfExists(models.Sensor{})
	a.DB.AutoMigrate(models.Sensor{})
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestGetSensorsHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/sensors", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	//Make the handler function satisfy http.Handler
	//https://lanreadelowo.com/blog/2017/04/03/http-in-go/
	http.HandlerFunc(a.GetSensors).
		ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	//Confirm the returned json is what we expected
	//Manually build up the expected json string
	expected := string(`[]`)

	// //The assert package checks if both JSON string are equal and for a plus, it actually confirms if our manually built JSON string is valid
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMetricsStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := promhttp.Handler()
	//r.Handle("/metrics", promhttp.Handler())
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetSensorHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/sensors/1", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	//Make the handler function satisfy http.Handler
	//https://lanreadelowo.com/blog/2017/04/03/http-in-go/
	http.HandlerFunc(a.GetSensor).
		ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
}

func TestCreateSensor(t *testing.T) {

	sensor := models.Sensor{
		Name: "SensorX",
		Type: "Pressure",
	}
	jsonSensor, _ := json.Marshal(sensor)

	req, err := http.NewRequest("POST", "/sensors", bytes.NewBuffer(jsonSensor))
	checkError(err, t)

	rr := httptest.NewRecorder()
	http.HandlerFunc(a.PostSensor).
		ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusCreated, status)
	}
}

func TestDeleteSensor(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/sensors/1", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()
	http.HandlerFunc(a.DeleteSensor).
		ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusNoContent, status)
	}
}
