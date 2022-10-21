package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func create_thread(w http.ResponseWriter, r *http.Request) {
	var thread_ *thread
	if r.Method == http.MethodPost {

		// // fmt.Println("create_thread1")

		topic := r.FormValue("topic")
		content := r.FormValue("content")
		id := uuid.New()
		set_get(w, r)
		split := strings.Split(cookie.Value, "|")
		fmt.Println(cookie.Value, split)

		// fmt.Println("create_thread2")

		thread_ = &thread{
			Id:       id,
			UserName: split[1],
			Topic:    topic,
			Content:  content,
		}
		// fmt.Println("create_thread3")
		create_time := thread_.Created_time()
		query := "INSERT INTO thread (Id, UserName, Topic, Content, CreatedAt) VALUES (?, ?, ?, ?, ?)"
		_, err := db.Exec(query, thread_.Id, thread_.UserName, thread_.Topic, thread_.Content, create_time)
		checkerr(err)
		// fmt.Println("create_thread4")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// fmt.Println("create_thread5")
		return
	}
}

// for main page rendering threads
// getting threads from database.
func show_thread() {
	threads = nil
	query := "SELECT * FROM thread"
	rows, err := db.Query(query)
	checkerr(err)
	defer rows.Close()
	for rows.Next() {
		var thread_ thread
		err := rows.Scan(&thread_.Id, &thread_.UserName, &thread_.Topic, &thread_.Content, &thread_.CreatedAt)
		checkerr(err)
		threads = append(threads, thread_)
	}
}

func read_thread(w http.ResponseWriter, r *http.Request) {

	var post post
	m, err := url.ParseQuery(r.URL.RawQuery)
	checkerr(err)
	var b both

	// getting the thread information
	query := "SELECT * FROM thread WHERE UserName=" + "'" + m["UserName"][0] + "'" + " AND " + "Id=" + "'" + m["Id"][0] + "'"
	rows, err := db.Query(query)
	checkerr(err)
	for rows.Next() {
		err := rows.Scan(&b.Thread.Id, &b.Thread.UserName, &b.Thread.Topic, &b.Thread.Content, &b.Thread.CreatedAt)
		checkerr(err)
	}
	// getting the post information
	query = "SELECT * FROM post WHERE ThreadUserName=" + "'" + m["UserName"][0] + "'" + " AND " + "Id=" + "'" + m["Id"][0] + "'"
	rows, err = db.Query(query)
	checkerr(err)
	for rows.Next() {
		err := rows.Scan(&post.Thread_username, &post.Thread_Id, &post.UserName, &post.Content, &post.Id)
		checkerr(err)
		b.Post = append(b.Post, post)
	}
	defer rows.Close()
	if !Requestcookie(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		rep := r.FormValue("reply")
		id := uuid.New()
		set_get(w, r)
		split := strings.Split(cookie.Value, "|")

		query := "INSERT INTO post (ThreadUserName, Id, PostUserName, Content, PostId) VALUES (?, ?, ?, ?, ?)"
		_, err := db.Exec(query, m["UserName"][0], m["Id"][0], split[1], rep, id)
		checkerr(err)

	}
	if Requestcookie(r) {
		tpl.ExecuteTemplate(w, "post.html", b)
	} else {
		tpl.ExecuteTemplate(w, "not_logged_post.html", b)
	}

}
