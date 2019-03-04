package loggernew

import (
	"encoding/json"
	"reflect"
)

func cleanInformationConstruct(str *InformationConstruct) {
	str.Operation = nil
	str.Labels = nil
}

func checklogTag(logTag *string) {
	if *logTag == "" {
		*logTag = logClient.Config.DefaultLogTag
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

func LogERROR(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "ERROR")
}
func LogEMERGENCY(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "EMERGENCY")
}
func LogCRITICAL(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "CRITICAL")
}
func LogALERT(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "ALERT")
}
func LogWARNING(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "WARNING")
}
func LogNOTICE(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "NOTICE")
}
func LogINFO(construct InformationConstruct, logTag string) {
	logit(construct, logTag, "INFO")
}

func logit(construct InformationConstruct, logTag string, severity string) {
	checklogTag(&logTag)
	//log.Println(logClient.Client)
	//panic("meow")
	logClient.Client.log(&construct, severity, logTag)
}
