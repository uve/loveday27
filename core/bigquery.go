package core

import (
	"container/list"
	"net/http"
	"errors"
    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
	newappengine "google.golang.org/appengine"
    //newurlfetch "google.golang.org/appengine/urlfetch"

	"encoding/json"
	//"appengine"
    //"fmt"
    //"strconv"

    "github.com/SeanDolphin/bqschema"
)

const (
	BIGQUERY_PROJECT = "cometiphrd"
	BIGQUERY_DATASET = "october"//"appstore"
	BIGQUERY_TABLE_DATA = "data"
	BIGQUERY_TABLE_PROCEED = "apps_proceed"
)


// Wraps the BigQuery service and dataset and provides some helper functions.
type bqDataset struct {
	ProjectId string
	DatasetId string
	TableId string
	bq      *bigquery.Service
	dataset *bigquery.Dataset
	jobsets map[string]*list.List
}


func newBQDataset(client *http.Client, projectId string, datasetId string, tableId string) (*bqDataset,
	error) {

	service, err := bigquery.New(client)
	if err != nil {
		return nil, err
	}

	return &bqDataset{
		ProjectId: projectId,
		DatasetId: datasetId,
		TableId:   tableId,
		bq:      service,
		dataset: &bigquery.Dataset{
			DatasetReference: &bigquery.DatasetReference{
				DatasetId: datasetId,
				ProjectId: projectId,
			},
		},
		jobsets: make(map[string]*list.List),
	}, nil
}

func (ds *bqDataset) Search(params *CampaignParams) (*[]App, error) {
  	/*rows := make([]*bigquery.TableDataInsertAllRequestRows, len(*apps))
 
    for i, app := range *apps {
    	rows[i] = new(bigquery.TableDataInsertAllRequestRows)
    	Json, err := app.getJson()
    	if err != nil {
    		return err
    	}
    	rows[i].Json = Json
    }

	insertRequest := &bigquery.TableDataInsertAllRequest{Rows: rows}
	//fmt.Println(ds.ProjectId, ds.DatasetId, ds.TableId)
    _, err := ds.bq.Tabledata.InsertAll(ds.ProjectId, ds.DatasetId, ds.TableId, insertRequest).Do()*/

    // prepare the list of ids.
   /* var ids []string
    var appID = BIGQUERY_PROJECT
    datasets, err := ds.Datasets.List(appID).Do()
    if err != nil {
        return nil, fmt.Errorf("could not list datasets for %q: %v", appID, err)
    }
    for _, d := range datasets.Datasets {
        ids = append(ids, d.Id)
    }
    return ids, nil
*/
    //query := `SELECT data.trackId trackId, data.trackName trackName, data.sellerUrl sellerUrl, data.trackViewUrl trackViewUrl FROM [appstore.data] data
    query := `SELECT 
                    artistId,
                    artistName,
                    artistViewUrl,
                    artworkUrl100,
                    artworkUrl512,
                    artworkUrl60,
                    averageUserRating,
                    averageUserRatingForCurrentVersion,
                    bundleId,
                    contentAdvisoryRating,
                    currency,
                    description,
                    /*features,*/
                    fileSizeBytes,
                    formattedPrice,
                    /*genreIds,
                    genres,
                    ipadScreenshotUrls,*/
                    isGameCenterEnabled,
                    kind,
                    /*languageCodesISO2A,*/
                    minimumOsVersion,
                    price,
                    primaryGenreId,
                    primaryGenreName,
                    releaseDate,
                    releaseNotes,
                    /*screenshotUrls,*/
                    sellerName,
                    sellerUrl,
                    /*supportedDevices,*/
                    trackCensoredName,
                    trackContentRating,
                    trackId,
                    trackName,
                    trackViewUrl,
                    userRatingCount,
                    userRatingCountForCurrentVersion,
                    version,
                    wrapperType
                    
                    FROM                            
                           /// FLATTEN(
                                [appstore.data]
                          //  , LanguageCodesISO2A)
                             as data
            /// LEFT JOIN (SELECT proceed.TrackId TrackId, proceed.Campaign Campaign FROM [october.tracks_proceed] proceed 
            // WHERE Campaign = 'CAMPAIGN_LOCALIZATION') as proceed
            // on data.trackId = proceed.TrackId
              WHERE TrackId NOT IN (SELECT TrackId FROM [october.tracks_proceed] proceed 
             WHERE Campaign = 'CAMPAIGN_LOCALIZATION')
            ORDER BY TrackId LIMIT 5`

    queryRequest := &bigquery.QueryRequest{
        //DefaultDataset: datasetRef,
        Query: query,
        MaxResults:     5,//int64(max),
        Kind:           "json",
        //Kind: "igquery#queryRequest",
        //TimeoutMs:      100000000,//defaultTimeOutMs,
    }

    jobsService := bigquery.NewJobsService(ds.bq)
    queryResponse, err := jobsService.Query(BIGQUERY_PROJECT, queryRequest).Do()
    if err != nil {
        return nil, err
    }
  
    var apps []App
    err = bqschema.ToStructs(queryResponse, &apps)
    if err != nil {
        return nil, err
    }
    return &apps, nil
    /*
    count := strconv.Itoa(len(apps))

    ids = append(ids, "c")
    ids = append(ids, count)
    ids = append(ids, strconv.Itoa(apps[0].TrackId))
    ids = append(ids, strconv.Itoa(apps[1].TrackId))
    ids = append(ids, strconv.Itoa(apps[2].TrackId))
    ids = append(ids, strconv.Itoa(apps[3].TrackId))
    ids = append(ids, strconv.Itoa(apps[4].TrackId))
/*
    b, err := json.Marshal(apps[1])
    if err != nil {
        return nil, err
    }
    ids = append(ids, string(b))

    return ids, nil
    */
}


func getString(row *bigquery.TableRow) (string, error) {
    b, err := json.Marshal(row)
    if err != nil {
        return "", err
    }
    return string(b), err
}

func getJson(row *bigquery.TableRow) (*App, error) {
    b, err := json.Marshal(row)
    if err != nil {
        return nil, err
    }

    var app *App
    err = json.Unmarshal(b, &app)
    if err != nil {
        return nil, err
    }
    return app, nil
}


func (ds *bqDataset) Insert(apps *[]AppProceed) error {
  	rows := make([]*bigquery.TableDataInsertAllRequestRows, len(*apps))
 
    for i, app := range *apps {
    	rows[i] = new(bigquery.TableDataInsertAllRequestRows)
    	Json, err := app.getJson()
    	if err != nil {
    		return err
    	}
    	rows[i].Json = Json
    }

	insertRequest := &bigquery.TableDataInsertAllRequest{Rows: rows}
	//fmt.Println(ds.ProjectId, ds.DatasetId, ds.TableId)
    _, err := ds.bq.Tabledata.InsertAll(ds.ProjectId, ds.DatasetId, ds.TableId, insertRequest).Do()
	return err
}


func connectBigQueryDB(r *http.Request, table string) (*bqDataset, error) {
	if len(table) < 0 {
	    return nil, errors.New("BigQuery table name is not defined")
	}

	var ctx context.Context = newappengine.NewContext(r)

	// Use oauth2.NoContext if there isn't a good context to pass in.
    //ctx := context.Background()
    ts, err := google.DefaultTokenSource(ctx, bigquery.BigqueryScope,
    										  //storage.DevstorageReadOnlyScope,
    										  "https://www.googleapis.com/auth/userinfo.profile")
	if err != nil {
	    return nil, err
	}
	client := oauth2.NewClient(ctx, ts)

	return newBQDataset(client, BIGQUERY_PROJECT, BIGQUERY_DATASET, table)
}
