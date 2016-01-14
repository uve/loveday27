package core

import (
    "time"
	"encoding/json"

	bigquery "google.golang.org/api/bigquery/v2"
)

type App struct {
	Advisories []string
	ArtistId int
	ArtistName string
	ArtistViewUrl string
	ArtworkUrl100 string
	ArtworkUrl512 string
	ArtworkUrl60 string
	AverageUserRating float32
	AverageUserRatingForCurrentVersion float32
	BundleId string
	ContentAdvisoryRating string
	Currency string
	Description string
	Features []string
	FileSizeBytes string
	FormattedPrice string
	GenreIds []string
	Genres []string
	IpadScreenshotUrls []string
	IsGameCenterEnabled bool
	//IsVppDeviceBasedLicensingEnabled bool `json:"isVppDeviceBasedLicensingEnabled"`
	Kind string
	LanguageCodesISO2A []string
	MinimumOsVersion string
	Price float32
	PrimaryGenreId int
	PrimaryGenreName string
	ReleaseDate time.Time
	ReleaseNotes string
	ScreenshotUrls []string
	SellerName string
	SellerUrl string
	SupportedDevices []string
	TrackCensoredName string
	TrackContentRating string
	TrackId int
	TrackName string
	TrackViewUrl string
	UserRatingCount float32
	UserRatingCountForCurrentVersion float32
	Version string
	WrapperType string
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

