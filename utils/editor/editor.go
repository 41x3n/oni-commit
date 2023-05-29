package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func openEditor(filename string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("vi", filename)
	case "darwin":
		cmd = exec.Command("vi", filename)
	default:
		cmd = exec.Command("vi", filename)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func TakeBodyAsInput() ([]byte, error) {
	// Create a temporary file to hold the user's input
	tmpFile, err := ioutil.TempFile("", "input")
	if err != nil {
		fmt.Println("Failed to create temporary file:", err)
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	// Open the temporary file with the default editor based on the OS
	err = openEditor(tmpFile.Name())
	if err != nil {
		fmt.Println("Failed to open editor:", err)
		return nil, err
	}

	// Read the contents of the temporary file
	inputData, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		fmt.Println("Failed to read temporary file:", err)
		return nil, err
	}

	return inputData, nil

}
