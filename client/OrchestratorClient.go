package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"orca/configs"
	"orca/constants"
	"orca/testStatusReport"
	"strconv"
	"strings"
)

type OrchestratorClient interface {
	StartTest(hostname, port, testName string)
	GetAndPrintTestStatus(hostname, port, name string)
	StopTest(hostname, port, testName string)
	StartTests(hostname, port string, tests []string)
}

type DefaultOrchestratorClient struct{}

func New() *DefaultOrchestratorClient {
	return &DefaultOrchestratorClient{}
}

func (client *DefaultOrchestratorClient) StartTest(hostname, port, testName string) {
	connString := fmt.Sprintf("%s/test/%s", createBaseUrl(hostname, port), testName)
	executeHttpPostRequest(connString)
}

func (client *DefaultOrchestratorClient) StopTest(hostname, port, testName string) {
	connString := fmt.Sprintf("%s/test/stop/%s", createBaseUrl(hostname, port), testName)
	executeHttpPostRequest(connString)
}

func (client *DefaultOrchestratorClient) StartTests(hostname, port string, tests []string) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("?test=%s", tests[0]))
	for index := range tests[1:] {
		buffer.WriteString(fmt.Sprintf("&test=%s", tests[index]))
	}
	connString := fmt.Sprintf("%s/test%s", createBaseUrl(hostname, port), buffer.String())
	connString = strings.ReplaceAll(connString, " ", "%20")
	executeHttpPostRequest(connString)
}

func (client *DefaultOrchestratorClient) GetAndPrintTestStatus(hostname, port, testName string) {
	connString := fmt.Sprintf("%s/test/status/%s", createBaseUrl(hostname, port), testName)
	response, err := http.Get(connString)
	if err != nil {
		log.Fatalln(err)
	}
	printStatusResponse(response)
}

func createBaseUrl(hostname, port string) string {
	return fmt.Sprintf("http://%s:%s", hostname, port)
}

func executeHttpPostRequest(url string) {
	response, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatalln(err)
	}
	printResponse(response)
}

func printResponseStatus(response *http.Response) {
	fmt.Print("Status: ")
	if response.StatusCode == 200 {
		printFunc := configs.ResponseCodeColors[response.StatusCode]
		printFunc("ok\n")
	} else {
		printFunc := configs.ResponseCodeColors[response.StatusCode]
		printFunc("error (%s)\n", strconv.Itoa(response.StatusCode))
	}
}

func printResponse(response *http.Response) {
	message, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	printResponseStatus(response)
	if response.StatusCode == 200 {
		fmt.Println(string(message))
	} else {
		var resp map[string]interface{}
		err = json.Unmarshal(message, &resp)
		if err != nil {
			log.Fatalln(err)
		}
		msg, exists := resp["message"]

		if !exists {
			log.Fatalln(resp)
		}
		fmt.Println(msg)
	}

}

func printStatusResponse(response *http.Response) {
	message, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	printResponseStatus(response)
	msgBody := testStatusReport.TestStatusReport{}
	err = json.Unmarshal(message, &msgBody)
	if err != nil {
		log.Fatalln(err)
	}
	printStatusReport(msgBody, 0)
}

func printWithIndent(msg string, indent int) {
	addIndentation(indent)
	fmt.Print(msg)
}

func addIndentation(indentLevel int) {
	fmt.Print(strings.Repeat(" ", indentLevel))
}

func printStatusReport(report testStatusReport.TestStatusReport, level int) {
	if level > 0 {
		fmt.Printf("\n")
	}
	printWithIndent(fmt.Sprintf("Name: %s\n", report.Name), level)
	printWithIndent(fmt.Sprintf("Start Time: %s\n", report.StartTime), level)
	addIndentation(level)
	fmt.Printf("Status: ")
	configs.OutputColors[report.Status]("%s\n", report.Status)
	if report.Status == constants.FAILED {
		printWithIndent(fmt.Sprintf("Message: %s", report.Message), level)
	}
	fmt.Print("\n")
	if report.Children != nil {
		if level == 0 {
			printWithIndent(fmt.Sprintf("Children:\n"), level)
		} else {
			printWithIndent(fmt.Sprintf("Steps:\n"), level)
		}

		for _, child := range report.Children {
			printStatusReport(child, level+1)
		}
	}
}
