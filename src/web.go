package web

import (
	"fmt"
	"log"
	"net/http"
)

type funcHandler func(http.ResponseWriter, *http.Request)

type Client interface {
	Xpto(route string, handler funcHandler, args ...ClientOption)
	Get(route string, handler funcHandler, args ...ClientOption)
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
	doMagicWithApp(args...)

	if appDefault.docs {
		http.Handle(appDefault.docs_url, get(*appDefault, docsHandler))
	}

	return &web{*appDefault}
}

func (web *web) Xpto(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, Handler(xptoHandler, WithContentType))
}

func (web *web) Get(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, get(web.config, handler, args...))
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
	doMagicWithRun(args...)

	loggerNew(
		web.config.debug,
		fmt.Sprintf(
			"Api running on http://%s:%d (Press CTRL+C to quit)",
			runDefault.host,
			runDefault.port,
		),
	)

	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", runDefault.host, runDefault.port), nil); err != nil {
		log.Fatalf("ListenAndServer: %v", err)
	}
}
