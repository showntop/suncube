package main

import (
	"fmt"
	"github.com/showntop/suncube/admin"
	"net/http"
	"os"
	"strings"

	// _ "github.com/showntop/suncube/db"
	_ "github.com/showntop/suncube/db/migrations"
	"github.com/showntop/suncube/routes"
)

func main() {
	mux := http.NewServeMux()
	admin.Admin.MountTo("admin", mux)
	mux.Handle("/", routes.Instrument())
	// mux.Handle("/api/", apiHandler{})
	RootPath := os.Getenv("GOPATH") + "/src/github.com/showntop/suncube"

	publicDir := http.Dir(strings.Join([]string{RootPath, "public"}, "/"))

	mux.Handle("/dist/", http.FileServer(publicDir))
	mux.Handle("/vendors/", http.FileServer(publicDir))
	mux.Handle("/stylesheets/", http.FileServer(publicDir))
	mux.Handle("/javascripts/", http.FileServer(publicDir))
	mux.Handle("/images/", http.FileServer(publicDir))
	mux.Handle("/fonts/", http.FileServer(publicDir))

	mux.HandleFunc("/log", func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintf(w, "Welcome to the log page!")
	})
	http.ListenAndServe(":6001", mux)
}
