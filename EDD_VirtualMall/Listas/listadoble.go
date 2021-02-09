package Listas

import (
	//"fmt"
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

type Node struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
	Next *Node
	Back *Node
}

type List struct {
	first, last *Node
	size int
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
}

