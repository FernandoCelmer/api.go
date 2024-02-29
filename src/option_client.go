package web

import "net/http"

type ClientConfig struct {
	contentType string
}

type ClientOption func(option *ClientConfig)

func ContentType(value string) ClientOption {
	return func(option *ClientConfig) {
		option.contentType = value
	}
}

func WithContentType(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
	}
}

var clientDefault = &ClientConfig{
	contentType: "application/json",
}
