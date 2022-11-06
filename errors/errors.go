package errors

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
)

const cross = "\u2718"

func NewErr(text string, elem ...string) {
	err := fmt.Errorf(text, elem)
	logAndExit(err, text)
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
	os.Exit(0)
}

func failed(text string) {
	red := color.New(color.FgRed).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	fmt.Println(red.Sprint(cross), white.Sprint(text))
}
