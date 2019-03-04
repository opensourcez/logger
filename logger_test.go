package loggernew

import (
	"fmt"
	"testing"
	"time"
)

func ExamplePrint() {
	fmt.Println("mowmoemw")
	// output: mowmoemw
}

func TestNewError(t *testing.T) {
	error := NewObject("some message", 400)
	if error.HTTPCode != 400 {
		t.Error("did not get the right code")
	}
	if error.Message != "some message" {
		t.Error("did not get the right message")
	}
}

func TestCloudShipping(t *testing.T) {
	client := Client{}
	err := client.InitCloudLogger(&LoggingConfig{
		ProjectID:      "heroic-truck-168212",
		DefaultLogName: "general",
		Logs:           []string{"transaction", "error", "activity"},
		WithTrace:      true,
		TraceAsJSON:    true,
		SimpleTrace:    false,
		Debug:          true,
	})

	if err != nil {
		panic(err)
	}

	op := &Operation{
		Producer: "GET: /some/path/to/awesomeness",
		ID:       "123123jb123b12",
		First:    true,
		Last:     false,
	}
	labels := make(map[string]string)
	labels["CUSTOMER"] = "Google"
	labels["SHORTCODE"] = "Goo"
	labels["PATH"] = "/some/path/to/awesomeness"
	labels["RANDOM"] = "whatDoYouNeedHere?"

	newError := GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong")
	newError.Operation = op
	newError.Labels = labels

	LogERROR(*newError, "transaction")

	newError2 := BadEmailOrPassword(nil)
	op.First = false
	op.Last = true
	newError2.Operation = op
	newError2.Labels = labels

	LogERROR(*newError2, "transaction")
	time.Sleep(time.Second * 5)

}

func TestStdOutShipping(t *testing.T) {

	//time.Sleep(time.Second * 10)
	functionThree()
}
func functionThree() {
	functionTwo()
}
func functionTwo() {
	functionOne()
}
func functionOne() {
	client := Client{}
	err := client.InitStdOutLogger(&LoggingConfig{
		DefaultLogName: "general",
		WithTrace:      true,
		TraceAsJSON:    false,
		SimpleTrace:    true,
		Debug:          true,
	})

	if err != nil {
		panic(err)
	}

	op := &Operation{
		Producer: "GET: /some/path/to/awesomeness",
		ID:       "123123jb123b12",
		First:    true,
		Last:     false,
	}
	labels := make(map[string]string)
	labels["CUSTOMER"] = "Google"
	labels["SHORTCODE"] = "Goo"
	labels["PATH"] = "/some/path/to/awesomeness"
	labels["RANDOM"] = "whatDoYouNeedHere?"

	newError := GenericErrorWithMessage(nil, "A very descriptive message telling others what went wrong")
	newError.Operation = op
	newError.Labels = labels

	LogERROR(*newError, "my-custom-log-tag")

	newError2 := BadEmailOrPassword(nil)
	op.First = false
	op.Last = true
	newError2.Operation = op
	newError2.Labels = labels

	LogERROR(*newError2, "my-custom-log-tag2")

	// redirect stdout to file?
	// read file and compare..
	// delete file ..
}
