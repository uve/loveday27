package core

import (

	"log"

	"text/template"

 	"fmt"
	"net/http"

	"appengine"
	"appengine/mail"
)


type Params struct {
	ClientId  string
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


func handleMapPage(w http.ResponseWriter, r *http.Request) {

	params := Params{
		//ClientId: config.Config.OAuthProviders.Google.ClientId,
		ClientId: "882975820932-q34i2m1lklcmv8kqqrcleumtdhe4qbhk.apps.googleusercontent.com",
	}


	var index = template.Must(template.ParseFiles("templates/map.html"))

	err := index.Execute(w, params)
    if err != nil {
    	log.Fatalf("template execution: %s", err)
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }	
}


const confirmMessage = `
Thank you for creating an account!
Please confirm your email address by clicking on the link below:

%s
`

func confirm(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        //addr := r.FormValue("email")
        //url := createConfirmationURL(r)

		addrs := []string{"check-auth@verifier.port25.com"}//, "nikita.grachev@gmail.com"}
		url := "mindale.com"

        msg := &mail.Message{
                Sender:  "Mindale Localization Services <mail@mindale.com>",
                To:      addrs,
                Subject: "Confirm your registration",
                Body:    fmt.Sprintf(confirmMessage, url),
        }
        if err := mail.Send(c, msg); err != nil {
                c.Errorf("Couldn't send email: %v", err)
        }
}


func init() {
	http.HandleFunc("/map", handleMapPage)
	http.HandleFunc("/send", confirm)
	http.HandleFunc("/setup", setup)
    http.HandleFunc("/reduce", reducePage)
	http.HandleFunc("/", handleMainPage)
}
