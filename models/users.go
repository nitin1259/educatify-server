package models

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/nitin1259/educatify-server/utils"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT claim structure
*/
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// a struct to represent user account
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

//validate incoming user details
func (user *User) Validate() (map[string]interface{}, bool) {

	err := u.ValidateHost(user.Email)
	if smtpErr, ok := err.(u.SmtpError); ok && err != nil {
		msg := fmt.Sprintf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		return u.Message(false, msg), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required and should not be less than 6 words"), false
	}

	// Email must be unique
	temp := &User{}

	// check for error and duplicate mail id
	dbErr := GetDB().Table("users").Where("email=?", user.Email).First(temp).Error
	if dbErr != nil && dbErr != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection err, Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address is already in use"), false
	}

	return u.Message(false, "Requirement Passed"), true

}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

func Login(email, password string) map[string]interface{} {

	user := &User{}

	err := GetDB().Table("users").Where("email=?", email).First(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error, Please try in sometime")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //password doesnot match
		return u.Message(false, "Invalid login credential, Please try again")
	}

	//Worked and logged in

	user.Password = ""

	//Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString // stored the token in response

	resp := u.Message(true, "Successfully logged in")
	resp["user"] = user
	return resp

}

func GetUser(u uint) *User {

	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Email == "" { //User not found!
		return nil
	}

	user.Password = ""
	return user
}
