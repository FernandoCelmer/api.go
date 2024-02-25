package web

type ClientConfig struct {
	contentType string
}

type ClientOption func(option *ClientConfig)

func ContentType(value string) ClientOption {
	return func(option *ClientConfig) {
		option.contentType = value
	}
}
