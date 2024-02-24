package web

import (
	"fmt"
	"net/http"
)

type funcHandler func(http.ResponseWriter, *http.Request)

type Client interface {
	Get(route string, handler funcHandler)
	Post(route string, handler funcHandler)
	Put(route string, handler funcHandler)
	Patch(route string, handler funcHandler)
	Delete(route string, handler funcHandler)
	Options(route string, handler funcHandler)
	Run(port int)
}

type web struct {
	config Config
}

func NewApp(args ...Option) Client {
	config := &Config{docs: true}

	for _, opt := range args {
		opt(config)
	}

	if config.docs {
		http.Handle("/docs", get(docsHandler))
	}

	return &web{*config}
}

func (web *web) Get(route string, handler funcHandler) {
	http.Handle(route, get(handler))
}

func (web *web) Post(route string, handler funcHandler) {
	http.Handle(route, post(handler))
}

func (web *web) Put(route string, handler funcHandler) {
	http.Handle(route, put(handler))
}

func (web *web) Patch(route string, handler funcHandler) {
	http.Handle(route, patch(handler))
}

func (web *web) Delete(route string, handler funcHandler) {
	http.Handle(route, delete(handler))
}

func (web *web) Options(route string, handler funcHandler) {
	http.Handle(route, options(handler))
}

func (web *web) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
