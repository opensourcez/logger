package logger

import (
	"testing"
)

// func CloudShipping(t *testing.T) {
// 	err, logger := Init(&LoggingConfig{
// 		ProjectID:     "heroic-truck-168212",
// 		DefaultLogTag: "general",
// 		// Logs:          []string{"transaction", "error", "activity"},
// 		WithTrace:   true,
// 		SimpleTrace: true,
// 		PrettyPrint: true,
// 		Type:        "google",
// 	})

// 	//log.Println(logger)

// 	if err != nil {
// 		panic(err)
// 	}

// 	op := Operation{
// 		Producer: "GET: /some/path/to/awesomeness",
// 		ID:       "123123jb123b12",
// 		First:    true,
// 		Last:     false,
// 	}
// 	labels := make(map[string]string)
// 	labels["CUSTOMER"] = "Google"
// 	labels["SHORTCODE"] = "Goo"
// 	labels["PATH"] = "/some/path/to/awesomeness"
// 	labels["RANDOM"] = "whatDoYouNeedHere?"

// 	newError := GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong")
// 	newError.Operation = op
// 	newError.Labels = labels

// 	logger.ERROR(*newError, "transaction")

// 	newError2 := BadEmailOrPassword(nil)
// 	op.First = false
// 	op.Last = true
// 	newError2.Operation = op
// 	newError2.Labels = labels

// 	logger.ERROR(*newError2, "transaction")
// 	time.Sleep(time.Second * 5)
// }

func TestOperationChain(t *testing.T) {
	err, logger := Init(&LoggingConfig{
		DefaultLogTag: "general",
		WithTrace:     true,
		SimpleTrace:   true,
		PrettyPrint:   true,
		Colors:        true,
		Type:          "stdout",
	})

	if err != nil {
		panic(err)
	}

	op := Operation{
		Producer: "requestid or namespace or anything really..",
		ID:       "123123jb123b12",
		First:    true,
		Last:     false,
	}

	newError := GenericErrorWithMessage(nil, "Error1")
	newError.Operation = op
	newError.Labels = make(map[string]string)
	newError.Labels["Key"] = "value for error 1"
	newError.LogLevel = "INFO"
	logger.AddToChain(newError.Operation.ID, *newError)

	op.First = false
	newError = GenericErrorWithMessage(nil, "Error2")
	newError.Operation = op
	newError.Labels = make(map[string]string)
	newError.Labels["Key"] = "value for error 2"
	newError.LogLevel = "ERROR"
	logger.AddToChain(newError.Operation.ID, *newError)

	newError = GenericErrorWithMessage(nil, "Error3")
	newError.Operation = op
	newError.Labels = make(map[string]string)
	newError.Labels["Key"] = "value for error 3"
	newError.LogLevel = "EMERGENCY"
	logger.AddToChain(newError.Operation.ID, *newError)

	op.Last = true
	newError = GenericErrorWithMessage(nil, "Error4")
	newError.Operation = op
	newError.Labels = make(map[string]string)
	newError.Labels["Key"] = "value for error 4"
	newError.LogLevel = "EMERGENCY"
	logger.AddToChain(newError.Operation.ID, *newError)

	logger.LogOperationChain(op.ID)
}
func TestStdOutShipping(t *testing.T) {
	var err error
	err, GlobalLogger = Init(&LoggingConfig{
		DefaultLogTag: "general",
		WithTrace:     true,
		SimpleTrace:   true,
		PrettyPrint:   true,
		Colors:        true,
		Type:          "stdout",
	})

	if err != nil {
		panic(err)
	}

	firstFunction()
}
func firstFunction() {
	secondFunction()
}

func secondFunction() {
	logX := GenericMessage("x")
	logX.Labels = make(map[string]string)
	logX.Labels["ID"] = "234234-324234-23423-4234234"

	logX.Operation = Operation{
		Producer: "/api/v1/getSomething",
		ID:       "22342343",
		First:    true,
		Last:     false,
	}

	logX.Log("user-get", "INFO")
	problemFunction()
}

var GlobalLogger *Logger
var errorX *InformationConstruct

func problemFunction() {

	errorX := GenericErrorWithMessage(nil, "Problem Function is missbehaving")
	errorX.Labels = make(map[string]string)
	errorX.Labels["CUSTOMER"] = "Google"
	errorX.Labels["SHORTCODE"] = "Goo"
	errorX.Labels["PATH"] = "/some/path/to/awesomeness"
	errorX.Labels["RANDOM"] = "AnythingYouWantBaby"

	errorX.Operation = Operation{
		Producer: "/api/v1/getSomething",
		ID:       "22342343",
		First:    false,
		Last:     true,
	}

	errorX.Log("user-error", "ERROR")
}
