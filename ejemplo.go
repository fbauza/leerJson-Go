package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Datos struct {
	success bool
	payload pay
}

type pay struct {
	id         int32
	partner_id int32
}

func main() {

	file, _ := os.Open("plates.JSON")
	defer file.Close()

	decode := json.NewDecoder(file)
	account := Datos{}
	err := decode.Decode(&account)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account)
}
