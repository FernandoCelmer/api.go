package web

type AppConfig struct {
	title       string
	description string
	version     string
	docs        bool
	docs_url    string
	debug       bool
}

type AppOption func(option *AppConfig)

func Title(value string) AppOption {
	return func(option *AppConfig) {
		option.title = value
	}
}

func Description(value string) AppOption {
	return func(option *AppConfig) {
		option.description = value
	}
}

func Version(value string) AppOption {
	return func(option *AppConfig) {
		option.version = value
	}
}

func Docs(value bool) AppOption {
	return func(option *AppConfig) {
		option.docs = value
	}
}

func DocsURL(value string) AppOption {
	return func(option *AppConfig) {
		option.docs_url = value
	}
}

func Debug(value bool) AppOption {
	return func(option *AppConfig) {
		option.debug = value
	}
}

var appDefault = &AppConfig{
	title:       "API.go",
	description: "Minimalist API Test",
	version:     "0.1.0",
	docs:        true,
	docs_url:    "/docs",
	debug:       true,
}
