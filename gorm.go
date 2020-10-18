package logger

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

type GORMLogger struct {
	Level  string
	LogTag string
}

func (g *GORMLogger) Print(value ...interface{}) {
	if value[0] == "sql" {
		newErr := GenericError(nil)
		newErr.Labels = make(map[string]string)
		data := GetSQLString(newErr.Labels, value...)
		newErr.LogTag = g.LogTag
		newErr.LogLevel = g.Level
		newErr.Query = data[1].(string)
		newErr.QueryTimingString = data[0].(string)
		newErr.Message = data[2].(string)
		newErr.Hint = value[3].(string)
		newErr.Log()
	}

}
func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
func GetSQLString(labels map[string]string, values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql string
			// formattedValues []string
			level = values[0]
		)

		messages = []interface{}{}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf("%.2fms", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for index, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						if t.IsZero() {
							labels[strconv.Itoa(index)] = fmt.Sprintf("'%v'", "0000-00-00 00:00:00")
							// formattedValues = append(formattedValues, fmt.Sprintf("'%v'", "0000-00-00 00:00:00"))
						} else {
							labels[strconv.Itoa(index)] = fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05"))
							// formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
						}
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							labels[strconv.Itoa(index)] = str
							// formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							labels[strconv.Itoa(index)] = "<binary>"
							// formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							labels[strconv.Itoa(index)] = fmt.Sprintf("'%v'", value)
							// formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							labels[strconv.Itoa(index)] = "NULL"
							// formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							// formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
							labels[strconv.Itoa(index)] = fmt.Sprintf("'%v'", value)
						default:
							labels[strconv.Itoa(index)] = fmt.Sprintf("'%v'", value)
							// formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					labels[strconv.Itoa(index)] = "NULL"
					// formattedValues = append(formattedValues, "NULL")
				}
			}

			// // differentiate between $n placeholders or else treat like ?
			// if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
			// 	sql = values[3].(string)
			// 	for index, value := range labels {
			// 		placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
			// 		sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
			// 	}
			// } else {
			// 	formattedValuesLength := len(labels)
			// 	for index, value := range sqlRegexp.Split(values[3].(string), -1) {
			// 		sql += value
			// 		if index < formattedValuesLength {
			// 			sql += formattedValues[index]
			// 		}
			// 	}
			// }

			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf("%v ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned"))
		}
	}

	return
}

func ParseGORM(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			currentTime     = "\033[33m[" + time.Now().Format("2006-01-02 15:04:05") + "]\033[0m"
			source          = fmt.Sprintf("\033[35m(%v)\033[0m", values[1])
		)

		messages = []interface{}{source, currentTime}

		if len(values) == 2 {
			//remove the line break
			currentTime = currentTime[1:]
			//remove the brackets
			source = fmt.Sprintf("\033[35m%v\033[0m", values[1])

			messages = []interface{}{currentTime, source}
		}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf(" \033[36;1m[%.2fms]\033[0m ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						if t.IsZero() {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", "0000-00-00 00:00:00"))
						} else {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
						}
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
						default:
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf(" \n\033[36;31m[%v]\033[0m ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned "))
		} else {
			messages = append(messages, "\033[31;1m")
			messages = append(messages, values[2:]...)
			messages = append(messages, "\033[0m")
		}
	}

	return
}

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)
