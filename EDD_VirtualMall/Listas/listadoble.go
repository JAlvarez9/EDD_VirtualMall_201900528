package Listas

import (
	"strconv"
	"fmt"
)

type Tiendas struct {
	Nombre string `json:Nombre`
	Descripcion string `json:Descripcion`
	Contacto string `json:Contacto`
	Calificacion int `json:Calificacion`
}

type Departamentos struct {
	Nombre string `json:Nombre`
	Tiendas []Tiendas `json:Tiendas`
}

type Datos struct {
	Indice string `json:Indice`
	Departamentos []Departamentos `json:Departamentos`
}

type Enlace struct {
	Datos []Datos `json:Datos`
}

type Shops struct {
	Tiendas []Tiendas
}

type contain struct {
	Tiendas []Node
}

type Node struct {
	Tienda Tiendas
	Departamento string
	Indice string
	Next *Node
	Back *Node
}


type NodeListas struct {
	Lista1 []*Node
	Lista2 []*Node
	Lista3 []*Node
	Lista4 []*Node
	Lista5 []*Node
}

type List struct {
	first, last *Node
	size int
}

type TrueNode struct {
	Stuff []*Node
	next *TrueNode
	back *TrueNode
}

func NuevaLista() *List{
	return &List{nil,nil,0}
}

func (this *List) GetSize() int{
	return this.size
}

func (this *List)Insertar(nuevo *Node){
	if this.first ==nil{
		this.first = nuevo
		this.last = nuevo
	}else{
		this.last.Next = nuevo
		nuevo.Back = this.last
		this.last = nuevo
	}
	this.size++
}

func (this *Node) To_string() string {
	cadena := "--> Indice: "+ this.Indice + " , Destino: "+ this.Departamento + ", NombreTienda: "+ this.Tienda.Nombre
	cadena += "CalificacionTienda: "+ strconv.Itoa(this.Tienda.Calificacion)

	return cadena
}

func (this *List) To_string() string {
	var cadena string
	aux := this.first
	for aux != nil {
		cadena += aux.To_string()
		aux = aux.Next
	}
	return cadena
}

func (this *List) Imprimir() {
	fmt.Println("Lista -------------")
	fmt.Println(this.To_string())
}


