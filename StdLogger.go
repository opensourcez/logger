package logger

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (g *StdClient) new(config *LoggingConfig) (err error) {

	g.Loggers = make(map[string]string)
	for _, v := range config.Logs {
		g.Loggers[v] = v
	}

	g.Config = config
	return nil
}

func (g *StdClient) log(object *InformationConstruct, severity string, logTag string) {
	// set the stack trace
	stacktrace, err := getStack(g.Config)
	if err != nil {
		log.Println(err) // handle this better
	}
	if stacktrace != "" {
		object.StackTrace = stacktrace
	}

	object.print(logTag, severity, g.Config.Debug)
}

func (g *StdClient) close() {
	// no op
}

func (e *InformationConstruct) print(logTag string, severity string, debug bool) {
	if debug {
		if e.Operation == nil {
			e.Operation = &Operation{ID: uuid.New().String()}
		}
		fmt.Println("============ LOG =======\nOperation.ID:" +
			e.Operation.ID +
			"\nMessage:" +
			e.Message + "\n" +
			e.Query +
			"\n--------------------------\n" +
			e.StackTrace +
			"\n==========================")
		e.StackTrace = ""
		e.Query = ""
	}

	log.Println(severity, logTag, e.JSON())
}
