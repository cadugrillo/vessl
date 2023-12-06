package apps_repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	db "vessl/modules/database"
)

type Response struct {
	Version   string        `json:"version"`
	Templates []db.Template `json:"templates"`
}

func GetApps(UserId string) Response {

	var responseObject Response
	var Client = &http.Client{
		Timeout: time.Second * 3,
	}

	url, udefTemplates := db.GetAppsTemplates(UserId)

	req, _ := http.NewRequest("GET", url, nil)

	//res, err := http.DefaultClient.Do(req)
	res, err := Client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return Response{Templates: udefTemplates}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return Response{Templates: udefTemplates}
	}

	json.Unmarshal(body, &responseObject)
	for i := 0; i < len(udefTemplates); i++ {
		responseObject.Templates = append(responseObject.Templates, udefTemplates[i])
	}
	return responseObject
}
