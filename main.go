package main

import (
	"fmt"
	"os"

	"github.com/hazubeep/github-user-activity-cli/api"
	"github.com/hazubeep/github-user-activity-cli/formatter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	events, err := api.FetchEvents(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	formatter.PrintEvents(events)

}
