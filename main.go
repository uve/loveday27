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

	"appengine"

	//"api"
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




func init() {



	m := martini.Classic()


	m.Use(sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123"))))

	var client_id string
	var client_secret string
	var redirect_url string


	if ( appengine.IsDevAppServer() ) {

		client_id     = "882975820932-ltfu1oa31f80o4v9tqfp8k5ghe45oibu.apps.googleusercontent.com"
		client_secret = "RPT5XVC_X0rl-z9jcrTV0aHk"
		redirect_url  = "http://localhost:8080/oauth2callback"

	} else {

		client_id     = "882975820932-q34i2m1lklcmv8kqqrcleumtdhe4qbhk.apps.googleusercontent.com"
		client_secret = "TCEAw8gTQBwyOKcYpqH8NTF4"
		redirect_url  = "http://mindale.com/oauth2callback"
	}


	m.Use(oauth2.Google(&oauth2.Options{
		ClientId:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  redirect_url,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},


	}))



	m.Get("/oauth2error", oauth2error)


	m.Get("/", handleMainPage)

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




	http.Handle("/", m)
}

