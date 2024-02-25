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
	config := &Config{
		docs:  true,
		debug: true,
	}

	for _, opt := range args {
		opt(config)
	}

	if config.docs {
		http.Handle("/docs", get(*config, docsHandler))
	}

	return &web{*config}
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

func (web *web) Run(port int) {
	loggerNew(
		web.config.debug,
		fmt.Sprintf(
			"Api running on http://127.0.0.1:%d (Press CTRL+C to quit)", port,
		),
	)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
