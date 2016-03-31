package core

import (

    "text/template"
    "appengine"
    "appengine/mail"
    "appengine/datastore"


    "time"
    "bytes"

    "io"
    "fmt"
    "net/http"
)

// Score is an entity to store campaign
type Email struct {
	Name    string
	Content string `datastore:",noindex"`

	Status   *datastore.Key
	Campaign *datastore.Key
	Created  time.Time
	Modified time.Time
}

type MailParams struct {
    Params *Params
    App *App
    AppIcon string
    Calculations Calculations
}

func handleMailPage(w http.ResponseWriter, r *http.Request) {

   c := appengine.NewContext(r)

   tickets, ticket_keys, err := getTickets(c, STATUS_EMAIL_READY, ""/*STATUS_EMAIL_BODY_GENERATING*/, MAIL_BODY_GENERATING_LIMIT)
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }

   c.Debugf("Tickets with Email Ready: ", ticket_keys)

   if len(ticket_keys) < 1 {
       return
   }

   processedTickets := []Ticket{}
   processedTicketKeys := []*datastore.Key{}

   for i, ticket := range tickets {
      content, err := generateEmail(r, ticket.App)
      if err != nil {
          c.Errorf("Error: ", err.Error())
          continue
      }

      ticket.Content = content.Bytes()

      processedTickets = append(processedTickets, ticket)
      processedTicketKeys = append(processedTicketKeys, ticket_keys[i])
   }

   //ToRemove
   for _, ticket := range tickets {
       io.WriteString(w, string(ticket.Content[:]))
   }
   err = setTicketsStatus(c, processedTickets, processedTicketKeys, STATUS_EMAIL_READY/*STATUS_EMAIL_BODY_READY*/)

   if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


func generateEmail(r *http.Request, appKey *datastore.Key) (*bytes.Buffer, error) {
    c := appengine.NewContext(r)
    app, err := getApp(c, appKey)
    if err != nil {
        return nil, err
    }

    c.Debugf("App Name: ", app.TrackName)

    iconBody, err := app.GetIcon(r)
    if err != nil {
        return nil, err
    }

    params, err := parseTemplateParams()
    if err != nil {
        return nil, err
    }

    calculations, err := app.GetCalculations()
    if err != nil {
        return nil, err
    }

    mailParams := MailParams{
      Params: params,
      App: app,
      AppIcon: iconBody,
      Calculations: *calculations,
    }

    var index = template.Must(template.ParseFiles("templates/mail.html"))
    var doc bytes.Buffer
    err = index.Execute(&doc, mailParams)
    if err != nil {
        return nil, err
    }

    return &doc, nil
}


const confirmMessage = `
Hi,
 This is Name, I'm localization manager
It's my first message!!!
Thank you for creating an account!

%s


Best Regards,
Name
`
const templateMailFrom = "%s <%s>"

func handleSendPage(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    //addr := r.FormValue("email")
    //url := createConfirmationURL(r)
    params, err := parseTemplateParams()
    if err != nil {
        c.Errorf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    c.Infof("Sender:", fmt.Sprintf(templateMailFrom, params.MailSenderName, params.MailSenderEmail))

    addrs := []string{"nikita.grachev@gmail.com"}

    msg := &mail.Message{
        Sender:  fmt.Sprintf(templateMailFrom, params.MailSenderName, params.MailSenderEmail),
        To:      addrs,
        Subject: "Confirm your registration",
        Body:    fmt.Sprintf(confirmMessage, "Test text"),
    }

    if err := mail.Send(c, msg); err != nil {
        c.Errorf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

/*

// timestamp formats date/time of the score.
func (s *Score) timestamp() string {
	return s.Played.Format(TIME_LAYOUT)
}

// put stores the score in the Datastore.
func (s *Score) put(c appengine.Context) (err error) {
	key := s.key
	if key == nil {
		key = datastore.NewIncompleteKey(c, SCORE_KIND, nil)
	}
	key, err = datastore.Put(c, key, s)
	if err == nil {
		s.key = key
	}
	return
}

// newScore returns a new Score ready to be stored in the Datastore.
func newScore(outcome string, u *user.User) *Score {
	return &Score{Outcome: outcome, Played: time.Now(), Player: userId(u)}
}



// newUserScoreQuery returns a Query which can be used to list all previous
// games of a user.
func newUserScoreQuery(u *user.User) *datastore.Query {
	return datastore.NewQuery(SCORE_KIND).Filter("player =", userId(u))
}

// fetchScores runs Query q and returns Score entities fetched from the
// Datastore.
func fetchScores(c appengine.Context, q *datastore.Query, limit int) (
	[]*Score, error) {

	scores := make([]*Score, 0, limit)
	keys, err := q.Limit(limit).GetAll(c, &scores)
	if err != nil {
		return nil, err
	}
	for i, score := range scores {
		score.key = keys[i]
	}
	return scores, nil
}

*/
