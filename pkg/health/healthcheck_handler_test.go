package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nstoker/stoker.dev/pkg/version"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var got HealthCheckStruct
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("failed to unmarshal %v", err)
	}

	expected := HealthCheckStruct{
		Alive:   true,
		Version: version.Version(),
	}

	if got.Alive != expected.Alive {
		t.Errorf("unexpected reponse for alive. Got %v - expected %v", got.Alive, expected.Alive)
	}
	if got.Version != expected.Version {
		t.Errorf("Unexpected response for version. Got '%s' - expected '%s'", got.Version, expected.Version)
	}

}
