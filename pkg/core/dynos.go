package core

import (
	"errors"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/api"
)

type Dyno struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updatedAt"`
	Command   string `json:"command"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	State     string `json:"state"`
	Type      string `json:"type"`
}

func GetAppDynos(appName string) ([]Dyno, error) {
	/*
	 */
	var dynos []Dyno

	dynosListUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/dynos", appName)

	resp, err := api.GetRequest(dynosListUrl)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAppDynos : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		for _, rawItem := range resp.Body.([]interface{}) {
			item := rawItem.(map[string]interface{})
			dyno := Dyno{
				Id:        ParseItem(item["id"]),
				CreatedAt: ParseItem(item["created_at"]),
				UpdatedAt: ParseItem(item["updated_at"]),
				Command:   ParseItem(item["command"]),
				Name:      ParseItem(item["name"]),
				Size:      ParseItem(item["size"]),
				State:     ParseItem(item["state"]),
				Type:      ParseItem(item["type"]),
			}

			dynos = append(dynos, dyno)

		}

		return dynos, nil
	}

	return dynos, errors.New(fmt.Sprintf("ERROR : GetAppDynos : status code %v", resp.StatusCode))
}
