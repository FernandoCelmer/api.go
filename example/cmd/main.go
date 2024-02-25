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
	io.WriteString(w, "Hello!")
}
