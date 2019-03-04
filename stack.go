package loggernew

import (
	"encoding/json"
	"runtime/debug"
	"strings"
)

func GetSimpleStack(asJSON bool) (string, error) {
	stackSplit := strings.Split(string(debug.Stack()), "\n")
	var filesAndLines []string

	for i, v := range stackSplit {
		if i == 0 {
			continue
		}
		if (i % 2) == 0 {
			fileAndLine := strings.Split(v, ":")
			final := fileAndLine[len(fileAndLine)-1 : len(fileAndLine)][0]
			filesAndLines = append(filesAndLines, final)
		}
	}

	var stackTrace []string
	count := 0
	for i, v := range stackSplit {
		if (i % 2) == 1 {
			var line string
			if count < len(filesAndLines) {
				line = strings.Split(filesAndLines[count], " ")[0]
			} else {
				continue
			}
			splitFunc := strings.Split(v, "(")
			stackTrace = append(stackTrace, splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"():"+line)
			count++
		}
	}

	var finalStack string
	stackTrace = append(stackTrace[:0], stackTrace[0+5:]...)
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

func getStack() (stacktrace string, err error) {

	if logClient.Config.WithTrace {
		if logClient.Config.TraceAsJSON {
			if logClient.Config.SimpleTrace {
				stacktrace, err = GetSimpleStack(true)
				if err != nil {
					return "", err
				}
				return
			}
			stacktrace = string(debug.Stack())
			return

		}
		if logClient.Config.SimpleTrace {
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
