package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbbool *pgxpool.Pool

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func validateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func getCommentCountHandler(w http.ResponseWriter, r *http.Request) {
	u, err := parseUrlFromRequest(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Retrieve count for url
	var count int64
	err = dbbool.QueryRow(context.Background(), "select count(comment_id) from comments where comments.url_origin = $1", u).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print(count)

	fmt.Fprintf(w, "%d Comments", count)
}

func getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	u, err := parseUrlFromRequest(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	var comment_body StringArray
	err = dbbool.QueryRow(
		context.Background(),
		"select comment_body from comments where comments.url_origin = $1 order by comment_create_time desc limit 50",
		u).Scan(comment_body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print(comment_body)
}

func updateCommentLikes(w http.ResponseWriter, r *http.Request) {

}

func updateCommentDislikes(w http.ResponseWriter, r *http.Request) {

}

func postCommentHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	dbpool, err := pgxpool.New(context.Background(), "postgres://max:@localhost:5432/max")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	dbbool = dbpool

	http.HandleFunc("/validateUser", validateUserHandler)
	http.HandleFunc("/getCommentCount", getCommentCountHandler)
	http.HandleFunc("/getComments", getCommentsHandler)
	http.HandleFunc("/updateCommentLikes", updateCommentLikes)
	http.HandleFunc("/updateCommentDislikes", updateCommentDislikes)
	http.HandleFunc("/postComment", postCommentHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseUrlFromRequest(r *http.Request) (string, error) {
	// Decode url
	encodedUrl := r.URL.Query().Get("url")
	decodedUrl, err := url.QueryUnescape(encodedUrl)
	if err != nil {
		return "", err
	}

	// Parse url origin
	u, err := url.Parse(decodedUrl)
	if err != nil {
		return "", err
	}
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""

	return u.String(), nil
}
