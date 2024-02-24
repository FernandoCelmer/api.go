# api.go

---

- **Documentation**: [https://github.com/FernandoCelmer/api.go](https://github.com/FernandoCelmer/api.go)
- **Source Code**: [https://github.com/FernandoCelmer/api.go](https://github.com/FernandoCelmer/api.go)

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
	web "api/app/web"
)

func main() {
	app := web.NewApp(
		web.Title("API"),
		web.Description("Minimalist API Test"),
		web.Version("0.1.0"),
	)

	app.Get("/hello", handler.helloHandler)
	app.Run(8080)

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
