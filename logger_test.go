package logger

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

// func TestStuff(t *testing.T) {
// 	host := "https://zeus.crashguard.io"

// 	FormData := url.Values{
// 		"api_key:": {"api_key:9bdbd23fe6e5"},
// 	}
// 	resp, err := http.PostForm(host, FormData)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var result map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&result)
// 	log.Println(result)
// }
func TestCloudShipping(t *testing.T) {
	logger := Logger{}
	err := logger.Init(&LoggingConfig{
		ProjectID:     "heroic-truck-168212",
		DefaultLogTag: "general",
		Logs:          []string{"transaction", "error", "activity"},
		WithTrace:     true,
		TraceAsJSON:   false,
		SimpleTrace:   true,
		Debug:         true,
		Type:          "google",
	})

	//log.Println(logger)

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

	logger.ERROR(*newError, "transaction")

	newError2 := BadEmailOrPassword(nil)
	op.First = false
	op.Last = true
	newError2.Operation = op
	newError2.Labels = labels

	logger.ERROR(*newError2, "transaction")
	time.Sleep(time.Second * 5)

}

func TestStdOutShipping(t *testing.T) {

	//time.Sleep(time.Second * 10)
	s3()
	time.Sleep(time.Second * 5)
}
func s3() {
	s2()
}
func s2() {
	s1()
}
func s1() {
	logger := Logger{}
	err := logger.Init(&LoggingConfig{
		DefaultLogTag: "general",
		WithTrace:     true,
		TraceAsJSON:   false,
		SimpleTrace:   true,
		Debug:         true,
		Type:          "stdout",
	})

	if err != nil {
		panic(err)
	}

	op := &Operation{
		Producer: "requestid or namespace or anything really..",
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

	logger.ERROR(*newError, "my-custom-log-tag")

	//newError2 := BadEmailOrPassword(nil)
	//op.First = false
	//op.Last = true
	//newError2.Operation = op
	//newError2.Labels = labels

	//LogERROR(*newError2, "my-custom-log-tag2")

	// redirect stdout to file?
	// read file and compare..
	// delete file ..
}
