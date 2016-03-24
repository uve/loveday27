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
   "regexp"
)

var emailRegexp = regexp.MustCompile(EMAIL_REGEXP)


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

   processedTickets := []Ticket{}
   processedTicketKeys := []*datastore.Key{}

	for i, ticket := range tickets {
		err = parseApp(r, ticket.App)
      if err != nil {
          c.Debugf("Error: ", err.Error())
          continue
	   }

      processedTickets = append(processedTickets, ticket)
      processedTicketKeys = append(processedTicketKeys, ticket_keys[i])
   }

   err = setTicketsStatus(c, processedTickets, processedTicketKeys, STATUS_EMAIL_READY)

	if err != nil {
      c.Debugf("Error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func parseApp(r *http.Request, appKey *datastore.Key) (error) {
    c := appengine.NewContext(r)
    c.Debugf("Parse App: ", appKey)
    app, err := getApp(c, appKey)
    if err != nil {
		return err
	 }

    base, err := url.Parse(app.SellerUrl)
    if err != nil {
        return err
    }

    foundLinks, err := parseUrl(r, base)
    if err != nil {
        return err
    }
    //c.Debugf("Links all: ", foundLinks)

    parsedLinks, err := filterUrls(c, base, foundLinks)
    if err != nil {
        return err
    }
    c.Debugf("Links parsed: ", parsedLinks)

    parsedEmails, err := parseEmail(r, parsedLinks)
    if err != nil {
        return err
    }

    if len(parsedEmails) < 1 {
        return fmt.Errorf("No emails found: %d", app.TrackId)
    }

    c.Debugf("Emails parsed: ", parsedEmails)

    app.Emails = parsedEmails
    app.Save(c, appKey)

    return nil
}

func parseUrl(r *http.Request, base *url.URL) ([]*url.URL, error) {
    var ctx context.Context = newappengine.NewContext(r)
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

    err := isDomainBlacklisted(base)
    if err == nil {
        newLinks = append(newLinks, base)
    }

    c.Debugf("Base url:", base.String())

    for _, link := range links {
        link.Fragment = ""
        /*c.Debugf("link:", link)
        c.Debugf("Host:", link.Host)
        c.Debugf("Empty domain?:", link.Host == "")
        c.Debugf("Same domain?:", link.Host == base.Host)*/
        err := validateDomain(base, link)
        if err != nil {
            c.Debugf(err.Error())
            continue
        }

        if inSlice(link, newLinks) {
            continue
        }

        if len(newLinks) > PARSER_MAX_URLS {
            break
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
    /*
    if (base.Host != link.Host) {
        return fmt.Errorf("Validated domain error: different domain: %s", link)
    }*/

    err := isDomainBlacklisted(link)
    if err != nil {
        return err
    }

    err = isExtenstionBlacklisted(link)
    if err != nil {
        return err
    }

    return nil
}


func isDomainBlacklisted(link *url.URL) (error) {
    for _, blacklist := range BLACKLIST_DOMAINS {
        if strings.Contains(link.String(), blacklist) {
            return fmt.Errorf("Validated domain error: blacklisted domain: %s", link.String())
        }
    }
    return nil
}


func isEmailBlacklisted(str string) (error) {
    for _, blacklist := range BLACKLIST_EMAILS {
        if strings.Contains(str, blacklist) {
            return fmt.Errorf("Validated email error: blacklisted email: %s", str)
        }
    }
    return nil
}

func isExtenstionBlacklisted(link *url.URL) (error) {
    for _, blacklist := range BLACKLIST_EXTENSTIONS {
        if strings.HasSuffix(link.String(), blacklist) {
            return fmt.Errorf("Validated domain error: blacklisted extension: %s", link.String())
        }
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


// Extract all emails links from a given url
func crawl(r *http.Request, link *url.URL, ch chan string, chFinished chan bool) {
    var ctx context.Context = newappengine.NewContext(r)
    defer func() {
        // Notify that we're done after this function
        chFinished <- true
    }()

    content, err := fetchUrl(ctx, link)
    if err != nil {
        return
    }

    newEmails := emailRegexp.FindAllString(content, -1)

    for _, email := range newEmails {
        ch <- email
    }
}



func parseEmail(r *http.Request, links []*url.URL) ([]string, error) {
    foundEmails := make(map[string]bool)

    // Channels
    chEmails := make(chan string)
    chFinished := make(chan bool)


    defer func() {
        // Notify that we're done after this function
        close(chEmails)
        close(chFinished)
    }()


    // Kick off the crawl process (concurrently)
    for _, link := range links {
        go crawl(r, link, chEmails, chFinished)
    }

    // Subscribe to both channels
    for c := 0; c < len(links); {
        select {
        case email := <-chEmails:
            foundEmails[email] = true
        case <-chFinished:
            c++
        }
    }

    results := make([]string, 0, len(foundEmails))

    for email, _ := range foundEmails {
        err := isEmailBlacklisted(email)
        if err != nil {
            continue
        }

        results = append(results, email)
    }

    return results, nil
}
