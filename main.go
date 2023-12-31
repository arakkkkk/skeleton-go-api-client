package main

import (
	"fmt"
	"go-api-client-sample/samplesdk/client"
	"go-api-client-sample/samplesdk/issue"
	"os"
)

func main() {
	cfg := &client.Config{EndPoint: "https://api.github.com", ApiToken: os.Getenv("APIKEY")}

	issueC, err := issue.New(cfg, nil)
	if err != nil {
		panic(err)
	}
	issueModel, err := issueC.ListIssue(&issue.ListIssueQuery{PerPage: 2})
	 if err != nil {
		panic(err)
	}
	fmt.Println(issueModel)

	// res, err := issue.CreateIssue()
	//  if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)

}
