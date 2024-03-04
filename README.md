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

## Required imports

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	web "github.com/FernandoCelmer/api.go/src"
)
```

## Simple Example

```go
func main() {
	app := web.NewApp()
	app.Get("/item", itemHandler)

	app.Run()
}
```

## Example with parameters

```go
func main() {
	app := web.NewApp(
		web.Title("API"),
		web.Description("Minimalist API Test"),
		web.Version("0.1.0"),
	)

	app.Get("/item", itemHandler)

	app.Run(
		web.Host("127.0.0.5"),
		web.Port(8080),
	)
}
```

## Handler implementation


```go

type Response struct {
	Message string `json:"message"`
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{Message: "Item"}
	response, _ := json.Marshal(data)
	fmt.Fprintf(w, string(response))
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
