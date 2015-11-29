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

	//"encoding/json"
	//"appengine"
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


func (ds *bqDataset) Search(apps *[]Ap) error {
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
	if len(table) < nil {
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
