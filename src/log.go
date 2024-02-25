package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func loggerNew(debug bool, message string) {
	if debug {
		logger.Printf(message)
		fmt.Print(&buf)
	}
}

func logRequest(debug bool, r *http.Request) {
	loggerNew(
		debug,
		fmt.Sprintf("%s: http://%s%s", r.Method, r.Host, r.URL.Path),
	)
}
