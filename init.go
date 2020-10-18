package logger

import (
	"log"
)

var internalLogger *Logger

// NewObject ...
func NewObject(message string, HTTPCode int) InformationConstruct {
	return InformationConstruct{
		Message:        message,
		HTTPCode:       HTTPCode,
		Temporary:      false,
		ReturnToClient: false,
	}
}
func NewDBRInterface(logtag string, showTiming, showErrors, showInfo, addToChain bool, opID string) DBREventReceiver {
	return DBREventReceiver{
		LogTag:     logtag,
		ShowInfo:   showInfo,
		ShowErrors: showErrors,
		ShowTiming: showTiming,
		AddToChain: addToChain,
		OPID:       opID,
	}
}

// Init ...
// This method gives you a new logging client
func Init(config *LoggingConfig) (err error, l *Logger) {
	l = &Logger{}
	if config.Type == "google" {
		if config.ProjectID == "" {
			return NewObject("You need to set a projectID", 0), nil
		}
	}

	if config.DefaultLogTag == "" {
		return NewObject("You need to set a default log tag", 0), nil
	}

	l.Config = config

	switch config.Type {
	case "google":
		client := GoogleClient{}
		err = client.new(config)
		l.Client = &client
		if err != nil {
			panic(err)
		}
		break
	case "stdout":
		client := StdClient{}
		err = client.new(config)
		l.Client = &client
	case "aws":
		log.Println("todo ...")
	default:
		client := StdClient{}
		err = client.new(config)
		l.Client = &client
	}
	l.Chain = make(map[string][]InformationConstruct)
	internalLogger = l
	return
}
