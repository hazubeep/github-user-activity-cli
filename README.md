# GitHub User Activity CLI

A command-line interface (CLI) tool written in Go that fetches public events of a GitHub user and displays them directly in the terminal.

This project is an implementation of the roadmap.sh project idea. For more details, you can visit the following link:
https://roadmap.sh/projects/github-user-activity

## Features

- Fetches the latest public activity of a GitHub user using the GitHub API in real-time.
- Groups PushEvents and displays the commit counts per repository cleanly.
- Displays WatchEvents (when a user stars a repository).
- Displays CreateEvents (when a user creates a new repository or branch).
- Displays IssuesEvents (when a user opens a new issue).
- Dynamically displays other types of events performed by the user.
- Provides informative error handling when a username is not found or connection issues occur.

## Prerequisites

Before getting started, make sure you have installed:
- Go 1.26.3 or a compatible newer version.

## Installation

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/hazubeep/github-user-activity-cli.git
   cd github-user-activity-cli
   ```

2. Tidy up the Go module dependencies:
   ```bash
   go mod tidy
   ```

## Usage

You can run the program directly without building it first by using the `go run` command:

```bash
go run main.go <username>
```

Example usage:

```bash
go run main.go hazubeep
```

## Compilation (Build)

To compile the application into a binary executable, you can use the provided `Makefile` or build it manually using the `go build` command.

### Using the Makefile

Run the following command:

```bash
make build
```

This will generate a binary executable named `github-activity` in your project directory.

### Building Manually

If you do not have the `make` utility installed, you can build manually by running:

```bash
go build -o github-activity
```

After compilation, you can run the generated binary executable:

```bash
./github-activity <username>
```

## Example Output

When you run the application with a valid and active GitHub username, the output will look like this:

```text
output:
- Pushed 3 commits to hazubeep/github-user-activity-cli
- Starred google/go-cmp
- Created a new branch in hazubeep/github-user-activity-cli
- Opened a new issue in hazubeep/some-repo
```

## Project Reference

This project was built as part of the programming challenges from roadmap.sh. You can explore the challenge details further through the link below:
https://roadmap.sh/projects/github-user-activity