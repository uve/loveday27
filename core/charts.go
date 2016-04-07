package core

import (
    "strings"
    "io/ioutil"
    "net/http"
    newappengine "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "encoding/base64"
)

type Charts struct {
    MapBody string
}

const (
    MAP_URL = "https://chart.googleapis.com/chart?cht=map:fixed=-60,-180,80,180&chs=600x500"
    MAP_BACKGROUND_COLOR = "DBDCDD"
    MAP_HIGHLIGHTED_COLOR = "8DC642"
)

func (calculations *Calculations) GetMap(r *http.Request) (string, error) {
    var mapCountries = []string{}
    var mapColors = []string{"DBDCDD"}

    for _, country := range calculations.CountryShare.Countries {
        mapCountries = append(mapCountries, country.Code)
        mapColors    = append(mapColors, MAP_HIGHLIGHTED_COLOR)
    }

    urlParams := []string{MAP_URL,
                    "chld=" + strings.Join(mapCountries, "|"),
                    "chco=" + strings.Join(mapColors, "|"),
    }

    ctx := newappengine.NewContext(r)
    client := urlfetch.Client(ctx)
    resp, err := client.Get(strings.Join(urlParams, "&"))
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    result := base64.StdEncoding.EncodeToString(body)
    result = "data:image/png;base64," + result
    return result, nil
}

func (calculations *Calculations) GetCharts(r *http.Request) (*Charts, error) {
    mapBody, err := calculations.GetMap(r)
    if err != nil {
        return nil, err
    }

    var charts = &Charts{
        MapBody: mapBody,
    }
    return charts, nil
}
