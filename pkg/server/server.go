package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/elliot40404/lv/pkg/db"
	"github.com/elliot40404/lv/pkg/utils"
	assets "github.com/elliot40404/lv/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(logChannel <-chan string, exitChannel <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	var db = db.Open()
	go startServer()
	for {
		select {
		case line := <-logChannel:
			db.AddLog(line)
			// if env variable DEBUG is set, print the log to the console
			if os.Getenv("DEBUG") == "true" {
				fmt.Println(line)
			}
		case <-exitChannel:
			log.Println("Exiting server")
			return
		}
	}
}

func startServer() {
	r := chi.NewRouter()
	var db = db.Open()
	if os.Getenv("DEBUG") == "true" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Compress(5, "gzip"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logs := db.GetLogs()
		t, err := template.New("index").Parse(assets.IndexHtml)
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, logs)
	})

	log.Println("Starting server on port :3737")
	utils.OpenBrowser("http://localhost:3737")
	http.ListenAndServe(":3737", r)
}
