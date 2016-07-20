package gotest

import "testing"

var testContext *testing.T
var htmlRepost *HtmlReport
var givens *Givens
var inputsAndOutputs *InputsAndOutputs

func InitialiseRodents(context *testing.T) *testing.T {
	testContext = context
	givens = NewGivens()
	inputsAndOutputs = NewInputsAndOutputs()

	return testContext
}

func Given(returnedGivens *Givens) {
	givens.Merge(returnedGivens)
}

func When(returnedInputsAndOutputs *InputsAndOutputs) {
	inputsAndOutputs.Merge(returnedInputsAndOutputs)
}

func Then(actual interface{}, matcher func(actual interface{}) (result bool, expected interface{})) {
	result, expected := matcher(actual)
	if !result {
		testContext.Fatal("Expected:", expected, "But was:", actual)
	}
}

func BuildReport() {

}
