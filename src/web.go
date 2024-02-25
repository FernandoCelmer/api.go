package web

import (
	"fmt"
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
	config := &AppConfig{
		title:       "API.go",
		description: "Minimalist API Test",
		version:     "0.1.0",
		docs:        true,
		docs_url:    "/docs",
		debug:       true,
	}

	for _, opt := range args {
		opt(config)
	}

	if config.docs {
		http.Handle(config.docs_url, get(*config, docsHandler))
	}

	return &web{*config}
}

func (web *web) Get(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, get(web.config, handler))
}

func (web *web) Post(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, post(web.config, handler))
}

func (web *web) Put(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, put(web.config, handler))
}

func (web *web) Patch(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, patch(web.config, handler))
}

func (web *web) Delete(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, delete(web.config, handler))
}

func (web *web) Options(route string, handler funcHandler, args ...ClientOption) {
	http.Handle(route, options(web.config, handler))
}

func (web *web) Run(args ...RunOption) {
	client := &RunConfig{
		host: "127.0.0.1",
		port: 8080,
	}

	for _, opt := range args {
		opt(client)
	}

	loggerNew(
		web.config.debug,
		fmt.Sprintf(
			"Api running on http://%s:%d (Press CTRL+C to quit)",
			client.host,
			client.port,
		),
	)

	http.ListenAndServe(
		fmt.Sprintf("%s:%d", client.host, client.port), nil)
}
