package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nitin1259/educatify-server/models"
	"github.com/nitin1259/educatify-server/utils"
)

var CreatePost = func (w http.ResponseWriter, r *http.Request)  {
	
	user:= r.Context().Value("user").(uint) ////Grab the id of the user that send the request

	post:= &models.Post{}

	err:= json.NewDecoder(r.Body).Decode(post)

	if err !=nil{
		utils.Respond(w, utils.Message(false, "Error while decoding request"))
		return
	}

	post.AutherID = user;

	resp:=post.Create();

	utils.Respond(w, resp)

}