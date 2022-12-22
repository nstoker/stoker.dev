package health

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	r.HandleFunc("/health/", HealthCheckHandler).Methods("GET")
}
