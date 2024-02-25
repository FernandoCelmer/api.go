package web

type RunConfig struct {
	host string
	port int
}

type RunOption func(option *RunConfig)

func Host(value string) RunOption {
	return func(option *RunConfig) {
		option.host = value
	}
}

func Port(value int) RunOption {
	return func(option *RunConfig) {
		option.port = value
	}
}
