package stub

import (
	"os/exec"

	"github.intuit.com/sjavadekar/smoke-test-operator/pkg/apis/smoketest/v1alpha1"

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

func (h *Handler) Handle(ctx types.Context, event types.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.SmokeTest:
		// Execute script here
		cmdStr := "kubectl get namespaces"
		cmd := exec.Command("/bin/sh", "-c", cmdStr)
		output, err := cmd.Output()

		if err != nil {
			logrus.Errorf(err.Error())
			return err
		}

		logrus.Infof("Completed command with output %s (%s)", output, o.Name)
	}
	return nil
}
