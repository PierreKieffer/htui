package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func Auth() error {

	if os.Getenv("HEROKU_API_KEY") == "" {

		home, err := os.UserHomeDir()

		if err != nil {
			return errors.New(fmt.Sprintf("ERROR : Auth : %v", err.Error()))
		}

		netrc := fmt.Sprintf("%v/.netrc", home)

		_, err = os.Stat(netrc)
		if err != nil {
			return errors.New(fmt.Sprintf("ERROR : Auth : %v", err.Error()))
		}

		token, err := ExtractHerokuToken(netrc)
		if err != nil {
			return errors.New(fmt.Sprintf("ERROR : Auth : %v", err.Error()))
		}

		os.Setenv("HEROKU_API_KEY", token)
	}

	return nil
}

func ReadFile(filePath string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ERROR : ReadFile : %v", err.Error()))
	}
	data := strings.Split(string(fileBytes), "\n")
	return data, nil
}

func ExtractHerokuToken(netrc string) (string, error) {

	var token string

	netrcData, err := ReadFile(netrc)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ERROR : ExtractHerokuToken : %v", err.Error()))
	}

	for _, v := range netrcData {
		tokenMatch, _ := regexp.MatchString("password", v)
		if tokenMatch == true {
			vSplit := strings.Split(v, " ")
			for _, v := range vSplit {
				if v != "password" && len(v) > 5 {
					token = v
					break
				}
			}
			break
		}
	}

	return token, nil
}

func main() {
	Auth()
}
