package app

import "go.temporal.io/sdk/client"

const RegistrationTaskQueue = "REGISTRATION_TASK_QUEUE"

type PatientDetails struct {
	ID      int
	Name    string
	Age     int
	Contact string
	Address string
}

var TemporalClient client.Client

func SetTemporalClient(temClient *client.Client) {
	TemporalClient = *temClient
}
