package main

import (
	// "go/constant"

	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

// authenticating login

func loginauth(w http.ResponseWriter, r *http.Request) {

	var uname string
	var pass string
	// parsing form variables

	funame := r.FormValue("rUname")
	fpass := r.FormValue("rp")

	// query for database
	str := "SELECT username,password FROM users WHERE username = " + "'" + funame + "'"

	rows, err := db.Query(str)
	checkerr(err)
	for rows.Next() {
		rows.Scan(&uname, &pass)
		checkerr(err)
	}
	defer rows.Close()
	checkerr(err)

	// encrypting form password to compare with database password
	err1 := bcrypt.CompareHashAndPassword([]byte(pass), []byte(fpass))

	fmt.Println(pass)

	// checking if password is correct or not then redirecting to home page if correct else redirecting to login page
	if err1 == nil && uname == funame {

		// setting cookie if password and username matches
		set_get(w, r)
		cookie.Value = cookie.Value + funame
		http.SetCookie(w, cookie)

		usersession[cookie.Value] = funame

		http.Redirect(w, r, "/indexexecute", http.StatusSeeOther)

	} else {
		passTempt = "true"

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

// signup function executes template as well as validate the user registration.

func signup(w http.ResponseWriter, r *http.Request) {

	type tempt struct {
		Name       string
		Email      string
		Isusername bool
	}
	var temp tempt
	temp.Email = ""
	temp.Name = ""
	temp.Isusername = false
	if r.Method == http.MethodPost {

		// parsing form variables
		user1.UserName = r.FormValue("rUname")
		user1.Password = r.FormValue("rp")
		user1.Name = r.FormValue("rname")
		user1.Email = r.FormValue("rem")
		cp := r.FormValue("rcp")

		query := "SELECT username FROM users WHERE username = " + "'" + user1.UserName + "'"
		rows, err := db.Query(query)
		checkerr(err)
		var uname string
		for rows.Next() {
			rows.Scan(&uname)
			checkerr(err)
			// fmt.Println(uname)
		}

		if uname == user1.UserName {
			temp.Isusername = true
			temp.Name = user1.Name
			temp.Email = user1.Email
		}

		if user1.Password == cp && !temp.Isusername {
			sendEmail(user1.Email)
			http.Redirect(w, r, "/verifyemail", http.StatusSeeOther)
			return
		}

	}
	err = tpl.ExecuteTemplate(w, "signup.html", temp)
	checkerr(err)
}

func verifyemail(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		rotp := r.FormValue("otp")

		if rotp == otp {
			encp, err := bcrypt.GenerateFromPassword([]byte(user1.Password), bcrypt.DefaultCost)
			checkerr(err)
			// fmt.Println("encrypted password before forget pass: ", string(encp))
			query := "INSERT INTO users (username,password,name,email) VALUES(" + "'" + user1.UserName + "'" + "," + "'" + string(encp) + "'" + "," + "'" + user1.Name + "'" + "," + "'" + user1.Email + "'" + ")"

			// _, err = db.Query("INSERT INTO users(username,name,email,password) VALUES('?','?','?','?')", user1.UserName, user1.Name, user1.Email, string(encp))
			_, err = db.Exec(query)
			checkerr(err)
			set_get(w, r)
			cookie.Value = cookie.Value + user1.UserName
			set_get(w, r)
			usersession[cookie.Value] = r.FormValue("rUname")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			http.Redirect(w, r, "/signup", http.StatusSeeOther)

		}
	}
	err = tpl.ExecuteTemplate(w, "otp.html", nil)
	checkerr(err)

}

func forgetpass(w http.ResponseWriter, r *http.Request) {
	Is_user_present := false

	if r.Method == http.MethodPost {
		uname := r.FormValue("uname")
		query := "SELECT username,email,password FROM users WHERE username = " + "'" + uname + "'"

		rows, err := db.Query(query)
		checkerr(err)
		var uname1, email string
		var temp string
		for rows.Next() {
			rows.Scan(&uname1, &email, &temp)
			Is_user_present = true
			checkerr(err)
		}
		defer rows.Close()
		if Is_user_present {

			sendEmail(email)

			url := "/passchange?uname=" + uname
			http.Redirect(w, r, url, http.StatusSeeOther)
			return
		}
	}

	tpl.ExecuteTemplate(w, "forgot.html", Is_user_present)
}

func passchange(w http.ResponseWriter, r *http.Request) {
	m, err := url.ParseQuery(r.URL.RawQuery) // parsing the url.
	uname := m["uname"][0]                   // getting username from url

	checkerr(err)
	if r.Method == http.MethodPost {
		rotp := r.FormValue("otp")
		np := r.FormValue("pass")
		ncp := r.FormValue("cpass")
		fmt.Println(ncp)

		if rotp == otp {

			if np == ncp {
				encp, err := bcrypt.GenerateFromPassword([]byte(np), bcrypt.DefaultCost)
				checkerr(err)

				_, err = db.Query("UPDATE users SET password = ? WHERE username = ?", string(encp), uname)

				checkerr(err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)

				return
			}
		} else {
			http.Redirect(w, r, "/signup", http.StatusSeeOther)

			return
		}
	}
	err = tpl.ExecuteTemplate(w, "createNewPass.html", nil)
	checkerr(err)
}

// logout function destroys the session and redirect to login page.

func logout(w http.ResponseWriter, r *http.Request) {
	set_get(w, r)
	delete(usersession, cookie.Value)
	// deleting cookie.
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
