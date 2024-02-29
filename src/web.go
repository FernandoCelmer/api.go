package web

import (
	"fmt"
	"log"
	"net/http"
)

type funcHandler func(http.ResponseWriter, *http.Request)

type Client interface {
	Get(route string, handler funcHandler, args ...ClientOption)
	Post(route string, handler funcHandler, args ...ClientOption)
	Put(route string, handler funcHandler, args ...ClientOption)
	Patch(route string, handler funcHandler, args ...ClientOption)
	Delete(route string, handler funcHandler, args ...ClientOption)
	Options(route string, handler funcHandler, args ...ClientOption)
	Run(args ...RunOption)
}

type web struct {
	config AppConfig
}

func NewApp(args ...AppOption) Client {
	_loadApp(args...)

	if appDefault.docs {
		http.Handle(appDefault.docs_url, _get(docsHandler))
	}

	return &web{*appDefault}
}

func (web *web) Get(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, _get(handler, args...))
}

func (web *web) Post(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, _post(handler, args...))
}

func (web *web) Put(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, _put(handler, args...))
}

func (web *web) Patch(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, _patch(handler, args...))
}

func (web *web) Delete(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, _delete(handler, args...))
}

func (web *web) Options(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, _options(handler, args...))
}

func (web *web) Run(args ...RunOption) {
	_loadRun(args...)
	addr := getAddr()

	_loggerNew(
		web.config.debug,
		fmt.Sprintf("Api running on http://%s (Press CTRL+C to quit)", addr),
	)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("ListenAndServer: %v", err)
	}
}
