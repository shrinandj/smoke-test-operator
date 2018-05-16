package main

import (
	"context"
	"os/exec"
	"runtime"

	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	stub "github.intuit.com/sjavadekar/smoke-test-operator/pkg/stub"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func copyTestsLocally() {
	cmdStr := "cp -fRL /smoke-tests/*.sh /tmp/"
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		logrus.Errorf("Failed to copy tests locally: %s", err.Error())
	}

	cmdStr = "chmod 777 /tmp/*.sh"
	_, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		logrus.Errorf("Failed to update file permissions: %s", err.Error())
	}
}

func main() {
	printVersion()
	copyTestsLocally()
	sdk.Watch("smoketest.k8s.io/v1alpha1", "SmokeTest", "default", 5)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
