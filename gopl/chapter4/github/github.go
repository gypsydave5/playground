package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issues = []Issue

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Create struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func CreateIssue(owner, repo, title, issueBody string) (issue Issue, err error) {
	const RepoIssuesURL = "https://api.github.com/repos/%s/%s/issues"
	create := Create{
		Body:  issueBody,
		Title: title,
	}
	body, _ := json.Marshal(create)

	b := bytes.NewReader(body)

	resp, err := http.Post(
		fmt.Sprintf(RepoIssuesURL, owner, repo),
		"application/json",
		b,
	)
	if err != nil {
		return issue, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Issue{}, fmt.Errorf("create issue failed for repo %s/%s failed: %s", owner, repo, resp.Status)
	}

	json.NewDecoder(resp.Body).Decode(&issue)
	return issue, nil
}

// GetIssues gets all the issues for a particular repo
func GetIssues(owner, repo string) (*Issues, error) {
	const RepoIssuesURL = "https://api.github.com/repos/%s/%s/issues"
	resp, err := http.Get(fmt.Sprintf(RepoIssuesURL, owner, repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issues for repo %s/%s failed: %s", owner, repo, resp.Status)
	}

	var result Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetIssue details
func GetIssue(owner, repo string, id int) (Issue, error) {
	const repoIssueURL = "https://api.github.com/repos/%s/%s/issues/%d"
	var result Issue
	resp, err := http.Get(fmt.Sprintf(repoIssueURL, owner, repo, id))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("get issue %d for repo %s/%s failed: %s", id, owner, repo, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
