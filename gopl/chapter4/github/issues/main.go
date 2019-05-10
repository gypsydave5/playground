package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gypsydave5/playground/gopl/chapter4/github"
)

const (
	lessThanAMonth int = iota
	lessThanAYear
	moreThanAYear
)

func lessThanMonth(t time.Time) bool {
	return time.Since(t) <= time.Hour*24*30
}

func lessThanYear(t time.Time) bool {
	return time.Since(t) <= time.Hour*24*365
}

func main() {
	ageCategories := [3][]*github.Issue{}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// Exercise 4.10: Modify issues to report the results in age categories, say
	// less than a month old, less than a year old, and more than a year old.
	for _, item := range result.Items {
		switch {
		case lessThanMonth(item.CreatedAt):
			ageCategories[lessThanAMonth] = append(ageCategories[lessThanAMonth], item)
		case lessThanYear(item.CreatedAt):
			ageCategories[lessThanAYear] = append(ageCategories[lessThanAYear], item)
		default:
			ageCategories[moreThanAYear] = append(ageCategories[moreThanAYear], item)
		}
	}

	if len(ageCategories[lessThanAMonth]) != 0 {
		fmt.Println("--- Less Than A Month Old ---")
		for _, item := range ageCategories[lessThanAMonth] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	if len(ageCategories[lessThanAYear]) != 0 {
		fmt.Println("--- Less Than A Year Old ---")
		for _, item := range ageCategories[lessThanAYear] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	if len(ageCategories[moreThanAYear]) != 0 {
		fmt.Println("--- More Than A Year Old ---")
		for _, item := range ageCategories[moreThanAYear] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
