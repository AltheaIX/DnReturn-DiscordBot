package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configStruct struct {
	Token  string
	Prefix string
}

var (
	Token  string
	Prefix string
	config *configStruct
)

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	Prefix = config.Prefix

	return nil
}
