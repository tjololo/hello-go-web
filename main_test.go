package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		responseString string
		appInfo        string
		wantStatus     int
		wantResponse   string
		wantInfo       string
	}{
		{
			name:           "default response",
			responseString: "Hello from go code",
			appInfo:        "",
			wantStatus:     http.StatusOK,
			wantResponse:   "Hello from go code",
			wantInfo:       "Application info; ",
		},
		{
			name:           "custom response with app info",
			responseString: "custom response",
			appInfo:        "v1.0.0",
			wantStatus:     http.StatusOK,
			wantResponse:   "custom response",
			wantInfo:       "Application info; v1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseString = tt.responseString
			t.Setenv("APP_INFO", tt.appInfo)

			req := httptest.NewRequest(http.MethodGet, "/hello", nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Errorf("status = %d, want %d", res.StatusCode, tt.wantStatus)
			}
			if ct := res.Header.Get("Content-Type"); ct != "application/json; charset=utf-8" {
				t.Errorf("Content-Type = %q, want %q", ct, "application/json; charset=utf-8")
			}
			var body WebResponse
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Fatalf("decode response: %v", err)
			}
			if body.Response != tt.wantResponse {
				t.Errorf("Response = %q, want %q", body.Response, tt.wantResponse)
			}
			if body.Info != tt.wantInfo {
				t.Errorf("Info = %q, want %q", body.Info, tt.wantInfo)
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
	}{
		{
			name:       "zero error rate always returns OK",
			query:      "/errors?500=0",
			wantStatus: http.StatusOK,
		},
		{
			name:       "101 error rate always returns internal server error",
			query:      "/errors?500=101",
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "no param defaults to zero error rate",
			query:      "/errors",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.query, nil)
			w := httptest.NewRecorder()

			errorHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
		})
	}
}
