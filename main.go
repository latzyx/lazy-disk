package main

import (
	"fmt"
	"lazy/server"
)

func main() {
	client := server.Client{
		Corpid:     "",
		CorPsecret: "",
	}
	gettoken := client.Gettoken()

	fmt.Printf("%v", gettoken)

}
