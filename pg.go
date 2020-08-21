package logger

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	pgx "github.com/lib/pq"
)

func ParsePGError(er error) (outError *InformationConstruct) {
	if er == nil {
		return nil
	}
	switch er.(type) {
	case *pgx.Error:
	default:
		// some errors are going to get triggered here...
		newErr := GenericError(er)
		newErr.Message = er.Error()
		newErr.HTTPCode = 404
		return newErr
	}

	err := er.(*pgx.Error)
	switch err.Code {
	case "42702":
		outError = BadRequest(err, err.Detail)
		outError.Hint = err.Hint
		outError.Message = err.Message
		outError.Code = "42702"
	case "23502":
		outError = BadRequest(err, err.Detail)
		outError.Hint = err.Hint
		outError.Message = err.Message
		outError.Code = "23502"
	case "23505":
		outError = BadRequest(err, err.Detail)
		outError.Hint = err.Hint
		outError.Message = err.Message
		outError.Code = "23505"
	case "42703": // Column not found error
		outError = BadRequest(err, err.Routine)
		outError.Hint = err.Hint
		outError.Message = "This column does not appear to exist: " + strings.Split(err.Message, " ")[2]
		outError.Code = "42307"
	case "42601": // bad syntax error
		outError = BadRequest(err, err.Routine)
		outError.Hint = "Your syntax might be off, review all your column and table references."
		outError.Message = err.Message
		outError.Code = "42601"
	case "22P02": // bad syntax error for UUID
		outError = BadRequest(err, err.Routine)
		outError.Hint = "You have an inalid PRIMARY ID in your database transaction"
		outError.Message = err.Message
		outError.Code = "22P02"
	case "42P01": // table not found
		outError = BadRequest(err, err.Routine)
		outError.Hint = "You are trying to interact with a table that does not exist, double check your table names"
		outError.Message = "This table does not appear to exist: " + strings.Split(err.Message, " ")[2]
		outError.Code = "42P01"
	case "42701":
		outError = BadRequest(err, err.Routine)
		outError.Message = err.Message
		outError.Code = "42701"
	default:
		// this is  away to catch errors that are not supported.
		// so that they can be added.
		PrintObject(err)
	}
	return
}
func PrintObject(Object interface{}) {
	fields := reflect.TypeOf(Object).Elem()
	values := reflect.ValueOf(Object).Elem()
	num := fields.NumField()
	parseFields(num, fields, values)
}

func parseFields(num int, fields reflect.Type, values reflect.Value) {
	log.Println("!!!!!!!!!! UN-HANDLED POSTGRES ERROR !!!!!!!!!!")
	for i := 0; i < num; i++ {
		value := values.Field(i)
		field := fields.Field(i)

		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueInt := strconv.FormatInt(value.Int(), 64)
			if valueInt != "" {
				fmt.Println(field.Name, valueInt)
			}
		case reflect.String:
			if value.String() != "" {
				fmt.Println(field.Name, value.String())
			}
		}
	}
	log.Println("!!!!!!!!!! UN-HANDLED POSTGRES ERROR !!!!!!!!!!")
}
