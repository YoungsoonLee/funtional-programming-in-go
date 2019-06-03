package utils

import (
	"fmt"
	"path"
	"runtime"
)

func HandlePanic(err error) {
	if err != nil {
		_, filePath, lineNo, _ := runtime.Caller(1)
		_, fileName := path.Split(filePath)
		msg := fmt.Sprintf("[file:%s line:%d]: %s", fileName, lineNo, err.Error())
		panic(msg)

	}
}
