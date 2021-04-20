package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/api"
	"net/http"
	"os"
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

type Formation struct {
	Id        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updatedAt"`
	Command   string  `json:"command"`
	Size      string  `json:"size"`
	Quantity  float64 `json:"quantity"`
	Type      string  `json:"type"`
	App       struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"app"`
}

func GetAppFormation(appName string) ([]Formation, error) {
	/*
	 */
	var formations []Formation

	formationsListUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/formation", appName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", formationsListUrl, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAppFormation : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAppFormation : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		err := json.NewDecoder(resp.Body).Decode(&formations)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("ERROR : GetAppFormation : %v", err.Error()))
		}

		return formations, nil
	}

	return formations, errors.New(fmt.Sprintf("ERROR : GetAppFormation : status code %v", resp.StatusCode))
}

func GetFormationInfo(appName string, formationType string) (Formation, error) {
	var formation Formation

	formationUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/formation/%v", appName, formationType)

	client := &http.Client{}
	req, err := http.NewRequest("GET", formationUrl, nil)
	if err != nil {
		return formation, errors.New(fmt.Sprintf("ERROR : GetFormationInfo : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return formation, errors.New(fmt.Sprintf("ERROR : GetFormationInfo : %v", err.Error()))
	}

	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&formation)
		if err != nil {
			return formation, errors.New(fmt.Sprintf("ERROR : GetFormationInfo : %v", err.Error()))
		}

		return formation, nil
	}

	return formation, errors.New(fmt.Sprintf("ERROR : GetFormationInfo : status code %v", resp.StatusCode))
}

func UpdateFormationQuantity(appName string, formationType string, quantity int) (string, error) {

	updateFormationUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/formation", appName)

	payload := fmt.Sprintf(`{"updates" : [{"quantity" : %v,"type" : "%v"}]}`, quantity, formationType)

	resp, err := api.PatchRequest(updateFormationUrl, payload)

	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : UpdateFormation : %v", err.Error()))
	}

	if resp.StatusCode != 200 {
		return fmt.Sprintf("ERROR : UpdateFormation : %v, %v", resp.StatusCode, resp.Body.(map[string]interface{})), nil
	}

	responsePayload, err := json.MarshalIndent(resp.Body.([]interface{}), "", "    ")
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : UpdateFormation : %v", err.Error()))
	}

	return string(responsePayload), nil
}

func UpdateFormationSize(appName string, formationType string, size string) (string, error) {

	updateFormationUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/formation", appName)

	payload := fmt.Sprintf(`{"updates" : [{"size" : "%v","type" : "%v"}]}`, size, formationType)

	resp, err := api.PatchRequest(updateFormationUrl, payload)

	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : UpdateFormation : %v", err.Error()))
	}

	if resp.StatusCode != 200 {
		return fmt.Sprintf("ERROR : UpdateFormation : %v, %v", resp.StatusCode, resp.Body.(map[string]interface{})), nil
	}

	responsePayload, err := json.MarshalIndent(resp.Body.([]interface{}), "", "    ")
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : UpdateFormation : %v", err.Error()))
	}

	return string(responsePayload), nil
}

func GetAppDynos(appName string) ([]Dyno, error) {
	/*
	 */
	var dynos []Dyno

	dynosListUrl := fmt.Sprintf("https://api.heroku.com/apps/%v/dynos", appName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", dynosListUrl, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAppDynos : %v", err.Error()))
	}
	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAppDynos : %v", err.Error()))
	}

	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&dynos)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("ERROR : GetAppDynos : %v", err.Error()))
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

	client := &http.Client{}
	req, err := http.NewRequest("GET", dynoInfoUrl, nil)
	if err != nil {
		return dyno, errors.New(fmt.Sprintf("ERROR : GetDynoInfo : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return dyno, errors.New(fmt.Sprintf("ERROR : GetDynoInfo : %v", err.Error()))
	}

	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&dyno)
		if err != nil {
			return dyno, errors.New(fmt.Sprintf("ERROR : GetDynoInfo : %v", err.Error()))
		}

		return dyno, nil
	}

	return dyno, errors.New(fmt.Sprintf("ERROR : GetDynoInfo : status code %v", resp.StatusCode))
}

func GetDynoState(appName string, dynoName string) (string, error) {

	dyno, err := GetDynoInfo(appName, dynoName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : GetDynoState : %v", err.Error()))
	}

	dynoState := dyno.State
	return dynoState, nil
}
