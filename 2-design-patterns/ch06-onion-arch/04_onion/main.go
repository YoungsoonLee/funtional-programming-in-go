package main

import (
	"io/ioutil"
	"os"

	u "github.com/YoungsoonLee/funtional-programming-in-go/2-design-patterns/ch06-onion-arch/04_onion/src/utils"
	"github.com/pkg/errors"
)

func init() {
	GetOptions()
	if Config.LogDebugInfo {
		InitLog("trace-debug-log.txt", os.Stdout, os.Stdout, os.Stderr)
	} else {
		InitLog("trace-log.txt", ioutil.Discard, os.Stdout, os.Stderr)
	}

	fileName := os.Getenv("TEST_FILENAME")
	if len(fileName) == 0 {
		fileName = defaultFileName
	}

	u.HandlePanic(os.Chdir(Config.ProjectRoot))
}

func main() {
	gcpi, err := infrastructure.GetGcpInteractor()
	u.HandlePanic(errors.Wrap(err, "unable to get gcp interactor"))
	li, err := infrastructure.GetLocalInteractor()
	u.HandlePanic(errors.Wrap(err, "unable to get local interactor"))

	wsh = WebserviceHandler{}
	wsh.GcpInteractor = gcpi
	wsh.LocalInteractor = li
}
