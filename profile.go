package main

import (
	"fmt"
	"net/http"
	"strings"
)

func profile(w http.ResponseWriter, r *http.Request) {
	var profile Profile

	if Requestcookie(r) {
		set_get(w, r)
		split := strings.Split(cookie.Value, "|")
		username := split[1]

		query := `SELECT UserName, Name, Email FROM demondb.users WHERE UserName=` + `"` + username + `"`
		rows, err := db.Query(query)
		checkerr(err)
		for rows.Next() {
			rows.Scan(&profile.UserName, &profile.Name, &profile.Email)
		}
		fmt.Println(profile.UserName, profile.Email, profile.Name)
		query = `SELECT * FROM demondb.thread WHERE UserName=` + `"` + username + `"`

		rows, err = db.Query(query)
		checkerr(err)
		for rows.Next() {
			profile.NoThread++
		}
		query = `SELECT * FROM demondb.post WHERE PostUserName=` + `"` + username + `"`

		rows, err = db.Query(query)
		checkerr(err)
		for rows.Next() {
			profile.NoPost++
		}

		tpl.ExecuteTemplate(w, "profile.html", profile)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
