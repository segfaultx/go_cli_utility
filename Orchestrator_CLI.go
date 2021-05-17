package main

import (
	"Orchestrator_CLI/client"
	"Orchestrator_CLI/constants"
	"flag"
	"fmt"
	"os"
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

	var orchestratorClient client.OrchestratorClient = client.New()

	switch action {

	case constants.START:
		orchestratorClient.StartTest(hostname, port, testName)
		break
	case constants.STATUS:
		orchestratorClient.GetAndPrintTestStatus(hostname, port, testName)
	default:
		flag.PrintDefaults()
		fmt.Printf("\nInvalid argument for 'action', possible values: %s, %s, %s, %s\n", constants.START, constants.STOP, constants.UPDATE, constants.STATUS)
		os.Exit(-1)
	}
}

func printErrorForFlag(flagName string) {
	flag.PrintDefaults()
	fmt.Println(fmt.Errorf("\nMissing value for argument %s\n", flagName))
	os.Exit(-1)
}
