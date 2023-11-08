package main

import (
	"log"
	"patient-registration/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	// Create client object only once per process
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}

	defer c.Close()

	w := worker.New(c, app.RegistrationTaskQueue, worker.Options{})

	w.RegisterWorkflow(app.PatientRegistrationWorkflow)
	w.RegisterActivity(app.CollectPatientInformationActivity)
	w.RegisterActivity(app.ValidatePatientInformationActivity)
	w.RegisterActivity(app.CreatePatientRecordActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to Start Worker", err)
	}

}
