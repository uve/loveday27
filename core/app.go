package core

import (
	"encoding/json"
	"time"
	"strings"

	bigquery "google.golang.org/api/bigquery/v2"
)

const (
	APP_DESCRIPTION_MAX_LENGTH = 1500
)

type AppProceed struct {
	TrackId  int
	Campaign string
	Created time.Time
}

func (appBuffer *AppBuffer) toApp() (App) {
	var app App

	app.ArtistId = appBuffer.ArtistId
	app.ArtistName = appBuffer.ArtistName
	app.ArtistViewUrl = appBuffer.ArtistViewUrl
	app.ArtworkUrl100 = appBuffer.ArtworkUrl100
	app.ArtworkUrl512 = appBuffer.ArtworkUrl512
	app.ArtworkUrl60 = appBuffer.ArtworkUrl60
	app.AverageUserRating = appBuffer.AverageUserRating
	app.AverageUserRatingForCurrentVersion = appBuffer.AverageUserRatingForCurrentVersion
	app.BundleId = appBuffer.BundleId
	app.ContentAdvisoryRating = appBuffer.ContentAdvisoryRating
	app.Currency = appBuffer.Currency
	app.Description = appBuffer.Description

	if len(app.Description) > APP_DESCRIPTION_MAX_LENGTH {
		app.Description = app.Description[0:APP_DESCRIPTION_MAX_LENGTH]
	}

	app.FileSizeBytes = appBuffer.FileSizeBytes
	app.FormattedPrice = appBuffer.FormattedPrice
	app.IsGameCenterEnabled = appBuffer.IsGameCenterEnabled
	app.Kind = appBuffer.Kind
	app.MinimumOsVersion = appBuffer.MinimumOsVersion
	app.Price = appBuffer.Price
	app.PrimaryGenreId = appBuffer.PrimaryGenreId
	app.PrimaryGenreName = appBuffer.PrimaryGenreName
	//app.ReleaseDate = appBuffer.ReleaseDate
	app.ReleaseNotes = appBuffer.ReleaseNotes
	app.SellerName = appBuffer.SellerName
	app.SellerUrl = appBuffer.SellerUrl
	app.TrackCensoredName = appBuffer.TrackCensoredName
	app.TrackContentRating = appBuffer.TrackContentRating
	app.TrackId = appBuffer.TrackId
	app.TrackName = appBuffer.TrackName
	app.TrackViewUrl = appBuffer.TrackViewUrl
	app.UserRatingCount = appBuffer.UserRatingCount
	app.UserRatingCountForCurrentVersion = appBuffer.UserRatingCountForCurrentVersion
	app.Version = appBuffer.Version
	app.WrapperType = appBuffer.WrapperType

	//app.ReleaseDate = time.Unix(appBuffer.ReleaseDate, 0)

	app.ReleaseDate, _ = time.Parse("2006-01-02 15:04:05", appBuffer.ReleaseDate)

	app.Advisories         = strings.Split(appBuffer.Advisories, ",")
	app.Features           = strings.Split(appBuffer.Features, ",")
	app.GenreIds           = strings.Split(appBuffer.GenreIds, ",")
	app.Genres             = strings.Split(appBuffer.Genres, ",")
	app.IpadScreenshotUrls = strings.Split(appBuffer.IpadScreenshotUrls, ",")
	app.LanguageCodesISO2A = strings.Split(appBuffer.LanguageCodesISO2A, ",")
	app.ScreenshotUrls     = strings.Split(appBuffer.ScreenshotUrls, ",")
	app.SupportedDevices   = strings.Split(appBuffer.SupportedDevices, ",")

	app.Created = time.Now()
	
	return app
}


type AppBuffer struct {
    Advisories                         string
	ArtistId                           int
	ArtistName                         string
	ArtistViewUrl                      string
	ArtworkUrl100                      string
	ArtworkUrl512                      string
	ArtworkUrl60                       string
	AverageUserRating                  float32
	AverageUserRatingForCurrentVersion float32
	BundleId                           string
	ContentAdvisoryRating              string
	Currency                           string
	Description                        string
	Features                           string
	FileSizeBytes                      string
	FormattedPrice                     string
	GenreIds                           string
	Genres                             string
	IpadScreenshotUrls                 string
	IsGameCenterEnabled                bool
	//IsVppDeviceBasedLicensingEnabled bool `json:"isVppDeviceBasedLicensingEnabled"`
	Kind                             string
	LanguageCodesISO2A               string
	MinimumOsVersion                 string
	Price                            float32
	PrimaryGenreId                   int
	PrimaryGenreName                 string
	ReleaseDate                      string//time.Time
	ReleaseNotes                     string
	ScreenshotUrls                   string
	SellerName                       string
	SellerUrl                        string
	SupportedDevices                 string
	TrackCensoredName                string
	TrackContentRating               string
	TrackId                          int
	TrackName                        string
	TrackViewUrl                     string
	UserRatingCount                  float32
	UserRatingCountForCurrentVersion float32
	Version                          string
	WrapperType                      string

	LangsCount         int
	Campaign           string
}


type App struct {
	Advisories                         []string
	ArtistId                           int
	ArtistName                         string
	ArtistViewUrl                      string
	ArtworkUrl100                      string
	ArtworkUrl512                      string
	ArtworkUrl60                       string
	AverageUserRating                  float32
	AverageUserRatingForCurrentVersion float32
	BundleId                           string
	ContentAdvisoryRating              string
	Currency                           string
	Description                        string
	Features                           []string
	FileSizeBytes                      string
	FormattedPrice                     string
	GenreIds                           []string
	Genres                             []string
	IpadScreenshotUrls                 []string
	IsGameCenterEnabled                bool
	//IsVppDeviceBasedLicensingEnabled bool `json:"isVppDeviceBasedLicensingEnabled"`
	Kind                             string
	LanguageCodesISO2A               []string
	MinimumOsVersion                 string
	Price                            float32
	PrimaryGenreId                   int
	PrimaryGenreName                 string
	ReleaseDate                      time.Time
	ReleaseNotes                     string
	ScreenshotUrls                   []string
	SellerName                       string
	SellerUrl                        string
	SupportedDevices                 []string
	TrackCensoredName                string
	TrackContentRating               string
	TrackId                          int
	TrackName                        string
	TrackViewUrl                     string
	UserRatingCount                  float32
	UserRatingCountForCurrentVersion float32
	Version                          string
	WrapperType                      string

	Created                          time.Time
	//Campaign           string
}

func (app *App) getJson() (map[string]bigquery.JsonValue, error) {
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
