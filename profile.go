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

		query := `SELECT username,name,email,no_of_threads,no_of_posts FROM users WHERE username=` + `'` + username + `'`
		rows, err := db.Query(query)
		checkerr(err)
		for rows.Next() {
			rows.Scan(&profile.UserName, &profile.Name, &profile.Email, &profile.NoThread, &profile.NoPost)
		}
		fmt.Println(profile)

		tpl.ExecuteTemplate(w, "profile.html", profile)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
