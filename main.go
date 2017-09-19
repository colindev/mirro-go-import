package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.HandleFunc("/{projectID}/{repo:.*}", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.String())
		vars := mux.Vars(r)
		projectID := vars["projectID"]
		repo := vars["repo"]

		log.Println(projectID, repo)
		w.Write([]byte(fmt.Sprintf(
			metaTemp,
			domain,
			projectID,
			repo,
			projectID,
			repo,
		)))
	}).Methods("GET")

	log.Println("listen on:", addr)
	log.Println("mirro:", domain)
	http.ListenAndServe(addr, router)
}
