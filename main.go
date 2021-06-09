package main

import (
	"fmt"
	"html/template"
	"net/http"

		"io/ioutil"
		"bytes"
		"encoding/json"
		
		
)
type CadenaAnalizar struct {
 Cadena string

}
func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "")
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/codemirror/", http.StripPrefix("/codemirror/", http.FileServer(http.Dir("codemirror/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/jstree/", http.StripPrefix("/jstree/", http.FileServer(http.Dir("jstree/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/TokensJS", TokensJS)
	http.HandleFunc("/AnalisisJS", AnalisisJS)
	http.HandleFunc("/AnalisisPY", AnalisisPY)
	
	fmt.Printf("Servidor escuchando en: http://localhost:7000/")
	http.ListenAndServe(":7000", nil)
}


func TokensJS(w http.ResponseWriter, r *http.Request) {

	
	ClientHTTP := &http.Client{}

	Peticion := "http://localhost:4001/TraduccionJS"

	Req, Error := http.NewRequest("GET", Peticion, nil)

	if Error != nil {

		print("Error Creando Petición: %v", Error)

	}

	Req.Header.Add("Content-Type", "application/json")
	
	
	Response, Error := ClientHTTP.Do(Req)

	if Error != nil {

		print("Error Haciendo Petición: %v", Error)

	}

	defer Response.Body.Close()

	RespuestaCuerpo, Error := ioutil.ReadAll(Response.Body)

	if Error != nil {

		print("Error no se pudo leer: %v", Error)

	}

	_, _ = fmt.Fprintf(w, "%s", RespuestaCuerpo)}


	func AnalisisJS(w http.ResponseWriter, r *http.Request) {
		
		ClientHTTP := &http.Client{}

		// Decodificador Entrada
		var Decoder = json.NewDecoder(r.Body)

		// Cadena A Analizar
		var Cadena CadenaAnalizar
        
		// Verficar Si Hay Error
		Error := Decoder.Decode(&Cadena)

		// Verificar Si Hay Error
		if Error != nil {

			print("Error Al Codificar La Entrada: %v", Error)

		}
		// Json Para El Request
		JsonFormat, Error := json.Marshal(Cadena)
		
		// Verificar Si Hay Error
		if Error != nil {

			print("Error Al Convertir La Cadena En Json: %v", Error)

		}

		// Peticion NodeJs
		Peticion := "http://localhost:4001/JSS"

		// Solicitar Peticion A NodeJs
		Req, Error := http.NewRequest("POST", Peticion, bytes.NewBuffer(JsonFormat))
        // Verificar Si Hay Errores
		if Error != nil {

			print("Error Creando Petición: %v", Error)

		}

		// Agregar Encabezados
		Req.Header.Add("Content-Type", "application/json")
		// Req.Header.Add("X-Hola-Mundo", "Ejemplo")

		// Respueta NodeJS
		Response, Error := ClientHTTP.Do(Req)

		// Verificcar Si Hay Error
		if Error != nil {

			print("Error Haciendo Petición: %v", Error)

		}

		// Cerrar Cuerpo Response
		defer Response.Body.Close()

		// Cuerpo De La Respuseta
		RespuestaCuerpo, Error := ioutil.ReadAll(Response.Body)

		// Verificar Si Hay Error
		if Error != nil {

			print("Error Leyendo Respuesta: %v", Error)

		}
            
		// Enviar Respuesta
		_, _ = fmt.Fprintf(w, "%s", RespuestaCuerpo)
		

	}

	func AnalisisPY(w http.ResponseWriter, r *http.Request) {
		
		ClientHTTP := &http.Client{}

		// Decodificador Entrada
		var Decoder = json.NewDecoder(r.Body)

		// Cadena A Analizar
		var Cadena CadenaAnalizar
        
		// Verficar Si Hay Error
		Error := Decoder.Decode(&Cadena)

		// Verificar Si Hay Error
		if Error != nil {

			print("Error Al Codificar La Entrada: %v", Error)

		}
		// Json Para El Request
		JsonFormat, Error := json.Marshal(Cadena)
		
		// Verificar Si Hay Error
		if Error != nil {

			print("Error Al Convertir La Cadena En Json: %v", Error)

		}

		// Peticion NodeJs
		Peticion := "http://localhost:4006/Python"

		// Solicitar Peticion A NodeJs
		Req, Error := http.NewRequest("POST", Peticion, bytes.NewBuffer(JsonFormat))
        // Verificar Si Hay Errores
		if Error != nil {

			print("Error Creando Petición: %v", Error)

		}

		// Agregar Encabezados
		Req.Header.Add("Content-Type", "application/json")
		// Req.Header.Add("X-Hola-Mundo", "Ejemplo")

		// Respueta NodeJS
		Response, Error := ClientHTTP.Do(Req)

		// Verificcar Si Hay Error
		if Error != nil {

			print("Error Haciendo Petición: %v", Error)

		}

		// Cerrar Cuerpo Response
		defer Response.Body.Close()

		// Cuerpo De La Respuseta
		RespuestaCuerpo, Error := ioutil.ReadAll(Response.Body)

		// Verificar Si Hay Error
		if Error != nil {

			print("Error Leyendo Respuesta: %v", Error)

		}
            print(string(RespuestaCuerpo))
		// Enviar Respuesta
		_, _ = fmt.Fprintf(w, "%s", RespuestaCuerpo)
		

	}