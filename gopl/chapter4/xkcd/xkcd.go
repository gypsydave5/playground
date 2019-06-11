// Package xkcd is part of a solution to Exercise 4.12: The popular web comic
// xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic
// 571, one of many favorites. Download each URL (once!) and build an offline
// index. Write a tool xkcd that, using this index, prints the URL and
// transcript of each comic that matches a search term provided on the command
// line.
package xkcd

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
	return fmt.Sprintf("%s\t%s", c.SafeTitle, c.Img)
}

func (c *Comic) Url() string {
	return fmt.Sprintf("https://xkcd.com/%d", c.Num)
}

// Search a collection of xkcd Comics for ones which match the search term
func Search(cc Comics, terms ...string) Comics {
	var result []*Comic
	for i := range terms {
		terms[i] = strings.ToLower(terms[i])
	}

	for _, c := range cc {
		match := true
		for _, term := range terms {
			match = match &&
				(strings.Contains(strings.ToLower(c.SafeTitle), term) ||
					strings.Contains(strings.ToLower(c.Transcript), term))
		}
		if match {
			result = append(result, c)
		}
	}

	return result
}

// Db of all saved xkcd comics
func Db() (Comics, error) {
	var comics Comics
	_, err := os.Stat(dbFilePath())
	if err != nil {
		err = FetchAndSaveAll()
		if err != nil {
			return comics, err
		}
	}
	db, err := os.Open(dbFilePath())
	if err != nil {
		return comics, err
	}
	err = json.NewDecoder(db).Decode(&comics)
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

// FetchAndSaveAll comic entries from xkcd, saving them in `.xkcd` in
// the user directory.
func FetchAndSaveAll() error {
	max, err := maxComicID()
	if err != nil {
		return err
	}
	cs := fetchComics(0, max)
	file := dbFilePath()
	db, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	json.NewEncoder(db).Encode(cs)
	return nil
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
