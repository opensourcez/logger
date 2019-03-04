package loggernew

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"cloud.google.com/go/logging"
)

var LogClient *Client

func NewObject(message string, HTTPCode int) InformationConstruct {
	return InformationConstruct{
		Message:        message,
		HTTPCode:       HTTPCode,
		Timestamp:      int32(time.Now().Unix()),
		Hint:           "",
		Temporary:      false,
		Retries:        0,
		MaxRetries:     0,
		ReturnToClient: false,
		OriginalError:  nil,
	}
}

// InitCloudLogger ...
// This method gives you a new logging client
func (lc *Client) InitCloudLogger(config *LoggingConfig) (err error) {

	if config.ProjectID == "" {
		return NewObject("You need to set a projectID", 0)
	}

	if config.DefaultLogName == "" {
		return NewObject("You need to set a default log name", 0)
	}

	lc.Config = config
	// Init the client
	ctx := context.Background()
	newClient, err := logging.NewClient(ctx, lc.Config.ProjectID)
	if err != nil {
		return
	}
	lc.Client = newClient

	// Add loggers
	lc.Loggers = make(map[string]Logger)
	for _, v := range lc.Config.Logs {
		lc.AddLogger(v)
	}

	// set the global
	// TODO: do we even need this ?
	LogClient = lc
	return
}

// InitStdOutLogger ...
// This method gives you a new logging client
func (lc *Client) InitStdOutLogger(config *LoggingConfig) (err error) {
	if config.DefaultLogName == "" {
		return NewObject("You need to set a default log name", 0)
	}

	lc.Config = config

	// set the global
	// TODO: do we even need this ?
	LogClient = lc
	return
}

// AddLogger ...
// Starts a new logger to a specific log file
func (lc *Client) AddLogger(logName string) {
	lc.Loggers[logName] = lc.Client.Logger(logName)
}

func (e InformationConstruct) Error() string {
	outJSON, err := json.Marshal(e)
	if err != nil {
		return JSONEncoding(err).Error()
	}
	return string(outJSON)
}

func (i InformationConstruct) JSON() string {
	outJSON, err := json.Marshal(i)
	if err != nil {
		return JSONEncoding(err).Error()
	}
	return string(outJSON)
}

func GetHTTPCode(err error) int {
	if reflect.TypeOf(err) == reflect.TypeOf(InformationConstruct{}) {
		return err.(InformationConstruct).HTTPCode
	}

	return 0
}

func (c *Client) Close() error {
	return c.Client.Close()
}
