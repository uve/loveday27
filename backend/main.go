// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (

	"net/http"
	"fmt"
	"github.com/go-martini/martini"	
)


func handleMainPage(w http.ResponseWriter, r *http.Request) {



	fmt.Fprint(w, "Running backend")
}


func init() {

	m := martini.Classic()
	
	m.Get("/backend/", handleMainPage)		
	
}
