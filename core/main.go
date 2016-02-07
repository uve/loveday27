package core

import (
	"log"

	"text/template"

	"fmt"
	"net/http"

	"appengine"
	"appengine/mail"
)

const (
   mailSenderName = "Anna Berry"
   mailSenderAddress = "anna@localization.expert"
   mailDomain = "localization.expert"
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

const confirmMessage = `
Hi,
 This is Anna Berry, I'm localization manager
It's my first message!!!
Thank you for creating an account!

%s


Best Regards,
Anna Berry
`
const templateMailFrom = "%s <%s>"

func confirm(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//addr := r.FormValue("email")
	//url := createConfirmationURL(r)
   c.Infof("Sender:", fmt.Sprintf(templateMailFrom, mailSenderName, mailSenderAddress))

	addrs := []string{"check-auth@verifier.port25.com", "nikita.grachev@gmail.com", "daria.esaulova@gmail.com"}

	msg := &mail.Message{
		Sender:  fmt.Sprintf(templateMailFrom, mailSenderName, mailSenderAddress),
		To:      addrs,
		Subject: "Confirm your registration",
		Body:    fmt.Sprintf(confirmMessage, mailDomain),
	}
	if err := mail.Send(c, msg); err != nil {
		c.Errorf("Couldn't send email: %v", err)
	}
}

func init() {
	http.HandleFunc("/map", handleMapPage)
	http.HandleFunc("/send", confirm)
	http.HandleFunc("/setup", setup)
	http.HandleFunc("/search", searchPage)
	http.HandleFunc("/parser", parserPage)
	http.HandleFunc("/", handleMainPage)
}
