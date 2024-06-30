package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GCSObject represents the payload received from a GCS event
type GCSObject struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// handleGCSNotification processes the GCS event
func handleGCSNotification(w http.ResponseWriter, r *http.Request) {
	var obj GCSObject
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Bad Request: Unable to parse JSON", http.StatusBadRequest)
		return
	}
	log.Printf("File finalized: %s in bucket %s\n", obj.Name, obj.Bucket)
	logrus.Infof("Received file: %s in bucket %s", obj.Name, obj.Bucket)
	fmt.Fprintf(w, "Received: %s in bucket %s", obj.Name, obj.Bucket)
}

func main() {
	http.HandleFunc("/events", handleGCSNotification)
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
