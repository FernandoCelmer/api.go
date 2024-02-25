package web

import (
	"net/http"
)

func get(config Config, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(config.debug, r)

		if r.Method == http.MethodGet {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func post(config Config, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(config.debug, r)

		if r.Method == http.MethodPost {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func put(config Config, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(config.debug, r)

		if r.Method == http.MethodPut {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func patch(config Config, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(config.debug, r)

		if r.Method == http.MethodPatch {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func delete(config Config, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(config.debug, r)

		if r.Method == http.MethodDelete {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func options(config Config, handler funcHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(config.debug, r)

		if r.Method == http.MethodOptions {
			handler(w, r)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}
