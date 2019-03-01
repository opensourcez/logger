package loggernew

import (
	"fmt"
	"log"
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

func logLevel(construct InformationConstruct, logName string, severity logging.Severity) {
	checkLogName(&logName)

	if !shipToCloud(logName) {
		construct.Print(logName, severity)
		return
	}

	labels := construct.Labels
	operation := construct.Operation
	cleanInformationConstruct(&construct)
	sendToGoogleCloud(construct, operation, labels, severity, logName)
}

func cleanInformationConstruct(str *InformationConstruct) {
	var nilStuff Operation
	str.Operation = nilStuff
	str.Labels = nil
}

func checkLogName(logName *string) {
	if *logName == "" {
		*logName = Client.Config.DefaultLogName
	}
}

func sendToGoogleCloud(construct interface{}, op Operation, labels map[string]string, severity logging.Severity, logName string) {
	log.Println(Client.Loggers[logName])
	Client.Loggers[logName].Log(logging.Entry{
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

func (e *InformationConstruct) Print(logName string, severity logging.Severity) {
	infoJSON := e.JSON()
	fmt.Println(severity.String(), logName, infoJSON)
}

func shipToCloud(logName string) bool {
	for i := range Client.Loggers {
		if i == logName {
			return true
		}
	}
	return false
}
