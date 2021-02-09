package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func example(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Example de mi Rest API")
}

func cargaArchivos(w http.ResponseWriter, r *http.Request){
	var newDocument listas.Entrelace

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}else {
		json.Unmarshal(reqBody, &newDocument)
		for _, message := range newDocument.Mensajes{
			aux2:= listas.NodoAll{
				message.Origen,
				message.Destino,
				message.Msg,
				nil,
				nil,
			}

			Mensajes.Insertar(&aux2)
		}
		//jSon = newDocument
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newDocument)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", example)
	router.HandleFunc("/", cargaArchivos)

	log.Fatal(http.ListenAndServe(":3000", router))
}
