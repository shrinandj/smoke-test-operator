package stub

import (
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
	return &Handler{}
}

// Handler object has the Handle() method that executes for every smoketest object.
type Handler struct {
	// Fill me
}

func updateCR(cr *v1alpha1.SmokeTest, testOutput string) {
	cr.Status.TestOutput = testOutput
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

		testToRun := "test.sh"
		if cr.Annotations != nil {
			if val, ok := cr.Annotations["testToRun"]; ok {
				testToRun = val
				logrus.Infof("Found test to run annotation: %s", testToRun)
			}
		}
		testFile := "/smoke-tests/" + testToRun
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			logrus.Infof("Test %s does not exist for %s.", testFile, cr.Name)
			updateCR(cr, "Test "+testFile+" does not exist")
			return nil
		}

		// Execute script here
		destFile := "/tmp/" + testToRun
		op, err := exec.Command("/bin/sh", "-c", destFile).Output()
		if err != nil {
			logrus.Errorf("Failed to execute script: %s", err.Error())
			return err
		}

		updateCR(cr, string(op))
	}
	return nil
}
