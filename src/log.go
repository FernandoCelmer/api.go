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

func _loggerNew(message string) {
	if appDefault.debug {
		logger.Printf(message)
		fmt.Print(&buf)
	}
}

func _trace(r *http.Request) {
	_loggerNew(fmt.Sprintf("%s: http://%s%s", r.Method, r.Host, r.URL.Path))
}
