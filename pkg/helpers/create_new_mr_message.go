package helpers

import (
	"fmt"
	"mr-bot/pkg/datastruct"
	"strings"
)

func CreateNewMRMessage(mr *datastruct.MR, duty string) string {
	return strings.ReplaceAll(fmt.Sprintf(
		"Новый merge request (%s)\n%s\n*Ответственный:* @%s",
		mr.Title, mr.Link, duty,
	), "_", "\\_")
}
