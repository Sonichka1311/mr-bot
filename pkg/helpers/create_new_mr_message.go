package helpers

import (
	"fmt"
	"mr-bot/pkg/datastruct"
	"strings"
)

func CreateNewMRMessage(mr *datastruct.MR) string {
	return strings.ReplaceAll(fmt.Sprintf(
		"Новый merge request (%s)\n%s\n*Ответственный:* %s",
		mr.Title, mr.Link, GetDuty(),
	), "_", "\\_")
}
