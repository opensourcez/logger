package loggernew

import (
	"fmt"
	"log"
)

func (g StdClient) new(config *LoggingConfig) (err error) {

	g.Loggers = make(map[string]string)
	for _, v := range config.Logs {
		g.Loggers[v] = v
	}

	g.Config = config
	return nil
}

func (g StdClient) log(object *InformationConstruct, severity string, logTag string) {
	// set the stack trace
	stacktrace, err := getStack()
	if err != nil {
		log.Println(err) // handle this better
	}
	if stacktrace != "" {
		object.StackTrace = stacktrace
	}

	object.print(logTag, severity)
}

func (g StdClient) close() {
	// no op
}

func (e *InformationConstruct) print(logTag string, severity string) {
	if logClient.Config.Debug {
		if e.Operation != nil {
			fmt.Println("========= OPERATION STACK: " + e.Operation.ID + " ==========")
		} else {
			fmt.Println("========= STACK ==========")
		}
		fmt.Println(e.StackTrace)
		fmt.Println("================================")
		e.StackTrace = ""
	}
	log.Println(severity, logTag, e.JSON())
}
