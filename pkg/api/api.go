package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andrleite/relayr-app/pkg/api/models"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// This is necessary for gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Api struct {
	Router *mux.Router
	DB     *models.Database
}

func (a *Api) Initialize(dbUser, dbPass, dbHost, dbPort, dbName string) {
	var err error
	a.DB, err = models.New(dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}

	models.AutoMigrations(a.DB)
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *Api) InitializeRoutes() {
	a.Router.HandleFunc("/sensors", a.GetSensors).Methods("GET")
	a.Router.HandleFunc("/sensors/{id}", a.GetSensor).Methods("GET")
	a.Router.HandleFunc("/sensors", a.PostSensor).Methods("POST")
	a.Router.HandleFunc("/sensors/{id}", a.PutSensor).Methods("PUT")
	a.Router.HandleFunc("/sensors/{id}", a.DeleteSensor).Methods("DELETE")
	a.Router.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	a.Router.Handle("/metrics", promhttp.Handler())
}

func (a *Api) middlewareHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.Proto, r.Host, r.URL, time.Since(time.Now()))
		a.Router.ServeHTTP(w, r)
	})
}

func (a *Api) listen(p int) {
	fmt.Printf("\n\nListening port %d...", p)
	port := fmt.Sprintf(":%d", p)
	http.HandleFunc("/", a.middlewareHandler())
	log.Fatal(http.ListenAndServe(port, nil))
}

// Run database migrations and call server listen
func (a *Api) Run() {
	defer a.DB.Close()
	a.listen(9000)
}
