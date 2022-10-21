package main

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	if passTempt == "false" {
		IsPassWrong = "false"

	} else {
		IsPassWrong = "true"

	}

	if Requestcookie(r) {
		http.Redirect(w, r, "/indexexecute", http.StatusSeeOther)
	} else {
		err := tpl.ExecuteTemplate(w, "login.html", IsPassWrong)
		checkerr(err)
	}

}

func indexexecute(w http.ResponseWriter, r *http.Request) {
	show_thread()

	if Requestcookie(r) {
		err = tpl.ExecuteTemplate(w, "logged_index.html", threads)
	} else {
		err = tpl.ExecuteTemplate(w, "not_logged_index.html", threads)
	}

	checkerr(err)
}

func create(w http.ResponseWriter, r *http.Request) {
	if !Requestcookie(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	// fmt.Println("create_thread0")

	err = tpl.ExecuteTemplate(w, "create_thread.html", nil)
	checkerr(err)
	// fmt.Println("create_thread10")
	create_thread(w, r)
	// fmt.Println("create_thread11")
}

func feedback(w http.ResponseWriter, r *http.Request) {
	if !Requestcookie(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	err = tpl.ExecuteTemplate(w, "feedback.html", nil)
	checkerr(err)
}
