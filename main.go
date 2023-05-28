/*
Copyright © 2023 Anindya Chowdhury <anindya.chowdhury@protonmail.com>
*/
package main

import (
	"oni-commit/cmd"
	"oni-commit/utils/console"
	"os/exec"
)

func main() {
	err := exec.Command("git", "rev-parse", "--git-dir").Run()
	if err != nil { // if error is not nil, then it is not a git repository
		console.BoldBgRed("Not a Git repository")
		return
	}

	cmd.Execute()
}
