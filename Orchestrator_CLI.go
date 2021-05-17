package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	var hostname, port, action, testName string
	flag.StringVar(&hostname, "hostname", "localhost", "hostname")
	flag.StringVar(&port, "port", "8080", "port number")
	flag.StringVar(&action, "action", "", "action to perform")
	flag.StringVar(&testName, "test_name", "", "name of test to perform action on")
	flag.Parse()

	outputColors[SUCCESS]("hi")
	if action == "" {
		printErrorForFlag("action")
	}
	if testName == "" {
		printErrorForFlag("test_name")
	}

	client := NewClient()

	switch action {

	case START:
		client.StartTest(hostname, port, testName)
		break
	case STATUS:
		getAndPrintTestStatus(hostname, port, testName)
	default:
		flag.PrintDefaults()
		fmt.Printf("\nInvalid argument for 'action', possible values: %s, %s, %s, %s\n", START, STOP, UPDATE, STATUS)
		os.Exit(-1)
	}
}

func getAndPrintTestStatus(hostname string, port string, name string) {
	connString := fmt.Sprintf("http://%s:%s/test/status", hostname, port)
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
	msgBody := TestStatusReport{}
	err = json.Unmarshal(message, &msgBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msgBody)
}

func printErrorForFlag(flagName string) {
	flag.PrintDefaults()
	fmt.Println(fmt.Errorf("\nMissing value for argument %s\n", flagName))
	os.Exit(-1)
}
