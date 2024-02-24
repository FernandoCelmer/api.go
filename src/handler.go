package web

import (
	"io"
	"net/http"
)

func docsHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Docs")
}
