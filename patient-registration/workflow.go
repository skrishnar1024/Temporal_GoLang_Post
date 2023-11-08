package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func PatientRegistrationWorkflow(ctx workflow.Context, input PatientDetails) (PatientDetails, error) {

	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        10 * time.Second,
		MaximumAttempts:        10, // unlimited Attempts
		NonRetryableErrorTypes: []string{"Error"},
	}

	options := workflow.ActivityOptions{

		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var patientDetailsOutput PatientDetails

	fmt.Printf("Input in Workflow :%s \n", input.Name)
	patientDetailsErr := workflow.ExecuteActivity(ctx, CollectPatientInformationActivity, input).Get(ctx, &patientDetailsOutput)
	fmt.Printf("patientDetailsOutput in Workflow after Execute Activity :%s \n", patientDetailsOutput.Name)

	if patientDetailsErr != nil {
		return PatientDetails{}, patientDetailsErr
	}

	var validationOutput bool

	validationErr := workflow.ExecuteActivity(ctx, ValidatePatientInformationActivity, input).Get(ctx, &validationOutput)

	if validationErr != nil {
		return PatientDetails{}, validationErr
	}
	// if !validationOutput {
	// 	return PatientDetails{}, errors.New("validation failed")
	// }

	var CreatePatientRecordOutput bool

	CreatePatientErr := workflow.ExecuteActivity(ctx, CreatePatientRecordActivity, input).Get(ctx, &CreatePatientRecordOutput)

	if CreatePatientErr != nil {
		return PatientDetails{}, CreatePatientErr
	}

	fmt.Println("Patient workflow complete")

	return patientDetailsOutput, nil

}
