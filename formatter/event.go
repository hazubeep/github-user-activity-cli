package formatter

import (
	"fmt"

	"github.com/hazubeep/github-user-activity-cli/model"
)

func PrintEvents(events []model.GithubEvent) {
	pushCounts := make(map[string]int)

	fmt.Println("output:")
	for _, event := range events {
		if event.Type == "PushEvent" {
			pushCounts[event.Repo.Name]++
		}
	}

	for repo, count := range pushCounts {
		fmt.Printf("- Pushed %d commits to %s\n", count, repo)
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			continue

		case "WatchEvent":
			if event.Payload.Action == "started" {
				fmt.Printf("- Starred %s\n", event.Repo.Name)
			}

		case "CreateEvent":
			if event.Payload.RefType == "branch" {
				fmt.Printf("- Created a new branch in %s\n", event.Repo.Name)
			} else {
				fmt.Printf("- Created a new repository in %s\n", event.Repo.Name)
			}

		case "IssuesEvent":
			if event.Payload.Action == "opened" {
				fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
			}

		default:
			fmt.Printf("- Did a %s in %s\n", event.Type, event.Repo.Name)
		}
	}
}
