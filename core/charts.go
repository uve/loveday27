package core

import (
    "strings"
    "io/ioutil"
    "net/http"
    newappengine "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "encoding/base64"
)

type LangsChart struct {
    Src string
    Body string
    Width int
    Height int
}

type Charts struct {
    MapBody string
    LangsChart *LangsChart
}

const (
    MAP_URL = "https://chart.googleapis.com/chart?cht=map:fixed=-60,-180,80,180&chs=600x500"
    MAP_BACKGROUND_COLOR = "DBDCDD"
    MAP_HIGHLIGHTED_COLOR = "8DC642"

    LANGS_URL = "https://chart.googleapis.com/chart?cht=bhs&chs=600x500&chxt=x,y"
    LANGS_BACKGROUND_COLOR = "CECECE"
    LANGS_HIGHLIGHTED_COLOR = "01B0F1"
    LangsChartWidth = 500
    LangsChartHeight = 600
)

func (calculations *Calculations) GetLangsChart(r *http.Request) (*LangsChart, error) {
    var mapLangs = []string{"1:"}
    var mapColors = []string{}
    var mapPopulations = []string{}

    for _, lang := range calculations.LangShare.Langs {
        mapLangs = append(mapLangs, lang.Name)
        mapColors = append(mapColors, LANGS_HIGHLIGHTED_COLOR)
        mapPopulations = append(mapPopulations, "50")
    }

    for _, lang := range ALL_LANGS {
        mapLangs = append(mapLangs, lang.Name)
        mapColors = append(mapColors, LANGS_BACKGROUND_COLOR)
        mapPopulations = append(mapPopulations, "20")
    }

    urlParams := []string{LANGS_URL,
                    "chxl=" + strings.Join(mapLangs, "|"),
                    "chd=t:" + strings.Join(mapPopulations, ","),
                    "chco=" + strings.Join(mapColors, "|"),
    }
    src := strings.Join(urlParams, "&")
    ctx := newappengine.NewContext(r)
    client := urlfetch.Client(ctx)
    resp, err := client.Get(src)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    bodySrc := "data:image/png;base64," + base64.StdEncoding.EncodeToString(body)

    result := &LangsChart{
        Src: src,
        Body: bodySrc,
        Height: LangsChartHeight,
        Width: LangsChartWidth,
    }
    return result, nil
}

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

    langsChart, err := calculations.GetLangsChart(r)
    if err != nil {
        return nil, err
    }

    var charts = &Charts{
        MapBody: mapBody,
        LangsChart: langsChart,
    }
    return charts, nil
}
