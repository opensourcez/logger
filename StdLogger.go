package logger

func (g *StdClient) new(config *LoggingConfig) (err error) {

	g.Loggers = make(map[string]string)
	// for _, v := range config.Logs {
	// 	g.Loggers[v] = v
	// }

	g.Config = config
	return nil
}

// func (g *StdClient) log(object *InformationConstruct, severity string, logTag string) {
// 	if object.StackTrace == "" {
// 		err := GetStack(g.Config, object)
// 		if err != nil {
// 			object.StackTrace = "Could not get stacktrace, error:" + err.Error()
// 		}
// 	}
// 	object.print()
// }

func (g *StdClient) close() {
	// no op
}
