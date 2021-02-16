package Listas

import (
	"fmt"
	"strings"
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

type Node struct {
	Tienda Tiendas
	Departamento string
	Indice string
	Next *Node
	Back *Node
}

type NodeListas struct {
	Lista1 List
	Lista2 List
	Lista3 List
	Lista4 List
	Lista5 List
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

type Pedidos struct {
	Categoria string `json:Categoria`
	Nombre string `json: Nombre`
	Calificacion int	`json:Calificacion`
}

func (this *List) GetSize() int{
	if this.first == nil{
		this.size=0
	}
	return this.size
}

func (this *List)GetFirst() *Node {
	return this.first
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

func (this *List)Search(ped *Pedidos) Tiendas {
	var aux = this.first
	for aux != nil{
		if aux.Departamento == ped.Categoria && aux.Tienda.Nombre == ped.Nombre && aux.Tienda.Calificacion == ped.Calificacion{
			return aux.Tienda
		}
		aux = aux.Next
	}
	nulo := Tiendas{
		Nombre:       "",
		Descripcion:  "",
		Contacto:     "",
		Calificacion: 0,
	}
	return nulo
}

func(this *List) Delete(ped *Pedidos) bool{
	var aux = this.first
	for aux != nil{
		if aux.Departamento == ped.Categoria && aux.Tienda.Nombre == ped.Nombre && aux.Tienda.Calificacion == ped.Calificacion{
			if aux == this.last{
				aux.Back.Next = nil
				this.last = aux.Back
				this.size--
				return true
			}else if aux == this.first{
				aux.Next.Back = nil
				this.first = aux.Next
				this.size--
				return true
			}
			aux.Next.Back = aux.Back
			aux.Back.Next = aux.Next

			aux.Next = nil
			aux.Back = nil
			this.size--
			return true
		}else {
			return false
		}
		aux = aux.Next
	}
	return false
}

func (this *List)Show() []Tiendas{
	var tiendas []Tiendas
	var aux = this.first
	for aux != nil{
		tiendas = append(tiendas,aux.Tienda)
		aux = aux.Next
	}
	return tiendas
}

func (this *List)Graphic(cadenita *strings.Builder)  {
	graphic(this.first, cadenita, nil)
}

func graphic(back *Node, s *strings.Builder, now *Node){
	if back != nil{
		fmt.Fprintf(s, "node%p[label=\"%v|%v \"]; \n ",&(*back),back.Tienda.Nombre, back.Tienda.Calificacion)
		if now != nil{
			fmt.Fprintf(s, "node%p->node%p; \n", &(*now),&(*back))
			fmt.Fprintf(s, "node%p->node%p; \n", &(*back),&(*now))
		}
		graphic(back.Next, s , back)
	}

}


