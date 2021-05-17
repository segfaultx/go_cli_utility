package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type OrchestratorClient interface {
	New() *OrchestratorClient
	StartTest(hostname, port, testName string)
}

type DefaultOrchestratorClient struct{}

func NewClient() *DefaultOrchestratorClient {
	return &DefaultOrchestratorClient{}
}

func (client *DefaultOrchestratorClient) StartTest(hostname, port, testName string) {
	connString := fmt.Sprintf("http://%s:%s/test/hello", hostname, port)
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
