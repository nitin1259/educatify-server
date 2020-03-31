package app

import (
	"context"
	u "educatify-server/utils"
	"educative-server/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		noAuth:= []string{"/api/user/new", "api/user/login"} //List of endpoints that doesn't require auth
		requestpath := r.URL.Path //current path


		//check if request does not need authentication, serve the request if it doesn't need it
		for _, path := range noAuth{

			if path == requestpath{
				next.ServeHTTP(w,r)
				return
			}
		}

		response:= make(map[string]interface{})

		tokenHeader := r.Header.Get("Authorization")// grab the token from header...

		if tokenHeader == ""{ // if token is missing, return with error code 403 Unauthorized.
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted:= strings.Split(tokenHeader, "") ////The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk:= &models.token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func (token *jwt.Token)(interface{}, err)  {
			return []byte(os.Getenv("token_password")), nil
		})

		if err !=nil{//Malformed token, returns with http code 403 as usual

			response = u.Message(false, "Malformed Authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid{ //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User: %", tk.Username) //useful for monitoring

		ctx:= context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		
		next.ServeHTTP(w, r) // proceed in the middleware chain
	}

	)
}
