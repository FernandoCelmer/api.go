package web

type Config struct {
	title       string
	description string
	version     string
}

type Option func(option *Config)

func Title(value string) Option {
	return func(option *Config) {
		option.title = value
	}
}

func Description(value string) Option {
	return func(option *Config) {
		option.description = value
	}
}

func Version(value string) Option {
	return func(option *Config) {
		option.version = value
	}
}
