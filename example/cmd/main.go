package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	web "github.com/FernandoCelmer/api.go/src"
)

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

type Response struct {
	Message string `json:"message"`
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{Message: "Item"}
	response, _ := json.Marshal(data)
	fmt.Fprintf(w, string(response))
}
