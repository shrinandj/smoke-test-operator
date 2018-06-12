package main

import (
	"context"
	"os"
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

func getWatchNamespace() string {
	value := os.Getenv("SMOKE_TEST_NAMESPACE")
	if len(value) == 0 {
		return "default"
	}
	return value
}

func main() {
	printVersion()
	watchNamespace := getWatchNamespace()
	logrus.Infof("Watching namespace %s", watchNamespace)
	sdk.Watch("smoketest.k8s.io/v1alpha1", "SmokeTest", watchNamespace, 5)
	sdk.Handle(run.NewHandler())
	sdk.Run(context.TODO())
}
