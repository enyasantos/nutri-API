package settings

import (
	"os"
	"time"
)

type Server struct {
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Api struct {
	ServerUrl  string
	ServerPort string
}

var ApiSetting = &Api{}

var ServerSetting = &Server{}

func Setup() {
	readTimeout := 60
	writeTimeout := 60
	ServerSetting.ReadTimeout = time.Duration(readTimeout) * time.Second
	ServerSetting.WriteTimeout = time.Duration(writeTimeout) * time.Second

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	url := os.Getenv("APPLICATION_URL")
	if len(url) == 0 {
		url = "http://localhost:" + port
	}

	ApiSetting.ServerUrl = url
	ApiSetting.ServerPort = port
}
