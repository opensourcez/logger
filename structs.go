package logger

import (
	gclogging "cloud.google.com/go/logging"
)

// Logger ...
// The base logging struct
type Logger struct {
	Config *LoggingConfig
	Client LoggingClient
	Chain  map[string][]InformationConstruct
}

// LoggingClient ...
type LoggingClient interface {
	new(config *LoggingConfig) error
	close()
}

// GoogleClient ...
type GoogleClient struct {
	Loggers map[string]*gclogging.Logger
	Config  *LoggingConfig
	Client  *gclogging.Client
}

// StdClient ...
type StdClient struct {
	Loggers map[string]string
	Config  *LoggingConfig
}

// LoggingConfig ...
type LoggingConfig struct {
	// Enable or disable colors
	Colors bool
	// The type of logging config.
	// 1. stdout
	// 2. .. in development
	Type string
	// The default tag used for your logs.
	DefaultLogTag string
	// The default log level
	DefaultLogLevel string
	// Do you want a stack trace with your log ?
	WithTrace bool
	// Do you want the simplified stack trace or the default one ?
	SimpleTrace bool
	// Enable pretty printing in the console
	PrettyPrint bool
	// Include file names when logging
	FilesInStack bool
	// Only used for Type:google
	ProjectID string

	// This field is only for google cloud logging, which is still in development
	Logs []string
}

// InformationConstruct ...
type InformationConstruct struct {
	// A log level specifically for this log entry
	LogLevel string `json:"Loglevel,omitempty" xml:"LogLevel"`
	// A custom log tag for this specific log entry
	LogTag string `json:"LogTag,omitempty" xml:"LogTag"`
	// The operation represents an execution chain, the ID of
	//the operation can be used to corralate log entries.
	Operation Operation `json:"Operation,omitempty" xml:"Operation"`
	// Key/value labels
	Labels map[string]string `json:"Labels,omitempty" xml:"Labels"`
	// A custom error message
	Message string `json:"Message,omitempty" xml:"Message"`
	// Internal error code
	Code string `json:"Code,omitempty" xml:"Code"`
	// HTTP error code
	HTTPCode int `json:"HTTPCode,omitempty" xml:"HTTPCode"`
	// A custom timestamp
	Timestamp int64 `json:"Timestamp,omitempty" xml:"Timestamp"`
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
	// indicates this is an internal error and only the code + message will be returned to the user
	ReturnToClient bool `json:"-" xml:"-"`
	// The original error that caused the problem.
	OriginalError error `json:"OriginalError,omitempty" xml:"OriginalError"`
	// A hint for developers on how to potentially fix thid problem
	Hint string `json:"Hint,omitempty" xml:"Hint"`
	// The current stack trace in slice format
	StackTrace string `json:"-" xml:"-"`
	// The current stack trace in string format
	SliceStack []string `json:"StackTrace,omitempty" xml:"StackTrace"`
	// If a database query or any kind of search parameters were in play they can be placed here.
	Query string `json:"Query,omitempty" xml:"Query"`
	// The timing of the before mentioned query
	QueryTiming       int64  `json:"QueryTiming,omitempty" xml:"QueryTiming"`
	QueryTimingString string `json:"QueryTimingString,omitempty" xml:"QueryTimingString"`
	// The current session
	Session string `json:"Session,omitempty" xml:"Session"`
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
