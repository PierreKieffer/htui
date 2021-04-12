package core

import (
	"errors"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/api"
)

func StreamLogs(appName string, logsBuffer chan string, signal chan bool) error {
	/*
	 */

	var apiSignal = make(chan bool)

	payload := `{"dyno" : "", "lines" : 300, "source" : "app", "tail" : true}`

	logSessionUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/log-sessions", appName)

	logSession, err := api.PostRequest(logSessionUrl, payload)

	if err != nil {
		return errors.New(fmt.Sprintf("ERROR : StreamLogs : %v", err.Error()))
	}

	if logSession.StatusCode != 201 {
		return errors.New(fmt.Sprintf("ERROR : StreamLogs : %v", logSession.StatusCode))
	}

	logplexUrl := logSession.Body.(map[string]interface{})["logplex_url"]

	go api.StreamRequest(logplexUrl.(string), logsBuffer, apiSignal)

	select {
	case <-signal:
		apiSignal <- true
		fmt.Println("end of StreamLogs")
		return nil
	}

	return nil
}
