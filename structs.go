package logger

import (
	"fmt"
	"log"
	"os"

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
type FileClient struct {
	BaseLogFolder          string
	fileChannels           map[string]chan []byte
	BaseLogFilePermissions os.FileMode // default is 0660
	Loggers                map[string]string
	Config                 *LoggingConfig
}

// InformationConstruct ...
type InformationConstruct struct {
	// The operation represents an execution chain, the ID of
	//the operation can be used to corralate log entries.
	Operation *Operation `json:"Operation,omitempty" xml:"Operation"`
	// Key/value labels
	Labels map[string]string `json:"Labels,omitempty" xml:"Labels"`
	// A custom error message
	Message string `json:"Message,omitempty" xml:"Message"`
	// Internal error code
	Code string `json:"Code,omitempty" xml:"Code"`
	// HTTP error code
	HTTPCode int `json:"HTTPCode,omitempty" xml:"HTTPCode"`
	// A custom timestamp
	Timestamp int32 `json:"Timestamp,omitempty" xml:"Timestamp"`
	// Indicates if the error is temporary. If a method fails with a temporary error
	// it can most of the time be retired within a certain time frame.
	Temporary bool `json:"Temporary,omitempty" xml:"Temporary"`
	// How many times has this error been retried
	Retries int `json:"Retries,omitempty" xml:"Retries"`
	// The interval of which to retry the method that caused this error.
	// Seconds, Milliseconds, Microseconds, Nanoseconds.. delers choice.
	RetryInterval int `json:"RetryInterval,omitempty" xml:"RetryInterval"`
	// How often should you retry the method that caused this error.
	MaxRetries int `json:"MaxRetries,omitempty" xml:"MaxRetries"`
	// Should this error be returned to the external client. This variable being set to false
	// indicates this is an internal error and that a different error should be returned to the
	// end user.
	ReturnToClient bool `json:"-" xml:"-"`
	// The original error that caused the problem.
	OriginalError error `json:"OriginalError,omitempty" xml:"OriginalError"`
	// A hint for developers on how to potentially fix thid problem
	Hint string `json:"Hint,omitempty" xml:"Hint"`
	// The current stack trace.
	StackTrace string `json:"StackTrace,omitempty" xml:"StackTrace"`
	// If a database query or any kind of search parameters were in play they can be placed here.
	Query string `json:"Query,omitempty" xml:"Query"`
	// The timing of the before mentioned query
	QueryTiming int64 `json:"QueryTiming,omitempty" xml:"QueryTiming"`
	// The current session
	Session string `json:"Session,omitempty" xml:"Session"`
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
		if e.Hint != "" {
			logString = logString + "Hint: " + e.Hint + "\n"
		}

		if e.StackTrace != "" {
			logString = logString + "--------------------------\n"
			logString = logString + e.StackTrace + "\n"
		}

		logString = logString + "=========================="

		fmt.Println(logString)
		// Remove fields we have already displayed
		e.StackTrace = ""
		e.Query = ""
		e.Hint = ""
		e.Message = ""
	}

	log.Println(severity, logTag, e.JSON())
}

// LoggingConfig ...
type LoggingConfig struct {
	// The type of logging config.
	// Available as of this moment:
	// 1. google ( in development )
	// 2. stdout
	Type string
	// The default tag or file used for your log entries.
	// For Type:google, this indicates the default logger used
	// for Type:stdout this is a tag that will be placed on the log as it's printed
	DefaultLogTag string
	// This is the list of available logs
	// for Type:google this indicates files in their log menu
	// for Type:stdout this indicates .. nothing, yet.
	Logs []string
	// Do you want a stack trace with your log ?
	WithTrace bool
	// Do you want your stacktrace as a json object ?
	TraceAsJSON bool
	// Do you want the simplified stack trace or the default one ?
	SimpleTrace bool
	// Are we in debug mode ?
	Debug bool
	// Only used for Type:google
	ProjectID string
}

// Operation ...
type Operation struct {
	ID string `json:"ID,omitempty" xml:"ID"`
	// The method, route, file, etc.. that profuced this error
	Producer string `json:"Producer,omitempty" xml:"Producer"`
	// If this is the first instance of logging for this operation this should be set to true
	First bool `json:"First,omitempty" xml:"First"`
	// If this is the last instance of logging for this opperation this should be set to false.
	Last bool `json:"Last,omitempty" xml:"Last"`
}
