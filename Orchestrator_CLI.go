package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	START = "start"
	STOP = "stop"
	UPDATE = "update"
	STATUS = "status"
)

func main() {
	var hostname, port, action, testName string
	flag.StringVar(&hostname, "hostname", "localhost", "hostname")
	flag.StringVar(&port, "port", "8080", "port number")
	flag.StringVar(&action, "action", "", "action to perform")
	flag.StringVar(&testName, "test_name", "", "name of test to perform action on")
	flag.Parse()

	if action == "" {
		printErrorForFlag("action")
	}
	if testName == "" {
		printErrorForFlag("test_name")
	}

	switch action {

	case START:
		startTest(hostname, port, testName)
		break
	default:
		flag.PrintDefaults()
		fmt.Printf("\nInvalid argument for 'action', possible values: %s, %s, %s, %s\n", START, STOP, UPDATE, STATUS)
		os.Exit(-1)
	}
}

func startTest(hostname, port, testName string) {
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

func printErrorForFlag(flagName string){
	flag.PrintDefaults()
	fmt.Println(fmt.Errorf("\nMissing value for argument %s\n", flagName))
	os.Exit(-1)
}
