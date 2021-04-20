package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type App struct {
	Acm                          bool   `json:"acm"`
	ArchivedAt                   string `json:"archived_at"`
	BuildpackProvidedDescription string `json:"buildpack_provided_description"`
	BuildStack                   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"build_stack"`
	CreatedAt       string `json:"created_at"`
	GitURL          string `json:"git_url"`
	ID              string `json:"id"`
	InternalRouting bool   `json:"internal_routing"`
	Maintenance     bool   `json:"maintenance"`
	Name            string `json:"name"`
	Owner           struct {
		Email string `json:"email"`
		ID    string `json:"id"`
	} `json:"owner"`
	Organization struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"organization"`
	Team struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"team"`
	Region struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"region"`
	ReleasedAt string `json:"released_at"`
	RepoSize   int    `json:"repo_size"`
	SlugSize   int    `json:"slug_size"`
	Space      struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Shield bool   `json:"shield"`
	} `json:"space"`
	Stack struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"stack"`
	UpdatedAt string `json:"updated_at"`
	WebURL    string `json:"web_url"`
}

func GetApps() ([]App, error) {
	/*
	 */

	var apps []App

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.heroku.com/apps", nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetApps : %v", err.Error()))
	}
	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetApps : %v", err.Error()))
	}

	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&apps)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("ERROR : GetApps : %v", err.Error()))
		}

		return apps, nil
	}

	return apps, errors.New(fmt.Sprintf("ERROR : GetApps : status code %v", resp.StatusCode))
}

func GetAppInfo(appName string) (App, error) {
	/*
	 */

	var app App

	appInfoUrl := fmt.Sprintf("https://api.heroku.com/apps/%v", appName)
	client := &http.Client{}
	req, err := http.NewRequest("GET", appInfoUrl, nil)
	if err != nil {
		return app, errors.New(fmt.Sprintf("ERROR : GetAppInfo : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return app, errors.New(fmt.Sprintf("ERROR : GetAppInfo : %v", err.Error()))
	}

	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&app)
		if err != nil {
			return app, errors.New(fmt.Sprintf("ERROR : GetAppInfo : %v", err.Error()))
		}

		return app, nil
	}

	return app, errors.New(fmt.Sprintf("ERROR : GetAppInfo : status code %v", resp.StatusCode))
}
