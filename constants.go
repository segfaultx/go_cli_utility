package main

type TestStatus string

const (
	SUCCESS   TestStatus = "SUCCESS"
	PENDING              = "PENDING"
	FAILED               = "FAILED"
	EXECUTING            = "EXECUTING"
)

const (
	START  = "start"
	STOP   = "stop"
	UPDATE = "update"
	STATUS = "status"
)
