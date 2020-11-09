package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mr-bot/pkg/datastruct"
	"net/http"
)

func GetMRsFromResponse(resp *http.Response, err error) []*datastruct.MR {
	if err != nil {
		log.Printf("Err: %s\n", err.Error())
		return nil
	} else if resp == nil {
		log.Printf("Body is nil")
		return nil
	}
	defer resp.Body.Close()
	jsonResp, _ := ioutil.ReadAll(resp.Body)

	var mrs []*datastruct.MR
	json.Unmarshal(jsonResp, &mrs)
	return mrs
}

func GetCommentsFromResponse(resp *http.Response, err error) []*datastruct.Comment {
	if err != nil {
		log.Printf("Err: %s\n", err.Error())
		return nil
	} else if resp == nil {
		log.Printf("Body is nil")
		return nil
	}
	defer resp.Body.Close()
	jsonResp, _ := ioutil.ReadAll(resp.Body)

	var comments []*datastruct.Comment
	json.Unmarshal(jsonResp, &comments)
	return comments
}