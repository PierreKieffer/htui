package core

import (
	"errors"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/api"
)

type App struct {
	Id           string `json:"id"`
	CreatedAt    string `json:"createdAt"`
	ReleasedAt   string `json:"releasedAt"`
	UpdatedAt    string `json:"updatedAt"`
	Organization string `json:"organization"`
	Team         string `json:"team"`
	WebUrl       string `json:"webUrl"`
	Name         string `json:"name"`
	Owner        string `json:"owner"`
	Region       string `json:"region"`
}

func GetApps() ([]App, error) {
	/*
	 */

	var apps []App

	resp, err := api.GetRequest("https://api.heroku.com/apps")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetApps : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		for _, rawItem := range resp.Body.([]interface{}) {
			item := rawItem.(map[string]interface{})
			app := App{
				Id:           ParseItem(item["id"]),
				CreatedAt:    ParseItem(item["created_at"]),
				ReleasedAt:   ParseItem(item["released_at"]),
				UpdatedAt:    ParseItem(item["updated_at"]),
				Organization: ParseItem(item["organization"]),
				Team:         ParseItem(item["team"]),
				WebUrl:       ParseItem(item["web_url"]),
				Name:         ParseItem(item["name"]),
				Owner:        ParseItem(item["owner"].(map[string]interface{})["email"]),
				Region:       ParseItem(item["region"].(map[string]interface{})["name"]),
			}

			apps = append(apps, app)

		}

		return apps, nil
	}

	return apps, errors.New(fmt.Sprintf("ERROR : GetApps : status code %v", resp.StatusCode))
}

func GetAppInfo(appName string) (App, error) {
	/*
	 */

	var app App

	resp, err := api.GetRequest(fmt.Sprintf("https://api.heroku.com/apps/%v", appName))

	if err != nil {
		return app, errors.New(fmt.Sprintf("ERROR : GetApps : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		item := resp.Body.(map[string]interface{})
		app = App{
			Id:           ParseItem(item["id"]),
			CreatedAt:    ParseItem(item["created_at"]),
			ReleasedAt:   ParseItem(item["released_at"]),
			UpdatedAt:    ParseItem(item["updated_at"]),
			Organization: ParseItem(item["organization"]),
			Team:         ParseItem(item["team"]),
			WebUrl:       ParseItem(item["web_url"]),
			Name:         ParseItem(item["name"]),
			Owner:        ParseItem(item["owner"].(map[string]interface{})["email"]),
			Region:       ParseItem(item["region"].(map[string]interface{})["name"]),
		}

		return app, nil
	}

	return app, errors.New(fmt.Sprintf("ERROR : GetApps : status code %v", resp.StatusCode))
}
