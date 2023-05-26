package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hospital-management/doctor"
	"github.com/hospital-management/patient"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/doctor", doctor.CreateDoctor).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/doctor/{name}", doctor.DeleteDoc).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/patient", patient.CreatePatient).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/patient/{name}", patient.DeletePatient).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/doctors", doctor.ListDoctors).Methods(http.MethodGet)
	
	log.Println(http.ListenAndServe(":8080", router))
}
