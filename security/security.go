package security

import (
	"log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "os"
)

//TODO should be moved elsewhere
type UserData struct {
    Azp string `json:"azp"`
    Aud string `json:"aud"`
    Sub string `json:"sub"`
    AtHash string `json:"at_hash"`
    Iss string `json:"iss"`
    Iat string `json:"iat"`
    Exp string `json:"exp"`
    Alg string `json:"alg"`
    Kid string `json:"kid"`
    Email string `json:"email"`
    EmailVerified string `json:"email_verified"`
    
    // Name string `json:"name"`
    // GivenName string `json:"given_name"`
    // FamilyName string `json:"family_name"`
    // Profile string `json:"profile"`
    // Picture string `json:"picture"`
    // Email string `json:"email"`
    // EmailVerified string `json:"email_verified"`
    // Gender string `json:"gender"`
}

//TODO should be moved elsewhere
type Config struct {
    Cid string `json:"cid"`
}

var User UserData

// Handle security middleware aims to implement a JWT authentication.
func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        url := "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=";
        
        tokenString := r.Header.Get("Authorization")

        if tokenString == "" {
            log.Printf("No token string")
            http.Error(w, "No token string", http.StatusBadRequest)
    
            return
        }

        tokenString = tokenString[7:]
        log.Printf("tokenstring found %v", tokenString)
        
        url = url + tokenString
        resp, err := http.Get(url)

        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&User)

		if err != nil {
            log.Printf("Erreur found in json response %v", err)
			return
        }
        
        if (checkUserAudEqualsClientId()) {
            next.ServeHTTP(w, r)

            return
        }
    
        log.Printf("User %v", User)
        http.Error(w, "", http.StatusBadRequest)
	})
}

/**
* 
*
**/
func checkUserAudEqualsClientId() bool {
    file, e := ioutil.ReadFile("./config/config.json")

    if e != nil {
        log.Printf("File error! %v", e)    
        os.Exit(1)
    }

    var config Config
    json.Unmarshal(file, &config)

    log.Printf("Comparison user Aud (%s) == Config ClientId (%s)", User.Aud, config.Cid)

    return User.Aud == config.Cid
}
