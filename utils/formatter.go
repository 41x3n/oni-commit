package utils

import (
	"github.com/fatih/color"
)

func BoldCyan(text string) {
	color.New(color.FgCyan, color.Bold).Println(text)
}
