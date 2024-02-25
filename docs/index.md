# api.go

---
- **Github**: [https://github.com/FernandoCelmer/api.go](https://github.com/FernandoCelmer/api.go)
- **Documentation**: [https://fernandocelmer.github.io/api.go](https://fernandocelmer.github.io/api.go)
- **Go Package**: [https://pkg.go.dev/github.com/FernandoCelmer/api.go/src](https://pkg.go.dev/github.com/FernandoCelmer/api.go/src)

---

## Installation

To install, simply run:

```shell
go get github.com/FernandoCelmer/api.go
```

## Example

```go
package main

import (
	"io"
	"net/http"

	web "github.com/FernandoCelmer/api.go/src"
)

func main() {
	app := web.NewApp(
		web.Title("API"),
		web.Description("Minimalist API Test"),
		web.Version("0.1.0"),
	)

	app.Get("/hello", helloHandler)
	app.Run(8080)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Test")
}
```

## Commit Style

- ⚙️ FEATURE
- 📝 PEP8
- 📌 ISSUE
- 🪲 BUG
- 📘 DOCS
- 📦 PyPI
- ❤️️ TEST
- ⬆️ CI/CD
- ⚠️ SECURITY

## License

This project is licensed under the terms of the MIT license.
