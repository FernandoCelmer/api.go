package web

func doMagicWithApp(args ...AppOption) {
	for _, opt := range args {
		opt(appDefault)
	}
}

func doMagicWithRun(args ...RunOption) {
	for _, opt := range args {
		opt(runDefault)
	}
}

func doMagicWithClient(args ...ClientOption) *ClientConfig {
	var clientDefault = &ClientConfig{}
	for _, opt := range args {
		opt(clientDefault)
	}
	return clientDefault
}
