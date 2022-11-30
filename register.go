package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Register struct {
	Data string `json:"data"`
}

type NameGen struct {
	Username string `json:"username"`
	Password string `json:"email_u"`
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

func doRandomRegister() (string, string) {
	rand.Seed(time.Now().UnixNano())
	var result string
	fakeName := nameGenApi()
	caser := cases.Title(language.English)
	username := caser.String(fakeName.Username)
	_password := fmt.Sprintf("%v%v", caser.String(fakeName.Password), (rand.Intn(9999))+100)
	password := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(_password, "")
	email := fmt.Sprintf("%v@gmail.com", caser.String(fakeName.Password))
	pin := fmt.Sprintf("%v", rand.Intn(1000000-100000)+100000)
	payload := fmt.Sprintf("%v:%v:%v:%v", username, password, pin, email)

	result = doRegister(username, password, pin, email)
	return result, payload
}

func nameGenApi() NameGen {
	client := &http.Client{}
	data := NameGen{}
	resp, err := client.Get("https://api.namefake.com/")
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	return data
}
