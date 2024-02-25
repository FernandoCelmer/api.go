package web

import (
	"net/http"
)

func execRequest(
	method string,
	handler funcHandler,
	w http.ResponseWriter,
	r *http.Request,
) {
	if method == r.Method {
		handler(w, r)
	} else {
		http.Error(
			w, http.StatusText(405),
			http.StatusMethodNotAllowed,
		)
	}
}

func get(config AppConfig, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceRequest(config.debug, r)
		execRequest(http.MethodGet, handler, w, r)
	})
}

func post(config AppConfig, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceRequest(config.debug, r)
		execRequest(http.MethodPost, handler, w, r)
	})
}

func put(config AppConfig, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceRequest(config.debug, r)
		execRequest(http.MethodPut, handler, w, r)
	})
}

func patch(config AppConfig, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceRequest(config.debug, r)
		execRequest(http.MethodPatch, handler, w, r)
	})
}

func delete(config AppConfig, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceRequest(config.debug, r)
		execRequest(http.MethodDelete, handler, w, r)
	})
}

func options(config AppConfig, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceRequest(config.debug, r)
		execRequest(http.MethodOptions, handler, w, r)
	})
}
