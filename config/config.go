package config

import (
	"fmt"
	"os"
	"encoding/json"
	//"common/oauth"
	
	"appengine"
	
	"github.com/martini-contrib/oauth2"
)


const config_prod = "config/config_prod.json"
const config_dev  = "config/config_dev.json"


type OAuthProviders struct{

	Google    oauth2.Options
	Facebook  oauth2.Options
	Github    oauth2.Options
	Twitter   oauth2.Options
	
}


var Config = struct {

	
	CookieName string
	CookieSecret string
	
	OAuthProviders OAuthProviders
	// Your OAuth configuration information for protected user data access.
	//OAuthConfig oauth.Config
	
	// The path in your application to which users will be redirected after they
	// allow or deny permission for your application to access their data.
	RedirectURL string
	// The scheme, hostname and port at which your application can be accessed
	// when running on App Engine.
	RootUrl string
}{}





	

func init() {	
	
	var configPath string
	
	if appengine.IsDevAppServer() {
	
		configPath = config_dev
		
	} else {
	
		configPath = config_prod	
	}
	
	
	configFile, err := os.Open(configPath)
	if err != nil {
		panic(fmt.Sprintf("Could not open %s: %s", configPath, err))
	}
	defer configFile.Close()


	if err = json.NewDecoder(configFile).Decode(&Config); err != nil {
		panic(fmt.Sprintf("Could not parse %s: %s", configPath, err))
	}


}