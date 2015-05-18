package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type appContext struct {
	config map[string]interface{}
}

func loadConfig(url string) (config map[string]interface{}) {
	file, _ := ioutil.ReadFile(url)
	json.Unmarshal(file, &config)
	return config
}

func (ctx *appContext) eventbriteOAuthHandler(w http.ResponseWriter, r *http.Request) {
	// Need JS on client to handle fragment (#) URL
	t, _ := template.ParseFiles("forwardToApp.html")
	t.Execute(w, ctx.config["ios"])
}

func main() {
	context := appContext{config: loadConfig("./appURI.json")}
	port := os.Getenv("PORT")

	http.HandleFunc("/oauth/eventbrite/callback", context.eventbriteOAuthHandler)

	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
