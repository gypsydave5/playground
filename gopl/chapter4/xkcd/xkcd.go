// Package xkcd is part of a solution to Exercise 4.12: The popular web comic
// xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic
// 571, one of many favorites. Download each URL (once!) and build an offline
// index. Write a tool xkcd that, using this index, prints the URL and
// transcript of each comic that matches a search term provided on the command
// line.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const xkcdURL = "https://xkcd.com/%d/info.0.json"

// A Comic from xkcd
type Comic struct {
	Alt        string `json:"alt"`
	Day        string `json:"day"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      string `json:"month"`
	News       string `json:"news"`
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
}

// Comics from xkcd
type Comics []*Comic

func (c *Comic) String() string {
	return fmt.Sprintf("Title: %s\nLink: %s\n\n%s\n", c.SafeTitle, c.Img, c.Transcript)
}

func (cs Comics) String() string {
	if len(cs) == 0 {
		return ""
	}

	var s strings.Builder
	s.WriteString(cs[0].String())
	for _, c := range cs[1:] {
		s.WriteString("%%\n")
		s.WriteString(c.String())
	}
	return s.String()
}

func main() {
	term := os.Args[1]
	comics, err := db()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(search(comics, term))
}

func search(cc Comics, term string) Comics {
	var result []*Comic
	term = strings.ToLower(term)
	for _, c := range cc {
		if strings.Contains(strings.ToLower(c.SafeTitle), term) {
			result = append(result, c)
		}
		if strings.Contains(strings.ToLower(c.Transcript), term) {
			result = append(result, c)
		}
	}

	return result
}

func db() (Comics, error) {
	var comics Comics
	db, _ := os.Open(dbFilePath())
	err := json.NewDecoder(db).Decode(&comics)
	return comics, err
}

func filter(cs Comics, term string) Comics {
	var result []*Comic

	return result
}

func fetchComic(id int) (*Comic, error) {
	var comic Comic
	url := fmt.Sprintf(xkcdURL, id)
	resp, err := http.Get(url)

	if err != nil {
		return &comic, err
	}

	if resp.StatusCode != http.StatusOK {
		return &comic, fmt.Errorf("status %d for comic %d", resp.StatusCode, id)
	}

	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return &comic, err
	}

	return &comic, nil
}

func fetchComics(start, end int) []*Comic {
	var comics []*Comic
	for i := start; i <= end; i++ {
		comic, err := fetchComic(i)
		if err != nil {
			log.Printf("Could not fetch comic %d\n", i)
			continue
		}
		comics = append(comics, comic)
	}

	return comics
}

func dbFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".xkcd")
}

func fetchAndSaveAll() {
	max, _ := maxComicID()
	cs := fetchComics(0, max)
	file := dbFilePath()
	db, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
	json.NewEncoder(db).Encode(cs)
}

func maxComicID() (int, error) {
	var comic Comic
	url := "https://xkcd.com/info.0.json"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("fetching latest comic Status Code %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}

	return comic.Num, nil
}
