package web

import (
	"fmt"
	"net/http"
)

type typeHandler func(http.ResponseWriter, *http.Request)

type Client interface {
	Get(route string, handler typeHandler)
	Post(route string, handler typeHandler)
	Put(route string, handler typeHandler)
	Patch(route string, handler typeHandler)
	Delete(route string, handler typeHandler)
	Options(route string, handler typeHandler)
	Run(port int)
}

type web struct {
	config Config
}

func NewApp(args ...Option) Client {
	config := &Config{}
	for _, opt := range args {
		opt(config)
	}
	return &web{*config}
}

func (web *web) Get(route string, handler typeHandler) {
	http.Handle(route, get(handler))
}

func (web *web) Post(route string, handler typeHandler) {
	http.Handle(route, post(handler))
}

func (web *web) Put(route string, handler typeHandler) {
	http.Handle(route, put(handler))
}

func (web *web) Patch(route string, handler typeHandler) {
	http.Handle(route, patch(handler))
}

func (web *web) Delete(route string, handler typeHandler) {
	http.Handle(route, delete(handler))
}

func (web *web) Options(route string, handler typeHandler) {
	http.Handle(route, options(handler))
}

func (web *web) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
