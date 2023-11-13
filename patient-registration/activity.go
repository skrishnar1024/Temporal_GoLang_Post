package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.temporal.io/sdk/activity"
)

func CollectPatientInformationActivity(ctx context.Context, data PatientDetails) (PatientDetails, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CollectPatientInformationActivity..started")

	log.Printf("Collecting Patient Information of Name from Activity: %s", data.Name)

	return data, nil
}

func ValidatePatientInformationActivity(ctx context.Context, data PatientDetails) (bool, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ValidatePatientInformationActivity..started")

	//return false, errors.New("validation failed")

	if data.Age <= 0 {
		return false, &InvalidAgeError{}
	}
	return true, nil
}

func CreatePatientRecordActivity(ctx context.Context, data PatientDetails) (bool, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CreatePatientRecordActivity..started")

	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, errors.New("error during marshal.")
	}
	endpointURL := "https://cam001appdev.optioncarehc.net:39443/DataServices/Insert/CPR/integration/test/InsertIntoPatientRegTemp"

	//SendPostRequest
	resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonData))

	fmt.Printf("Body:%s ", jsonData)
	if err != nil {
		return false, errors.New("Error during Post Request")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Success
		return true, nil
	}
	return false, fmt.Errorf("failed to create patient record. Status: %v", resp.Status)
}

// Sample implementation of SendRegistrationConfirmationActivity
func SendRegistrationConfirmationActivity(ctx context.Context, data PatientDetails) (bool, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SendRegistrationConfirmationActivity..started")

	// Simulate sending a confirmation email to the patient
	log.Printf("Registration Confirmation : %s", data.Name)
	return true, nil
}

type InvalidAgeError struct{}

func (m *InvalidAgeError) Error() string {
	return "Age provided is not valid"
}
