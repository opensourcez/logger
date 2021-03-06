package logger

import (
	"errors"
	"log"
)

func TypeCastRecoverInterface(r interface{}) (err error) {
	if r == nil {
		return errors.New("No recovery detected")
	}
	switch x := r.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		log.Println(r)
		err = errors.New("could not turn panic into error.. see system logs")
	}
	return
}
