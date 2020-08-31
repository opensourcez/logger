package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func cleanInformationConstruct(str *InformationConstruct) {
	str.Operation = Operation{}
	str.Labels = nil
}
func (l *Logger) checklogTag(logTag *string) {
	if *logTag == "" {
		*logTag = l.Config.DefaultLogTag
	}
}

func (e InformationConstruct) Error() string {
	if !e.ReturnToClient {
		return e.Code + ":" + e.Message
	}
	if internalLogger.Config.PrettyPrint {
		outJSON, err := json.Marshal(e)
		if err != nil {
			return JSONEncoding(err).Error()
		}
		return string(outJSON)
	}

	e.OriginalError = nil
	e.Hint = ""
	e.StackTrace = ""
	e.Query = ""
	e.QueryTiming = 0
	e.Labels = nil
	e.Session = ""

	outJSON, err := json.Marshal(e)
	if err != nil {
		return JSONEncoding(err).Error()
	}
	return string(outJSON)

}

func (i InformationConstruct) JSON() string {
	if i.StackTrace != "" && internalLogger.Config.SimpleTrace {
		i.SliceStack = strings.Split(i.StackTrace, "\n")
	}
	outJSON, err := json.Marshal(i)
	if err != nil {
		return JSONEncoding(err).Error()
	}
	return string(outJSON)
}

func GetHTTPCode(err error) int {
	if reflect.TypeOf(err) == reflect.TypeOf(InformationConstruct{}) {
		return err.(InformationConstruct).HTTPCode
	}

	return 0
}

func (e *InformationConstruct) Log() {
	if e.StackTrace == "" {
		err := GetStack(internalLogger.Config, e)
		if err != nil {
			e.StackTrace = "Could not get stacktrace, error:" + err.Error()
		}
	}
	if e.LogTag == "" {
		e.LogTag = internalLogger.Config.DefaultLogTag
	}
	if e.LogLevel == "" {
		e.LogLevel = internalLogger.Config.DefaultLogLevel
	}
	e.Timestamp = time.Now().Unix()
	e.log()
}

func (e *InformationConstruct) log() {

	if internalLogger.Config.PrettyPrint {
		// if we do not have an opperation we add one.
		if e.Operation.ID == "" {
			e.Operation = Operation{ID: uuid.New().String(), Producer: "Debug logger", First: true, Last: true}
		}
		var logString string
		if internalLogger.Config.Colors {
			logString = color.MagentaString("============ DEBUG ENTRY ===================================\n")
		} else {
			logString = "============ DEBUG ENTRY ===================================\n"
		}

		if internalLogger.Config.Colors {
			if e.Operation.First {
				logString = logString + color.RedString("FIRST\n")
			}
			if e.Operation.Last {
				logString = logString + color.RedString("LAST\n")
			}
		} else {
			logString = logString + "FIRST\n"
		}

		if internalLogger.Config.Colors {

			logString = logString + "OPID: " + color.GreenString(e.Operation.ID) + "\n"
			logString = logString + "PRODUCER: " + color.GreenString(e.Operation.Producer) + "\n"
		} else {
			logString = logString + "OPID: " + e.Operation.ID + "\n"
			logString = logString + "OPID: " + e.Operation.Producer + "\n"
		}
		if internalLogger.Config.Colors {
			logString = logString + "MSG: " + color.GreenString(e.Message) + "\n"
		} else {
			logString = logString + "MSG: " + e.Message + "\n"
		}

		if e.QueryTimingString != "" {
			logString = logString + "QueryTiming: " + e.QueryTimingString + "\n"
		}
		if e.QueryTiming != 0 {
			logString = logString + "QueryTiming: " + strconv.FormatInt(e.QueryTiming, 10) + "\n"
		}
		if e.Query != "" {
			logString = logString + "Query: " + e.Query + "\n"
		}

		if e.OriginalError != nil {
			logString = logString + "OriginalError: " + e.OriginalError.Error() + "\n"
		}
		if e.Hint != "" {
			logString = logString + "Hint: " + e.Hint + "\n"
		}
		if e.Labels != nil {
			if internalLogger.Config.Colors {
				logString = logString + color.YellowString("---------- Labels ---------\n")
			} else {
				logString = logString + "---------- Labels ---------\n"
			}
			for i, v := range e.Labels {
				if internalLogger.Config.Colors {
					logString = logString + i + " > " + color.GreenString(v) + "\n"
				} else {
					logString = logString + i + " > " + v + "\n"
				}
			}
		}

		if e.StackTrace != "" {
			if internalLogger.Config.Colors {
				logString = logString + color.YellowString("---------- StackTrace ---------\n")
				logString = logString + e.StackTrace + "\n"
				logString = logString + color.YellowString("---------- JSON Object ----------")
			} else {
				logString = logString + "---------- StackTrace ---------\n"
				logString = logString + e.StackTrace + "\n"
				logString = logString + "---------- JSON Object ----------"
			}
		}

		fmt.Println(logString)
		// Remove fields we have already displayed
		e.StackTrace = ""
		e.Query = ""
		e.Hint = ""
		// e.Message = ""
	}
	if internalLogger.Config.Colors {
		log.Println(color.YellowString(e.LogLevel), color.YellowString(e.LogTag), color.GreenString(e.JSON()))
	} else {
		fmt.Println(e.JSON())
	}
	if internalLogger.Config.PrettyPrint {
		if internalLogger.Config.Colors {
			fmt.Println(color.MagentaString("=========================="))
		} else {
			fmt.Println("==========================")
		}
	}

}
func (e *InformationConstruct) AddToChain() {
	internalLogger.Chain[e.Operation.ID] = append(internalLogger.Chain[e.Operation.ID], *e)
}

func (l *Logger) LogOperationChain(id string) {
	l.Lock()
	defer l.Unlock()
	for _, v := range l.Chain[id] {
		v.Log()
	}
	delete(l.Chain, id)
}
func (l *Logger) DeleteChain(id string) {
	l.Lock()
	defer l.Unlock()
	delete(l.Chain, id)
}
