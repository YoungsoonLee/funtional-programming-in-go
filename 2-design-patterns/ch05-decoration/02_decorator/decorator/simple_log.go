package decorator

import (
	"io"
	"log"
	"os"
)

var (
	Debug       *log.Logger
	Info        *log.Logger
	Error       *log.Logger
	InfoHandler io.Writer
)

func InitLog(
	traceFileName string,
	debugHandler io.Writer,
	infohandler io.Writer,
	errorHandler io.Writer) {
	if len(traceFileName) > 0 {
		_ = os.Remove(traceFileName)
		file, err := os.OpenFile(traceFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("failed to create log file: %s", traceFileName)
		}
		debugHandler = io.MultiWriter(file, debugHandler)
		infohandler = io.MultiWriter(file, infohandler)
		errorHandler = io.MultiWriter(file, errorHandler)
	}

	InfoHandler = infohandler
	Debug = log.New(debugHandler, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infohandler, "INFO: ", log.Ltime)
	Error = log.New(errorHandler, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
