package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "branch", "--format=%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	branches := strings.Split(string(output), "\n")
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if branch != "" && branch != "main" {
			fmt.Println("Deleting branch:", branch)
			deleteCmd := exec.Command("git", "branch", "-D", branch)
			deleteOutput, deleteErr := deleteCmd.Output()
			if deleteErr != nil {
				fmt.Println("Error deleting branch:", branch, deleteErr)
			} else {
				fmt.Println(string(deleteOutput))
			}
		}
	}
}
