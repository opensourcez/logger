# ZkynetIO logger
<img src="gopher.png" alt="drawing" width="250"/>

# About
This logger was made to wrap a generic golang error with richer informational content, cleaner stack trace and more capabilities.Additionally this module provides error shortcuts for time saving. This module is still in development but I would love any and all help with it's progress. <br><br>Poke me at sveinn@zkynet.io if you want to contribute, or just fork and PR <3

# Usage
Look at the test file to see usage: logger_test.go 

# Installation 
 $ go get github.com/zkynetio/logger

# Settings
### Shortcuts
https://github.com/zkynetio/logger/blob/master/errorCollection.go
## Global logger settings
These can also be edited at runtime.
```golang
// LoggingConfig ...
type LoggingConfig struct {
	// The type of logging config.
	// Available as of this moment:
	// 1. google ( in development )
	// 2. stdout
	// 3.file
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
```
## The Information Construct ( aka logger object )
```golang
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

```

# Initialization
```golang
import L "github.com/zkynetio/logger"

var GlobalLogger *L.Logger

err, GlobalLogger := Init(&L.LoggingConfig{
    DefaultLogTag: "your-default-log-tag",
    WithTrace:     true,
    TraceAsJSON:   false,
    SimpleTrace:   true,
    Debug:         true,
    Type:          "stdout",
})
```
 
# Usage and output
here is an example of how I use this logger in another product<br>
1. I get an opperation ID. This ID is specific for each runtime or requests. 
2. I initialize an error object and a log object. Treating these as seperate things gives me more usability later. 
```golang
opID := helpers.GetOPID()
runtimeLog := L.GenericMessage("No message set..")
runtimeLog.Operation = &L.Operation{ID: opID, Last: false, First: true}
runtimeLog.Labels = make(map[string]string)
runtimeError := L.GenericError(nil)
runtimeError.Operation = &L.Operation{ID: opID, Last: false, First: true}
runtimeError.Labels = make(map[string]string)
```
## Here is an example of this code being used later within a goroutine
```golang
func (nl *NodeList) Connect(runtimeError L.InformationConstruct) {
    defer func() {
        r := recover()
        runtimeError.OriginalError = L.GetRecoverError(r)
        runtimeError.Message = "Crash in goroutine"
        runtimeError.Labels["type"] = "goroutine"
        runtimeError.Labels["nr"] = "8"
        runtimeError.Labels["method"] = "(*NodeList)Connect()"
        GlobalLogger.CRITICAL(runtimeError, "")
    }()
```
## Here is an all-in-one example from a real world cenario

```golang
defer func() {
    r := recover()
    GlobalLogger.CRITICAL(L.InformationConstruct{
        Message:       "Node crashed on ip/port: " + ms.ThisNode.ServerIP + ":" + port,
        OriginalError: L.GetRecoverError(r),
        Labels: map[string]string{"ListenerType":strconv.Itoa(listenerType)}
        Operation: L.Operation{
            ID:       opID.String(),
            Last:     true,
            First:    false,
            Producer: "main.openConnection()",
        },
    }, "startup-crash")
}()
```
## Example output in debug mode.

============ DEBUG ENTRY =======<br>
OperationID: d70e2dfe-b84b-11e9-846f-e86a64785fb9<br>
Message: Node crashed on ip/port: 0.0.0.0:8890<br>
OriginalError: listen tcp 0.0.0.0:8890: bind: address already in use<br>
--------------------------<br>
main.openConnection.func1():258<br>
panic():522<br>
main.listenToUsers():303<br>
main.openConnection():286<br>
==========================<br>
2019/08/06 16:12:28 CRITICAL startup-crash {"Operation":{"ID":"d70e2dfe-b84b-11e9-846f-e86a64785fb9","Producer":"main.openConnection()","Last":true},"Labels":{"ListenerType":"1"},"OriginalError":{"Op":"listen","Net":"tcp","Source":null,"Addr":{"IP":"0.0.0.0","Port":8890,"Zone":""},"Err":{"Syscall":"bind","Err":98}}}

## JSON after prettifying
```json
{  
    "Operation":{  
      "ID":"d70e2dfe-b84b-11e9-846f-e86a64785fb9",
      "Producer":"main.openConnection()",
      "Last":true
   },
   "Labels":{  
      "ListenerType":"1"
   },
   "OriginalError":{  
      "Op":"listen",
      "Net":"tcp",
      "Source":null,
      "Addr":{  
         "IP":"0.0.0.0",
         "Port":8890,
         "Zone":""
      },
      "Err":{  
         "Syscall":"bind",
         "Err":98
      }
   }
}
```
## Example output production mode
```json
2019/08/06 16:16:59 CRITICAL startup-crash {"Operation":{"ID":"7880c15c-b84c-11e9-a90b-e86a64785fb9","Producer":"main.openConnection()","Last":true},"Labels":{"ListenerType":"1"},"Message":"Node crashed on ip/port: 0.0.0.0:8890","OriginalError":{"Op":"listen","Net":"tcp","Source":null,"Addr":{"IP":"0.0.0.0","Port":8890,"Zone":""},"Err":{"Syscall":"bind","Err":98}},"StackTrace":"[\"main.openConnection.func1():258\",\"panic():522\",\"main.listenToUsers():303\",\"main.openConnection():286\"]"}
```
## After prettifying
```json
    {
        "Operation":{
            "ID":"7880c15c-b84c-11e9-a90b-e86a64785fb9",
            "Producer":"main.openConnection()",
            "Last":true
        },
        "Labels":{"ListenerType":"1"},
        "Message":"Node crashed on ip/port: 0.0.0.0:8890",
        "OriginalError":{
            "Op":"listen",
            "Net":"tcp",
            "Source":null,
            "Addr":{"IP":"0.0.0.0","Port":8890,"Zone":""},
            "Err":{"Syscall":"bind","Err":98}
            },
        "StackTrace":"[\"main.openConnection.func1():258\",\"panic():522\",\"main.listenToUsers():303\",\"main.openConnection():286\"]"
    }
```
## You can also get the stack trace in its original format
Here is an example from the test file:
```json
============ DEBUG ENTRY =======
OperationID: 123123jb123b12
Message: A very descriptive message telling others what went wrong
--------------------------
goroutine 7 [running]:
runtime/debug.Stack(0xefc500, 0x0, 0xefd1c0)
	/usr/lib/golang/src/runtime/debug/stack.go:24 +0x9d
github.com/zkynetio/logger.GetStack(0xc0000c67d0, 0xc0001080c0, 0x40e0d8, 0xc0)
	/home/zkynet/go/src/github.com/zkynetio/logger/stack.go:74 +0x1aa
github.com/zkynetio/logger.(*StdClient).log(0xc00023c700, 0xc0001080c0, 0xa1680a, 0x5, 0xa1f620, 0x11)
	/home/zkynet/go/src/github.com/zkynetio/logger/StdLogger.go:16 +0x93
github.com/zkynetio/logger.(*Logger).logit(0xc00000e900, 0xc000258c00, 0xc000258c30, 0xa36f8e, 0x39, 0x0, 0x0, 0x1f4, 0x5d497f30, 0x0, ...)
	/home/zkynet/go/src/github.com/zkynetio/logger/shipping.go:66 +0xde
github.com/zkynetio/logger.(*Logger).ERROR(...)
	/home/zkynet/go/src/github.com/zkynetio/logger/shipping.go:43
github.com/zkynetio/logger.s1()
	/home/zkynet/go/src/github.com/zkynetio/logger/logger_test.go:129 +0x3c2
github.com/zkynetio/logger.s2(...)
	/home/zkynet/go/src/github.com/zkynetio/logger/logger_test.go:96
github.com/zkynetio/logger.s3(...)
	/home/zkynet/go/src/github.com/zkynetio/logger/logger_test.go:93
github.com/zkynetio/logger.TestStdOutShipping(0xc000284100)
	/home/zkynet/go/src/github.com/zkynetio/logger/logger_test.go:89 +0x24
testing.tRunner(0xc000284100, 0xa41680)
	/usr/lib/golang/src/testing/testing.go:865 +0xc0
created by testing.(*T).Run
	/usr/lib/golang/src/testing/testing.go:916 +0x35a

==========================
2019/08/06 16:22:56 ERROR my-custom-log-tag {"Operation":{"ID":"123123jb123b12","Producer":"requestid or namespace or anything really..","First":true},"Labels":{"CUSTOMER":"Google","PATH":"/some/path/to/awesomeness","RANDOM":"whatDoYouNeedHere?","SHORTCODE":"Goo"},"HTTPCode":500,"Timestamp":1565097776}
```

