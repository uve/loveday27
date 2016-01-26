package core

import (
	"appengine"
	"appengine/datastore"
	"time"
)

const (
	DATASTORE_STATUS        = "STATUS"
	STATUS_NEW_APP          = "STATUS_NEW_APP"
	STATUS_EMAIL_SEARCHING  = "STATUS_EMAIL_SEARCHING"
	STATUS_EMAIL_READY      = "STATUS_EMAIL_READY"
	STATUS_EMAIL_BODY_READY = "STATUS_EMAIL_BODY_READY"
	STATUS_EMAIL_SENT       = "STATUS_EMAIL_SENT"
	STATUS_EMAIL_VIEWED     = "STATUS_EMAIL_VIEWED"
	STATUS_EMAIL_CLICKED    = "STATUS_EMAIL_CLICKED"
)

type Status struct {
	Name    string
	Created time.Time
}

func getStatusByType(c appengine.Context, name string) (*datastore.Key, error) {
	key := datastore.NewKey(c, DATASTORE_STATUS, name, 0, nil)
	e := new(Status)
	if err := datastore.Get(c, key, e); err != nil {
		return nil, err
	}
	return key, nil
}

func getStatusCount(c appengine.Context) (int, error) {

	q := datastore.NewQuery(DATASTORE_STATUS)

	count, err := q.Count(c)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func createAllStatusTypes(c appengine.Context) error {
	names := [...]string{
		STATUS_NEW_APP,
		STATUS_EMAIL_SEARCHING,
		STATUS_EMAIL_READY,
		STATUS_EMAIL_BODY_READY,
		STATUS_EMAIL_SENT,
		STATUS_EMAIL_VIEWED,
		STATUS_EMAIL_CLICKED,
	}

	var keys []*datastore.Key
	var values []*Status

	for _, value := range names {
		status := &Status{
			Name:    value,
			Created: time.Now(),
		}
		key := datastore.NewKey(c, DATASTORE_STATUS, value, 0, nil)
		keys = append(keys, key)
		values = append(values, status)
	}

	_, err := datastore.PutMulti(c, keys, values)
	return err
}

func (status *Status) String() string {
	return status.Name
}
