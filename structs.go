package logger

import gclogging "cloud.google.com/go/logging"

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
	OriginalError  error             `json:"-" xml:"-"`
	Hint           string            `json:"Hint,omitempty" xml:"Hint"`
	StackTrace     string            `json:"StackTrace,omitempty" xml:"StackTrace"`
	Query          string            `json:"Query,omitempty" xml:"Query"`
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
