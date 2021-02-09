package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"EDD_VirtualMall/Listas"
)

var cubix [][][5] string

var tiendas = Listas.NuevaLista()

func example(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Example de mi Rest API")
}

func cargaArchivos(w http.ResponseWriter, r *http.Request){
	var newDoc Listas.Enlace

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}else {
		json.Unmarshal(reqBody, &newDoc)
		for _, datos := range newDoc.Datos{
			for _, departamentos := range datos.Departamentos{
				for _, tienda := range departamentos.Tiendas{
					aux2:= Listas.Node{
						tienda.Nombre,
						tienda.Descripcion,
						tienda.Contacto,
						tienda.Calificacion,
						nil,
						nil,
					}
					tiendas.Insertar(&aux2)
				}
			}

		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newDoc)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", example)
	router.HandleFunc("/cargartienda", cargaArchivos).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
