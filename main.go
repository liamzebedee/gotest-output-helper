package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// TestEvent represents a single event in the JSON output from go test -json.
type TestEvent struct {
	Time    string  `json:"Time"`
	Action  string  `json:"Action"`
	Package string  `json:"Package"`
	Test    string  `json:"Test,omitempty"`
	Elapsed float64 `json:"Elapsed,omitempty"`
	Output  string  `json:"Output,omitempty"`
}

var testOutputs map[string]string

func main() {
	testOutputs = make(map[string]string)
	decoder := json.NewDecoder(os.Stdin)
	for {
		var event TestEvent
		if err := decoder.Decode(&event); err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
			os.Exit(1)
		}
		processEvent(event)
	}
}

func processEvent(event TestEvent) {
	testOutputs[event.Test] += event.Output

	switch event.Action {
	case "start":
	case "run":
	case "pass":
		if event.Test == "" {
			return
		}
		fmt.Printf("PASS: %s (%.2fs)\n", event.Test, event.Elapsed)
	case "fail":
		fmt.Printf("FAIL: %s (%.2fs)\n", event.Test, event.Elapsed)
		fmt.Printf("\n%s\n", testOutputs[event.Test])
	case "skip":
	case "output":
	default:
		fmt.Printf("Unknown action: %s\n", event.Action)
	}
}
