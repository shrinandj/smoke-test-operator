package main

import (
	"context"
	"os"
	"os/exec"
	"runtime"

	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	run "github.intuit.com/sjavadekar/smoke-test-operator/pkg/run"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func copyTestsLocally() {
	// There must be a better way to do this!
	cmdStr := "cp -fRL /smoke-tests/*.sh /tmp/"
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		logrus.Panic("Failed to copy tests locally: %s", err.Error())
	}

	cmdStr = "chmod 777 /tmp/*.sh"
	_, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		logrus.Panic("Failed to update file permissions: %s", err.Error())
	}
}

func getWatchNamespace() string {
	value := os.Getenv("SMOKE_TEST_NAMESPACE")
	if len(value) == 0 {
		return "default"
	}
	return value
}

func main() {
	printVersion()
	copyTestsLocally()
	watchNamespace := getWatchNamespace()
	logrus.Infof("Watching namespace %s", watchNamespace)
	sdk.Watch("smoketest.k8s.io/v1alpha1", "SmokeTest", watchNamespace, 5)
	sdk.Handle(run.NewHandler())
	sdk.Run(context.TODO())
}
