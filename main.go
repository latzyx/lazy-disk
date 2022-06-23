package main

import (
	"fmt"
	"lazy/config"
)

func main() {
	a := config.Client{
		Corpid:     "",
		CorPsecret: "",
	}
	fmt.Println(a)
}
