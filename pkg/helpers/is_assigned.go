package helpers

import (
	"mr-bot/pkg/datastruct"
	"strings"
)

func IsAssigned(comments []*datastruct.Comment) bool {
	for _, comment := range comments {
		if strings.HasPrefix(comment.Text, "Reviewers:") {
			return true
		}
	}
	return false
}
