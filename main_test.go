package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHomeRouteReturnsOK(t *testing.T) {
	// Create temp static directory and file
	tmp := t.TempDir()
	staticDir := filepath.Join(tmp, "static")
	if err := os.MkdirAll(staticDir, 0o755); err != nil {
		t.Fatal(err)
	}

	homePath := filepath.Join(staticDir, "home.html")
	if err := os.WriteFile(homePath, []byte("<h1>home</h1>"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Change working dir so "/static/home.html" can be found as "./static/home.html"
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if chErr := os.Chdir(oldWd); chErr != nil {
			// Can't call t.Fatal inside Cleanup; log it.
			t.Logf("cleanup: failed to chdir back to %s: %v", oldWd, chErr)
		}
	})

	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}


	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d, body=%s", http.StatusOK, rr.Code, rr.Body.String())
	}
}
