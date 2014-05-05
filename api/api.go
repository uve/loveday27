// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package api

import (


	"common/endpoints"
	
	"appengine/datastore"
	"time"
	
	"net/http"

	"appengine/user"
	"errors"
	"config"


)


//const clientId = "YOUR-CLIENT-ID"
var clientId = config.Config.OAuthProviders.Google.ClientId

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.ApiExplorerClientId}
	// in case we'll want to use TicTacToe API from an Android app
	audiences = []string{clientId}
)



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





  u, err := getCurrentUser(c)
  if err != nil {
 	return err
  }


  c.Infof("GreetingService list execute...")

  c.Infof(userId(u))



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



// getCurrentUser retrieves a user associated with the request.
// If there's no user (e.g. no auth info present in the request) returns
// an "unauthorized" error.
func getCurrentUser(c endpoints.Context) (*user.User, error) {
	u, err := endpoints.CurrentUser(c, scopes, audiences, clientIds)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("Unauthorized: Please, sign in.")
	}
	c.Debugf("Current user: %#v", u)
	return u, nil
}


// userId returns a string ID of the user u to be used as Player of Score.
func userId(u *user.User) string {
	return u.String()
}



// RegisterService exposes TicTacToeApi methods as API endpoints.
//
// The registration/initialization during startup is not performed here but
// in app package. It is separated from this package (tictactoe) so that the
// service and its methods defined here can be used in another app,
// e.g. http://github.com/crhym3/go-endpoints.appspot.com.
func RegisterService() (*endpoints.RpcService, error) {

	api := &GreetingService{}
	rpcService, err := endpoints.RegisterService(api,
		"greeting", "v1", "Greetings API", true)
	if err != nil {
		return nil, err
	}

	info := rpcService.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc =
		"greets.list", "GET", "greetings", "List most recent greetings."
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	return rpcService, nil
}
