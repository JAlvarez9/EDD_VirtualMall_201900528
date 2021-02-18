package main

import (
	"EDD_VirtualMall/Listas"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)


var tiendas2 [] Listas.List

var cubix [][] Listas.NodeListas
var sizedep int
var sizeindex int

func example(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to my REST API of EDD, hopefully you enjoy it! :)")

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
		sup := make([]Listas.NodeListas, sizedep)
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
					convertAscii(tienda.Nombre),
					nil,
					nil,
				}

				sup = putStore(aux2,sup,j)
			}
		}
		cubix[i] = sup

	}

	for i:= 0; i< sizedep; i++{
		for j:=0; j< sizeindex; j++{

			tiendas2 = append(tiendas2, cubix[j][i].Lista1)
			tiendas2 = append(tiendas2, cubix[j][i].Lista2)
			tiendas2 = append(tiendas2, cubix[j][i].Lista3)
			tiendas2 = append(tiendas2, cubix[j][i].Lista4)
			tiendas2 = append(tiendas2, cubix[j][i].Lista5)

		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newDoc)
}

func Deletition(w http.ResponseWriter, r *http.Request){
	var newDoc Listas.Pedidos
	reqBody, err := ioutil.ReadAll(r.Body)
	var contain bool
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)

	for i:= 0; i< len(tiendas2) ; i++{
		if tiendas2[i].Delete(&newDoc) == true{
			contain = true
			break
		}else {
			contain = false
		}
	}
	if contain{
		fmt.Fprintf(w,"The Store was deleted succesfully")
	}else {
		fmt.Fprintf(w,"We don't found that store please check your values")
	}
	var solo = len(tiendas2)
	fmt.Println(solo)

}

func Search(w http.ResponseWriter, r *http.Request){
	var newDoc Listas.Pedidos
	reqBody, err := ioutil.ReadAll(r.Body)
	var contain bool
	var finded Listas.Tiendas
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)

	for i:= 0; i < len(tiendas2);i++{
		finded = tiendas2[i].Search(&newDoc)
		if finded.Nombre != ""{
			contain = true
			break
		}else {
			contain = false
		}
	}
	if contain{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(finded)
	}else {
		fmt.Fprintf(w,"We don't found that store please check your values")
	}
}

func ShowList (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	PosVector, err :=strconv.Atoi(vars["id"])
	var stores []Listas.Tiendas
	if err != nil {
		fmt.Fprintf(w,"Invalid ID")
		return
	}
	var list = tiendas2[PosVector-1]
	stores = list.Show()
	if len(stores)==0{
		fmt.Fprintf(w,"In this list don't exist a store :(")
	}else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stores)
	}
}

func Graphviz(w http.ResponseWriter, r *http.Request){
	var cadenita strings.Builder
	fmt.Fprintf(&cadenita, "digraph G{ \n")
	fmt.Fprintf(&cadenita, "rankdir= \" LR \" \n")
	fmt.Fprintf(&cadenita, "node[shape=record]; \n")
	for i:= 0; i< len(tiendas2); i++{
		if tiendas2[i].GetSize() == 0{
			fmt.Fprintf(&cadenita, "node%d[label=\"vacio | %d \"]; \n ", i,i)
			fmt.Fprintf(&cadenita, "node%dv[label=\" \", color=\"white\"] \n",i)
			fmt.Fprintf(&cadenita, "node%d->node%dv; \n",i,i )

		}else{
			fmt.Fprintf(&cadenita, "node%d[label=\"%v|%v \"]; \n ", i,tiendas2[i].GetFirst().Departamento,tiendas2[i].GetFirst().Indice)
			tiendas2[i].Graphic(&cadenita)
			fmt.Fprintf(&cadenita, "node%d->node%p; \n",i,&(*tiendas2[i].GetFirst()) )

		}
		//fmt.Fprintf(&cadenita, "node%d[label=\"%v|%v \"]; \n ", i,tiendas2[i].GetFirst().Departamento,tiendas2[i].GetFirst().Indice)

	}
	fmt.Fprintf(&cadenita, "} \n")
	saveDot(cadenita.String())


}

func putStore(aux2 Listas.Node, sup []Listas.NodeListas, depa int,) []Listas.NodeListas{

	switch aux2.Tienda.Calificacion {
	case 1:
		sup[depa].Lista1.Insertar(&aux2)
		break
	case 2:sup[depa].Lista2.Insertar(&aux2)
		break
	case 3:
		sup[depa].Lista3.Insertar(&aux2)
		break
	case 4:
		sup[depa].Lista4.Insertar(&aux2)
		break
	case 5:
		sup[depa].Lista5.Insertar(&aux2)
		break
	}
	return sup
}

func saveDot(s string){
	f, err := os.Create("lista.dot")
	if err != nil{
		fmt.Println("There was an error!")
	}
	l, err := f.WriteString(s)
	if err != nil{
		fmt.Println("There was an error!")
		f.Close()
		return
	}
	fmt.Println(l,"The graphic was created succesfully")
	err = f.Close()
	if err!=nil{
		fmt.Println("There was an erro!")
		return
	}
}

func convertAscii(s string)int{
	ascii:= int(0)
	runes := []rune(s)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	for i := 0; i < len(result);i++ {
		ascii = ascii + result[i]
	}

	return ascii
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", example).Methods("GET")
	router.HandleFunc("/cargartienda", cargaArchivos).Methods("POST")
	router.HandleFunc("/Eliminar", Deletition).Methods("DELETE")
	router.HandleFunc("/TiendaEspecifica", Search).Methods("POST")
	router.HandleFunc("/id/{id}", ShowList).Methods("GET")
	router.HandleFunc("/getArreglo", Graphviz).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
