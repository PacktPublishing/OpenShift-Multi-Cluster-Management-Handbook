package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Existing code from above
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/cloud", returnAllClouds).Methods("GET")
	myRouter.HandleFunc("/cloud/{id}", returnSingleCloud).Methods("GET")
	myRouter.HandleFunc("/cloud", createNewCloud).Methods("POST")
	myRouter.HandleFunc("/cloud/{id}", deleteCloud).Methods("DELETE")
	myRouter.HandleFunc("/cloud/{id}", updateCloud).Methods("PUT")

	srv := &http.Server{
		Handler:      myRouter,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting Server")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func main() {

	// Configure Logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	initCloudsObject()
	handleRequests()
}

func returnSingleCloud(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, cloud := range Clouds {
		if cloud.Id == key {
			json.NewEncoder(w).Encode(cloud)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Cloud{Id: "", Title: "Not Found", Desc: "Not Found", Content: "Not Found"})
}

func createNewCloud(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var cloud Cloud
	json.Unmarshal(reqBody, &cloud)
	Clouds = append(Clouds, cloud)

	json.NewEncoder(w).Encode(cloud)
}

func deleteCloud(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, cloud := range Clouds {
		if cloud.Id == id {
			Clouds = append(Clouds[:index], Clouds[index+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Cloud{Id: "", Title: "Not Found", Desc: "Not Found", Content: "Not Found"})
}

func updateCloud(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, cloud := range Clouds {
		if cloud.Id == id {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var cloud Cloud
			json.Unmarshal(reqBody, &cloud)
			Clouds[index] = cloud

			json.NewEncoder(w).Encode(cloud)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Cloud{Id: "", Title: "Not Found", Desc: "Not Found", Content: "Not Found"})
}

type Cloud struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Clouds []Cloud

func returnAllClouds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllClouds")
	json.NewEncoder(w).Encode(Clouds)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage! - version:alpha")
	fmt.Println("Endpoint Hit: homePage")
}

func initCloudsObject() {
	Clouds = []Cloud{
		Cloud{Id: "1", Title: "AWS", Desc: "Amazon Web Services", Content: "Amazon Web Services, Inc. is a subsidiary of Amazon providing on-demand cloud computing platforms and APIs to individuals, companies, and governments, on a metered pay-as-you-go basis."},
		Cloud{Id: "2", Title: "Azure", Desc: "Microsoft Azure", Content: "Microsoft Azure, often referred to as Azure, is a cloud computing service operated by Microsoft for application management via Microsoft-managed data centers."},
		Cloud{Id: "3", Title: "GCP", Desc: "Google Cloud Platform", Content: "Google Cloud Platform, offered by Google, is a suite of cloud computing services that runs on the same infrastructure that Google uses internally for its end-user products, such as Google Search, Gmail, Google Drive, and YouTube."},
		Cloud{Id: "4", Title: "IBM Cloud", Desc: "IBM Cloud", Content: "IBM cloud computing is a set of cloud computing services for business offered by the information technology company IBM."},
	}
}
