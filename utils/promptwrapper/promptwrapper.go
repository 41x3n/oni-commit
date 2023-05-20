package promptwrapper

import (
	"github.com/manifoldco/promptui"
)

type CommitType struct {
	ID   int
	Name string
}

var Scopes = []CommitType{
	{0, "build: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"},
	{1, "ci: Changes to our CI configuration files and scripts (examples: CircleCi, SauceLabs)"},
	{2, "docs: Documentation only changes"},
	{3, "feat: A new feature"},
	{4, "fix: A bug fix"},
	{5, "perf: A code change that improves performance"},
	{6, "refactor: A code change that neither fixes a bug nor adds a feature"},
	{7, "test: Adding missing tests or correcting existing tests"},
}

func Prompt() promptui.Select {
	promptui.IconInitial = "👹"
	prompt := promptui.Select{
		HideHelp:     true,
		HideSelected: true,
		Label:        promptui.Styler(promptui.FGBlue)("Oni is here to help you. Select the type of your commit:"),
		Items:        Scopes,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Name }}",
			Active:   "-> {{ .Name | cyan }}",
			Inactive: "   {{ .Name }}",
			Selected: "{{ .Name | green }}",
		},
	}

	return prompt
}
