# Logger

TODO: TRACE AND CONFIG

## Installing
1. go get github.com/zkynetio/logger
2. Set up google cloud credentials on your laptop
3. profit

## Disclaimer
This logger will log to stdOut if you want it to. But it's main function is to format and log to Google Cloud Logging.

# Initialization
<b>Example usage of the google cloud logging mechanism: </b>
```
client := logger.Client{}
err := client.InitCloudLogger(&LoggingConfig{
    // Your google cloud project id
    ProjectID:      "heroic-truck-XXXXXX",
    // The default log is used if no log name is given when logging an entry
    DefaultLogName: "general",
    // A list of available logs on the google cloud logging service
    Logs:           []string{"transaction", "error", "activity"},
})

if err != nil {
    panic("never run your app without a logger..")
}
```
<b>Example usage when logging to stdOut:</b>
```
client := Client{}
err := client.InitStdOutLogger(&LoggingConfig{
    // The default log tag is used if no log name is given when logging an entry
    DefaultLogName: "general",
})

if err != nil {
        panic("never run your app without a logger..")
}
```

## Uncought errors and logs
Use a system like rsyslog og Google Fluentd to ship your logs to their final destination.
```
 $ ./path/to/my/app &>> /path/to/log/file
```

# THE THINGS
## errorCollection.go
This file contains a broad collection of pre-set errors that can be used to make ones life easier.<br>
Some Example usage:
```
if err != nil {
    return logger.GenericErrorWithMessage(err, "A very descriptive message telling others what went wrong")
    return logger.BadRequest(err, "You need to check your inputs..")
    return logger.MissingCookie(err)
    return logger.MissingJWTHeader(err)
    return logger.RouteNotFound(err)
    return logger.InvalidUniqueIdentifier(err)
    return logger.BadEmailOrPassword(err)
    return logger.Unauthorized(err)
    return logger.LoginExpired(err)
    return logger.TooManyPasswordResetAttemps(err)
    return logger.DatabaseConnectionErrror(err)
    // etc ... Check the file for more shortcuts
}

```

Some Example usage with printing:
```
if err != nil {
    logger.LogERROR(*logger.GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong"),"errors")
}

```

## The Information Construct
```
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
}
```
## The logging functions

func <b>LogERROR</b>(construct InformationConstruct, logName string)<br>
func <b>LogEMERGENCY</b>(construct InformationConstruct, logName string)<br>
func <b>LogCRITICAL</b>(construct InformationConstruct, logName string)<br>
func <b>LogALERT</b>(construct InformationConstruct, logName string)<br>
func <b>LogWARNING</b>(construct InformationConstruct, logName string)<br>
func <b>LogNOTICE</b>(construct InformationConstruct, logName string)<br>
func <b>LogINFO</b>(construct InformationConstruct, logName string)<br>  




# Basic Usage
For initialization, see examples above
```
logger.LogERROR(*logger.GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong"),"")
```
Example output:
- 2019/03/01 15:32:13 Error general {"Message":"A very descriptive message telling others what went wrong","HTTPCode":500,"Timestamp":1551454333}

<br>
<b>With a log file tag:<br></b>

```
logger.LogERROR(*logger.GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong"),"my-custom-log-tag")
```
Example output:
- 2019/03/01 15:32:13 Error my-custom-log-tag {"Message":"A very descriptive message telling others what went wrong","HTTPCode":500,"Timestamp":1551454333}

<br>
<b>With log file tag and labels:<br></b>

```
newError := logger.GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong")

newError.Labels = make(map[string]string)
newError.Labels["CUSTOMER"] = "My-Customer"
newError.Labels["SERIAL"] = "1232fsdf23422"

logger.LogERROR(*newError, "my-custom-log-tag")
```
Example output:
 - 2019/03/01 16:06:14 Error my-custom-log-tag {"Labels":{"CUSTOMER":"My-Customer","SERIAL":"1232fsdf23422"},"Message":"A very descriptive message telling others what went wrong","HTTPCode":500,"Timestamp":1551456374}

<br>
<b>With log file tag, labels and operation chain:<br></b>

```
// initialize and operation
op := &Operation{
    Producer: "CustomController",
    ID:       "operation-id-324234234",
    First:    true,
    Last:     false,
}

// make some labels
labels := make(map[string]string)
labels["CUSTOMER"] = "My-Customer"
labels["SERIAL"] = "1232fsdf23422"

// send the first error
newError := GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong")
newError.Operation = op
newError.Labels = labels
LogERROR(*newError, "my-custom-log-tag")

// send the last error
newError2 := BadEmailOrPassword(nil)
op.First = false
op.Last = true
newError2.Operation = op
newError2.Labels = labels
LogERROR(*newError2, "my-custom-log-tag")
```
Example output:
- 2019/03/01 16:36:58 Error my-custom-log-tag <b>{"Operation":{"ID":"123123jb123b12","Producer":"GET: /some/path/to/awesomeness","First":true}</b>,"Labels":{"CUSTOMER":"Google","PATH":"/some/path/to/awesomeness","RANDOM":"whatDoYouNeedHere?","SHORTCODE":"Goo"},"Message":"A very descriptive message telling others what went wrong","HTTPCode":500,"Timestamp":1551458218}
- 2019/03/01 16:36:58 Error my-custom-log-tag <b>{"Operation":{"ID":"123123jb123b12","Producer":"GET: /some/path/to/awesomeness","Last":true}</b>,"Labels":{"CUSTOMER":"Google","PATH":"/some/path/to/awesomeness","RANDOM":"whatDoYouNeedHere?","SHORTCODE":"Goo"},"Message":"Email or Password incorrect","HTTPCode":401,"Timestamp":1551458218}


# Notes on how the errors will look in google cloud
I do not want to include some clunky screenshots in here so you get notes instead <3
1. The operation will be extracted into googles Protobuffer Operation struct and displayed as the "operation" in the google Log Entry struct
2. The labels will be extracted to be the top level labels in the google Log Entry struct
3. The log file name will be the log file in your google logging UI/Console.