package stub

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.intuit.com/sjavadekar/smoke-test-operator/pkg/apis/smoketest/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk/action"
	"github.com/operator-framework/operator-sdk/pkg/sdk/handler"
	"github.com/operator-framework/operator-sdk/pkg/sdk/types"
	"github.com/sirupsen/logrus"
)

func NewHandler() handler.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func fileCopy(sourceFile string, destFile string) {
	from, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) Handle(ctx types.Context, event types.Event) error {
	switch cr := event.Object.(type) {
	case *v1alpha1.SmokeTest:
		// Execute script here

		// 1. Copy the file to /tmp.
		destFile := "/tmp/test.sh"
		fileCopy("/smoke-tests/test.sh", destFile)
		defer os.Remove(destFile)
		logrus.Infof("Successfully copied script to %s", destFile)

		// 2. Execute the copied file.
		op, err := exec.Command("/bin/sh", "-c", destFile).Output()
		if err != nil {
			logrus.Errorf("Failed to execute script: %s", err.Error())
			return err
		}

		cr.Status.TestOutput = string(op)
		err = action.Update(cr)
		if err != nil {
			logrus.Errorf("Failed to update cr: %v", err)
		}
		logrus.Infof("Successfully executed script for smoketest %s", cr.Name)
	}
	return nil
}
