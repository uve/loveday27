package core

import (
    "log"
	"text/template"
    "errors"
    "time"
    "fmt"
	"strconv"

  "io/ioutil"
  "net/url"


	"net/http"
   "golang.org/x/oauth2"
   "golang.org/x/oauth2/vk"

   "appengine"
   newappengine "google.golang.org/appengine"

)

const secretCookieName = "invitation_id"
const secretCookieValue = 3453400000
const htmlIndex = `<html><body>
Login in with <a href="/auth">vk</a>
</body></html>
`


const htmlError = `<html><link rel="stylesheet" href="/static/css/main.css">
<body class="error-image">
  <script>
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
    })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

    ga('create', 'UA-9798363-12', 'auto');
    ga('send', 'pageview');

  </script>
</body></html>
`


type Guest struct {
    Id int
    Name string
    Description string
    IsAdmin bool
}


var Guests = []Guest{
    Guest{
            Id: 1184396,
            Name: "Никита",
            Description: "",
            IsAdmin: true,
        },
    Guest{
            Id: 22298116,
            Name: "Саша",
            Description: "",
            IsAdmin: true,
        },
    Guest{
            Id: 193578951,
            Name: "Юля",
            Description: "",
            IsAdmin: true,
        },
    Guest{
            Id: 133500597,
            Name: "Люба",
            Description: "",
            IsAdmin: true,
        },
    Guest{
            Id: 1873981,
            Name: "Юля",
            Description: "",
        },
    Guest{
            Id: 2297896,
            Name: "Даша",
            Description: "",
            IsAdmin: true,
        },
    Guest{
            Id: 2682290,
            Name: "Катя",
            Description: "",
        },
    Guest{
            Id: 112806371,
            Name: "Саша",
            Description: "",
        },
    Guest{
            Id: 256500244,
            Name: "Миша",
            Description: "",
        },
    Guest{
            Id: 49065541,
            Name: "Наталья",
            Description: "",
        },
    Guest{
            Id: 2806312,
            Name: "Лена",
            Description: "",
        },
    Guest{
            Id: 1724037,
            Name: "Аня",
            Description: "",
        },
    Guest{
            Id: 3018289,
            Name: "Лена",
            Description: "",
        },
    Guest{
            Id: 4548762,
            Name: "Маша",
            Description: "",
        },
    Guest{
            Id: 88237469,
            Name: "Света",
            Description: "",
        },
    Guest{
            Id: 2384460,
            Name: "Аня",
            Description: "",
        },
    Guest{
            Id: 3876751,
            Name: "Таня",
            Description: "",
        },
    Guest{
            Id: 7531487,
            Name: "Юля",
            Description: "",
        },
    Guest{
            Id: 14410032,
            Name: "Алёна",
            Description: "",
        },
    Guest{
            Id: 924223,
            Name: "Ира",
            Description: "",
        },
    Guest{
            Id: 2884436,
            Name: "Катя",
            Description: "",
        },
    Guest{
            Id: 2444334,
            Name: "Алла",
            Description: "",
        },
    Guest{
            Id: 279409437,
            Name: "Елена",
            Description: "",
        },
    Guest{
            Id: 9333181,
            Name: "Наташа",
            Description: "",
        },
    Guest{
            Id: 61034886,
            Name: "Таня",
            Description: "",
        },
    Guest{
            Id: 1410190,
            Name: "Саша",
            Description: "",
        },
    Guest{
            Id: 585208,
            Name: "Лиана",
            Description: "",
        },
    Guest{
            Id: 5137830,
            Name: "Ярик",
            Description: "",
        },
    Guest{
            Id: 16506866,
            Name: "Надя",
            Description: "",
        },
    Guest{
            Id: 183883702,
            Name: "Юля",
            Description: "",
        },
    Guest{
            Id: 11951555,
            Name: "Максим",
            Description: "",
        },
    Guest{
            Id: 5258309,
            Name: "Аня",
            Description: "",
        },
    Guest{
            Id: 55888743,
            Name: "Женя",
            Description: "",
        },
    Guest{
            Id: 3428873,
            Name: "Артём",
            Description: "",
        },
    }

var (
  oauthStateString = "success"
)




type Params struct {
    Guests []Guest
    GuestName string
}

func getConfig(r *http.Request) (*oauth2.Config) {
    redirectURL := "https://loveday27.ru/oauth2callback"
    if newappengine.IsDevAppServer() {
        redirectURL = "http://localhost:8080/oauth2callback"
    }

    return &oauth2.Config{
        ClientID:     "5453424",
        ClientSecret: "0L6m3M9jHzIbUxuPGoCr",
        RedirectURL:  redirectURL,
        Scopes:       []string{},
        Endpoint: vk.Endpoint,
    }
}


func setCookie(w http.ResponseWriter, user_id int) {
   value := strconv.Itoa(user_id)

   expiration := time.Now().Add(60 * 24 * time.Hour)
   cookie :=  http.Cookie{Name: secretCookieName, Value: value, Expires:expiration}
   http.SetCookie(w, &cookie)
}


func parseTemplateParams() (*Params, error) {
    params := Params{
        Guests: Guests,
    }

    return &params, nil
}

func handleAuth(w http.ResponseWriter, r *http.Request) {

    c := appengine.NewContext(r)

    conf := getConfig(r)

    // Redirect user to consent page to ask for permission
    // for the scopes specified above.
    url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline,
                oauth2.SetAuthURLParam("response_type", "code"),
                oauth2.SetAuthURLParam("v", "5.52"),
                oauth2.SetAuthURLParam("state", oauthStateString),
                oauth2.SetAuthURLParam("grant_type", "client_credentials"))

    c.Debugf("Visit the URL for the auth dialog: ", url)

    http.Redirect(w, r, url, 302)
    /*
    // Use the authorization code that is pushed to the redirect URL.
    // NewTransportWithCode will do the handshake to retrieve
    // an access token and initiate a Transport that is
    // authorized and authenticated by the retrieved token.
    var code string
    if _, err := fmt.Scan(&code); err != nil {
        log.Fatal(err)
    }
    tok, err := conf.Exchange(oauth2.NoContext, code)
    if err != nil {
        log.Fatal(err)
    }

    client := conf.Client(oauth2.NoContext, tok)
    client.Get("...")
    */
}


func handleError(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(htmlError))
}

func handleLoginPage(w http.ResponseWriter, r *http.Request) {
    params, err := parseTemplateParams()
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    var index = template.Must(template.ParseFiles("templates/login.html"))

    err = index.Execute(w, params)
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func isAuthorized(w http.ResponseWriter, r *http.Request) (error) {

    // Get the key from the URL
    specialKey := r.FormValue("sp")
    if specialKey == "0207" {
       setCookie(w, secretCookieValue)
       return nil
    }

    cookie, err := r.Cookie(secretCookieName)
    if err != nil {
        return err
    }

    user_id, err := strconv.Atoi(cookie.Value)
    if err != nil {
        return err
    }

    return checkUserPermissions(user_id)
}

func checkUserPermissions(user_id int) (error) {
    for _, guest := range Guests {
        if guest.Id == user_id {
            return nil
        }
    }

    if secretCookieValue == user_id {
        return nil
    }

    return errors.New("User is not in the list")
}


func handleMainPage(w http.ResponseWriter, r *http.Request) {

    err := isAuthorized(w, r)
    if err != nil {
        //log.Fatalf("isAuthorized execution: %s", err)
        http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
        return
    }

    params, err := parseTemplateParams()
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    var index = template.Must(template.ParseFiles("templates/index.html"))

    err = index.Execute(w, params)
    if err != nil {
        log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func handleCallback(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  ctx := newappengine.NewContext(r)

  state := r.FormValue("state")
  if state != oauthStateString {
     c.Debugf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
     http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
    return
  }

  code := r.FormValue("code")

  conf := getConfig(r)

  token, err := conf.Exchange(ctx, code)
  if err != nil {
     c.Debugf("conf.Exchange() failed with '%s'\n", err)
     http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
    return
  }


  user_id_float := token.Extra("user_id")

  s := fmt.Sprintf("%.0f", user_id_float)
  user_id, err := strconv.Atoi(s)
  if err != nil {
    c.Debugf("err: ", err)

    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
  }


  c.Debugf("user_id: ", user_id)
  err = checkUserPermissions(user_id)
  if err != nil {
    c.Errorf("Error: ", err)
    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
  }

  client := conf.Client(ctx, token)

  resp, err := client.Get("https://api.vk.com/method/users.get?user_id=" + s + "&v=5.52&access_token=" +
    url.QueryEscape(token.AccessToken))
  if err != nil {
     c.Debugf("Get: %s\n", err)
    http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
    return
  }
  defer resp.Body.Close()

  response, err := ioutil.ReadAll(resp.Body)
  if err != nil {
     c.Debugf("ReadAll: %s\n", err)
     http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
     return
  }

   c.Debugf("parseResponseBody: %s\n", string(response))

   setCookie(w, user_id)

   http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}



func init() {
   http.HandleFunc("/error", handleError)
   http.HandleFunc("/login", handleLoginPage)
   http.HandleFunc("/auth", handleAuth)
   http.HandleFunc("/oauth2callback", handleCallback)
   http.HandleFunc("/", handleMainPage)
}
