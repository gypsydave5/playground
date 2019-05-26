package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const issuesURL = "https://api.github.com/search/issues"

// The IssuesSearchResult from GitHub
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issues from GitHub
type Issues = []Issue

// An Issue on GitHub
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// A User on GitHub
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// The IssueDetails from GitHub
type IssueDetails struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

// DeleteIssue is unimplemented as it is not supported on the current GitHub API
func DeleteIssue() {
}

// UpdateIssue on GitHub
func UpdateIssue(owner, repo string, id int, title, issueBody string) (issue Issue, err error) {
	const RepoIssuesURL = "https://api.github.com/repos/%s/%s/issues/%d"
	create := IssueDetails{
		Body:  issueBody,
		Title: title,
	}
	body, _ := json.Marshal(create)
	b := bytes.NewReader(body)

	url := fmt.Sprintf(RepoIssuesURL, owner, repo, id)
	req, _ := http.NewRequest("PATCH", url, b)
	token := os.Getenv("GITHUB_TOKEN")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return issue, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Issue{}, fmt.Errorf("update issue failed for repo %s/%s failed: %s\n\n%s", owner, repo, resp.Status, url)
	}

	json.NewDecoder(resp.Body).Decode(&issue)
	return issue, nil
}

// CreateIssue for a repo on GitHub
func CreateIssue(owner, repo, title, issueBody string) (issue Issue, err error) {
	const RepoIssuesURL = "https://api.github.com/repos/%s/%s/issues"
	create := IssueDetails{
		Body:  issueBody,
		Title: title,
	}
	body, _ := json.Marshal(create)
	b := bytes.NewReader(body)

	url := fmt.Sprintf(RepoIssuesURL, owner, repo)
	req, _ := http.NewRequest("POST", url, b)
	token := os.Getenv("GITHUB_TOKEN")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return issue, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return Issue{}, fmt.Errorf("create issue failed for repo %s/%s failed: %s\n\n%s", owner, repo, resp.Status, url)
	}

	json.NewDecoder(resp.Body).Decode(&issue)
	return issue, nil
}

// GetIssues for a repo on GitHub
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

// GetIssue details from GitHub
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

// SearchIssues on GitHub
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issuesURL + "?q=" + q)
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
