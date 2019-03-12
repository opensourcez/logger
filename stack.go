package logger

import (
	"encoding/json"
	"runtime/debug"
	"strings"
)

func GetSimpleStack(asJSON bool) (string, error) {
	stackSplit := strings.Split(string(debug.Stack()), "\n")
	var stackTrace []string
	count := 0

	var currentLine string
	for i, v := range stackSplit {
		if (i % 2) == 0 {
			lineNumberIndex := i + 2
			if lineNumberIndex > len(stackSplit)-1 {
				continue
			}
			currentLine = strings.Split(strings.Split(stackSplit[lineNumberIndex], ":")[1], " ")[0]
		}

		if (i % 2) == 1 {
			splitFunc := strings.Split(v, "(")
			if len(splitFunc) <= 1 {
				continue
			}
			stackTrace = append(stackTrace, splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"():"+currentLine)
			count++
		}
	}

	var finalStack string
	stackTrace = append(stackTrace[:0], stackTrace[0+6:]...)
	if asJSON {
		jsonSTACK, err := json.Marshal(stackTrace)
		if err != nil {
			return "", err
		}
		finalStack = string(jsonSTACK)
	} else {
		finalStack = strings.Join(stackTrace, "\n")
	}

	return finalStack, nil
}

func getStack(config *LoggingConfig) (stacktrace string, err error) {

	if config.WithTrace {
		if config.TraceAsJSON {
			if config.SimpleTrace {
				stacktrace, err = GetSimpleStack(true)
				if err != nil {
					return "", err
				}
				return
			}
			stacktrace = string(debug.Stack())
			return

		}
		if config.SimpleTrace {
			stacktrace, err = GetSimpleStack(false)
			if err != nil {
				return "", err
			}
			return
		}
		stacktrace = string(debug.Stack())
		return
	}

	// no trace
	return "", nil
}
