package logger

import (
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
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
			stackSplit[lineNumberIndex] = stackSplit[lineNumberIndex][1:]
			currentLine = strings.Replace(strings.Split(stackSplit[lineNumberIndex], " ")[0], "\t", "", 0)

		}

		if (i % 2) == 1 {
			splitFunc := strings.Split(v, "(")
			if len(splitFunc) <= 1 {
				continue
			}
			if internalLogger.Config.Colors {
				stackTrace = append(stackTrace, color.GreenString(splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"(): ")+currentLine)
			} else {
				stackTrace = append(stackTrace, splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"(): "+currentLine)
			}
			count++
		}
	}

	var finalStack string
	stackTrace = append(stackTrace[:0], stackTrace[0+4:]...)
	finalStack = strings.Join(stackTrace, "\n")

	return finalStack, nil
}

func GetStack(config *LoggingConfig, object *InformationConstruct) (err error) {

	if config.WithTrace {
		if config.SimpleTrace {
			stacktrace, err := GetSimpleStack(false)
			if err != nil {
				return err
			}
			object.StackTrace = stacktrace
			return nil
		}
		stacktrace := string(debug.Stack())
		object.StackTrace = stacktrace
		return nil
	}

	// no trace
	return nil
}
