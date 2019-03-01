package loggernew

import (
	"encoding/json"
	"reflect"
	"time"
)

func NewInfo(message string, HTTPCode int) InformationConstruct {
	return InformationConstruct{
		BaseConstruct{
			Message:   message,
			HTTPCode:  HTTPCode,
			Code:      "0",
			Timestamp: int32(time.Now().Unix()),
		},
	}
}

func NewError(message string, HTTPCode int) ErrorConstruct {
	return ErrorConstruct{
		BaseConstruct: BaseConstruct{
			Message:   message,
			HTTPCode:  HTTPCode,
			Code:      "0",
			Timestamp: int32(time.Now().Unix()),
		},
		Hint:           "",
		Temporary:      false,
		Retries:        0,
		MaxRetries:     0,
		ReturnToClient: false,
		OriginalError:  nil,
	}
}

func (e ErrorConstruct) Error() string {
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
	if reflect.TypeOf(err) == reflect.TypeOf(ErrorConstruct{}) {
		return err.(ErrorConstruct).HTTPCode
	}

	return 0
}
