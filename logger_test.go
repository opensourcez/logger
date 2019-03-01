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
	client := LoggingClient{}
	err := client.InitCloudLogger(&LoggingConfig{
		ProjectID:      "heroic-truck-168212",
		DefaultLogName: "general",
		Logs:           []string{"transaction", "error", "activity"},
	})

	if err != nil {
		panic(err)
	}

	newError := BadEmailOrPassword(nil)
	newError.Operation = Operation{
		Producer: "Transaction Controller",
		ID:       "current-transaction",
		First:    true,
		Last:     false,
	}

	newError.Labels = make(map[string]string)
	newError.Labels["CUSTOMER"] = "ISB"
	newError.Labels["STUFF"] = "isgreat"

	LogERROR(*newError, "transaction")
	time.Sleep(time.Second * 5)

}

func TestStdOutShipping(t *testing.T) {
	client := LoggingClient{}
	err := client.InitStdOutLogger(&LoggingConfig{
		DefaultLogName: "general",
	})

	if err != nil {
		panic(err)
	}

	newError := BadEmailOrPassword(nil)
	newError.Operation = Operation{
		Producer: "Transaction Controller",
		ID:       "current-transaction",
		First:    true,
		Last:     false,
	}

	newError.Labels = make(map[string]string)
	newError.Labels["CUSTOMER"] = "ISB"
	newError.Labels["STUFF"] = "isgreat"

	LogERROR(*newError, "transaction")
	//time.Sleep(time.Second * 10)

}
