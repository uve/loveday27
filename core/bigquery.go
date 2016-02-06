package core

import (
	"container/list"
	"errors"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
	newappengine "google.golang.org/appengine"
	"net/http"
	//newurlfetch "google.golang.org/appengine/urlfetch"

	"encoding/json"
	//"appengine"
	//"fmt"
	//"strconv"

	//"github.com/SeanDolphin/bqschema"

    //"appengine"
    "bytes"
    "html/template"
)

const (
	BIGQUERY_PROJECT       = "cometiphrd"
	BIGQUERY_DATASET       = "appstore"
	BIGQUERY_TABLE_DATA    = "data"
	BIGQUERY_TABLE_PROCEED = "apps_proceed2"
    BIGQUERY_QUERY_LIMIT   = 5
)

// Wraps the BigQuery service and dataset and provides some helper functions.
type bqDataset struct {
	ProjectId string
	DatasetId string
	TableId   string
	bq        *bigquery.Service
	dataset   *bigquery.Dataset
	jobsets   map[string]*list.List
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
		bq:        service,
		dataset: &bigquery.Dataset{
			DatasetReference: &bigquery.DatasetReference{
				DatasetId: datasetId,
				ProjectId: projectId,
			},
		},
		jobsets: make(map[string]*list.List),
	}, nil
}

    const queryNewAppsTmpl = `SELECT                    

                    FIRST(data.trackId) TrackId,

                    COUNT(data.languageCodesISO2A) LangsCount,
                    proceed.Campaign Campaign,

                    FORMAT_UTC_USEC(data.releaseDate) ReleaseDate,
                    GROUP_CONCAT(data.advisories, ",") Advisories,
                    GROUP_CONCAT(data.features, ",") Features,
                    GROUP_CONCAT(data.genreIds, ",") GenreIds,
                    GROUP_CONCAT(data.genres, ",") Genres,
                    GROUP_CONCAT(data.ipadScreenshotUrls, ",") IpadScreenshotUrls,
                    GROUP_CONCAT(data.languageCodesISO2A, ",") LanguageCodesISO2A,
                    GROUP_CONCAT(data.screenshotUrls, ",") ScreenshotUrls,
                    GROUP_CONCAT(data.supportedDevices, ",") SupportedDevices,
                    
                    data.artistId ArtistId,
                    data.artistName ArtistName,
                    data.artistViewUrl ArtistViewUrl,
                    data.artworkUrl100 ArtworkUrl100,
                    data.artworkUrl512 ArtworkUrl512,
                    data.artworkUrl60 ArtworkUrl60,
                    data.averageUserRating AverageUserRating,
                    data.averageUserRatingForCurrentVersion AverageUserRatingForCurrentVersion,
                    data.bundleId BundleId,
                    data.contentAdvisoryRating ContentAdvisoryRating,
                    data.currency Currency,
                    data.description Description,
                    data.fileSizeBytes FileSizeBytes,
                    data.formattedPrice FormattedPrice,
                    data.isGameCenterEnabled IsGameCenterEnabled,
                    data.kind Kind,
                    data.minimumOsVersion MinimumOsVersion,
                    data.price Price,
                    data.primaryGenreId PrimaryGenreId,
                    data.primaryGenreName PrimaryGenreName,
                    data.releaseNotes ReleaseNotes,
                    data.sellerName SellerName,
                    data.sellerUrl SellerUrl,
                    data.trackCensoredName TrackCensoredName,
                    data.trackContentRating TrackContentRating,
                    
                    data.trackName TrackName,
                    data.trackViewUrl TrackViewUrl,
                    data.userRatingCount UserRatingCount,
                    data.userRatingCountForCurrentVersion UserRatingCountForCurrentVersion,
                    data.version Version,
                    data.wrapperType WrapperType

                    FROM  [{{.DATASET}}.{{.TABLE_DATA}}] as data
            LEFT JOIN [{{.DATASET}}.{{.TABLE_PROCEED}}] as proceed
            on data.trackId = proceed.TrackId
            GROUP BY
                    ArtistId,
                    ArtistName,
                    ArtistViewUrl,
                    ArtworkUrl100,
                    ArtworkUrl512,
                    ArtworkUrl60,
                    AverageUserRating,
                    AverageUserRatingForCurrentVersion,
                    BundleId,
                    ContentAdvisoryRating,
                    Currency,
                    Description,
                    FileSizeBytes,
                    FormattedPrice,
                    IsGameCenterEnabled,
                    Kind,
                    MinimumOsVersion,
                    Price,
                    PrimaryGenreId,
                    PrimaryGenreName,
                    ReleaseDate,
                    ReleaseNotes,
                    SellerName,
                    SellerUrl,
                    TrackCensoredName,
                    TrackContentRating,
                    TrackName,
                    TrackViewUrl,
                    UserRatingCount,
                    UserRatingCountForCurrentVersion,
                    Version,
                    WrapperType,
                    Campaign
            HAVING LangsCount > 0
                  AND Campaign IS NULL
            ORDER BY LangsCount ASC
            LIMIT {{.LIMIT}}`

func (ds *bqDataset) Search(params *CampaignParams) (*[]AppBuffer, error) {
	
    data := map[string]interface{}{
        "LIMIT":         BIGQUERY_QUERY_LIMIT,
        "DATASET":       BIGQUERY_DATASET,
        "TABLE_DATA":    BIGQUERY_TABLE_DATA,
        "TABLE_PROCEED": BIGQUERY_TABLE_PROCEED,
    }

    t := template.Must(template.New("email").Parse(queryNewAppsTmpl))
    buf := &bytes.Buffer{}
    if err := t.Execute(buf, data); err != nil {
        panic(err)
    }

	queryRequest := &bigquery.QueryRequest{
		//DefaultDataset: datasetRef,
		Query:      buf.String(),
		MaxResults: BIGQUERY_QUERY_LIMIT, //int64(max),
		Kind:       "json",
		//Kind: "igquery#queryRequest",
		TimeoutMs: 60000, //defaultTimeOutMs,
	}

	jobsService := bigquery.NewJobsService(ds.bq)
	queryResponse, err := jobsService.Query(BIGQUERY_PROJECT, queryRequest).Do()
	if err != nil {
		return nil, err
	}

    //c.Debugf(strconv.Itoa(len(queryResponse.Rows)))

	var appsBuffer []AppBuffer
	err = toStructs(queryResponse, &appsBuffer)
	if err != nil {
		return nil, err
	}

	return &appsBuffer, nil
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
