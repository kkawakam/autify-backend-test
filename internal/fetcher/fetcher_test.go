package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHtml(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/html" {
			w.Header().Add("content-type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<!DOCTYPE html>
			<html>
			<head>
			<title>Hello World</title>
			</head>
			<body>
			<p>I am HTML</p>
			</body>
			</html>`))
		} else if r.URL.Path == "/other" {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"hello":"world"}`))
		} else {
			t.Errorf("Unsupported path %s", r.URL.Path)
		}
	}))
	defer server.Close()

	t.Run("Positive - Get HTML", func(t *testing.T) {
		html, theTime, err := getHtml(server.URL + "/html")
		assert.Nil(t, err, "No errors were expected")
		assert.NotNil(t, theTime, "Expected a time to be returned")
		assert.Contains(t, html, "I am HTML")
	})

	t.Run("Negative - Returned content type isn't html", func(t *testing.T) {
		html, theTime, err := getHtml(server.URL + "/other")
		assert.ErrorContains(t, err, "does not contain text/html response content type")
		assert.Nil(t, theTime, "Expected a time to be returned")
		assert.Empty(t, html, "No content returned")
	})

}
