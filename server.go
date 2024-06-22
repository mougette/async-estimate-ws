package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

/*
func getCommentCountHandler(w http.ResponseWriter, r *http.Request) {
	u, err := parseUrlFromRequest(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Retrieve count for url
	var count int64
	//err = dbbool.QueryRow(context.Background(), "select count(comment_id) from comments where comments.url_origin = $1", u).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print(count)

	fmt.Fprintf(w, "%d Comments", count)
}
*/

func main() {
    /*
	dbpool, err := pgxpool.New(context.Background(), "postgres://max:@localhost:5432/max")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	dbbool = dbpool
    */

    fmt.Fprintf(os.Stderr, "Hello world!")
    http.HandleFunc("/handler", handler)

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
