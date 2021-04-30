package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	ar "github.com/vivasaayi/cloudrover/repositories"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/public/layout.html", "./web/public/index.html")
	t.ExecuteTemplate(w, "layout", "")
}

func AlertsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/public/layout.html", "./web/public/alerts.html")
	t.ExecuteTemplate(w, "layout", "")
}

func AlertsDataHandler(w http.ResponseWriter, r *http.Request) {
	alerts := ar.GetAllAlerts()
	json.NewEncoder(w).Encode(alerts)
}

func InitHttpServer() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web/public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	router.HandleFunc("/", HomeHandler)

	router.HandleFunc("/alerts", AlertsHandler)
	router.HandleFunc("/data/alerts", AlertsDataHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}
