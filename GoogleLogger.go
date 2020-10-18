package logger

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"cloud.google.com/go/logging"
	gclogging "cloud.google.com/go/logging"
	"github.com/google/uuid"
	option "google.golang.org/api/option"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

func (g *GoogleClient) new(config *LoggingConfig) (err error) {

	// Init the client
	ctx := context.Background()
	newClient, err := logging.NewClient(ctx, config.ProjectID, option.WithCredentialsFile(config.CredentialsPath))
	if err != nil {
		return err
	}

	g.Loggers = make(map[string]*gclogging.Logger)
	g.Loggers[config.DefaultLogTag] = newClient.Logger(config.DefaultLogTag)
	for _, v := range config.Logs {
		g.Loggers[v] = newClient.Logger(v)
	}

	g.Client = newClient
	g.Config = config
	return nil
}

func (g *GoogleClient) Log(object *InformationConstruct) {

	defer func(object *InformationConstruct, severity string, logTag string) {
		if r := recover(); r != nil {
			// if object.Operation.ID != "" {
			// 	log.Println("GOOGLE CLOUD LOGGER FAILED, OP ID:", object.Operation.ID, "\n", r)
			// } else {
			// 	object.Operation = Operation{ID: uuid.New().String()}
			// 	log.Println("GOOGLE CLOUD LOGGER FAILED, OP ID:", object.Operation.ID, "\n", r)
			// }
			// object.log()
			log.Println(string(debug.Stack()), r)
		}
	}(object, object.LogLevel, object.LogTag)
	// set the stack trace
	// deconstruct labels and op from the construct
	labels := object.Labels
	op := object.Operation
	// cleanup
	cleanInformationConstruct(object)

	if object.LogTag == "" {
		object.LogTag = internalLogger.Config.DefaultLogTag
	}
	// ship
	g.Loggers[object.LogTag].Log(logging.Entry{
		InsertID: uuid.New().String(),
		//InsertID:  "sadasdasd",
		Timestamp: time.Now(),
		Labels:    labels,
		Payload:   object,
		Severity:  getSeverity(object.LogLevel),
		Operation: &logpb.LogEntryOperation{
			Id:       op.ID,
			Producer: op.Producer,
			First:    op.First,
			Last:     op.Last,
		}})
}

func (g *GoogleClient) close() {
	g.Client.Close()
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
