package core

import (
	"log"

	"text/template"

	//"fmt"
	"net/http"

   "encoding/base64"
)

type Params struct {
   MailSenderName string
   MailSenderEmail string
	ClientId string
}

func parseTemplateParams() (*Params, error) {
    MailSenderName, err := base64.StdEncoding.DecodeString(MAIL_SENDER_NAME)
    if err != nil {
       return nil, err
    }
    MailSenderEmail, err := base64.StdEncoding.DecodeString(MAIL_SENDER_EMAIL)
    if err != nil {
       return nil, err
    }
    CliendId, err := base64.StdEncoding.DecodeString(GOOGLE_CLIENT_ID)
    if err != nil {
       return nil, err
    }

    params := Params{
      MailSenderName: string(MailSenderName),
      MailSenderEmail: string(MailSenderEmail),
      ClientId: string(CliendId),
    }

    return &params, nil
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
    params, err := parseTemplateParams()
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    var index = template.Must(template.ParseFiles("templates/index.html"))

    err = index.Execute(w, params)
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func init() {
	http.HandleFunc("/map", handleMapPage)
   http.HandleFunc("/mail", handleMailPage)
	http.HandleFunc("/send", handleSendPage)
	http.HandleFunc("/setup", setup)
	http.HandleFunc("/search", searchPage)
	http.HandleFunc("/parser", parserPage)
	http.HandleFunc("/", handleMainPage)
}
