package frontend

import (

	"io"
	"log"
	//"fmt"

	"net/http"
	"text/template"


	//"github.com/martini-contrib/oauth2"
	//"github.com/martini-contrib/sessions"
    //"github.com/go-martini/martini"
	//"github.com/martini-contrib/cors"

	//"core/user"

	//"config"
	//"github.com/crhym3/go-endpoints/endpoints"


	//"api"

)



func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found")
}


type Params struct {
	ClientId  string
}



func handleMainPage(w http.ResponseWriter, r *http.Request) {




	params := Params{
		//ClientId: config.Config.OAuthProviders.Google.ClientId,
		ClientId: "882975820932-q34i2m1lklcmv8kqqrcleumtdhe4qbhk.apps.googleusercontent.com",
	}


	var index = template.Must(template.ParseFiles("templates/index.html"))

	err := index.Execute(w, params)
    if err != nil {
    	log.Fatalf("template execution: %s", err)
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

	
}


func handleMapPage(w http.ResponseWriter, r *http.Request) {




	params := Params{
		//ClientId: config.Config.OAuthProviders.Google.ClientId,
		ClientId: "882975820932-q34i2m1lklcmv8kqqrcleumtdhe4qbhk.apps.googleusercontent.com",
	}


	var index = template.Must(template.ParseFiles("polymer/index.html"))

	err := index.Execute(w, params)
    if err != nil {
    	log.Fatalf("template execution: %s", err)
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

	
}

/*
func oauth2error(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Auth error")
}



func PaymentsPage(u *user.User, s sessions.Session, c martini.Context, w http.ResponseWriter, r *http.Request) {



	fmt.Fprint(w, c)

}

*/





func init() {

	
	http.HandleFunc("/map", handleMapPage)
	http.HandleFunc("/", handleMainPage)
	


	/*

	m := martini.Classic()
	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"https://storage.googleapis.com"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//w.Header().Set("Access-Control-Allow-Origin", "https://storage.googleapis.com")
	

	//params := oauth2.Options( config.Config.OAuthProviders.Google )
	//params.RedirectURL = config.Config.RedirectURL
	//m.Use(oauth2.Google(&params))


	store := sessions.NewCookieStore([]byte(config.Config.CookieSecret))
	m.Use(sessions.Sessions(config.Config.CookieName, store))


	m.Get("/oauth2error", oauth2error)


	//m.Get("/", testPage)




	m.Get("/", handleMainPage)

	m.Get("/payments", oauth2.LoginRequired, PaymentsPage)
	/*

	// Tokens are injected to the handlers
	m.Get("/", func(tokens oauth2.Tokens) string {
		if tokens.IsExpired() {




			return "not logged in, or the access token is expired2"
		}
		return "logged in"

	})



	m.Get("/restrict", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
			return tokens.Access()
	})
	
	

	http.Handle("/", m)
	*/


   /*

	if _, err := api.RegisterService(); err != nil {
		panic(err.Error())
	}

	endpoints.HandleHttp()
	*/

}

