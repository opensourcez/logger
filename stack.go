package logger

import (
	"runtime/debug"
	"strings"

	// "github.com/fatih/color"
	color "github.com/logrusorgru/aurora"
)

func GetSimpleStack(asJSON bool) (string, error) {
	stackSplit := strings.Split(string(debug.Stack()), "\n")
	var stackTrace []string
	count := 0

	// log.Println(stackSplit)
	// var currentLine string
	var fileAndLineSplit []string
	for i, v := range stackSplit {
		if (i % 2) == 0 {
			lineNumberIndex := i + 2
			if lineNumberIndex > len(stackSplit)-1 {
				continue
			}
			stackSplit[lineNumberIndex] = stackSplit[lineNumberIndex][1:]
			// currentLine = )
			fileAndLineSplit = strings.Split(strings.Replace(strings.Split(stackSplit[lineNumberIndex], " ")[0], "\t", "", -1), ":")
		}
		if (i % 2) == 1 {
			splitFunc := strings.Split(v, "(")
			if len(splitFunc) <= 1 {
				continue
			}
			if internalLogger.Config.FilesInStack {
				finalFilePrint := strings.Join(fileAndLineSplit, ":")
				if internalLogger.Config.Colors {
					stackTrace = append(stackTrace, color.Green(splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"(): ").String()+finalFilePrint)
				} else {
					stackTrace = append(stackTrace, splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"(): "+finalFilePrint)
				}
			} else {
				if internalLogger.Config.Colors {
					stackTrace = append(stackTrace, color.Green(splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"(): ").String()+fileAndLineSplit[len(fileAndLineSplit)-1])
				} else {
					stackTrace = append(stackTrace, splitFunc[0]+strings.Split(splitFunc[1], ")")[1]+"(): "+fileAndLineSplit[len(fileAndLineSplit)-1])
				}
			}

			count++
		}
	}

	var finalStack string
	stackTrace = append(stackTrace[:0], stackTrace[0+3:]...)
	finalStack = strings.Join(stackTrace, "\n")

	return finalStack, nil
}

func (object *InformationConstruct) Stack() {

	if internalLogger.Config.WithTrace {
		if internalLogger.Config.SimpleTrace {
			stacktrace, err := GetSimpleStack(false)
			if err != nil {
				return
			}
			object.StackTrace = stacktrace
			return
		}
		stacktrace := string(debug.Stack())
		object.StackTrace = stacktrace
		return
	}

	// no trace
	return
}
