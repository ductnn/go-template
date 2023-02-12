package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/olekukonko/tablewriter"
)

const (
	GitDirectory = ".git"
)

var name plumbing.ReferenceName

func CloneGitUrl(topics []string) error {
	repos, err := getReposFromTopics(topics)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		url := repo.CloneUrl
		branchName := repo.DefaultBranch
		ProjectName := repo.Name

		if len(strings.TrimSpace(branchName)) != 0 {
			name = plumbing.NewBranchReferenceName(branchName)
		}

		cloneOptions := &git.CloneOptions{
			URL:           url,
			Progress:      os.Stdout,
			ReferenceName: name,
		}

		_, err := git.PlainClone(ProjectName, false, cloneOptions)
		if err != nil {
			return err
		}

		gitDirectory := ProjectName + string(filepath.Separator) + GitDirectory
		if err = os.RemoveAll(gitDirectory); err != nil {
			return err
		}
	}

	return err
}

func CloneGitUrlMaxStar(topics []string) error {
	var starArray []int

	repos, err := getReposFromTopics(topics)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		starArray = append(starArray, repo.StargazersCount)
		url := repo.CloneUrl
		branchName := repo.DefaultBranch
		ProjectName := repo.Name

		if len(strings.TrimSpace(branchName)) != 0 {
			name = plumbing.NewBranchReferenceName(branchName)
		}

		cloneOptions := &git.CloneOptions{
			URL:           url,
			Progress:      os.Stdout,
			ReferenceName: name,
		}

		if repo.StargazersCount == max(starArray) {
			_, err := git.PlainClone(ProjectName, false, cloneOptions)
			if err != nil {
				return err
			}
		}

		gitDirectory := ProjectName + string(filepath.Separator) + GitDirectory
		if err = os.RemoveAll(gitDirectory); err != nil {
			return err
		}
	}

	return err
}

func getInfoRepo(topics []string) {
	var a []int

	repos, err := getReposFromTopics(topics)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", "Repositories:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Name", "Repository URL", "Description", "Star"})
	table.SetRowLine(true)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
	)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	for i, repo := range repos {
		a = append(a, repo.StargazersCount)
		data := [][]string{{
			strconv.Itoa(i),
			repo.Name,
			repo.RepoUrl,
			repo.Description,
			strconv.Itoa(repo.StargazersCount),
		}}

		for _, v := range data {
			table.Append(v)
		}
	}

	table.Render()

	for i := 0; i < len(repos); i++ {
		if repos[i].StargazersCount == max(a) {
			fmt.Println()
			color.HiRed("- Dumaaaaaaaaaaa: " + color.HiGreenString(repos[i].RepoUrl) + " - " + color.HiYellowString(strconv.Itoa(repos[i].StargazersCount)))
			break
		}
	}

}
