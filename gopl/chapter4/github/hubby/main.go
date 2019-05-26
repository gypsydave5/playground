package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/gypsydave5/playground/gopl/chapter4/github"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}
	command := args[0]
	switch command {
	case "list":
		listIssues(args[1:])
	case "details":
		details(args[1:])
	case "create":
		create(args[1:])
	case "update":
		update(args[1:])
	default:
		usage()
		os.Exit(1)
	}

	os.Exit(0)
}

func usage() {
	fmt.Println("hubby usage")
	fmt.Println("hubby [list | details | create | update]")
}

func update(args []string) {
	if len(args) != 2 {
		detailsUsage()
		os.Exit(1)
	}
	user, repo := ownerAndRepo(args[0])
	id, err := strconv.Atoi(args[1])
	if err != nil {
		detailsUsage()
		os.Exit(1)
	}
	issue, _ := github.GetIssue(user, repo, id)

	content, err := editorInput(fmt.Sprintf("%s\n\n%s", issue.Title, issue.Body))
	if err != nil {
		log.Fatal(err)
	}
	title, body := getTitleAndBody(content)
	issue, err = github.UpdateIssue(user, repo, id, title, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printIssue(issue)
}

func create(args []string) {
	user, repo := ownerAndRepo(args[0])
	content, err := editorInput("")
	if err != nil {
		log.Fatal(err)
	}
	title, body := getTitleAndBody(content)
	issue, err := github.CreateIssue(user, repo, title, body)
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}

func getTitleAndBody(content []byte) (title, body string) {
	s := bytes.SplitN(content, []byte{'\n', '\n'}, 2)
	if len(s) == 2 {
		body = string(s[1])
	}
	return string(s[0]), body
}

func editorInput(initial string) (input []byte, err error) {
	editor := os.Getenv("EDITOR")
	tmpfile, err := ioutil.TempFile("", "editor-input")
	if err != nil {
		return input, err
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.WriteString(initial)
	tmpfile.Close()

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Run()

	return ioutil.ReadFile(tmpfile.Name())
}

func details(args []string) {
	if len(args) != 2 {
		detailsUsage()
		os.Exit(1)
	}
	user, repo := ownerAndRepo(args[0])
	id, err := strconv.Atoi(args[1])
	if err != nil {
		detailsUsage()
		os.Exit(1)
	}
	issue, _ := github.GetIssue(user, repo, id)
	printIssue(issue)
}

func printIssue(i github.Issue) {
	fmt.Printf("Number:\t\t%d\nTitle:\t\t%s\nState:\t\t%s\nUser:\t\t%s\nCreatedAt:\t%v\nURL:\t\t%s\n\n",
		i.Number,
		i.Title,
		i.State,
		i.User.Login,
		i.CreatedAt,
		i.HTMLURL,
	)
	fmt.Printf("%s\n", i.Body)
}

func listIssues(args []string) {
	if len(args) != 1 {
		listUsage()
		os.Exit(1)
	}
	user, repo := ownerAndRepo(args[0])

	result, err := github.GetIssues(user, repo)
	if err != nil {
		panic(err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "id\tuser\tdate\ttitle")
	for _, i := range *result {
		fmt.Fprintf(w, "%d\t%s\t%v\t%s\n",
			i.Number,
			i.User.Login,
			i.CreatedAt.Format("01/02/2006"),
			i.Title,
		)
	}

	w.Flush()
}

func ownerAndRepo(s string) (owner, repo string) {
	x := strings.Split(s, "/")
	return x[0], x[1]
}

func detailsUsage() {
	fmt.Println("hubby details USAGE")
	fmt.Println("hubby details [owner/repo] [issueNumber]")
}

func listUsage() {
	fmt.Println("hubby list USAGE")
	fmt.Println("hubby list [owner]/[repo]")
}
