package web

import (
	"encoding/json"
	"net/http"
)

func main() {
	app := NewApp(
		Title("API"),
		Description("Minimalist API Test"),
		Version("0.1.0"),
	)

	app.Get("/item", itemHandler)

	app.Run(
		Host("127.0.0.5"),
		Port(8080),
	)
}

type Response struct {
	Message string `json:"message"`
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{Message: "Item"}
	response, _ := json.Marshal(data)
	w.Write(response)
}
