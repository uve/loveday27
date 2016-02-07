package core

import (
    "appengine"
    "appengine/datastore"
    "time"
)

const (
	DATASTORE_TICKET        = "TICKET"
)

type Ticket struct {
    App      *datastore.Key
    Compaign *datastore.Key
    Status   *datastore.Key

    Created  time.Time
    Modified time.Time
}

func getTickets(c appengine.Context, status string, limit int) ([]Ticket, []*datastore.Key, error) {

    statusKey, err := getStatusByType(c, STATUS_NEW_APP)
    if err != nil {
        return nil, nil, err
    }

    q := datastore.NewQuery(DATASTORE_TICKET).
                Filter("Status =", statusKey).
                Order("Modified").
                Limit(limit)

    var tickets []Ticket
    ticket_keys, err := q.GetAll(c, &tickets)
    if err != nil {
        return nil, nil, err
    }

    return tickets, ticket_keys, nil
}