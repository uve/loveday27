package core

import (
	"net/http"

	"appengine"
	"log"
	"text/template"
)

func handleMapPage(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    c.Debugf("map Page")

    /*
       campaign, nil := getCampaignByType(c, CAMPAIGN_LOCALIZATION)
       campaign.getAppByStatus(APP_STATUS_)
    */
    params, err := parseTemplateParams()
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    var index = template.Must(template.ParseFiles("templates/map.html"))

    err = index.Execute(w, params)
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

/*
   apps, err := searchApps(r)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
   }
*/
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

/*
func searchApps(r *http.Request) (*[]App, error) {
    c := appengine.NewContext(r)
    c.Debugf("Searching new apps")

    db, err := connectBigQueryDB(r, BIGQUERY_TABLE_DATA)
    if err != nil {
        return nil, err
    }
    apps, err := db.Search(&apps)
    if err != nil {
        return nil, err
    }

    return apps, nil
}



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
