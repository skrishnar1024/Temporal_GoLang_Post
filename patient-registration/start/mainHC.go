package main

import (
	"context"
	"log"
	"patient-registration/app"
	"strconv"

	"go.temporal.io/sdk/client"
)

func mainHC() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	var data app.PatientDetails

	data.ID = 1
	data.Name = "Alex"
	data.Age = -5
	data.Contact = "123456789"
	data.Address = "Texas"

	options := client.StartWorkflowOptions{
		ID:        "Patient-" + strconv.Itoa(data.ID),
		TaskQueue: app.RegistrationTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.PatientRegistrationWorkflow, data)

	if err != nil {
		log.Fatalln("Unable to complete workflow\n", err)
	}

	var workflowOutput app.PatientDetails
	err = we.Get(context.Background(), &workflowOutput)

	log.Fatalln("workflow workflowOutput: ", workflowOutput)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

}
