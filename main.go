package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	addr     string
	domain   = "colindev.io"
	metaTemp = `<meta name="go-import" content="%s/%s/%s git https://source.developers.google.com/p/%s/r/%s">`
)

func init() {
	flag.StringVar(&addr, "addr", ":8000", "http listen on")
	flag.StringVar(&domain, "domain", "colindev.io", "repo domain")
}

func main() {

	flag.Parse()

	router := httprouter.New()

	router.GET("/:projectID/:repo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		projectID := p.ByName("projectID")
		repo := p.ByName("repo")
		w.Write([]byte(fmt.Sprintf(
			metaTemp,
			domain,
			projectID,
			repo,
			projectID,
			repo,
		)))
	})
	http.ListenAndServe(addr, router)
}
