package main

import "github.com/fatih/color"

type PrintFuncType func(string, ...interface{})

var outputColors = map[TestStatus]PrintFuncType{
	SUCCESS:   color.Green,
	FAILED:    color.Red,
	PENDING:   color.Yellow,
	EXECUTING: color.Blue,
}
