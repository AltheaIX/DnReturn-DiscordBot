package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestDoRegister(t *testing.T) {
	payload := "suis!prefix HaelaCi:Haela123:777777:HaelaChiii@gmail.com"
	regex, _ := regexp.Compile(`[^\s]*$`)
	regexPayload := regex.FindAllString(payload, -1)
	splitPayload := strings.Split(regexPayload[0], ":")
	username, password, pin, email := splitPayload[0], splitPayload[1], splitPayload[2], splitPayload[3]
	register := doRegister(username, password, pin, email)
	fmt.Println(register)
	fmt.Println(splitPayload.len())
}
