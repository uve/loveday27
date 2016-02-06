package core

import (
    "appengine/datastore"
    "time"
)

type Ticket struct {
    App      *datastore.Key
    Compaign *datastore.Key
    Status   *datastore.Key

    Created  time.Time
    Modified time.Time
}
