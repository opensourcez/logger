package logger

import (
	"log"
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
	if object.StackTrace == "" {
		stacktrace, err := GetStack(g.Config)
		if err != nil {
			log.Println(err) // handle this better
		}
		object.StackTrace = stacktrace
	}

	object.print(logTag, severity, g.Config.Debug)
}

func (g *StdClient) close() {
	// no op
}
