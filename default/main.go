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

	"github.com/go-martini/martini"	
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/sessions"
	

	//"api"
	"core/user"
	
	"config"
	
	"github.com/crhym3/go-endpoints/endpoints"
	

	"default/tictactoe"

)



func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found")
}


type Params struct {
	ClientId  string
}




func handleMainPage(w http.ResponseWriter, r *http.Request) {



	t := template.Must(template.New("main.html").ParseGlob("default/templates/*.html"))


	params := Params{
		//ClientId: config.Config.OAuthProviders.Google.ClientId,
		ClientId: "882975820932-q34i2m1lklcmv8kqqrcleumtdhe4qbhk.apps.googleusercontent.com",
	}



	err := t.Execute(w, params)
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







func init() {


	m := martini.Classic()


	store := sessions.NewCookieStore([]byte(config.Config.CookieSecret))
    m.Use(sessions.Sessions(config.Config.CookieName, store))
		
	
	
	//params := oauth2.Options( config.Config.OAuthProviders.Google )
	//params.RedirectURL = config.Config.RedirectURL
	//m.Use(oauth2.Google(&params))



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

	/*
	if _, err := api.RegisterService(); err != nil {
		panic(err.Error())
	}
	*/


	if _, err := tictactoe.RegisterService(); err != nil {
		panic(err.Error())
	}


	endpoints.HandleHttp()

	//api.Start()
	
	
	
	
}
