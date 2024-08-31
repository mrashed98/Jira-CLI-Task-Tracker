package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

type SearchResult struct {
	Issues []struct {
		Link string `json:"self"`
	} `json:"issues"`
}

type Issue struct {
	Key    string `json:"key"`
	Fields struct {
		Parent     Parent        `json:"parent"`
		Labels     []string      `json:"labels"`
		Issuelinks []interface{} `json:"issuelinks"`
		Project    struct {
			Name string `json:"name"`
		} `json:"project"`
		Timetracking struct {
			OriginalEstimate  string `json:"originalEstimate"`
			RemainingEstimate string `json:"remainingEstimate"`
		} `json:"timetracking"`
		Summary  string `json:"summary"`
		Priority struct {
			Name string `json:"name"`
		} `json:"priority"`
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}

type Parent struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Status  struct {
			Name string `json:"name"`
		} `json:"status"`
		Priority struct {
			Name string `json:"name"`
		} `json:"priority"`
		Issuetype struct {
			Name string `json:"name"`
		} `json:"issuetype"`
	} `json:"fields"`
}

var AllIssues []Issue
var GroupedTasks = make(map[string][]Issue)
var Groups = make(map[string]Parent)

func groupTasks() {
	if len(AllIssues) <= 0 {
		fmt.Println("No Tasks Found While Grouping")
		fmt.Println(AllIssues)
		return
	}

	for _, task := range AllIssues {

		_, ok := Groups[task.Fields.Parent.Fields.Summary]
		if !ok {
			Groups[task.Fields.Parent.Fields.Summary] = task.Fields.Parent
		}

		GroupedTasks[task.Fields.Parent.Fields.Summary] = append(GroupedTasks[task.Fields.Parent.Fields.Summary], task)
	}
}

func GetAllTasks() {

	if len(GroupedTasks) <= 0 {
		fmt.Println("No Tasks Found in the Grouping")
		fmt.Println(GroupedTasks)
		return
	}
	clearScreen()
	currentIdx := 0
	for _, val := range Groups {
		currentIdx = printEpic(val, currentIdx)
	}

}

func printEpic(epic Parent, lastIndex int, status ...string) int {
	maxWidths := calculateMaxWidths(GroupedTasks[epic.Fields.Summary])
	clearScreen()
	fmt.Printf("Epic: %s Number of tasks: %d\n", epic.Fields.Summary, len(GroupedTasks[epic.Fields.Summary]))
	fmt.Println(strings.Repeat("=", 150))
	fmt.Printf("%-*s\t%-*s\t%-*s\t%-*s\n", maxWidths[0], "Task id", maxWidths[1], "Task name", maxWidths[2], "Status", maxWidths[3], "Link")
	fmt.Println(strings.Repeat("=", 150))
	for _, task := range GroupedTasks[epic.Fields.Summary] {
		JIRA_URL := os.Getenv("JIRA_URL") + "/" + task.Key
		lastIndex += 1
		if len(status) > 0 && task.Fields.Status.Name == status[0] {
			fmt.Printf("%-*d\t%-*s\t%-*s\t%-*s\n", maxWidths[0], lastIndex, maxWidths[1], task.Fields.Summary, maxWidths[2], task.Fields.Status.Name,
				maxWidths[3], JIRA_URL)
		} else if len(status) == 0 {
			lastIndex += 1
		}

	}
	fmt.Println(strings.Repeat("=", 150))
	fmt.Printf("\n\n")
	return lastIndex
}

func calculateMaxWidths(tasks []Issue) []int {
	maxWidths := []int{0, 0, 0, 0}
	currentIndex := 0
	for _, task := range tasks {
		currentIndex += 1
		if len(fmt.Sprintf("%d", currentIndex)) > maxWidths[0] {
			maxWidths[0] = len(fmt.Sprintf("%d", currentIndex))
		}
		if len(task.Fields.Summary) > maxWidths[1] {
			maxWidths[1] = len(task.Fields.Summary)
		}
		if len(task.Fields.Status.Name) > maxWidths[2] {
			maxWidths[2] = len(task.Fields.Status.Name)
		}
		JIRA_URL := os.Getenv("JIRA_URL") + "/" + task.Key
		if len(JIRA_URL) > maxWidths[3] {
			maxWidths[3] = len(JIRA_URL)
		}
	}
	return maxWidths
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("tput", "clear")
	default:
		// Fallback to ANSI escape codes
		print("\033[H\033[2J")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}

func GetOpenTasks() {
	var groupsWithOpenTasks []string
	for key := range Groups {
		for _, task := range GroupedTasks[key] {
			if task.Fields.Status.Name == "Open" {
				groupsWithOpenTasks = append(groupsWithOpenTasks, key)
				break
			}
		}
	}
	if len(groupsWithOpenTasks) != 0 {
		currentIndex := 0
		for _, group := range groupsWithOpenTasks {
			printEpic(Groups[group], currentIndex, "Open")
		}
	}
}

func GetInProgressTasks() {
	var groupsWithInProgressTasks []string
	for key := range Groups {
		for _, task := range GroupedTasks[key] {
			if task.Fields.Status.Name == "In Progress" {
				groupsWithInProgressTasks = append(groupsWithInProgressTasks, key)
				break
			}
		}
	}
	if len(groupsWithInProgressTasks) != 0 {
		currentIndex := 0
		for _, group := range groupsWithInProgressTasks {
			printEpic(Groups[group], currentIndex, "In Progress")
		}
	}
}

func GetCompletedTasks() {
	var groupsWithCompletedTasks []string
	for key := range Groups {
		for _, task := range GroupedTasks[key] {
			if task.Fields.Status.Name == "Done" {
				groupsWithCompletedTasks = append(groupsWithCompletedTasks, key)
				break
			}
		}
	}
	if len(groupsWithCompletedTasks) != 0 {
		currentIndex := 0
		for _, group := range groupsWithCompletedTasks {
			printEpic(Groups[group], currentIndex, "Done")
		}
	}
}

func GetTasksWithSpecificFlag(flag string) {
	//TODO
}

func ChangeTaskStatus() {
	//TODO
}

func populateIssues(issuelinks []string, username string, apiToken string) {
	client := &http.Client{}
	for _, link := range issuelinks {
		req, err := http.NewRequest("GET", link, nil)
		if err != nil {
			log.Fatalf("Error occurred while generating task request: %v", err)
		}
		req.SetBasicAuth(username, apiToken)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error occurred while getting task info: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error occurred while reading task response body: %v", err)
		}
		var issue Issue
		err = json.Unmarshal(body, &issue)
		if err != nil {
			log.Fatalf("Error occurred while parsing task response body: %v", err)
		}
		AllIssues = append(AllIssues, issue)
	}
	groupTasks()
}

func populateSearchResults() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERROR OCCURRED LOADING ENV FILE: %v", err)
	}
	JIRA_URL := os.Getenv("JIRA_URL")
	USERNAME := os.Getenv("USERNAME")
	API_TOKEN := os.Getenv("API_TOKEN")
	LABEL := os.Getenv("LABEL")

	client := &http.Client{}

	jql := url.QueryEscape(fmt.Sprintf("assignee = currentUser() AND labels = %s AND issuetype = Sub-task", LABEL))
	url := fmt.Sprintf("%s/rest/api/3/search?jql=%s", JIRA_URL, jql)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.SetBasicAuth(USERNAME, API_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while executing the search: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error while reading jira response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected Status Code: %d, body: %s", resp.StatusCode, body)
	}

	var searchResult SearchResult
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		log.Fatalf("Error while deoding Response Body: %v", err)
	}

	var issuelinks []string
	for _, issue := range searchResult.Issues {
		issuelinks = append(issuelinks, issue.Link)
	}
	populateIssues(issuelinks, USERNAME, API_TOKEN)
}

func init() {
	populateSearchResults()
}
