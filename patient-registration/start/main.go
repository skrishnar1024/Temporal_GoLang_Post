package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"patient-registration/app"
	"strconv"

	"github.com/gorilla/mux" // You need to install the "gorilla/mux" package
	"go.temporal.io/sdk/client"
)

func main() {
	r := mux.NewRouter()

	// Initialize Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	r.HandleFunc("/register-patient", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var data app.PatientDetails
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Start the Temporal workflow
		options := client.StartWorkflowOptions{
			ID:        "Patient-" + strconv.Itoa(data.ID),
			TaskQueue: app.RegistrationTaskQueue,
		}
		we, err := c.ExecuteWorkflow(context.Background(), options, app.PatientRegistrationWorkflow, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Wait for the workflow to complete
		var workflowOutput app.PatientDetails
		err = we.Get(context.Background(), &workflowOutput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(workflowOutput)
	})

	http.Handle("/", r)

	log.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}
