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
	"os/exec"
	"strconv"
	"strings"
)


var tiendas2 [] Listas.List

var cubix [][] Listas.NodeListas
var sizedep int
var sizeindex int
var departa [] string
var indice [] string

var prueba Listas.Enlace

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
	for i:= 0; i < len(newDoc.Datos[0].Departamentos); i++{
		departa = append(departa, newDoc.Datos[0].Departamentos[i].Nombre )
	}
	for i, datos := range newDoc.Datos{

		indice = append(indice, datos.Indice)

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
	var newDoc Listas.PedidosE
	reqBody, err := ioutil.ReadAll(r.Body)
	var contain bool
	var position int
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)
	position = searchingVectorE(&newDoc)
	contain = tiendas2[position].Delete(&newDoc)
	if contain && position >= 0{
		fmt.Fprintf(w,"The Store was deleted succesfully")
	}else {
		fmt.Fprintf(w,"We don't found that store please check your values")
	}
	var solo = len(tiendas2)
	fmt.Println(solo)

}

func Search(w http.ResponseWriter, r *http.Request){
	var newDoc Listas.PedidosS
	reqBody, err := ioutil.ReadAll(r.Body)
	var finded Listas.Tiendas
	var position int
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)

	position = searchingVectorS(&newDoc)
	finded = tiendas2[position].Search(&newDoc)

	if position >= 0 && finded.Nombre != ""{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(finded)
	}else {
		fmt.Fprintf(w,"We don't found that store please check your values")
	}
	fmt.Print("asd")
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
	cont := int(0)
	fmt.Fprintf(&cadenita, "digraph G{ \n")
	fmt.Fprintf(&cadenita, "rankdir= \"LR\" \n")
	fmt.Fprintf(&cadenita, "node[fontname=\"Arial\" style=\"filled\" shape=record color=\"blue\" fillcolor=\"mediumspringgreen\"]; \n")
	for i:= 0; i<= len(tiendas2);{
		if cont < 5{
			if i != len(tiendas2){
				if tiendas2[i].GetSize() == 0{
					fmt.Fprintf(&cadenita, "node%d[label=\"vacio | %d \"]; \n ", i,i)
					fmt.Fprintf(&cadenita, "node%dv[label=\" \", color=\"white\" fillcolor=\"white\"] \n",i)
					fmt.Fprintf(&cadenita, "node%d->node%dv; \n",i,i )
				}else{
					fmt.Fprintf(&cadenita, "node%d[label=\"Indice%v |%v|{%d|No.%d}\"]; \n ", i,tiendas2[i].GetFirst().Indice,tiendas2[i].GetFirst().Departamento,tiendas2[i].GetFirst().Tienda.Calificacion,i)
					tiendas2[i].Graphic(&cadenita)
					fmt.Fprintf(&cadenita, "node%d->node%p; \n",i,&(*tiendas2[i].GetFirst()) )

				}
				cont++
			}
			i++
		}else{
			fmt.Fprintf(&cadenita, "} \n")
			saveDot(cadenita.String(), i)
			cadenita.Reset()
			fmt.Fprintf(&cadenita, "digraph G{ \n")
			fmt.Fprintf(&cadenita, "rankdir= \"LR\" \n")
			fmt.Fprintf(&cadenita, "node[fontname=\"Arial\" style=\"filled\" shape=record color=\"blue\" fillcolor=\"mediumspringgreen\"]; \n")
			cont = 0
		}

	}


}

func SaveStuff(w http.ResponseWriter, r *http.Request){
	var stuffsvaing []Listas.Datos

	 for i:= 0; i< len(indice); i++{
		 var stuffdepa []Listas.Departamentos
	 	aux := Listas.Datos{
			Indice:        indice[i],
			Departamentos: stuffdepa,
		}
	 	stuffsvaing = append(stuffsvaing, aux)
	 }

	 for i:= 0; i<len(indice); i++{
	 	for j:= 0; j<len(departa) ;j++{
			stores := []Listas.Tiendas{}
			aux2:= Listas.Departamentos{
				Nombre:  departa[j],
				Tiendas: stores,
			}
			stuffsvaing[i].Departamentos = append(stuffsvaing[i].Departamentos, aux2)
		 }

	 }

	for _, listas := range tiendas2{
		t := listas.GetStores()
		if listas.GetSize()> 0 {
			for i, indices := range stuffsvaing{
				if indices.Indice == listas.GetFirst().Indice{
					for j, depars := range indices.Departamentos{
						if depars.Nombre == listas.GetFirst().Departamento{
							stuffsvaing[i].Departamentos[j].Tiendas = append(stuffsvaing[i].Departamentos[j].Tiendas, t...)
							break
						}
					}
				}
			}
		}
	}
	var sending Listas.Enlace
	sending.Datos = stuffsvaing
	f,_ := json.MarshalIndent(sending,""," ")
	_ = ioutil.WriteFile("NewStuff.json",f,0644)


}

func searchingVectorE(pedido *Listas.PedidosE) int{
	var indicefound, depafound bool
	var first, second, result int

	for i, s := range indice{
		if s[0] == pedido.Nombre[0]{
			first = i
			indicefound = true
		}
	}
	for i, s := range departa{
		if s == pedido.Categoria{
			second = i
			depafound = true
		}
	}

	if !indicefound {
		return -1
	}

	if !depafound{
		return -1
	}

	f := second - 0
	s := f * len(indice) + first
	result = s * 5 + pedido.Calificacion-1
	return result
	/*[i][j][w]
	primero = j-0
	segundo = primero * cantidad filas + i
	tercero = segundo*5 + w*/

}

func searchingVectorS(pedido *Listas.PedidosS) int{
	var indicefound, depafound bool
	var first, second, result int

	for i, s := range indice{
		if s[0] == pedido.Nombre[0]{
			first = i
			indicefound = true
		}
	}
	for i, s := range departa{
		if s == pedido.Departamento{
			second = i
			depafound = true
		}
	}

	if !indicefound {
		return -1
	}

	if !depafound{
		return -1
	}

	f := second - 0
	s := f * len(indice) + first
	result = s * 5 + pedido.Calificacion-1
	return result
	/*[i][j][w]
	primero = j-0
	segundo = primero * cantidad filas + i
	tercero = segundo*5 + w*/

}

func putStore(aux2 Listas.Node, sup []Listas.NodeListas, depa int,) []Listas.NodeListas{

	switch aux2.Tienda.Calificacion {
	case 1:
		sup[depa].Lista1.SortedInsert(&aux2)
		break
	case 2:sup[depa].Lista2.SortedInsert(&aux2)
		break
	case 3:
		sup[depa].Lista3.SortedInsert(&aux2)
		break
	case 4:
		sup[depa].Lista4.SortedInsert(&aux2)
		break
	case 5:
		sup[depa].Lista5.SortedInsert(&aux2)
		break
	}
	return sup
}

func saveDot(s string,i int){
	nombre := string("lista"+strconv.Itoa(i)+".pdf")
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
	fmt.Println(l,"Created Succesfully")
	p := "dot -Tpdf lista.dot -o " + nombre

	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)

	}
	fmt.Printf("%s\n", b)
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
	router.HandleFunc("/guardar", SaveStuff).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
