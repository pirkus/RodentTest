package gotest

func With(expected interface{}) func(actual interface{}) (result bool, expected interface{}) {
	return func(actual interface{}) (bool, interface{}) {
		return expected == actual, expected
	}
}
