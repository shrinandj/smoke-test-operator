package run

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.intuit.com/sjavadekar/smoke-test-operator/pkg/apis/smoketest/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk/action"
	"github.com/operator-framework/operator-sdk/pkg/sdk/handler"
	"github.com/operator-framework/operator-sdk/pkg/sdk/types"
	"github.com/sirupsen/logrus"
)

// NewHandler creates and returns a new Handler object.
func NewHandler() handler.Handler {
	return &Handler{
		defaultTest: "test.sh",
	}
}

// Handler object has the Handle() method that executes for every smoketest object.
type Handler struct {
	defaultTest string
}

// TestOutput captures the stdout and stderr for each test.
type TestOutput struct {
	Stdout       string
	Stderr       string
	OutputFormat string
}

func updateCR(cr *v1alpha1.SmokeTest, testOutput TestOutput) {
	output := ""
	if testOutput.OutputFormat == "json" {
		op, err := json.Marshal(testOutput)
		if err != nil {
			logrus.Errorf("Failed to create json of output: %v", err)
		}

		output = string(op)
	} else {
		output = "stdout:\n"
		output = output + testOutput.Stdout
		output = output + "\nstderr:\n"
		output = output + testOutput.Stderr
	}
	cr.Status.TestOutput = string(output)
	err := action.Update(cr)
	if err != nil {
		logrus.Errorf("Failed to update cr: %v", err)
	}
	logrus.Infof("Successfully updated TestOutput for smoketest %s", cr.Name)
}

// Handle is invoked everytime a smoketest custom resource is created/updated.
func (h *Handler) Handle(ctx types.Context, event types.Event) error {
	// Would be good if there was a way to disinguish between "create" and "update".
	// Updates to the custom resource could be ignored. Creates should result
	// in the test getting executed.
	switch cr := event.Object.(type) {
	case *v1alpha1.SmokeTest:
		if cr.Status.TestOutput != "" {
			// SmokeTest has been processed previously.
			return nil
		}

		outputFormat := "text"
		testToRun := h.defaultTest
		if cr.Spec.TestToRun != "" {
			testToRun = cr.Spec.TestToRun
		}

		if cr.Spec.OutputFormat != "" {
			outputFormat = cr.Spec.OutputFormat
		}

		testFile := "/smoke-tests/" + testToRun
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			errMsg := "Test " + testFile + " does not exist"
			testOutput := TestOutput{Stdout: "", Stderr: errMsg, OutputFormat: outputFormat}
			updateCR(cr, testOutput)
			return nil
		}

		// Execute script here
		destFile := "/tmp/" + testToRun
		op, err := exec.Command("/bin/sh", "-c", destFile).Output()
		stdErr := ""
		if err != nil {
			stdErr = err.Error()
		}
		testOutput := TestOutput{Stdout: string(op), Stderr: stdErr, OutputFormat: outputFormat}
		updateCR(cr, testOutput)
	}
	return nil
}
