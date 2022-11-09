package apps_repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	users "vessl/modules/users"
)

type Response struct {
	Version   string     `json:"version"`
	Templates []Template `json:"templates"`
}

type Template struct {
	Type           int      `json:"type"`
	Title          string   `json:"title"`
	Name           string   `json:"name"`
	Hostname       string   `json:"hostname"`
	Description    string   `json:"description"`
	Info_url       string   `json:"info_url"`
	Categories     []string `json:"categories"`
	Platform       string   `json:"platform"`
	Logo           string   `json:"logo"`
	Image          string   `json:"image"`
	Restart_policy string   `json:"restart_policy"`
	Network        string   `json:"network"`
	Env            []string `json:"env"`
	Cmd            []string `json:"cmd"`
	Ports          []string `json:"ports"`
	Volumes        []string `json:"volumes"`
}

func GetApps(UserId string) Response {

	var responseObject Response

	url := users.GetAppsRepositoryUrl(UserId)

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return Response{}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return Response{}
	}

	json.Unmarshal(body, &responseObject)
	return responseObject
}
