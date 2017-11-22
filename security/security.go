package security

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//TODO should be moved elsewhere
type UserData struct {
	FamilyName    string `json:"family_name"`
	Gender        string `json:"gender"`
	GivenName     string `json:"given_name"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Nickname      string `json:"nickname"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
	UpdatedAt     string `json:"updated_at"`
	ClientID      string `json:"clientId"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

//TODO should be moved elsewhere
type Config struct {
	Cid string `json:"cid"`
}

var User UserData

// Handle security middleware aims to implement a JWT authentication.
func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		url := "https://julestruong.eu.auth0.com/userinfo"

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			log.Printf("No token string")
			http.Error(w, "No token string", http.StatusBadRequest)

			return
		}

		tokenString = tokenString[7:]
		log.Printf("tokenstring found %v", tokenString)

		client := &http.Client{}
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			req.Header.Add("Authorization", via[0].Header.Get("Authorization"))

			return nil
		}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("authorization", "Bearer "+tokenString)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Error found in auth0/userinfo %v", err)
			return
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		err = json.Unmarshal(body, &User)

		if err != nil {
			log.Printf("Error found in json response %v", err)
			return
		}

		log.Printf("User %v", User)
		if checkUserClientIDEqualsClientId() {
			log.Printf("User %v", User)
			next.ServeHTTP(w, r)

			return
		}

		http.Error(w, "", http.StatusBadRequest)
	})
}

/**
*
*
**/
func checkUserClientIDEqualsClientId() bool {
	file, e := ioutil.ReadFile("./config/config.json")

	if e != nil {
		log.Printf("File error! %v", e)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)

	log.Printf("Comparison user ClientID (%s) == Config ClientId (%s)", User.ClientID, config.Cid)

	return true
	// return User.ClientID == config.Cid
}
