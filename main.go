package main

import (
	"fmt"
)

func main() {
	var a []int

	topics := []string{"golang", "boilerplate"}

	repos, err := getReposFromTopics(topics)
	if err != nil {
		panic(err)
	}

	fmt.Println("Repositories:")
	for _, repo := range repos {
		a = append(a, repo.StargazersCount)
		// count := 1
		// fmt.Printf("- %s/%s: %d\n", repo.Owner.Login, repo.Name, repo.StargazersCount)
		fmt.Println("-", repo.RepoUrl)

	}

	for i := 0; i < len(repos); i++ {
		if repos[i].StargazersCount == max(a) {
			// CloneGitUrl()
			fmt.Println()
			fmt.Println("Dumaaaaaaaaaaa:", repos[i].RepoUrl, "-", repos[i].StargazersCount)
			break
		}
	}

	CloneGitUrlMaxStar()
}
