package health

import (
	"encoding/json"
	"net/http"

	"github.com/nstoker/stoker.dev/pkg/version"
)

type HealthCheckStruct struct {
	Alive   bool
	Version string
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	health := HealthCheckStruct{
		Alive:   true,
		Version: version.Version(),
	}

	json.NewEncoder(w).Encode(health)
}
