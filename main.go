// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (

	"io"
	"log"
	"fmt"

	"net/http"
	"text/template"

	"common/martini"
	"common/oauth2"
	"common/sessions"

	//"appengine"

	//"api"
	"core/user"
	
	"config"
	
	"common/endpoints"
	
	"appengine/datastore"
	"time"
	
	"api"
)



func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found")
}



func handleMainPage(w http.ResponseWriter, r *http.Request) {

	t := template.New("main.html")
	t = template.Must(t.ParseGlob("templates/*.html"))
	
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
		
	//fmt.Fprint(w, result)
}


func oauth2error(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Auth error")
}



func PaymentsPage(u *user.User, s sessions.Session, c martini.Context, w http.ResponseWriter, r *http.Request) {

	
	
	fmt.Fprint(w, c)	
	 
}






// Greeting is a datastore entity that represents a single greeting.
// It also serves as (a part of) a response of GreetingService.
type Greeting struct {
  Key     *datastore.Key `json:"id" datastore:"-"`
  Author  string         `json:"author"`
  Content string         `json:"content" datastore:",noindex" endpoints:"req"`
  Date    time.Time      `json:"date"`
}

// GreetingsList is a response type of GreetingService.List method
type GreetingsList struct {
  Items []*Greeting `json:"items"`
}

// Request type for GreetingService.List
type GreetingsListReq struct {
  Limit int `json:"limit" endpoints:"d=10"`
}










// GreetingService can sign the guesbook, list all greetings and delete
// a greeting from the guestbook.
type GreetingService struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *GreetingService) List(
  r *http.Request, req *GreetingsListReq, resp *GreetingsList) error {

  if req.Limit <= 0 {
    req.Limit = 10
  }

  c := endpoints.NewContext(r)
  q := datastore.NewQuery("Greeting").Order("-Date").Limit(req.Limit)
  greets := make([]*Greeting, 0, req.Limit)
  keys, err := q.GetAll(c, &greets)
  if err != nil {
    return err
  }

  for i, k := range keys {
    greets[i].Key = k
  }
  resp.Items = greets
  return nil
}







func init() {


	m := martini.Classic()


	store := sessions.NewCookieStore([]byte(config.Config.CookieSecret))
    m.Use(sessions.Sessions(config.Config.CookieName, store))
		
	
	
	params := oauth2.Options( config.Config.OAuthProviders.Google )
	params.RedirectURL = config.Config.RedirectURL
	m.Use(oauth2.Google(&params))	



	m.Get("/oauth2error", oauth2error)


	//m.Get("/", testPage)

	

	

	
	
	m.Get("/", handleMainPage)
	
	m.Get("/payments", oauth2.LoginRequired, PaymentsPage)
	/*
	
	// Tokens are injected to the handlers
	m.Get("/", func(tokens oauth2.Tokens) string {
		if tokens.IsExpired() {



	
			return "not logged in, or the access token is expired2"
		}
		return "logged in"

	})



	m.Get("/restrict", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
			return tokens.Access()
	})
	*/

	http.Handle("/", m)
	
	api.Start()
	
	
	
	
}