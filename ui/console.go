package ui

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	e "github.com/delineateio/go-neat/errors"
	"github.com/fatih/color"
)

const TICK = "\u2713"
const SKIP = "\u2012"

func Successful(text string, elem ...any) {
	text = strings.ToLower(fmt.Sprintf(text, elem...))
	green := color.New(color.FgGreen).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	fmt.Println(green.Sprint(TICK), white.Sprint(text))
}

func Checklist(question string, options []string) []string {
	prompt := &survey.MultiSelect{
		Message: strings.ToLower(question),
		Help:    "Branches",
		Options: options,
	}
	var answers []string
	err := survey.AskOne(prompt, &answers)
	e.CheckIfError(err, "failed to prompt the user")
	return answers
}

func Skipped(text string, elem ...any) {
	text = strings.ToLower(fmt.Sprintf(text, elem...))
	yellow := color.New(color.FgYellow).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	fmt.Println(yellow.Sprint(SKIP), white.Sprint(text))
}
