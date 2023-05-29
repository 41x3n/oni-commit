package gitcommand

import (
	"fmt"
	"os"
	"os/exec"
)

func CommitIt(commitMessage, commitBody string) {
	// Create the git commit command
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	if commitBody != "" {
		cmd.Args = append(cmd.Args, "-m", commitBody)
	}

	// Set the command's standard input, output, and error to the current process's standard streams
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running git commit:", err)
		return
	}

	fmt.Println("Git commit completed successfully.")
}
