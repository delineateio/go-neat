package ui

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	e "github.com/delineateio/go-neat/errors"
	"github.com/fatih/color"
)

const tick = "\u2713"
const skip = "\u2012"

func Successful(text string, elem ...any) {
	text = strings.ToLower(fmt.Sprintf(text, elem...))
	green := color.New(color.FgGreen).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	fmt.Println(green.Sprint(tick), white.Sprint(text))
}

func Skipped(text string, elem ...any) {
	text = strings.ToLower(fmt.Sprintf(text, elem...))
	yellow := color.New(color.FgYellow).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	fmt.Println(yellow.Sprint(skip), white.Sprint(text))
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
