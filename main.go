package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\x1b[34;1m%s\x1b[0m", "Enter topics: ")
	text, _ := reader.ReadString('\n')
	topics := strings.Split(strings.TrimSpace(text), ",")

	getInfoRepo(topics)

	var choice int
	fmt.Println()
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "Enter your choices:")
	fmt.Println("[1] Clone all repos <Be careful with this command>")
	fmt.Println("[2] Clone repo with the highest score")
	fmt.Println("[3] End ...")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		CloneGitUrl(topics)
	case 2:
		CloneGitUrlMaxStar(topics)
	default:
		break
	}
}
