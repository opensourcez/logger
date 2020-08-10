package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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

func (l *Logger) ERROR(construct InformationConstruct, logTag string) {
	construct.LogLevel = "ERROR"
	l.logit(&construct, logTag, "ERROR")
}
func (l *Logger) EMERGENCY(construct InformationConstruct, logTag string) {
	construct.LogLevel = "EMERGENCY"
	l.logit(&construct, logTag, "EMERGENCY")
}
func (l *Logger) CRITICAL(construct InformationConstruct, logTag string) {
	construct.LogLevel = "CRITICAL"
	l.logit(&construct, logTag, "CRITICAL")
}
func (l *Logger) ALERT(construct InformationConstruct, logTag string) {
	construct.LogLevel = "ALERT"
	l.logit(&construct, logTag, "ALERT")
}
func (l *Logger) WARNING(construct InformationConstruct, logTag string) {
	construct.LogLevel = "WARNING"
	l.logit(&construct, logTag, "WARNING")
}
func (l *Logger) NOTICE(construct InformationConstruct, logTag string) {
	construct.LogLevel = "NOTICE"
	l.logit(&construct, logTag, "NOTICE")
}
func (l *Logger) INFO(construct InformationConstruct, logTag string) {
	construct.LogLevel = "INFO"
	l.logit(&construct, logTag, "INFO")
}

func (l *Logger) logit(construct *InformationConstruct, logTag string, severity string) {
	l.checklogTag(&logTag)
	l.Client.log(construct, severity, logTag)
}

func (e *InformationConstruct) Log(logTag string, severity string) {
	e.LogLevel = severity
	if e.StackTrace == "" {
		err := GetStack(internalLogger.Config, e)
		if err != nil {
			e.StackTrace = "Could not get stacktrace, error:" + err.Error()
		}
	}
	e.Timestamp = time.Now().Unix()
	e.print(logTag, severity, internalLogger.Config.PrettyPrint)
}

func (e *InformationConstruct) print(logTag string, severity string, pretty bool) {

	if pretty {
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
		e.Message = ""
	}
	if internalLogger.Config.Colors {
		log.Println(color.YellowString(severity), color.YellowString(logTag), color.GreenString(e.JSON()))
	} else {
		fmt.Println(e.JSON())
	}
	if pretty {
		if internalLogger.Config.Colors {
			fmt.Println(color.MagentaString("=========================="))
		} else {
			fmt.Println("==========================")
		}
	}

}

func (l *Logger) AddToChain(id string, logItem InformationConstruct) {
	l.Lock()
	defer l.Unlock()
	_ = GetStack(l.Config, &logItem)
	l.Chain[id] = append(l.Chain[id], logItem)
}
func (l *Logger) LogOperationChain(id string) {
	l.Lock()
	defer l.Unlock()
	for _, v := range l.Chain[id] {
		if v.LogLevel == "" {
			v.LogLevel = "INFO"
		}

		l.logit(&v, v.LogLevel, v.LogTag)
	}
	delete(l.Chain, id)
}
func (l *Logger) DeleteChain(id string) {
	l.Lock()
	defer l.Unlock()
	delete(l.Chain, id)
}
