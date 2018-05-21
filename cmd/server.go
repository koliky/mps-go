package main

import (
	"fmt"
	"mps-go/pkg/app"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	app.CreateRouter(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
