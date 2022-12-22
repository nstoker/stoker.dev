package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nstoker/stoker.dev/pkg/health"
	homepage "github.com/nstoker/stoker.dev/pkg/homesite"
	"github.com/sirupsen/logrus"
)

var (
	r   *mux.Router
	api *mux.Router
)

func init() {
	logrus.Infof("Initializing router")
	r = mux.NewRouter()
	r.StrictSlash(true)
	api = r.PathPrefix("/api").Subrouter()
	defaultAPIRoutes(api)

	// Default routes come last
	homepage.ConnectToRouter(r, "/", "static", "index.html")

	r.Use(loggingMiddleware)
}

func GetRouter() (*mux.Router, error) {
	if r == nil {
		return nil, fmt.Errorf("router not initialized")
	}

	return r, nil
}

func defaultAPIRoutes(r *mux.Router) {
	health.Routes(r)
}

func Run(address string) error {
	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logrus.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
