package util

import (
	"bytes"
	b64 "encoding/base64"
	"net/http"
	"net/http/httptest"

	"github.com/spf13/viper"
)

var authString = b64.StdEncoding.EncodeToString([]byte(viper.GetString("admin.name") + ":" + viper.GetString("admin.pwd")))

// PerformRequest ... Perform Request (without body) util method. Token is optional(pass "")
func PerformRequest(method string, r http.Handler, path string, auth bool) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	if auth {
		req.Header.Set("Authorization", "Basic "+authString)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// PerformRequestWithBody ... PerformRequest With Body(POST,PUT,DELETE). Token is optional(pass "")
func PerformRequestWithBody(method string, r http.Handler, path string, body []byte, auth bool) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if auth {
		req.Header.Set("Authorization", "Basic "+authString)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
