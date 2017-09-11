package security

import (
	"fmt"
	"log"
    "net/http"
    "encoding/json"
)

type User struct {
    Azp string `json:"azp"`
    Aud string `json:"aud"`
    Sub string `json:"sub"`
    AtHash string `json:"at_hash"`
    Iss string `json:"iss"`
    Iat string `json:"iat"`
    Exp string `json:"exp"`
    Alg string `json:"alg"`
    Kid string `json:"kid"`
    // Name string `json:"name"`
    // GivenName string `json:"given_name"`
    // FamilyName string `json:"family_name"`
    // Profile string `json:"profile"`
    // Picture string `json:"picture"`
    // Email string `json:"email"`
    // EmailVerified string `json:"email_verified"`
    // Gender string `json:"gender"`
}

const CLIENTID = "646759497501-g0rqsvkcfpn92ugit4qhadkjpeausq11.apps.googleusercontent.com"

// Handle security middleware aims to implement a JWT authentication.
func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        url := "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=";
        
		tokenString := r.Header.Get("Authorization")[7:]

        fmt.Println("No token string")
        // if tokenString == nil {

        //     return
        // }

        log.Printf("tokenstring %v", tokenString)
        
        url = url + tokenString

        resp, err := http.Get(url)

        defer resp.Body.Close()
        var user User
        err = json.NewDecoder(resp.Body).Decode(&user)

		if err != nil {
            //TODO 
            fmt.Println("%v", err)
            // fmt.Printf("%T\n%s\n%#v\n",err, err, err)
            // switch v := err.(type){
            // case *json.SyntaxError:
            //     fmt.Println(string(resp.Body[v.Offset-40:v.Offset]))
            // }
            
			return
        }
        
        fmt.Println("%v", user.Aud)

        if (user.Aud != CLIENTID) {
            fmt.Println("CLIENT ID WRONG")
            return 
        }
        
        next.ServeHTTP(w, r)
	})
}
