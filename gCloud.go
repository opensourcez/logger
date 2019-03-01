package loggernew

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
)

var defaultLog string
var Client *LoggingClient

// AddLogger ...
// Starts a new logger to a specific log file
func (lc *LoggingClient) AddLogger(logName string) {
	lc.Loggers[logName] = lc.Client.Logger(logName)
}

// InitCloudLogger ...
// This method gives you a new logging client
func (lc *LoggingClient) InitCloudLogger(config *LoggingConfig) (err error) {

	if config.ProjectID == "" {
		return NewError("You need to set a projectID", 0)
	}

	if config.DefaultLogName == "" {
		return NewError("You need to set a default log name", 0)
	}

	lc.Config = config
	log.Println(lc.Config)
	// Init the client
	ctx := context.Background()
	newClient, err := logging.NewClient(ctx, lc.Config.ProjectID)
	if err != nil {
		return
	}
	lc.Client = newClient

	// Add loggers
	lc.Loggers = make(map[string]*logging.Logger)
	defaultLog = lc.Config.DefaultLogName
	for _, v := range lc.Config.Logs {
		log.Println("adding logger")
		lc.AddLogger(v)
	}

	log.Println("initalized a client")
	log.Println(lc)
	// set the global
	// TODO: do we even need this ?
	Client = lc
	return
}
