package main

import (
	"app/pkg/service/router"
	"app/settings"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	settings.Setup()
}

func main() {
	//create_database()
	//test()

	routersInit := router.InitRouter()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	port := settings.ApiSetting.ServerPort

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routersInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Default().Printf("[info] Server running on http://localhost:%s", port)
	server.ListenAndServe()

}
