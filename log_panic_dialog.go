package main

import (
	"log"

	"github.com/sqweek/dialog"
)

type LogDialog struct{}

func (lpd LogDialog) Panicln(title string, message string) {
	if title == "" {
		title = "Error"
	}

	dialog.Message(message).Title(title).Error()
	log.Panicln(message)
}
