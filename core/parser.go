package core

import (
	"net/http"

	"appengine"
   "appengine/datastore"
)

const (
    PARSER_APPS_LIMIT = 1
)

func parserPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Debugf("parser Page")

	tickets, ticket_keys, err := getTickets(c, STATUS_NEW_APP, PARSER_APPS_LIMIT)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Debugf("Tickets with New App: ", ticket_keys)

   if len(ticket_keys) < 1 {
       return
   }

	for _, ticket := range tickets {
		err = parseApp(c, ticket.App)
      if err != nil {
	       http.Error(w, err.Error(), http.StatusInternalServerError)
	   }
	}
}

func parseApp(c appengine.Context, appKey *datastore.Key) (error) {
    c.Debugf("Parse App: ", appKey)
    app, err := getApp(c, appKey)
    if err != nil {
		return err
	 }
    c.Debugf("SellerUrl: ", app.SellerUrl)

    return nil
}