package promptwrapper

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type CommitType struct {
	ID   int
	Type string
	Name string
}

var Scopes = []CommitType{
	{0, "build", "build: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"},
	{1, "ci", "ci: Changes to our CI configuration files and scripts (examples: CircleCi, SauceLabs)"},
	{2, "docs", "docs: Documentation only changes"},
	{3, "feat", "feat: A new feature"},
	{4, "fix", "fix: A bug fix"},
	{5, "perf", "perf: A code change that improves performance"},
	{6, "refactor", "refactor: A code change that neither fixes a bug nor adds a feature"},
	{7, "test", "test: Adding missing tests or correcting existing tests"},
}

func SelectPrompt() promptui.Select {
	prompt := promptui.Select{
		HideHelp:     true,
		HideSelected: true,
		Label:        promptui.Styler(promptui.FGBlue)("👹 Oni is here to help you. Select the type of your commit:"),
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

var Validate = func(selected CommitType) func(string) error {
	return func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("message cannot be empty")
		}
		if len(input)+len(selected.Type)+2 > 72 {
			return fmt.Errorf("message cannot exceed 72 characters")
		}
		return nil
	}
}

func TypePrompt(selected CommitType) promptui.Prompt {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:       "😁 Enter commit message (up to 72 characters):",
		HideEntered: true,
		Templates:   templates,
		Validate:    Validate(selected),
	}

	return prompt
}
