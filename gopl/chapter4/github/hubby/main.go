package main

import (
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
	command := args[0]
	switch command {
	case "list":
		listIssues(args[1:])
	case "details":
		details(args[1:])
	case "create":
		create(args[1:])
	default:
		usage()
		os.Exit(1)
	}

	os.Exit(0)
}

func usage() {
	fmt.Println("hubby usage")
	fmt.Println("hubby [list]")
}

func create(args []string) {
	editor := os.Getenv("EDITOR")
	tmpfile, err := ioutil.TempFile("", "example*.markdown")
	if err != nil {
		log.Fatal(err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Run()

	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
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
	fmt.Printf("Number:\t\t%d\nTitle:\t\t%s\nState:\t\t%s\nUser:\t\t%s\nCreatedAt:\t%v\nURL:\t\t%s\n\n",
		issue.Number,
		issue.Title,
		issue.State,
		issue.User.Login,
		issue.CreatedAt,
		issue.HTMLURL,
	)
	fmt.Printf("%s\n", issue.Body)
}

func listIssues(args []string) {
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
