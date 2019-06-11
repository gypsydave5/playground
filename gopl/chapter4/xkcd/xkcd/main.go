// xkcd is a tool to query the xkcd backcatalogue.
//
//	xkcd [terms | -http=url]
//
// By default the command takes a series of arguments as terms and returns a tab
// separated information about each comic that matches all the terms. The fields
// are title, url and image url.
//
// If a port/url is supplied to the `-http` flag, a webserver is started on that
// port. The root URL (`/`) will accept multiple query parameters of `term`, passing
// them to the engine in with the same logic as the command line version. The
// results are returned as a JSON array.
//
// In order to avoid spamming the website a  copy of the index is stored locally
// a file called `.xkcd`.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gypsydave5/playground/gopl/chapter4/xkcd"
)

var httpUrl = flag.String("http", "", "URL to open the xkcd query server on")

func main() {
	flag.Parse()
	comics, err := xkcd.Db()
	if err != nil {
		log.Fatal(err)
	}

	if *httpUrl != "" {
		server(*httpUrl, comics)
	}

	terms := os.Args[1:]

	cs := xkcd.Search(comics, terms...)
	for _, c := range cs {
		fmt.Printf("%s\t%s\t%s\n", c.Title, c.Url(), c.Img)
	}
}

func server(url string, comics xkcd.Comics) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		terms := r.URL.Query()["term"]
		cs := xkcd.Search(comics, terms...)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cs)
	})

	http.ListenAndServe(url, nil)
}
