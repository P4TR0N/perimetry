package services

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"perimetry/v1/services/eye_data"
	"perimetry/v1/services/fushion"
)

func HttpServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/login", fushion.FushionLogin).Methods("POST")
	r.HandleFunc("/api/v1/register", fushion.FushionRegistration).Methods("POST")
	r.HandleFunc("/api/v1/send_data", eye_data.ProcessingData).Methods("POST")
	r.HandleFunc("/api/v1/get_data", eye_data.GetData).Methods("POST")
	r.HandleFunc("/api/v1/appkey", AppKey).Methods("GET")
	r.HandleFunc("/api/status", SystemStatus).Methods("GET")
	r.HandleFunc("/api/get_version", PerimetryAppVersion).Methods("GET")
	log.Fatal(http.ListenAndServe(":2007", r))
}
