package loggernew

import (
	"cloud.google.com/go/logging"
)

type BaseConstruct struct {
	Operation Operation         `json:"Operation,omitempty" xml:"Operation"`
	Labels    map[string]string `json:"Labels,omitempty" xml:"Labels"`
	Message   string            `json:"Message,omitempty" xml:"Message"`
	Code      string            `json:"Code,omitempty" xml:"Code"`
	HTTPCode  int               `json:"HTTPCode,omitempty" xml:"HTTPCode"`
	Timestamp int32             `json:"Timestamp,omitempty" xml:"Timestamp"`
}

type InformationConstruct struct {
	BaseConstruct
	Temporary      bool   `json:"Temporary,omitempty" xml:"Temporary"`
	Retries        int    `json:"Retries,omitempty" xml:"Retries"`
	MaxRetries     int    `json:"MaxRetries,omitempty" xml:"MaxRetries"`
	ReturnToClient bool   `json:"-" xml:"-"`
	OriginalError  error  `json:"-" xml:"-"`
	Hint           string `json:"Hint,omitempty" xml:"Hint"`
}

// LoggingClient ...
// A logging client for Google Cloud
type LoggingClient struct {
	Client  *logging.Client
	Loggers map[string]*logging.Logger
	Config  *LoggingConfig
}

type LoggingConfig struct {
	Costumer       string
	ProjectID      string
	DefaultLogName string
	Logs           []string
	// TODO: what else do we want ?
}

type Operation struct {
	ID       string `json:"ID,omitempty" xml:"ID"`
	Producer string `json:"Producer,omitempty" xml:"Producer"`
	First    bool   `json:"First,omitempty" xml:"First"`
	Last     bool   `json:"Last,omitempty" xml:"Last"`
}
