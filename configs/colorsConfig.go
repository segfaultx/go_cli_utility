package configs

import (
	"Orca/constants"
	"github.com/fatih/color"
)

type PrintFuncType func(string, ...interface{})

var OutputColors = map[constants.TestStatus]PrintFuncType{
	constants.SUCCESS:   color.Green,
	constants.FAILED:    color.Red,
	constants.PENDING:   color.Yellow,
	constants.EXECUTING: color.Blue,
}

var ResponseCodeColors = map[int]PrintFuncType{
	200: color.Green,
	400: color.Red,
	404: color.Red,
	403: color.Red,
	500: color.Red,
}
