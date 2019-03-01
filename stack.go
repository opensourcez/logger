package loggernew

import (
	"encoding/json"
	"runtime/debug"
	"strings"
)

func GetSimpleStackAsJSON() (string, error) {
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
			firstLine := splitFunc[0]
			secondLine := strings.Split(splitFunc[1], ")")[1]
			stackTrace = append(stackTrace, firstLine+secondLine+"():"+line)
			count++
		}

	}

	stackTrace = append(stackTrace[:0], stackTrace[0+4:]...)
	jsonSTACK, err := json.Marshal(stackTrace)
	if err != nil {
		return "", err
	}

	return string(jsonSTACK), nil
}

func GetSimpleStack() string {
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
			firstLine := splitFunc[0]
			secondLine := strings.Split(splitFunc[1], ")")[1]
			stackTrace = append(stackTrace, firstLine+secondLine+"():"+line)
			count++
		}

	}

	stackTrace = append(stackTrace[:0], stackTrace[0+4:]...)
	return strings.Join(stackTrace, "\n")
}
