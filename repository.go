package main

type Repository struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`

	RepoUrl         string `json:"html_url"`
	CloneUrl        string `json:"clone_url"`
	StargazersCount int    `json:"stargazers_count"`
	DefaultBranch   string `json:"default_branch"`
}

type SearchResult struct {
	Items []Repository `json:"items"`
}
