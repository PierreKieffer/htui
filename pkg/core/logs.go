package core

import (
	"errors"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/api"
)

func StreamLogs(appName string) error {
	/*

	 */

	payload := `{"dyno" : "", "lines" : 10000, "source" : "app", "tail" : true}`

	logSessionUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/log-sessions", appName)

	logSession, err := api.PostRequest(logSessionUrl, payload)

	if err != nil {
		return errors.New(fmt.Sprintf("ERROR : StreamLogs : %v", err.Error()))
	}

	if logSession.StatusCode == 201 {
		logplexUrl := logSession.Body.(map[string]interface{})["logplex_url"]

		api.StreamRequest(logplexUrl.(string))
	}

	return nil
}
