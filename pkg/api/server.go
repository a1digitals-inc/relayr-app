package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/andrleite/relayr-app/pkg/api/routes"
)

func middlewareHandler(handler *mux.Router) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.Proto, r.Host, r.URL, time.Since(time.Now()))
		handler.ServeHTTP(w, r)
	})
}

func listen(p int) {
	fmt.Printf("\n\nListening port %d...", p)
	port := fmt.Sprintf(":%d", p)
	r := routes.NewRouter()
	http.HandleFunc("/", middlewareHandler(r))
	log.Fatal(http.ListenAndServe(port, nil))
}
