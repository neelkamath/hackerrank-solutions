package main

import (
	"fmt"
	"regexp"
	"sort"
)

type email struct {
	firstName string
	emailID   string
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	emails := make([]email, n)
	for i := range emails {
		fmt.Scanf("%s %s", &emails[i].firstName, &emails[i].emailID)
	}
	var firstNames []string
	pattern := regexp.MustCompile(`@gmail\.com`)
	for _, v := range emails {
		if pattern.MatchString(v.emailID) {
			firstNames = append(firstNames, v.firstName)
		}
	}
	sort.Strings(firstNames)
	for _, v := range firstNames {
		fmt.Println(v)
	}
}
