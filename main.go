package main

import "fmt"

var baseUrl = "https://dnreturn.com/"

func main() {
	err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	Run()
	<-make(chan struct{})
}
