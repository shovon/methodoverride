package methodoverride

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var example = XHTTPMethodOverrideHandler{
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "HAHA" {
			w.Write([]byte("Custom method"))
		} else {
			w.Write([]byte("Typical method"))
		}
	}),
}

func TestNoHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/abc/testfile.txt", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	example.ServeHTTP(rr, req)
	str := rr.Body.String()
	if req.Method != "GET" {
		t.Errorf("Expected method to be GET, but got %s", req.Method)
	}
	override := req.Header.Get("X-HTTP-Method-Override")
	if override != "" {
		t.Error("The X-HTTP-Method-Override header should have been empty")
	}
	if str != "Typical method" {
		t.Errorf("Expected response body to be %s, but got %s", "Typical method", str)
	}
}

func TestMethodOverride(t *testing.T) {
	req, err := http.NewRequest("POST", "/abc/testfile.txt", nil)
	req.Header.Add("X-HTTP-Method-Override", "HAHA")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	example.ServeHTTP(rr, req)
	str := rr.Body.String()
	if req.Method != "HAHA" {
		t.Errorf("Expected method to be GET, but got %s", req.Method)
	}
	override := req.Header.Get("X-HTTP-Method-Override")
	if override != "HAHA" {
		t.Error("The X-HTTP-Method-Override header should have been HAHA")
	}
	if str != "Custom method" {
		t.Errorf("Expected response body to be %s, but got %s", "Custom method", str)
	}
}
