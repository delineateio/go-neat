package errors

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
)

const CROSS = "\u2718"
const NEAT_AUTOMATION = "NEAT_AUTOMATION"

func NewErr(text string, elem ...any) {
	err := fmt.Errorf(text, elem...)
	logAndExit(err, err.Error())
}

func CheckIfError(err error, text string, elem ...any) {
	if err != nil {
		logAndExit(err, fmt.Sprintf(text, elem...))
	}
}

func logAndExit(err error, text string) {
	log.Err(err).Send()
	failed(text)
	fmt.Println()

	value := os.Getenv(NEAT_AUTOMATION)
	isAutomation, _ := strconv.ParseBool(value)
	if isAutomation {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func failed(text string) {
	red := color.New(color.FgRed).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	fmt.Println(red.Sprint(CROSS), white.Sprint(text))
}
