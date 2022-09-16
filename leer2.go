package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PayloadPlate struct {
	Id        int `json:"id"`
	PartnerId int `json:"partner_id"`
}

type Plate struct {
	Estado  bool           `json:"success"`
	Payload []PayloadPlate `json:"payload"`
}

func main() {

	fp, err := os.ReadFile("plates.JSON")
	if err != nil {
		log.Fatal(err)
	}

	var plates Plate
	json.NewDecoder(bytes.NewBuffer(fp)).Decode(&plates)

	fmt.Printf("%#v", plates)
	//fmt.Println(plates.Payload)
	//fmt.Println(plates.Estado)

}
