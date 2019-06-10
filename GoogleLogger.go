package logger

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/logging"
	gclogging "cloud.google.com/go/logging"
	"github.com/google/uuid"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

func (g *GoogleClient) new(config *LoggingConfig) (err error) {

	// Init the client
	ctx := context.Background()
	newClient, err := logging.NewClient(ctx, config.ProjectID)
	if err != nil {
		return err
	}

	g.Loggers = make(map[string]*gclogging.Logger)
	for _, v := range config.Logs {
		g.Loggers[v] = newClient.Logger(v)
	}

	g.Client = newClient
	g.Config = config
	return nil
}

func (g *GoogleClient) log(object *InformationConstruct, severity string, logTag string) {

	defer func(object *InformationConstruct, severity string, logTag string) {
		if r := recover(); r != nil {
			if object.Operation != nil {
				log.Println("GOOGLE CLOUD LOGGER FAILED, OP ID:", object.Operation.ID, "\n", r)
			} else {
				object.Operation = &Operation{ID: uuid.New().String()}
				log.Println("GOOGLE CLOUD LOGGER FAILED, OP ID:", object.Operation.ID, "\n", r)
			}
			object.print(logTag, severity, g.Config.Debug)
		}
	}(object, severity, logTag)
	// set the stack trace
	stacktrace, err := getStack(g.Config)
	if err != nil {
		// TODO: handle better ??
		log.Println(err) // handle this better
	}
	if stacktrace != "" {
		object.StackTrace = stacktrace
	}

	// deconstruct labels and op from the construct
	labels := object.Labels
	op := object.Operation
	// cleanup
	cleanInformationConstruct(object)

	// ship
	g.Loggers[logTag].Log(logging.Entry{
		InsertID: uuid.New().String(),
		//InsertID:  "sadasdasd",
		Timestamp: time.Now(),
		Labels:    labels,
		Payload:   object,
		Severity:  getSeverity(severity),
		Operation: &logpb.LogEntryOperation{
			Id:       op.ID,
			Producer: op.Producer,
			First:    op.First,
			Last:     op.Last,
		}})
}

func getSeverity(severity string) logging.Severity {
	switch severity {
	case "EMERGENCY":
		return logging.Emergency
	case "ERROR":
		return logging.Error
	case "CRITICAL":
		return logging.Critical
	case "ALERT":
		return logging.Alert
	case "WARNING":
		return logging.Warning
	case "NOTICE":
		return logging.Notice
	case "INFO":
		return logging.Info
	default:
		return logging.Info
	}
}

func (g *GoogleClient) close() {
	g.Client.Close()
}
