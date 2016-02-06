package core

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"time"
	"errors"
    "strconv"
    "strings"
)

const (
    DATASTORE_APP = "APP"
    DATASTORE_STORE = "APPSTORE"
    DATASTORE_TICKET = "TICKET"
	DATASTORE_CAMPAIGN    = "CAMPAIGN"
	CAMPAIGN_LOCALIZATION = "CAMPAIGN_LOCALIZATION"
)

type Campaign struct {
	Name    string
	Created time.Time
	Apps *[]App
    Keys []*datastore.Key
}

type CampaignParams struct {
	Name  string
	Limit int
}

func getCampaignByType(c appengine.Context, name string) (*Campaign, *datastore.Key, error) {
	key := datastore.NewKey(c, DATASTORE_CAMPAIGN, name, 0, nil)
	e := new(Campaign)
	if err := datastore.Get(c, key, e); err != nil {
		return nil, nil, err
	}
	return e, key, nil
}

func getCampaignCount(c appengine.Context) (int, error) {

	q := datastore.NewQuery(DATASTORE_CAMPAIGN)

	count, err := q.Count(c)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func createAllCampaign(c appengine.Context) error {
	names := [...]string{
		CAMPAIGN_LOCALIZATION,
	}

	var keys []*datastore.Key
	var values []*Campaign

	for _, value := range names {
		status := &Campaign{
			Name:    value,
			Created: time.Now(),
		}
		key := datastore.NewKey(c, DATASTORE_CAMPAIGN, value, 0, nil)
		keys = append(keys, key)
		values = append(values, status)
	}

	_, err := datastore.PutMulti(c, keys, values)
	return err
}

func (campaign *Campaign) String() string {
	return campaign.Name
}

func (campaign *Campaign) createCampaignParams() *CampaignParams {

	params := &CampaignParams{
		Name: "Param1",
	}

	return params
}

func (campaign *Campaign) searchNewApps(r *http.Request) (error) {
	db, err := connectBigQueryDB(r, BIGQUERY_TABLE_DATA)
	if err != nil {
		return err
	}

	params := campaign.createCampaignParams()
	appsBuffer, err := db.Search(params)
	if err != nil {
		return err
	}

    //с := appengine.NewContext(r)

    apps := make([]App, len(*appsBuffer))
    for i, appBuffer := range *appsBuffer {
        apps[i] = appBuffer.toApp()
        /*
        с.Debugf("TrackId: ",     apps[i].TrackId)
        с.Debugf("ReleaseDate: ", apps[i].ReleaseDate)
        c.Debugf("SupportedDevices: ", items[i].SupportedDevices)
        c.Debugf("ArtistName: ", items[i].ArtistName)
        c.Debugf("Description: ", items[i].Description)
        c.Debugf("LanguageCodesISO2A: ", items[i].LanguageCodesISO2A)
        */
    }

    campaign.Apps = &apps

	return nil
}

func (campaign *Campaign) save(r *http.Request) (error) {
    var err error
    err = campaign.saveApps(r)
    if err != nil {
        return err
    }

    err = campaign.saveProceed(r)
    if err != nil {
        return err
    }

    err = campaign.saveTickets(r)
    if err != nil {
        return err
    }

    return nil
}

func (campaign *Campaign) saveApps(r *http.Request) (error) {
    c := appengine.NewContext(r)

    keys := make([]*datastore.Key, len(*campaign.Apps))

    for i, app := range *campaign.Apps {
        key_name := []string{ DATASTORE_APP,
                              DATASTORE_STORE,
                              strconv.Itoa(app.TrackId),
                            }
        keys[i] = datastore.NewKey(c, DATASTORE_APP, strings.Join(key_name, "_"), 0, nil)
    }

    _, err := datastore.PutMulti(c, keys, *campaign.Apps)

    campaign.Keys = keys
    return err
}


func (campaign *Campaign) saveProceed(r *http.Request) (error) {
    db, err := connectBigQueryDB(r, BIGQUERY_TABLE_PROCEED)
    if err != nil {
        return err
    }

    appsProceed, err := campaign.getProceed()
    if err != nil {
        return err
    }

    err = db.Insert(appsProceed)
    if err != nil {
        return err
    }
    return nil
}

func (campaign *Campaign) getProceed() (*[]AppProceed, error) {
	if campaign.Apps == nil {
		return nil, errors.New("No apps to be proceed")
	}

    items := make([]AppProceed, len(*(campaign.Apps)))
    for i, app := range *(campaign.Apps) {
        items[i].TrackId = app.TrackId
        items[i].Campaign = campaign.Name
        items[i].Created = time.Now()
    }
    return &items, nil
}


func (campaign *Campaign) saveTickets(r *http.Request) (error) {
    c := appengine.NewContext(r)

    _, сompaignKey, err := getCampaignByType(c, campaign.Name)
    if err != nil {
        return err
    }

    statusKey, err := getStatusByType(c, STATUS_NEW_APP)
    if err != nil {
        return err
    }

    keys := make([]*datastore.Key, len(*campaign.Apps))
    tickets := make([]Ticket, len(*campaign.Apps))

    for i, app := range *campaign.Apps {
        key_name := []string{ 
                              DATASTORE_TICKET,
                              campaign.Name,
                              DATASTORE_STORE,
                              strconv.Itoa(app.TrackId),
                            }
        keys[i] = datastore.NewKey(c, DATASTORE_TICKET, strings.Join(key_name, "_"), 0, nil)

        ticket := &Ticket{
            App: campaign.Keys[i],
            Compaign: сompaignKey,
            Status: statusKey,
            Created: time.Now(),
            Modified: time.Now(),
        }
        tickets[i] = *ticket
    }

    _, err = datastore.PutMulti(c, keys, tickets)

    return err
}

