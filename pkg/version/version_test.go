package version

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	expected := fmt.Sprintf("v%s", version)
	got := Version()
	if expected != got {
		t.Errorf("expected '%s', got '%s'", expected, got)
	}
}
