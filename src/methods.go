package web

import (
	"net/http"
)

func _request(
	method string,
	handler funcHandler,
	args ...ClientOption,
) http.Handler {
	client := _loadClient(args...)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method == r.Method {
			_trace(r)
			w.Header().Set("Content-Type", client.contentType)
			handler(w, r)
		} else {
			http.Error(
				w, http.StatusText(405),
				http.StatusMethodNotAllowed,
			)
		}
	})

}

func _get(handler funcHandler, args ...ClientOption) http.Handler {
	return _request(http.MethodGet, handler, args...)
}

func _post(handler funcHandler, args ...ClientOption) http.Handler {
	return _request(http.MethodPost, handler, args...)
}

func _put(handler funcHandler, args ...ClientOption) http.Handler {
	return _request(http.MethodPut, handler, args...)
}

func _patch(handler funcHandler, args ...ClientOption) http.Handler {
	return _request(http.MethodPatch, handler, args...)
}

func _delete(handler funcHandler, args ...ClientOption) http.Handler {
	return _request(http.MethodDelete, handler, args...)
}

func _options(handler funcHandler, args ...ClientOption) http.Handler {
	return _request(http.MethodOptions, handler, args...)
}
