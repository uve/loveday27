package core

import (
	"log"

	"text/template"

	//"fmt"
	"net/http"

)

type Params struct {
	ClientId string
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {

	params := Params{
		ClientId: "882975820932-q34i2m1lklcmv8kqqrcleumtdhe4qbhk.apps.googleusercontent.com",
   }

	var index = template.Must(template.ParseFiles("templates/index.html"))

	err := index.Execute(w, params)
	if err != nil {
		log.Fatalf("template execution: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/map", handleMapPage)
	http.HandleFunc("/send", handleSendPage)
	http.HandleFunc("/setup", setup)
	http.HandleFunc("/search", searchPage)
	http.HandleFunc("/parser", parserPage)
	http.HandleFunc("/", handleMainPage)
}
