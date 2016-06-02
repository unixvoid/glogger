package glogger

import (
	"io"
	"log"
)

var (
	// initialize logger types
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

func LogInit(infoHandle io.Writer, debugHandle io.Writer, errorHandle io.Writer) {
	// read loglevel
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
