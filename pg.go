package logger

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	msql "github.com/go-sql-driver/mysql"
	pgx "github.com/lib/pq"
)

func ParsePG(err *pgx.Error) (outError *InformationConstruct) {
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
	case "42P18":
		outError = BadRequest(err, err.Routine)
		outError.Message = err.Message
		outError.Code = "42P18"
		outError.Hint = "Some of your query parameters might be invalid."
	default:
		// this is  away to catch errors that are not supported.
		// so that they can be added.
		PrintObject(err)
	}
	return
}

// ParseSQL ...
func ParseSQL(err *msql.MySQLError) (outError *InformationConstruct) {

	switch err.Number {
	// ambiguous col
	case 1052:
		outError = BadRequest(err, err.Message)
		// outError.Hint = err.Hint
		outError.Message = err.Message
		outError.Code = "1052"
	// NULL supplied to not null
	case 1263:
		outError = BadRequest(err, err.Message)
		// outError.Hint = err.Hint
		outError.Message = err.Message
		outError.Code = "1263"
	// 	Can't write, because of unique constraint, to table '%s'
	case 1169:
		outError = BadRequest(err, err.Message)
		outError.Message = err.Message
		outError.Code = "1169"
	// Can’t connect to local MySQL server through socket ‘/var/run/mysqld/mysqld.sock’ (2)”
	case 2002:
		outError = BadRequest(err, err.Message)
		outError.Message = err.Message
		outError.Code = "2002"
	// Access denied for user '%s'@'%s' to database '%s'
	case 1044:
		outError = BadRequest(err, err.Message)
		outError.Message = err.Message
		outError.Code = "1044"
	// Unknown table '%s'
	case 1051:
		outError = BadRequest(err, err.Message)
		outError.Message = err.Message
		outError.Code = "1051"
	default:
		PrintObject(err)
	}

}

func ParseDBError(er error) (outError *InformationConstruct) {
	if er == nil {
		return nil
	}

	fmt.Println(er)

	switch er.(type) {
	case *pgx.Error:
		outError = ParsePG(er.(*pgx.Error))
	case *msql.MySQLError:
		// todo for emil
		fmt.Println("hello world")
	default:
		// some errors are going to get triggered here...
		newErr := GenericError(er)
		newErr.Message = er.Error()
		newErr.HTTPCode = 404
		return newErr
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
