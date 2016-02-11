package core

import (
    "appengine"
    "appengine/datastore"
    "time"
)

const (
	DATASTORE_TICKET = "Ticket"
)

type Ticket struct {
    App      *datastore.Key
    Compaign *datastore.Key
    Status   *datastore.Key

    Created  time.Time
    Modified time.Time
}

func getTickets(c appengine.Context, status string, newStatus string, limit int) ([]Ticket, []*datastore.Key, error) {

    c.Debugf("getTickets: ", status)
    statusKey, err := getStatusByType(c, status)
    if err != nil {
        return nil, nil, err
    }

    c.Debugf("statusKey: ", statusKey)

    q := datastore.NewQuery(DATASTORE_TICKET).
                Filter("Status =", statusKey).
                //Filter("SellerUrl >", "").
                Order("Modified").
                Limit(limit)

    var tickets []Ticket
    ticket_keys, err := q.GetAll(c, &tickets)
    if err != nil {
        return nil, nil, err
    }
	 /*
    if newStatus != "" {
        err = setTicketsStatus(c, tickets, ticket_keys, newStatus)
        if err != nil {
            return nil, nil, err
        }
    }*/

    return tickets, ticket_keys, nil
}

func setTicketsStatus(c appengine.Context, tickets []Ticket, keys []*datastore.Key, status string) (error) {

    c.Debugf("setTicketsStatus: ", status)
    statusKey, err := getStatusByType(c, status)
    if err != nil {
        return err
    }

    c.Debugf("new statusKey: ", statusKey)

    for i, _ := range tickets {
        tickets[i].Status = statusKey
    }

    _, err = datastore.PutMulti(c, keys, tickets)
    return err
}