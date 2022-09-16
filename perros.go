package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Raza struct {
		Nombre, Pais string
	}

	type Mascota struct {
		Nombre string
		Edad   int
		Raza   Raza
		Amigos []string // Arreglo de strings
	}

	// Vamos a probar...
	mascotaComoJson := []byte(`{"Nombre":"Maggie","Edad":3,"Raza":{"Nombre":"Caniche","Pais":"Francia"},"Amigos":["Bichi","Snowball","Coqueta","Cuco","Golondrino"]}`)

	// Recuerda, primero se define la variable
	var mascota Mascota

	// Y luego se manda su dirección de memoria
	err := json.Unmarshal(mascotaComoJson, &mascota)
	if err != nil {
		fmt.Printf("Error decodificando: %v\n", err)
	} else {
		// Listo. Ahora podemos imprimir
		fmt.Printf("El nombre: %s\n", mascota.Nombre)
		fmt.Printf("País de Raza: %s\n", mascota.Raza.Pais)
		fmt.Printf("Primer amigo: %v\n", mascota.Amigos)
	}
}
