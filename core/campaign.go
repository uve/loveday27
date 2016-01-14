package core

import (
	"appengine"	
	"appengine/datastore"
	"time"
	"net/http"
)

const (
	DATASTORE_CAMPAIGN      = "CAMPAIGN"
    CAMPAIGN_LOCALIZATION   = "CAMPAIGN_LOCALIZATION"
)

type Campaign struct {
	Name  string
	Created  time.Time
}

type CampaignParams struct {
	Name  string
	Limit  int
}

//func getCampaignByType(c appengine.Context, name string) (*datastore.Key, error) {
func getCampaignByType(c appengine.Context, name string) (*Campaign, error) {
	key := datastore.NewKey(c, DATASTORE_CAMPAIGN, name, 0, nil)
	e := new(Campaign)
	if err := datastore.Get(c, key, e); err != nil {
		return nil, err
	}
	return e, nil
}


func getCampaignCount(c appengine.Context) (int, error) {

	q := datastore.NewQuery(DATASTORE_CAMPAIGN)

	count, err := q.Count(c)
	if err != nil {	
		return 0, err
	}
	return count, nil
}


func createAllCampaign(c appengine.Context) (error) {
	names := [...]string { 
		CAMPAIGN_LOCALIZATION,
	}

	var keys []*datastore.Key
	var values []*Campaign

	for _, value := range names {
    		status := &Campaign{
		        Name: value,
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


func (campaign *Campaign) createCampaignParams() (*CampaignParams) {

	params := &CampaignParams{
		Name: "Param1",
	}

    return params
}


func (campaign *Campaign) searchNewApps(r *http.Request) (*[]App, error) {
    c := appengine.NewContext(r)
	c.Debugf("Searching new apps")

    db, err := connectBigQueryDB(r, BIGQUERY_TABLE_DATA)
    if err != nil {
        return nil, err
    }

    c.Debugf("db", db)

    params := campaign.createCampaignParams()

    apps, err := db.Search(params)
    if err != nil {
        return nil, err
    }

    return apps, nil
}
