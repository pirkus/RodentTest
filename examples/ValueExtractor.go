package main

import (
	"encoding/json"
	"encoding/xml"
	"launchpad.net/xmlpath"
	"bytes"
	"log"
	"fmt"
)

type ActivateSubscriberFailureResponse struct {
	XMLName xml.Name `xml:"ActivateSubscriberFailureResponse"`
	//Version     string   `xml:"version,attr"`
	//Svs         []server `xml:"server"`
	//Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type ValueExtractorInterface interface {
	Extract() string
}

type ValueExtractor struct{}

func (v ValueExtractor) Extract(from string, elementName string) interface{} {

	var f interface{}
	json.Unmarshal([]byte(from), &f)
	m := f.(map[string]interface{})
	value, ok := m[elementName]
	if ok {
		return value.(interface{})
	}
	return "Error"
}

func (v ValueExtractor) ExtractXml(fromXml string, xpath string) string {
	//log.Printf("Recieved '%s' Xml and '%s' XPATH", fromXml, xpath)
	buffer := bytes.NewBuffer([]byte(fromXml))
	node, err := xmlpath.Parse(buffer)

	if (err == nil) {
		compile := xmlpath.MustCompile(xpath)
		value, isOk := compile.String(node)

		log.Printf("Value '%s', isOk '%b'", value, isOk)
		if (isOk) {
			return value
		}
	} else {
		fmt.Println("There was an error")
		log.Fatal(err)
	}

	return "Error"
}