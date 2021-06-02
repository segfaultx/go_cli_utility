package main

import (
	"flag"
	"fmt"
	"orca/client"
	"orca/constants"
	"os"
	"regexp"
	"strings"
)

func main() {

	var hostname, port, action, testName string
	flag.StringVar(&hostname, "hostname", "localhost", "hostname")
	flag.StringVar(&port, "port", "8080", "port number")
	flag.StringVar(&action, "action", "", "action to perform")
	flag.StringVar(&testName, "test_name", "", "name of test to perform action on, use 'ALL' for all tests")
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
		{
			if matched, _ := regexp.MatchString("([\\w\\W]+,[\\w\\W]+)+", testName); matched {
				orchestratorClient.StartTests(hostname, port, strings.Split(testName, ","))
			} else {
				orchestratorClient.StartTest(hostname, port, testName)
			}
		}
	case constants.STOP:
		orchestratorClient.StopTest(hostname, port, testName)

	case constants.STATUS:
		orchestratorClient.GetAndPrintTestStatus(hostname, port, testName)

	default:
		flag.PrintDefaults()
		fmt.Printf("\nInvalid argument for 'action', possible values: %s, %s, %s, %s\n",
			constants.START, constants.STOP, constants.UPDATE, constants.STATUS)
		os.Exit(-1)
	}
}

func printErrorForFlag(flagName string) {
	flag.PrintDefaults()
	fmt.Println(fmt.Errorf("\nMissing value for argument %s\n", flagName))
	os.Exit(-1)
}
