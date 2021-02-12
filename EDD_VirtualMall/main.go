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



var tiendas = Listas.NuevaLista()

var cubix [][] Listas.NodeListas
var sizedep int
var sizeindex int

func example(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Example de mi Rest API")
	tiendas.Imprimir()

}

func cargaArchivos(w http.ResponseWriter, r *http.Request){
	var newDoc Listas.Enlace

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)
	for i, indice := range newDoc.Datos{
		for j, _ := range indice.Departamentos{
			sizedep = j+1
		}
		sizeindex = i+1
	}
	cubix = make([][] Listas.NodeListas, sizeindex)
	for i, datos := range newDoc.Datos{

		for j, departamentos := range datos.Departamentos{

			for _, tienda := range departamentos.Tiendas{
				aux3 :=Listas.Tiendas{
					tienda.Nombre,
					tienda.Descripcion,
					tienda.Contacto,
					tienda.Calificacion,
				}
				aux2:= Listas.Node{
					aux3,
					departamentos.Nombre,
					datos.Indice,
					nil,
					nil,
				}
				putStore(aux2,i,j )
			}
		}

	}

	for i:= 0; i< sizeindex; i++{
		for j:=0; j< sizedep; j++{
			for k:= 0; k< len(cubix[i][j].Lista1); k++{
				tiendas.Insertar(cubix[i][j].Lista1[k])
			}
			for k:= 0; k< len(cubix[i][j].Lista2); k++{
				tiendas.Insertar(cubix[i][j].Lista2[k])
			}
			for k:= 0; k< len(cubix[i][j].Lista3); k++{
				tiendas.Insertar(cubix[i][j].Lista4[k])
			}
			for k:= 0; k< len(cubix[i][j].Lista4); k++{
				tiendas.Insertar(cubix[i][j].Lista4[k])
			}
			for k:= 0; k< len(cubix[i][j].Lista5); k++{
				tiendas.Insertar(cubix[i][j].Lista5[k])
			}
		}
	}
	solo := tiendas.GetSize()
	fmt.Println(solo)
	tiendas.Imprimir()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newDoc)


}

func putStore(aux2 Listas.Node, indice int, depa int,){
	//salir := bool()

		sup := make([]Listas.NodeListas, sizedep)

			switch aux2.Tienda.Calificacion {
			case 1:
				sup[depa].Lista1 = append(sup[depa].Lista1, &aux2)
				break
			case 2:
				sup[depa].Lista2 = append(sup[depa].Lista2, &aux2)
				break
			case 3:
				sup[depa].Lista3 = append(sup[depa].Lista3, &aux2)
				break
			case 4:
				sup[depa].Lista4 = append(sup[depa].Lista4, &aux2)
				break
			case 5:
				sup[depa].Lista5 = append(sup[depa].Lista5, &aux2)
				break
			}


		cubix[indice] = sup



}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", example).Methods("GET")
	router.HandleFunc("/cargartienda", cargaArchivos).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
