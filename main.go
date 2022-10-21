package main

import (
	"database/sql"
	"html/template"
	"os"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	port := ":" + os.Getenv("PORT")
	// url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("mysql", "root:Godslayer12@tcp(127.0.0.1:3306)/demondb?parseTime=true")
	checkerr(err)
	defer db.Close()
	styles := http.FileServer(http.Dir("./css"))
	js := http.FileServer(http.Dir("./js"))

	http.HandleFunc("/", indexexecute)
	http.HandleFunc("/create", create)

	http.HandleFunc("/read/", read_thread)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginauth", loginauth)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/verifyemail", verifyemail)
	http.HandleFunc("/forgetpass", forgetpass)
	http.HandleFunc("/passchange", passchange)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/feedback", feedback)

	http.HandleFunc("/favicon.ico", faviconHandler)
	http.Handle("/css/", http.StripPrefix("/css/", styles))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.ListenAndServe(port, nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./css/logo1.jpg")
}
