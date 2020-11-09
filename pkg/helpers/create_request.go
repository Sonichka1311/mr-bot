package helpers

import (
	"fmt"
	"net/url"
)

const GitlabAPI = "https://gitlab.com/api/v4/projects/"

func CreateGetMRsRequest(projectId, token string) string {
	return fmt.Sprintf(
		"%s%s/merge_requests?private_token=%s&state=opened&wip=no",
		GitlabAPI, projectId, token,
	)
}

func CreateGetCommentsRequest(projectId, token, mrId string) string {
	return fmt.Sprintf(
		"%s%s/merge_requests/%s/notes?private_token=%s",
		GitlabAPI, projectId, mrId, token,
	)
}

func CreateAddCommentRequest(projectId, token, mrId, duty string) string {
	return fmt.Sprintf(
		"%s%s/merge_requests/%s/notes?private_token=%s&body=%s",
		GitlabAPI, projectId, mrId, token, url.QueryEscape(fmt.Sprintf("Reviewers: @%s", duty)),
	)
}