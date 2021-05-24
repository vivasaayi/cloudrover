package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	ar "github.com/vivasaayi/cloudrover/repositories"

	rice "github.com/GeertJohan/go.rice"
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

func MonitorsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/public/layout.html", "./web/public/monitors.html")
	t.ExecuteTemplate(w, "layout", "")
}

func MonitorsDataHandler(w http.ResponseWriter, r *http.Request) {
	alerts := ar.GetAllAlerts()
	json.NewEncoder(w).Encode(alerts)
}

func InitHttpServer() {
	router := mux.NewRouter()

	box := rice.MustFindBox("./public/")
	fs := http.FileServer(box.HTTPBox())

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	router.HandleFunc("/", HomeHandler)

	router.HandleFunc("/alerts", AlertsHandler)
	router.HandleFunc("/data/alerts", AlertsDataHandler)

	router.HandleFunc("/monitors", MonitorsHandler)
	router.HandleFunc("/data/monitors", MonitorsDataHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",
	}

	log.Fatal(srv.ListenAndServe())
}
