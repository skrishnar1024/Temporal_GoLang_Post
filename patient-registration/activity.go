package app

import (
	"context"
	"errors"
	"log"
)

func CollectPatientInformationActivity(cts context.Context, data PatientDetails) (PatientDetails, error) {
	log.Printf("Collecting Patient Information of Name from Activity: %s", data.Name)

	return data, nil
}

func ValidatePatientInformationActivity(cts context.Context, data PatientDetails) (bool, error) {

	//return false, nil
	return false, errors.New("validation failed")
}

func CreatePatientRecordActivity(cts context.Context, data PatientDetails) (bool, error) {

	return SendRegistrationConfirmationActivity(cts, data)
}

// Sample implementation of SendRegistrationConfirmationActivity
func SendRegistrationConfirmationActivity(ctx context.Context, data PatientDetails) (bool, error) {
	// Simulate sending a confirmation email to the patient

	return true, nil
}
