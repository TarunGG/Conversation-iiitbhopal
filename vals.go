package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"

	// "github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

var tpl *template.Template
var db *sql.DB
var err error
var otp string
var usersession = make(map[string]string)
var cookie *http.Cookie
var passTempt = "false"
var IsPassWrong string

type thread struct {
	Id        uuid.UUID // id of specific thread
	UserName  string    // username of the user who created the thread
	Topic     string    // topic of the thread
	Content   string    // content of the thread
	CreatedAt time.Time // time when the thread was created
}

var threads []thread

type post struct {
	Thread_username string    // username of the user who created the thread
	Content         string    // content of the post / reply.
	UserName        string    // username of the user who created the post / reply.
	Thread_Id       uuid.UUID // id of the thread to which the post / reply belongs.
	Id              uuid.UUID // id of the post / reply.
}

type both struct {
	Thread thread // thread struct
	Post   []post // slice of post structs
}

// // GoogleClaims -
// type GoogleClaims struct {
// 	Email         string `json:"email"`
// 	EmailVerified bool   `json:"email_verified"`
// 	FirstName     string `json:"given_name"`
// 	LastName      string `json:"family_name"`
// 	jwt.StandardClaims
// }

type user struct {
	UserName string
	Name     string
	Email    string
	Password string
}

var user1 user

type Profile struct {
	UserName string
	Name     string
	Email    string
	NoPost   int32
	NoThread int32
}

type page_thread struct {
	Id        uuid.UUID // id of specific thread
	UserName  string    // username of the user who created the thread
	Topic     string    // topic of the thread
	Content   string    // content of the thread
	CreatedAt string
}

var page_threads []page_thread
