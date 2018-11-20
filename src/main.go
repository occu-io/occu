package main

import (
	"core/db"
	"log"
	"net/http"
	"settings"
	"util/router"
)

func main() {
	/* TODO: print logo */

	// read config parameters
	settings.Init()

	// connect to database
	db.VaultConnect()

	// http API router
	router := router.InitRoutes()

	listen := settings.C.Service.Host + ":" + settings.C.Service.Port

	log.Print("Running Occu service on " + listen)
	log.Fatal(http.ListenAndServe(listen, router))
}
