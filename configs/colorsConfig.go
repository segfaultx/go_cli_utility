package configs

import (
	"Orchestrator_CLI/constants"
	"github.com/fatih/color"
)

type PrintFuncType func(string, ...interface{})

var outputColors = map[constants.TestStatus]PrintFuncType{
	constants.SUCCESS:   color.Green,
	constants.FAILED:    color.Red,
	constants.PENDING:   color.Yellow,
	constants.EXECUTING: color.Blue,
}
