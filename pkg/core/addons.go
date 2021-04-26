package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Addon struct {
	ID         string        `json:"id"`
	CreatedAt  string        `json:"created_at"`
	UpdatedAt  string        `json:"updated_at"`
	ConfigVars []interface{} `json:"config_vars"`
	Name       string        `json:"name"`
	ProviderID string        `json:"provider_id"`
	State      string        `json:"state"`
	WebURL     string        `json:"web_url"`
	Actions    []struct {
		ID            string `json:"id"`
		Label         string `json:"label"`
		Action        string `json:"action"`
		URL           string `json:"url"`
		RequiresOwner bool   `json:"requires_owner"`
	} `json:"actions"`
	AddonService struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"addon_service"`
	BillingEntity struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"billing_entity"`
	App struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"app"`
	BilledPrice struct {
		Cents    float64 `json:"cents"`
		Contract bool    `json:"contract"`
		Unit     string  `json:"unit"`
	} `json:"billed_price"`
	Plan struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"plan"`
}

func GetAddons() ([]Addon, error) {
	var addons []Addon

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.heroku.com/addons", nil)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAddons : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : GetAddons : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		err := json.NewDecoder(resp.Body).Decode(&addons)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("ERROR : GetAddons : %v", err.Error()))
		}

		return addons, nil
	}

	return addons, errors.New(fmt.Sprintf("ERROR : GetAddons : status code %v", resp.StatusCode))

}

func GetAddonInfo(addonName string) (Addon, error) {
	var addon Addon

	client := &http.Client{}

	addonInfoUrl := fmt.Sprintf("https://api.heroku.com/addons/%v", addonName)
	req, err := http.NewRequest("GET", addonInfoUrl, nil)

	if err != nil {
		return addon, errors.New(fmt.Sprintf("ERROR : GetAddonInfo : %v", err.Error()))
	}

	req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("HEROKU_API_KEY")))

	resp, err := client.Do(req)
	if err != nil {
		return addon, errors.New(fmt.Sprintf("ERROR : GetAddonInfo : %v", err.Error()))
	}

	if resp.StatusCode == 200 {

		err := json.NewDecoder(resp.Body).Decode(&addon)
		if err != nil {
			return addon, errors.New(fmt.Sprintf("ERROR : GetAddonInfo : %v", err.Error()))
		}

		return addon, nil
	}

	return addon, errors.New(fmt.Sprintf("ERROR : GetAddonInfo : status code %v", resp.StatusCode))

}
