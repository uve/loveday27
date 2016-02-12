package core

import (
	"net/http"
   "io/ioutil"
   "fmt"
   "strings"

   "appengine"
   "appengine/datastore"

   "golang.org/x/net/html"
   "golang.org/x/net/context"
	newappengine "google.golang.org/appengine"
   "golang.org/x/oauth2"
   "time"
   "net/url"
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

    base, err := url.Parse(app.SellerUrl)
    if err != nil {
        return err
    }

    foundLinks, err := parseUrl(ctx, base)
    if err != nil {
        return err
    }
    c.Debugf("Links all: ", foundLinks)


    parsedLinks, err := filterUrls(c, base, foundLinks)
    if err != nil {
        return err
    }
    c.Debugf("Links parsed: ", parsedLinks)


    return nil
}

func parseUrl(ctx context.Context, base *url.URL) ([]*url.URL, error) {
    body, err := fetchUrl(ctx, base)
    if err != nil {
        return nil, err
    }

    links := []*url.URL{}

    doc, err := html.Parse(strings.NewReader(body))
    if err != nil {
        return nil, err
    }
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
                if a.Key == "href" {
                    //fmt.Println(a.Val)
                    link, err := resolveUrl(base, a.Val)
                    if err == nil {
                        links = append(links, link)
                    }
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)

    return links, nil
}


func resolveUrl(base *url.URL, href string) (*url.URL, error) {
    u, err := url.Parse(href)
	 if err != nil {
		 return nil, err
	 }

    return base.ResolveReference(u), nil
}

func inSlice(link *url.URL, links []*url.URL) bool {
    for _, u := range links {
        if link.String() == u.String() {
            return true
        }
    }
    return false
}

func filterUrls(c appengine.Context, base *url.URL, links []*url.URL) ([]*url.URL, error) {
    newLinks := []*url.URL{}

    for _, link := range links {
	     link.Fragment = ""
        /*c.Debugf("link:", link)
        c.Debugf("Host:", link.Host)
        c.Debugf("Empty domain?:", link.Host == "")
        c.Debugf("Same domain?:", link.Host == base.Host)*/
        err := validateDomain(base, link)
        if err != nil {
            continue
        }
        if inSlice(link, newLinks) {
            continue
        }
        newLinks = append(newLinks, link)
    }

    return newLinks, nil
}

func validateDomain(base *url.URL, link *url.URL) (error) {
    if !link.IsAbs() {
        return fmt.Errorf("Validated domain error: url is not absolute: %s", link)
    }

    if (base.Host == "") {
        return fmt.Errorf("Validated domain error: empty domain: %s", link)
    }

    if (base.Host != link.Host) {
        return fmt.Errorf("Validated domain error: different domain: %s", link)
    }

    return nil
}

func fetchUrl(ctx context.Context, link *url.URL) (string, error) {
    /*
    if url == "" {
       return "", fmt.Errorf("No SellerUrl found")
    }*/

    ctx_with_deadline, _ := context.WithTimeout(ctx, 2*time.Minute)
    client := oauth2.NewClient(ctx_with_deadline, nil)
    response, err := client.Get(link.String())
    if err != nil {
        return "", err
    }

    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", err
    }
    return string(contents), nil
}