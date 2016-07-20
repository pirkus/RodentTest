package gotest

import (
	"runtime"
	"strings"
)

func GetTestName() string {
	pathTokens := strings.Split(GetPathToTestFile(), "/")

	// now we got FileName.go, get rid of .go
	return strings.Replace(pathTokens[len(pathTokens)-1], ".go", "", -1)
}

func GetPathToTestFile() string {
	// runtime.Caller(0) - return a similar thing -> /Users/filip/gopath/package/FileName.go
	for i := 0; i < 10; i++ {
		_, filename, _, ok := runtime.Caller(i) // to get the name skip 2 stack frames (2 functions backwards)
		if !ok {
			panic("No caller information")
		}

		if strings.Contains(filename, "_test.go") {
			return filename
		}
	}

	panic("Method was not called from a test.")
}

func GetGivenMethodName() string {
	method, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("No caller information")
	}

	return runtime.FuncForPC(method).Name()
}
