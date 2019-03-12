package logger

import (
	"encoding/json"
	"reflect"
)

func cleanInformationConstruct(str *InformationConstruct) {
	str.Operation = nil
	str.Labels = nil
}

func (l *Logger) checklogTag(logTag *string) {
	if *logTag == "" {
		*logTag = l.Config.DefaultLogTag
	}
}

func (e InformationConstruct) Error() string {
	outJSON, err := json.Marshal(e)
	if err != nil {
		return JSONEncoding(err).Error()
	}
	return string(outJSON)
}

func (i InformationConstruct) JSON() string {
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
	l.logit(construct, logTag, "ERROR")
}
func (l *Logger) EMERGENCY(construct InformationConstruct, logTag string) {
	l.logit(construct, logTag, "EMERGENCY")
}
func (l *Logger) CRITICAL(construct InformationConstruct, logTag string) {
	l.logit(construct, logTag, "CRITICAL")
}
func (l *Logger) ALERT(construct InformationConstruct, logTag string) {
	l.logit(construct, logTag, "ALERT")
}
func (l *Logger) WARNING(construct InformationConstruct, logTag string) {
	l.logit(construct, logTag, "WARNING")
}
func (l *Logger) NOTICE(construct InformationConstruct, logTag string) {
	l.logit(construct, logTag, "NOTICE")
}
func (l *Logger) INFO(construct InformationConstruct, logTag string) {
	l.logit(construct, logTag, "INFO")
}

func (l *Logger) logit(construct InformationConstruct, logTag string, severity string) {
	l.checklogTag(&logTag)
	l.Client.log(&construct, severity, logTag)
}
