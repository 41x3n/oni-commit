/*
Copyright © 2023 Anindya Chowdhury <anindya.chowdhury@protonmail.com>
*/
package main

import (
	"oni-commit/cmd"
	"oni-commit/utils/console"
	"os/exec"
	"strings"
)

func main() {
	err := exec.Command("git", "rev-parse", "--git-dir").Run()
	if err != nil { // if error is not nil, then it is not a git repository
		console.BoldBgRed("Not a Git repository")
		return
	}

	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		console.BoldBgRed("Error running git status:")
		return
	}

	lines := strings.Split(string(out), "\n")
	allModifiedFilesStaged := true

	for _, line := range lines {
		if strings.HasPrefix(line, " M") { // if file is modified but not staged
			allModifiedFilesStaged = false
			break
		}
		if strings.HasPrefix(line, "MM") { // if file is staged and then modified
			allModifiedFilesStaged = false
			break
		}
	}

	if !allModifiedFilesStaged {
		console.BoldBgRed("Not all modified files are staged.")
		return
	}

	cmd.Execute()
}
