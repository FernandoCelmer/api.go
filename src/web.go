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
	Run(args ...RunOption)
}

type web struct {
	config AppConfig
}

func NewApp(args ...AppOption) Client {
	for _, opt := range args {
		opt(appDefault)
	}

	if appDefault.docs {
		http.Handle(appDefault.docs_url, get(*appDefault, docsHandler))
	}

	return &web{*appDefault}
}

func (web *web) Get(route string, handler funcHandler) {
	http.Handle(route, get(web.config, handler))
}

func (web *web) Post(route string, handler funcHandler) {
	http.Handle(route, post(web.config, handler))
}

func (web *web) Put(route string, handler funcHandler) {
	http.Handle(route, put(web.config, handler))
}

func (web *web) Patch(route string, handler funcHandler) {
	http.Handle(route, patch(web.config, handler))
}

func (web *web) Delete(route string, handler funcHandler) {
	http.Handle(route, delete(web.config, handler))
}

func (web *web) Options(route string, handler funcHandler) {
	http.Handle(route, options(web.config, handler))
}

func (web *web) Run(args ...RunOption) {
	for _, opt := range args {
		opt(runDefault)
	}

	loggerNew(
		web.config.debug,
		fmt.Sprintf(
			"Api running on http://%s:%d (Press CTRL+C to quit)",
			runDefault.host,
			runDefault.port,
		),
	)

	http.ListenAndServe(
		fmt.Sprintf("%s:%d", runDefault.host, runDefault.port), nil)
}
