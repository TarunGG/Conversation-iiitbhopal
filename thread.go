package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func create_thread(w http.ResponseWriter, r *http.Request) {
	var thread_ *thread
	if r.Method == http.MethodPost {

		topic := r.FormValue("topic")
		content := r.FormValue("content")
		id := uuid.New()
		set_get(w, r)
		split := strings.Split(cookie.Value, "|")
		fmt.Println(cookie.Value, split)

		thread_ = &thread{
			Id:       id,
			UserName: split[1],
			Topic:    topic,
			Content:  content,
		}

		// updating number of threads made by user.
		var No_of_threads int

		query := "SELECT no_of_threads FROM user WHERE username=" + "'" + split[1] + "'"
		rows, err := db.Query(query)
		checkerr(err)
		for rows.Next() {
			err := rows.Scan(&No_of_threads)
			checkerr(err)
		}
		No_of_threads++
		query = "UPDATE user SET no_of_threads=" + strconv.Itoa(No_of_threads) + "WHERE username=" + "'" + split[1] + "'"
		_, err = db.Query(query)
		checkerr(err)

		create_time := thread_.Created_time()
		query = "INSERT INTO thread(id, username, topic, content, created_at) VALUES (?, ?, ?, ?, ?)"
		_, err = db.Exec(query, thread_.Id, thread_.UserName, thread_.Topic, thread_.Content, create_time)
		checkerr(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)

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
	query := "SELECT * FROM thread WHERE username=" + "'" + m["UserName"][0] + "'" + " AND " + "Id=" + "'" + m["Id"][0] + "'"
	rows, err := db.Query(query)
	checkerr(err)
	for rows.Next() {
		err := rows.Scan(&b.Thread.Id, &b.Thread.UserName, &b.Thread.Topic, &b.Thread.Content, &b.Thread.CreatedAt)
		checkerr(err)
	}
	// getting the post information
	query = "SELECT * FROM post WHERE thread_user_name=" + "'" + m["UserName"][0] + "'" + " AND " + "Id=" + "'" + m["Id"][0] + "'"
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

		// updating number of posts made by user.
		var No_of_posts int
		query := "SELECT no_of_posts FROM user WHERE username=" + "'" + split[1] + "'"
		rows, err := db.Query(query)
		checkerr(err)
		for rows.Next() {
			err := rows.Scan(&No_of_posts)
			checkerr(err)
		}
		No_of_posts++
		query = "UPDATE user SET no_of_posts=" + strconv.Itoa(No_of_posts) + "WHERE username=" + "'" + split[1] + "'"
		_, err = db.Query(query)
		checkerr(err)

		// inserting reply into database.
		query = "INSERT INTO post(thread_user_name, thread_id, post_user_name, Content, post_id) VALUES(?, ?, ?, ?, ?)"
		_, err = db.Exec(query, m["UserName"][0], m["Id"][0], split[1], rep, id)
		checkerr(err)

	}
	if Requestcookie(r) {
		tpl.ExecuteTemplate(w, "post.html", b)
	} else {
		tpl.ExecuteTemplate(w, "not_logged_post.html", b)
	}

}
