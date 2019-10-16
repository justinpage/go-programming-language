package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/justinpage/go-programming-language/4-composite-types/exercises/4.10.json/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	yearAgo := now.AddDate(-1, 0, 0)
	monthAgo := now.AddDate(0, -1, 0)

	itemAgeCategory := func(created time.Time) string {
		if created.After(monthAgo) {
			return "less-than-a-month-old"
		}

		if created.After(yearAgo) {
			return "less-than-a-year-old"
		}

		return "more-than-a-year-old"
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %-21.21s (%s)\n",
			item.Number, itemAgeCategory(item.Created), item.Created)
	}
}
