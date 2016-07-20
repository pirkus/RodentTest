package main

import (
	"flag"
	. "gotest"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var body string
var testContext *testing.T

func TestMain(m *testing.M) {
	flag.Parse()
	//setUp
	result := m.Run()
	//tearDown
	htmlReport := NewHtmlReport(GetTestName())
	htmlReport.Save(result)
	// END: tearDown
	os.Exit(result)
}

func TestGetRequest(t *testing.T) {
	testContext = InitialiseRodents(t)

	Given(theHTTPServerIsRunning(onPort("8080")))
	When(someoneFiresAGetRequest())
	Then(theUserGetsPresented(), With("Jump jump"))

	BuildReport()
}

func theHTTPServerIsRunning(port string) *Givens {
	go StartServer(port)
	givens := NewGivens()
	givens.Add("port", port)

	return givens
}

func onPort(port string) string {
	return port
}

func someoneFiresAGetRequest() *InputsAndOutputs {
	response, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		testContext.Fatal("Get request failed. Reason: ", err)
	}

	defer response.Body.Close()
	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		testContext.Fatal("Failed to retrieve body from the request. Reason: ", err)
	}
	body = string(bodyByte[:])

	return NewInputsAndOutputs()
}

func theUserGetsPresented() string {
	return body
}
