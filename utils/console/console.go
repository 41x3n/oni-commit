package console

import (
	"fmt"

	"github.com/fatih/color"
)

func BoldCyan(text string) {
	color.New(color.FgCyan, color.Bold).Println(text)
}

func BoldBgRed(text string) {
	color.New(color.BgHiRed, color.Bold).Println(text)
}

func PrintSelection(option string, text string) {
	dictionary := make(map[string]string)
	dictionary["selected"] = "Selected type of commit"
	dictionary["message"] = "Your Commit message"

	head := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	tail := color.New(color.FgHiBlue, color.Bold).SprintFunc()

	fmt.Printf("%s : %s\n", head(dictionary[option]), tail(text))

}
