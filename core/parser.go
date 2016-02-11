package core

import (
	"net/http"
   "fmt"

   "appengine"
   "appengine/datastore"

   "golang.org/x/net/context"
	newappengine "google.golang.org/appengine"
   "golang.org/x/oauth2"

   //"appengine/urlfetch"

   "time"
)

const (
    PARSER_APPS_LIMIT = 1
)

func parserPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Debugf("parser Page")

	tickets, ticket_keys, err := getTickets(c, STATUS_NEW_APP, STATUS_EMAIL_SEARCHING, PARSER_APPS_LIMIT)
	if err != nil {
      c.Debugf("Error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Debugf("Tickets with New App: ", ticket_keys)

   if len(ticket_keys) < 1 {
       return
   }

	for _, ticket := range tickets {
		err = parseApp(r, ticket.App)
      if err != nil {
          c.Debugf("Error: ", err.Error())
	       http.Error(w, err.Error(), http.StatusInternalServerError)
	   }
	}
}

func parseApp(r *http.Request, appKey *datastore.Key) (error) {
    c := appengine.NewContext(r)
    c.Debugf("Parse App: ", appKey)
    app, err := getApp(c, appKey)
    if err != nil {
		return err
	 }

    var ctx context.Context = newappengine.NewContext(r)

    emails, err := parseUrl(ctx, app.SellerUrl)
    if err != nil {
		return err
	 }
    c.Debugf("Emails parsed: ", emails)
    return nil
}

func parseUrl(ctx context.Context, url string) ([]string, error) {

    //c.Debugf("parseUrl: ", url)

    if url == "" {
       return nil, fmt.Errorf("No SellerUrl found")
    }

    ctx_with_deadline, _ := context.WithTimeout(ctx, 2*time.Minute)
    client := oauth2.NewClient(ctx_with_deadline, nil)
    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }

    result := []string{resp.Status}
    return result, nil
}