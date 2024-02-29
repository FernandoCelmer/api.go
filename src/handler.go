package web

import (
	"io"
	"net/http"
)

func docsHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Docs")
}

func xptoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "XPTO")
}

type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i]
		h = d(h)
	}
	return h
}
