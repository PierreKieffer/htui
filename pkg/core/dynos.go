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

func RestartDyno(appName string, dynoName string) error {

	dynoRestartUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/dynos/%v", appName, dynoName)

	resp, err := api.DeleteRequest(dynoRestartUrl)

	if err != nil {
		return errors.New(fmt.Sprintf("ERROR : RestartDyno : %v", err.Error()))
	}

	if resp.StatusCode == 202 {
		if err != nil {
			return errors.New(fmt.Sprintf("ERROR : RestartDyno : %v", err.Error()))
		}
		return nil
	}

	return errors.New(fmt.Sprintf("ERROR : RestartDyno : status code %v", resp.StatusCode))
}

func GetDynoInfo(appName string, dynoName string) (Dyno, error) {

	var dyno Dyno

	dynoInfoUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/dynos/%v", appName, dynoName)

	resp, err := api.GetRequest(dynoInfoUrl)

	if err != nil {
		return dyno, errors.New(fmt.Sprintf("ERROR : GetDynoInfo : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		item := resp.Body.(map[string]interface{})
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

		return dyno, nil
	}

	return dyno, errors.New(fmt.Sprintf("ERROR : GetDynoInfo : status code %v", resp.StatusCode))

}

func GetDynoState(appName string, dynoName string) (string, error) {

	dyno, err := core.GetDynoInfo(appName, dynoName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : GetDynoState : %v", err.Error()))
	}

	dynoState := dyno.State
	return dynoState, nil
}
