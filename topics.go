package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getReposFromTopics(topics []string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=topic:%s", strings.Join(topics, "+topic:"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResult SearchResult
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return nil, err
	}

	return searchResult.Items, nil
}
