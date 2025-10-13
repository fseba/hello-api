package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fseba/hello-api/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	tests := []struct {
		name          string
		Endpoint      string
		StatusCode    int
		expectedLang  string
		expectedTrans string
	}{
		{
			name:          "default hello translation",
			Endpoint:      "/hello",
			StatusCode:    http.StatusOK,
			expectedLang:  "english",
			expectedTrans: "hello",
		},
		{
			name:          "hello to german",
			Endpoint:      "/hello?language=german",
			StatusCode:    http.StatusOK,
			expectedLang:  "german",
			expectedTrans: "hallo",
		},
		{
			name:          "missing language returns 404",
			Endpoint:      "/hello?language=unknown",
			StatusCode:    http.StatusNotFound,
			expectedLang:  "",
			expectedTrans: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.Endpoint, nil)

			handler := http.HandlerFunc(rest.TranslateHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.StatusCode {
				t.Errorf("expected status %d but received %d", tt.StatusCode, rr.Code)
			}

			var resp rest.Resp
			json.Unmarshal(rr.Body.Bytes(), &resp)

			if resp.Language != tt.expectedLang {
				t.Errorf(`expected language %q but received %q`, tt.expectedLang, resp.Language)
			}

			if resp.Translation != tt.expectedTrans {
				t.Errorf(`expected translation %q but received %q`, tt.expectedTrans, resp.Translation)
			}
		})
	}
}
