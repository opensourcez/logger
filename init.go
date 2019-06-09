package logger

import (
	"time"
)

// NewObject ...
func NewObject(message string, HTTPCode int) InformationConstruct {
	return InformationConstruct{
		Message:        message,
		HTTPCode:       HTTPCode,
		Timestamp:      int32(time.Now().Unix()),
		Temporary:      false,
		ReturnToClient: false,
	}
}

// Init ...
// This method gives you a new logging client
func (l *Logger) Init(config *LoggingConfig) (err error) {

	if config.Type == "google" {
		if config.ProjectID == "" {
			return NewObject("You need to set a projectID", 0)
		}
	}

	if config.DefaultLogTag == "" {
		return NewObject("You need to set a default log tag", 0)
	}

	l.Config = config

	switch config.Type {
	case "google":
		client := GoogleClient{}
		err = client.new(config)
		l.Client = &client
		break
	// case "crashguard":
	// 	client := CrashGuardClient{}
	// 	err = client.new(config)
	// 	l.Client = &client
	// 	break
	case "stdout":
		client := StdClient{}
		err = client.new(config)
		l.Client = &client
	case "aws":
		panic("aws logger has not been implemented yet")
	case "file":
		panic("file logger has not been implemented yet")
	default:
		client := StdClient{}
		err = client.new(config)
		l.Client = &client
	}

	return
}
