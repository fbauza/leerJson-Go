package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	clientHttp := &http.Client{}

	url := "http://localhost:3300/vehicles/plate?plate=FEU0807"
	peticion, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("Error creando peticion %v", err)
	}

	// Podemos agregar encabezados
	peticion.Header.Add("Content-Type", "application/json")
	//peticion.Header.Add("X-Hola-Mundo", "Ejemplo")
	respuesta, err := clientHttp.Do(peticion)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		log.Fatalf("Error haciendo petición: %v", err)
	}
	// No olvides cerrar el cuerpo al terminar
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Encabezados: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("El tipo de contenido: '%s'", contentType)
	// Aquí puedes decodificar la respuesta si es un JSON, o convertirla a cadena
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
}
