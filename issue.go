package GoJira

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Session) Issue(issueKey string) (issue *Issue, err error) {
	req, err := http.NewRequest("GET", s.url+"/issue/"+issueKey, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.username, s.password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var i Issue
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&i)
	if err != nil {
		fmt.Println(err)
	}
	return &i, nil

}

type Issue struct {
	Id     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields map[string]interface{}
	/*
		Implement this later?
		Fields struct {
			Summary  string `json:"summary"`
			Progress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
				Percent  int `json:"percent"`
			}
			TimeTracking struct {
				OriginalEstimate         string `json:"originalEstimate"`
				RemainingEstimate        string `json:"remainingEstimate"`
				OriginalEstimateSeconds  int    `json:"originalEstimateSeconds"`
				RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
			}
			IssueType struct {
				Self        string `json:"self"`
				Id          string `json:"id"`
				Description string `json:"description"`
				IconUrl     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
			}
			TimeSpent int `json:"timespent"`
			.
			.
			.
		}
	*/
}
