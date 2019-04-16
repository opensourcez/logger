package logger

import (
	"fmt"
	"log"

	gclogging "cloud.google.com/go/logging"
	"github.com/google/uuid"
)

// Logger ...
// The base logging struct
type Logger struct {
	Config *LoggingConfig
	Client LoggingClient
}

// LoggingClient ...
type LoggingClient interface {
	new(config *LoggingConfig) error
	log(object *InformationConstruct, severity string, logTag string)
	close()
}

// GoogleClient ...
type GoogleClient struct {
	Loggers map[string]*gclogging.Logger
	Config  *LoggingConfig
	Client  *gclogging.Client
}

// CrashGuardClient ...
type CrashGuardClient struct {
	Config *LoggingConfig
}

// StdClient ...
type StdClient struct {
	Loggers map[string]string
	Config  *LoggingConfig
}

// InformationConstruct ...
type InformationConstruct struct {
	Operation      *Operation        `json:"Operation,omitempty" xml:"Operation"`
	Labels         map[string]string `json:"Labels,omitempty" xml:"Labels"`
	Message        string            `json:"Message,omitempty" xml:"Message"`
	Code           string            `json:"Code,omitempty" xml:"Code"`
	HTTPCode       int               `json:"HTTPCode,omitempty" xml:"HTTPCode"`
	Timestamp      int32             `json:"Timestamp,omitempty" xml:"Timestamp"`
	Temporary      bool              `json:"Temporary,omitempty" xml:"Temporary"`
	Retries        int               `json:"Retries,omitempty" xml:"Retries"`
	MaxRetries     int               `json:"MaxRetries,omitempty" xml:"MaxRetries"`
	ReturnToClient bool              `json:"-" xml:"-"`
	OriginalError  error             `json:"OriginalError,omitempty" xml:"OriginalError"`
	Hint           string            `json:"Hint,omitempty" xml:"Hint"`
	StackTrace     string            `json:"StackTrace,omitempty" xml:"StackTrace"`
	Query          string            `json:"Query,omitempty" xml:"Query"`
}

func (e *InformationConstruct) print(logTag string, severity string, debug bool) {

	if debug {
		// if we do not have an opperation we add one.
		if e.Operation == nil {
			e.Operation = &Operation{ID: uuid.New().String(), Producer: "Debug logger", First: true, Last: true}
		}
		logString := "============ DEBUG ENTRY =======\nOperationID: " + e.Operation.ID + "\nMessage: " + e.Message + "\n"

		if e.Query != "" {
			logString = logString + "Query: " + e.Query + "\n"
		}

		if e.OriginalError != nil {
			logString = logString + "OriginalError: " + e.OriginalError.Error() + "\n"
		}

		if e.StackTrace != "" {
			logString = logString + "--------------------------\n"
			logString = logString + e.StackTrace + "\n"
		}

		logString = logString + "=========================="

		fmt.Println(logString)
		// remove trace and query from object when debugging
		e.StackTrace = ""
		e.Query = ""
	}

	log.Println(severity, logTag, e.JSON())
}

// LoggingConfig ...
type LoggingConfig struct {
	ProjectID     string
	DefaultLogTag string
	Logs          []string
	WithTrace     bool
	TraceAsJSON   bool
	SimpleTrace   bool
	Debug         bool
	Type          string // google, aws?, stdout, file?
}

// Operation ...
type Operation struct {
	ID       string `json:"ID,omitempty" xml:"ID"`
	Producer string `json:"Producer,omitempty" xml:"Producer"`
	First    bool   `json:"First,omitempty" xml:"First"`
	Last     bool   `json:"Last,omitempty" xml:"Last"`
}
