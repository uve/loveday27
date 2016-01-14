package core

import (
	"net/http"

	"appengine"
	bigquery "google.golang.org/api/bigquery/v2"

	"encoding/json"
)


type AppProceed struct {
	TrackId int
	Campaign string
}


func searchPage(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    c.Debugf("search Page")

    campaign, err := getCampaignByType(c, CAMPAIGN_LOCALIZATION)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    c.Debugf("campaign: ", campaign)

    apps, err := campaign.searchNewApps(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for _, app := range *apps {
        c.Debugf("app TrackId: ", app.TrackId)
        c.Debugf("app TrackName: ", app.TrackName)
    }
    //apps.save()

    //appProceed := apps.proceed()
    //appProceed.save()
/*
    var apps []AppProceed

    sample_app := AppProceed{
    	TrackId: 281656475,
    	Campaign: CAMPAIGN_LOCALIZATION,
    }
    apps = append(apps, sample_app)
*/
}

/*

func (apps *[]App) proceed() ([]AppProceed, error) {
	c.Debugf("Saving %d apps to AppEngine", len(*apps))

    items := make([]AppProceed, len(*apps))
    for i, app := range apps {
        items[i].TrackId = app.TrackId
        items[i].Campaign = CAMPAIGN_LOCALIZATION
    }
    return items
}*/
/*
func (apps *[]App) save() (error) {
	c.Debugf("Saving %d apps to AppEngine", len(*apps))

	return nil
}


func (apps *[]AppProceed) save(r *http.Request) (error) {
	c.Debugf("Saving %d apps to BigQuery", len(*apps))

    db, err := connectBigQueryDB(r, BIGQUERY_TABLE_PROCEED)
    if err != nil {
        return err
    }
    err = db.Insert(&apps)
    if err != nil {
        return err
    }

    return nil
}
*/

func (app *AppProceed) getJson() (map[string]bigquery.JsonValue, error) {
	b, err := json.Marshal(app)
    if err != nil {
        return nil, err
    }

    var Json map[string]bigquery.JsonValue
	err = json.Unmarshal(b, &Json)
    if err != nil {
        return nil, err
    }
	return Json, nil
}

