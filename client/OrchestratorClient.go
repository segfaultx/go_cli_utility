package client

import (
	"Orca/testStatusReport"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type OrchestratorClient interface {
	StartTest(hostname, port, testName string)
	GetAndPrintTestStatus(hostname, port, name string)
}

type DefaultOrchestratorClient struct{}

func New() *DefaultOrchestratorClient {
	return &DefaultOrchestratorClient{}
}

func createBaseUrl(hostname, port string) string {
	return fmt.Sprintf("http://%s:%s", hostname, port)
}

func (client *DefaultOrchestratorClient) StartTest(hostname, port, testName string) {
	connString := fmt.Sprintf("%s/test/hello", createBaseUrl(hostname, port))
	response, err := http.Get(connString)
	if err != nil {
		log.Fatalln(err)
	}
	printResponse(response)
}

func printResponse(response *http.Response) {
	message, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(message))
}

func (client *DefaultOrchestratorClient) GetAndPrintTestStatus(hostname, port, name string) {
	connString := fmt.Sprintf("%s/test/status", createBaseUrl(hostname, port))
	response, err := http.Get(connString)
	if err != nil {
		log.Fatalln(err)
	}
	printStatusResponse(response)
}

func printStatusResponse(response *http.Response) {
	message, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(message))
	msgBody := testStatusReport.TestStatusReport{}
	err = json.Unmarshal(message, &msgBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msgBody)
}
