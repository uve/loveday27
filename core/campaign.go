package core

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"time"
	"errors"
)

const (
	DATASTORE_CAMPAIGN    = "CAMPAIGN"
	CAMPAIGN_LOCALIZATION = "CAMPAIGN_LOCALIZATION"
)

type Campaign struct {
	Name    string
	Created time.Time
	Apps *[]App
}

type CampaignParams struct {
	Name  string
	Limit int
}

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

