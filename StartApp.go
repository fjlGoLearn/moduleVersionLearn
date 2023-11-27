package main

import (
	"fmt"
	"github.com/fjlGoLearn/moduleVersionLearn/moduleVersionLearn/InnerVersion"
	"gopl.io/ch4/github"
	"log"
	"time"
)

func main() {
	fmt.Println(InnerVersion.GetInnerVersion())
	issues, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}

	issuesBefore2020 := make([]github.Issue, 0, 20)
	issuesAfter2020 := make([]github.Issue, 0, 20)

	timeThreshold := "2021-10-10 00:00:00"
	dateThreshold, err := time.Parse(time.DateTime, timeThreshold)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range issues.Items {
		fmt.Printf("CreatedAt=%s\n", item.CreatedAt.Format(time.DateTime))
		if dateThreshold.After(item.CreatedAt) {
			issuesBefore2020 = append(issuesBefore2020, *item)
		} else {
			issuesAfter2020 = append(issuesAfter2020, *item)
		}

	}

	fmt.Println(len(issues.Items))
	fmt.Println(len(issuesBefore2020))
	fmt.Println(len(issuesAfter2020))

	//jsonStr, _ := json.Marshal(issues)
	//
	//
	//
	//fmt.Printf("%s\n", jsonStr)

	//fmt.Printf("%d issues:\n", issues.TotalCount)
	//for _, item := range issues.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
}
