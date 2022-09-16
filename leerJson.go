package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PhotoPlate struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type GenericPlate struct {
	Id      int         `json:"id"`
	Name    interface{} `json:"name"`
	Country int         `json:"country"`
}

type TravelledPlate struct {
	Total int    `json:"total"`
	Unit  string `json:"unit"`
}

type PricePlate struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

type PayloadPlate struct {
	Id        int            `json:"id"`
	PartnerId int            `json:"partner_id"`
	Hash      int            `json:"hash"`
	Entity    int            `json:"entity"`
	Plate     string         `json:"plate"`
	Stock     string         `json:"stock"`
	Doors     int            `json:"doors"`
	Armored   int            `json:"armored"`
	Engine    interface{}    `json:"engine"`
	Status    string         `json:"status"`
	VideoUrl  string         `json:"videourl"`
	Price     PricePlate     `json:"price"`
	Travelled TravelledPlate `json:"travelled"`
	Make      GenericPlate   `json:"make"`
	Model     GenericPlate   `json:"model"`
	Version   GenericPlate   `json:"version"`
	Year      struct {
		Assembly int `json:"assembly"`
		Model    int `json:"model"`
	} `json:"year"`
	Fuel         GenericPlate   `json:"fuel"`
	Transmission GenericPlate   `json:"transmission"`
	Color        GenericPlate   `json:"color"`
	BodyType     GenericPlate   `json:"body_type"`
	PhotoHash    int32          `json:"photo_hash"`
	Photo        []PhotoPlate   `json:"photos"`
	OptionalHash int32          `json:"optional_hash"`
	Optionals    []GenericPlate `json:"optionals"`
	ItemsHash    int            `json:"items_hash"`
}

//Price     PricePlate  `json:"price"`

type Plate struct {
	Estado  bool           `json:"success"`
	Payload []PayloadPlate `json:"payload"`
}

func main() {

	fp, err := ioutil.ReadFile("plates.JSON")
	if err != nil {
		log.Fatal(err)
	}

	var plates Plate

	err = json.Unmarshal(fp, &plates)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%#v", plates)
	//fmt.Printf("%#v", plates.Payload[0])
	for index := range plates.Payload {
		fmt.Printf("%#v", plates.Payload[index])
	}
	/*
		for index := range plates.Payload {
			fmt.Printf("%#v", plates.Payload[index].Model.Name)
		}
	*/
}
