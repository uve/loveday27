package core

import (

   "appengine"
	"appengine/mail"
	"appengine/datastore"
	"time"

	"fmt"
	"net/http"
)


const (
   mailSenderName = "Anna Berry"
   mailSenderAddress = "anna@localization.expert"
   mailDomain = "localization.expert"
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


const confirmMessage = `
Hi,
 This is Anna Berry, I'm localization manager
It's my first message!!!
Thank you for creating an account!

%s


Best Regards,
Anna Berry
`
const templateMailFrom = "%s <%s>"

func handleSendPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//addr := r.FormValue("email")
	//url := createConfirmationURL(r)
   c.Infof("Sender:", fmt.Sprintf(templateMailFrom, mailSenderName, mailSenderAddress))

	addrs := []string{"check-auth@verifier.port25.com", "nikita.grachev@gmail.com", "daria.esaulova@gmail.com"}

	msg := &mail.Message{
		Sender:  fmt.Sprintf(templateMailFrom, mailSenderName, mailSenderAddress),
		To:      addrs,
		Subject: "Confirm your registration",
		Body:    fmt.Sprintf(confirmMessage, mailDomain),
	}
	if err := mail.Send(c, msg); err != nil {
		c.Errorf("Couldn't send email: %v", err)
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
