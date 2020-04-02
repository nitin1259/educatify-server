package models

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/nitin1259/educatify-server/utils"
)

type Post struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Tags     string `json:"tags"`
	Status   uint   `json:"status"` // 1- Pending/In Draft 2- Approved 3- outdated post
	AutherID uint   `json:"autherID"`
}

func (post *Post) Validate() (map[string]interface{}, bool) {
	log.Printf("Posts validate method post: %s", post.ID)

	if !utils.ValidateStringInput(post.Title) {
		return utils.Message(false, "Title should not be empty"), false
	}

	if !utils.ValidateStringInput(post.Content) && len(strings.Split(post.Content, " ")) < 250 {
		return utils.Message(false, "Content should not be less than 250 words"), false
	}

	if !utils.ValidateStringInput(post.Tags) {
		return utils.Message(false, "Atleast one tag is required"), false
	}

	if post.AutherID < 0 {
		return utils.Message(false, "AutherID is required"), false
	}

	return utils.Message(true, "success"), true
}

func (post *Post) Create() map[string]interface{} {
	log.Printf("Posts validate method post: %s, %s, %s", post.Title, post.Tags, post.AutherID)
	if resp, ok := post.Validate(); !ok {
		log.Fatalf("Validation failed")
		return resp
	}

	post.Status = 1 // Pending
	GetDB().Create(post)

	resp := utils.Message(true, "success")
	resp["post"] = post

	return resp
}

func GetPostById(id uint) *Post {

	post := &Post{}

	err := GetDB().Table("posts").Where("id=?", id).First(post).Error

	if err != nil {
		return nil
	}

	return post
}

func GetAllPosts(userid uint) []*Post {

	posts := make([]*Post, 0)

	err := GetDB().Table("posts").Where("auther_id=?", userid).Find(&posts).Error

	if err != nil {
		log.Printf("Error wile getting posts from db: %s", err)
		return nil
	}

	return posts
}
