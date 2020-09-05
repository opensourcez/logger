package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	color "github.com/logrusorgru/aurora"
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
		err := e.Stack()
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
			logString = color.Magenta("============ DEBUG ENTRY ===================================\n").String()
		} else {
			logString = "============ DEBUG ENTRY ===================================\n"
		}

		if internalLogger.Config.Colors {
			if e.Operation.First {
				logString = logString + color.Red("FIRST\n").String()
			}
			if e.Operation.Last {
				logString = logString + color.Red("LAST\n").String()
			}
		} else {
			logString = logString + "FIRST\n"
		}

		if internalLogger.Config.Colors {

			logString = logString + "OPID: " + color.Green(e.Operation.ID).String() + "\n"
			logString = logString + "PRODUCER: " + color.Green(e.Operation.Producer).String() + "\n"
		} else {
			logString = logString + "OPID: " + e.Operation.ID + "\n"
			logString = logString + "OPID: " + e.Operation.Producer + "\n"
		}
		if internalLogger.Config.Colors {
			logString = logString + "MSG: " + color.Green(e.Message).String() + "\n"
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
				logString = logString + color.Yellow("---------- Labels ---------\n").String()
			} else {
				logString = logString + "---------- Labels ---------\n"
			}
			for i, v := range e.Labels {
				if internalLogger.Config.Colors {
					logString = logString + i + " > " + color.Green(v).String() + "\n"
				} else {
					logString = logString + i + " > " + v + "\n"
				}
			}
		}

		if e.StackTrace != "" {
			if internalLogger.Config.Colors {
				logString = logString + color.Yellow("---------- StackTrace ---------\n").String()
				logString = logString + e.StackTrace + "\n"
				logString = logString + color.Yellow("---------- JSON Object ----------").String()
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
		log.Println(color.Yellow(e.LogLevel), color.Yellow(e.LogTag), color.Green(e.JSON()))
	} else {
		fmt.Println(e.JSON())
	}
	if internalLogger.Config.PrettyPrint {
		if internalLogger.Config.Colors {
			fmt.Println(color.Magenta("==========================").String())
		} else {
			fmt.Println("==========================")
		}
	}

}
func (e *InformationConstruct) AddToChain() {
	internalLogger.Chain[e.Operation.ID] = append(internalLogger.Chain[e.Operation.ID], *e)
}

var ChainMutex = sync.Mutex{}

func (l *Logger) LogOperationChain(id string) {
	ChainMutex.Lock()
	defer ChainMutex.Unlock()
	for _, v := range l.Chain[id] {
		v.Log()
	}
	delete(l.Chain, id)
}
func (l *Logger) DeleteChain(id string) {
	ChainMutex.Lock()
	defer ChainMutex.Unlock()
	delete(l.Chain, id)
}
