package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"flag"

	"github.com/gorilla/mux"
)

type AdmissionRequest struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type AdmissionRequestRequest struct {
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Operation string `json:"operation"`
}

type AdmissionResponse struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Response   AdmissionResponseResponse `json:"response"`
}

type AdmissionResponseResponse struct {
	UID     string `json:"uid"`
	Allowed bool   `json:"allowed"`
}

func main() {
	cert := flag.String("cert", "", "")
	key := flag.String("key", "", "")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", AdmissionService)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServeTLS(*cert, *key))
}

func AdmissionService(w http.ResponseWriter, r *http.Request) {
	admissionReq := AdmissionRequest{}

	err := json.NewDecoder(r.Body).Decode(&admissionReq)
	if err != nil {
		log.Printf("%v", err)
	}

	admissionRes := AdmissionHandler(&admissionReq)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&admissionRes)
}

func AdmissionHandler(r *AdmissionRequest) *AdmissionResponse {
	log.Printf("%+v", r)
	return &AdmissionResponse{
		ApiVersion: "admission.k8s.io/v1",
		Kind:       "",
		Response: AdmissionResponseResponse{
			Allowed: true,
			UID:     "abc",
		},
	}
}
