package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Register struct {
	Data string `json:"data"`
}

func doRegister(username string, password string, pin string, email string) string {
	client := &http.Client{}
	var data map[string]interface{}
	var result string
	_payload := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
		"pin1":     pin,
		"DNR22":    "DNReturn",
	}

	payload, _ := json.Marshal(_payload)
	resp, err := client.Post("https://dnreturn.com/api/id/register", "application/json", bytes.NewBuffer(payload))

	if resp.StatusCode != 200 {
		result = fmt.Sprintf("Request failed with response code %d. Please do contact Mika sayang.", resp.StatusCode)
		return result
	}

	if err != nil {
		fmt.Println(err.Error())
		return result
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err.Error())
		return result
	}

	result = data["data"].(string)
	return result
}
