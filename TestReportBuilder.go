package gotest

import (
	"io/ioutil"
	"strings"
)

type HtmlReport struct {
	Name string
	Body string
}

func NewHtmlReport(name string) *HtmlReport {
	return &HtmlReport{name, ""}
}

func (report *HtmlReport) BuildReport(result int) {
	report.Body = strings.Replace(htmlTemplate, "{TEST_NAME}", report.Name, -1)
	if result == 0 {
		report.Body = strings.Replace(report.Body, "{TEST_RESULT}", testPassed, -1)
	} else {
		report.Body = strings.Replace(report.Body, "{TEST_RESULT}", testFailed, -1)
	}
	report.Body = strings.Replace(report.Body, "{SPECIFICATION}", extractSpecification(), -1)
}

func (report *HtmlReport) Save(result int) error {
	report.BuildReport(result)
	filename := report.Name + ".html"
	return ioutil.WriteFile(filename, []byte(report.Body), 0600)
}

func extractSpecification() string {
	var result string
	byteFileContents, err := ioutil.ReadFile(GetPathToTestFile())
	if err != nil {
		panic("Failed to read the source file.")
	}

	fileContents := string(byteFileContents)
	split := strings.Split(fileContents, "\n")
	for _, v := range split {
		if strings.Contains(v, "Given(") || strings.Contains(v, "When(") || strings.Contains(v, "Then(") {
			result += v + "<br />"
		}
	}
	return result
}

var htmlTemplate = "<html>" +
	"<head>" +
	"<style type=\"text/css\">" +
	"/* <![CDATA[ */" +
	"html, body {" +
	"  margin: 1em 2ex 2em;" +
	"  padding: 0;" +
	"  background-color: #fff;" +
	"  color: #000;" +
	"  font-family: Arial, Helvetica, Verdana, sans-serif;" +
	"  font-size: 10pt;" +
	"}" +
	"" +
	"h1, h2, h3, h4, h5, h6, th {" +
	"  text-transform: capitalize;" +
	"}" +
	"" +
	"h1 {" +
	"  font-size: 170%;" +
	"}" +
	"" +
	"h2 {" +
	"  font-size: 150%;" +
	"}" +
	"" +
	"h3 {" +
	"  font-size: 140%;" +
	"}" +
	"" +
	"h4 {" +
	"  font-size: 130%;" +
	"}" +
	"" +
	"h5 {" +
	"  font-size: 120%;" +
	"}" +
	"" +
	"h6 {" +
	"  font-size: 110%;" +
	"}" +
	"" +
	".highlight {" +
	"  background-color: #E8EEF7;" +
	"  border: 1px solid #C3D9FF;" +
	"  padding: 5px;" +
	"  font: 10pt Arial, sans-serif;" +
	"}" +
	"" +
	".highlight .literal {" +
	"  color: #066;" +
	"}" +
	"" +
	".highlight .keyword {" +
	"  color: #008;" +
	"}" +
	"" +
	".highlight .constant {" +
	"  color: #606;" +
	"}" +
	"" +
	".highlight .quote {" +
	"  color: #080;" +
	"}" +
	"" +
	".highlight.specification .literal," +
	".highlight.specification .keyword," +
	".highlight.specification .constant," +
	".highlight.specification .quote {" +
	"  font-weight: bold;" +
	"}" +
	"" +
	"pre.test-not-run, .test-not-run, .interestingGiven {" +
	"  background-color: #FFF4CC;" +
	"  border: 2px solid #FFBB66;" +
	"}" +
	"" +
	".interestingGiven {" +
	"  padding: 0 2px;" +
	"}" +
	"" +
	"pre.test-failed, .test-failed {" +
	"  background-color: #FFF6FF;" +
	"  border: 2px solid #FDA8A8;" +
	"}" +
	"" +
	"pre.test-passed, .test-passed {" +
	"  background-color: #F2FFEE;" +
	"  border: 2px solid #A5DDAD;" +
	"}" +
	"" +
	"li.test-not-run, li.test-passed {" +
	"  border: none;" +
	"}" +
	"" +
	"ul.contents a {" +
	"  text-transform: capitalize;" +
	"}" +
	"" +
	"a, table a {" +
	"  color: #660066;" +
	"  text-decoration: none;" +
	"}" +
	"" +
	"a:hover, table a:hover {" +
	"  text-decoration: underline;" +
	"}" +
	"" +
	"table {" +
	"  background-color: #F5F9FD;" +
	"  border: 1px solid #C3D9FF;" +
	"  border-collapse: collapse;" +
	"  empty-cells: show;" +
	"}" +
	"" +
	"th {" +
	"  background-color: #E8EEF7;" +
	"  border: 1px solid #C3D9FF;" +
	"}" +
	"" +
	"td {" +
	"  border: 1px solid #C3D9FF;" +
	"}" +
	"" +
	"div.contents, div.testmethod {" +
	"  margin-bottom: 40px;" +
	"}" +
	"" +
	".scenarios td {" +
	"  padding: 0 20px;" +
	"  text-align: center;" +
	"}" +
	"" +
	".scenarios, .scenario {" +
	"  margin-bottom: 20px;" +
	"}" +
	"" +
	".scenario {" +
	"  background-color: #F5F9FD;" +
	"  border: 1px solid #C3D9FF;" +
	"  padding: 0 10px;" +
	"}" +
	"" +
	"h2 {" +
	"  margin-top: 10px;" +
	"}" +
	"" +
	"pre {" +
	"  margin-bottom: 32px;" +
	"}" +
	"" +
	"div.example {" +
	"  margin-bottom: 64px;" +
	"}" +
	"" +
	".hide {" +
	"  display: none;" +
	"}" +
	"" +
	".logKey {" +
	"  cursor: pointer;" +
	"}" +
	"" +
	".logKey:hover {" +
	"  text-decoration: underline;" +
	"}" +
	"" +
	".interestingGivens th {" +
	"  text-align: left;" +
	"}" +
	"" +
	".interestingGivens td {" +
	"  text-align: right;" +
	"  padding: 0 5px;" +
	"  vertical-align: middle;" +
	"}" +
	"" +
	".logValue {" +
	"  white-space: pre-wrap;" +
	"  margin-bottom: 32px;" +
	"}" +
	"" +
	".package-name {" +
	"  cursor: pointer;" +
	"}" +
	"" +
	".package-name:hover {" +
	"  text-decoration: underline;" +
	"}" +
	"" +
	"dt {" +
	"  font-weight: bold;" +
	"  margin: 0px;" +
	"  padding: 2px;" +
	"}" +
	"" +
	"dd {" +
	"  font-weight: normal;" +
	"  margin-top:  0px;" +
	"  margin-bottom:  1px;" +
	"  margin-left:  20px;" +
	"  margin-right:  0px;" +
	"  padding: 2px;" +
	"  padding-right: 0px;" +
	"}" +
	"" +
	"dl {" +
	"  margin: 0px;" +
	"}" +
	"dl.index-result {" +
	"  padding: 10px;" +
	"}" +
	".index-result dt {" +
	"  margin-bottom: 5px;" +
	"}" +
	".index-result dd {" +
	"  margin-left: 0px;" +
	"  margin-right: 5px;" +
	"}" +
	"/* ]]> */" +
	"</style>" +
	"</head>" +
	"<body>" +
	"  " +
	"  <div class=\"testmethod\">" +
	"    " +
	"    <a id=\"returns404WhenSimbaRespondsWithSimNotFoundForIccid\"></a>" +
	"    <h2>{TEST_NAME}</h2>" +
	"    " +
	"    <div class=\"scenario\" id=\"866699721_1034879960\">" +
	"      <a id=\"\"></a>" +
	"      <h2>Specification</h2>" +
	"      <span class=\"keyword\"></span>" +
	"      {SPECIFICATION}" +
	"    </pre>" +
	"    <h2>Test results:</h2>" +
	"" +
	"    {TEST_RESULT}" +
	"  </div>" +
	"</div>" +
	"</body>" +
	"</html>"

var testPassed = "<pre class=\"highlight results test-passed highlighted\">Test passed</pre>"
var testFailed = "<pre class=\"highlight results test-failed highlighted\">Test failed</pre>"
