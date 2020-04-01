package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nitin1259/educatify-server/models"
	u "github.com/nitin1259/educatify-server/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Controller method: Create user body: %s", r.Body)
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur

	if err != nil {

		log.Fatalf("error wile decoding json err: %s", err)
		u.Respond(w, u.Message(false, "Invalid Request"))
		return
	}

	log.Printf("in Controller user :%s", user)
	resp := user.Create() //create user
	u.Respond(w, resp)

}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)

}
