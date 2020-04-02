package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitin1259/educatify-server/models"
	"github.com/nitin1259/educatify-server/utils"
)

var CreatePost = func (w http.ResponseWriter, r *http.Request)  {
	
	user:= r.Context().Value("user").(uint)  //Grab the id of the user that send the request

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

var GetAllPosts = func (w http.ResponseWriter, r *http.Request)  {

	id := r.Context().Value("user").(uint)
	data := models.GetAllPosts(id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)


}

var GetPost = func (w http.ResponseWriter, r *http.Request)  {

	params := mux.Vars(r);
	id, err:= strconv.Atoi(params["id"])
	if err !=nil{
		utils.Respond(w, utils.Message(false, "Error in your request"))
		return
	}

	post := models.GetPostById(uint(id))
	resp := utils.Message(true, "success")
	resp["post"] = post
	utils.Respond(w, resp)
}