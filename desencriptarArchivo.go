package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func leerArchivo(ruta string) string {

	content, err := ioutil.ReadFile(ruta)
	if err != nil {
		log.Fatal(err)
	}
	respuesta := string(content)
	return respuesta

}

func decodeB64(contenidoArchivo string) string {
	data, err := base64.StdEncoding.DecodeString(contenidoArchivo)
	if err != nil {
		fmt.Println("error:", err)
		return "failed"
	}
	respuesta := string(data)
	return respuesta
}

func main() {

	ruta := "b64.txt"
	cantidad := 50
	contenidoArchivo := leerArchivo(ruta)

	decodificado := contenidoArchivo

	for i := 0; i < cantidad; i++ {
		decodificado = decodeB64(decodificado)
	}
	fmt.Print("\n La respuesta final es: \n")
	fmt.Print(decodificado)

}
