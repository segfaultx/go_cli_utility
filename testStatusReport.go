package main

type TestStatusReport struct {
	Name      string                      `json:"name"`
	Status    TestStatus                  `json:"status"`
	Message   string                      `json:"message,omitempty"`
	StartTime string                      `json:"startTime"`
	Children  map[string]TestStatusReport `json:"children,omitempty"`
}
