package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mps-go/pkg/app"
	"net/http"
)

var jsonData []byte

func main() {
	loadConfig()
	app.SetAPIURL(getConfig("api_url"))
	mux := http.NewServeMux()
	app.CreateRouter(mux)
	fmt.Println("server start => localhost:" + getConfig("port"))
	err := http.ListenAndServe(":"+getConfig("port"), mux)
	if err != nil {
		fmt.Println(err)
	}
}

func loadConfig() {
	loadConfg, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	jsonData = loadConfg
}

func getConfig(key string) string {
	var objmap map[string]*json.RawMessage
	json.Unmarshal(jsonData, &objmap)
	var value string
	json.Unmarshal(*objmap[key], &value)
	return value
}
