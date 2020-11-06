package helpers

import (
	"fmt"
	"mr-bot/pkg/constants"
	"time"
)

func CreateGetMRsRequest(projectId, token string, now time.Time) string {
	now = now.Add(-constants.TimeDelta)
	return fmt.Sprintf(
		"https://gitlab.com/api/v4/projects/%s/merge_requests?private_token=%s&view=simple&created_after=%s",
		projectId, token, now.Format("2006-1-2T15:04:05Z"),
	)
}
