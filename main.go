package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	jira "github.com/andygrunwald/go-jira"
)

func main() {
	summary := os.Args[1]
	description := os.Args[2]
	jiraURL := os.Getenv("SQUIRREL_JIRA_URL")
	jiraAPIKey := os.Getenv("SQUIRREL_JIRA_APIKEY")
	jiraUsername := os.Getenv("SQUIRREL_JIRA_USERNAME")

	if len(summary) == 0 {
		fmt.Println("Squirrel error: summary cannot be empty")
		return
	}

	if len(description) == 0 {
		fmt.Println("Squirrel warning: description is empty, but we'll create it anyways")
	}

	fmt.Printf("Creating story '%s' with description '%s'\n", summary, description)
	tp := jira.BasicAuthTransport{
		Username: jiraUsername,
		Password: jiraAPIKey,
	}
	client, _ := jira.NewClient(tp.Client(), jiraURL)

	user, b, err := client.User.GetSelf()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if b.StatusCode != 200 {
		fmt.Printf("%#v\n", b.Response)
	}
	boards, b, err := client.Board.GetAllBoards(&jira.BoardListOptions{
		Name: "BOSS",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if b.StatusCode != 200 {
		fmt.Printf("%#v\n", b.Response)
	}
	if len(boards.Values) == 0 {
		fmt.Println("Squirrel error: no board found with name 'boss'")
		return
	}
	if len(boards.Values) > 1 {
		fmt.Println("Squirrel error: more than one board found with name 'boss'")
		return
	}

	sprints, _, err := client.Board.GetAllSprints(strconv.Itoa(boards.Values[0].ID))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var activeSprint jira.Sprint
	for _, sprint := range sprints {
		if sprint.State == "active" {
			activeSprint = sprint
			break
		}
	}
	if activeSprint.ID == 0 || activeSprint.Name == "" {
		fmt.Println("Squirrel error: no active sprint found")
		return
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				AccountID: user.AccountID,
			},
			Reporter: &jira.User{
				AccountID: user.AccountID,
			},
			Description: description,
			Type: jira.IssueType{
				Name: "Story",
			},
			Project: jira.Project{
				Key: "BOSS",
			},
			Parent: &jira.Parent{
				Key: "BOSS-576",
			},
			Summary: summary,
		},
	}

	issue, b2, err := client.Issue.Create(&i)
	if b2.StatusCode != 200 {
		aaa, _ := io.ReadAll(b2.Body)
		fmt.Println(string(aaa))
	}
	if err != nil {
		panic(err)
	}

	client.Sprint.MoveIssuesToSprint(activeSprint.ID, []string{issue.ID})

	// Move it to done
	var transitionID string
	possibleTransitions, _, _ := client.Issue.GetTransitions(issue.ID)
	for _, v := range possibleTransitions {
		if v.Name == "Done" {
			transitionID = v.ID
			break
		}
	}
	client.Issue.DoTransition(issue.ID, transitionID)
	fmt.Printf("Created story %s\n", issue.Key)
}
