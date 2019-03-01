package loggernew

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"cloud.google.com/go/logging"
	"github.com/google/uuid"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

func LogERROR(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Error)
}
func LogEMERGENCY(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Emergency)
}
func LogCRITICAL(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Critical)
}
func LogALERT(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Alert)
}
func LogWARNING(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Warning)
}
func LogNOTICE(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Notice)
}
func LogINFO(construct InformationConstruct, logName string) {
	logLevel(construct, logName, logging.Info)
}

func logLevel(construct InformationConstruct, logName string, severity logging.Severity) {
	checkLogName(&logName)
	var stackTrace string
	var err error

	if LogClient.Config.WithTrace {
		if LogClient.Config.TraceAsJSON {
			if LogClient.Config.SimpleTrace {
				stackTrace, err = GetSimpleStackAsJSON()
				if err != nil {
					panic(err)
				}
			} else {
				stackTrace = string(debug.Stack())
			}

		} else {
			if LogClient.Config.SimpleTrace {
				stackTrace = GetSimpleStack()
			} else {
				stackTrace = string(debug.Stack())
			}
		}
	}
	construct.StackTrace = stackTrace

	if !shipToCloud(logName) {
		construct.print(logName, severity)
		return
	}

	labels := construct.Labels
	operation := construct.Operation
	cleanInformationConstruct(&construct)
	sendToGoogleCloud(construct, *operation, labels, severity, logName)
}

func cleanInformationConstruct(str *InformationConstruct) {
	str.Operation = nil
	str.Labels = nil
}

func checkLogName(logName *string) {
	if *logName == "" {
		*logName = LogClient.Config.DefaultLogName
	}
}

func sendToGoogleCloud(construct interface{}, op Operation, labels map[string]string, severity logging.Severity, logName string) {
	LogClient.Loggers[logName].Log(logging.Entry{
		InsertID: uuid.New().String(),
		//InsertID:  "sadasdasd",
		Timestamp: time.Now(),
		Labels:    labels,
		Payload:   construct,
		Severity:  severity,
		Operation: &logpb.LogEntryOperation{
			Id:       op.ID,
			Producer: op.Producer,
			First:    op.First,
			Last:     op.Last,
		}})
}

func (e *InformationConstruct) print(logName string, severity logging.Severity) {
	if LogClient.Config.Debug {
		fmt.Println("========= DEBUG STACK ==========")
		fmt.Println(e.StackTrace)
		fmt.Println("================================")
		e.StackTrace = ""
	}
	log.Println(severity.String(), logName, e.JSON())
}

func shipToCloud(logName string) bool {
	for i := range LogClient.Loggers {
		if i == logName {
			return true
		}
	}
	return false
}
