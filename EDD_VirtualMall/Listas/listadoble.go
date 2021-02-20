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

type JsonErrors struct {
	Mensaje string `json:Mensajer`
}

type Departamentos struct {
	Nombre string `json:Nombre`
	Tiendas []Tiendas `json:Tiendas`

}

type Departamentos2 struct {
	Nombre string `json:Nombre`
	Tiendas []Tiendas `json:Tiendas`
	Indice string
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
	Ascii int
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
	first *Node
	size int
}

type TrueNode struct {
	Stuff []*Node
	next *TrueNode
	back *TrueNode
}

type PedidosE struct {
	Categoria string `json:Categoria`
	Nombre string `json: Nombre`
	Calificacion int	`json:Calificacion`
}

type PedidosS struct {
	Departamento string `json:Departamento`
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

func (this *List)SortedInsert(nuevo *Node) {
	var aux *Node
	if this.first == nil {
		this.first = nuevo

	}else if this.first.Ascii >= nuevo.Ascii{
		nuevo.Next = this.first

		nuevo.Next.Back = nuevo
		this.first = nuevo
	}else {
		aux = this.first
		for aux.Next != nil && aux.Next.Ascii < nuevo.Ascii{
			aux = aux.Next
		}
		nuevo.Next = aux.Next
		if aux.Next != nil{
			nuevo.Next.Back = nuevo
		}
		aux.Next = nuevo
		nuevo.Back = aux
	}
	this.size++
}

func (this *List)Search(ped *PedidosS) Tiendas {
	var aux = this.first
	for aux != nil{
		if aux.Departamento == ped.Departamento && aux.Tienda.Nombre == ped.Nombre && aux.Tienda.Calificacion == ped.Calificacion{
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

func(this *List) Delete(ped *PedidosE) bool{
	var finded bool
	var aux = this.first
	for aux != nil{
		if aux.Departamento == ped.Categoria && aux.Tienda.Nombre == ped.Nombre && aux.Tienda.Calificacion == ped.Calificacion{
			if aux == this.first && aux.Next == nil {
				this.first = nil
				this.size --
				return true
			}else if aux.Next == nil{
				aux.Back.Next = nil
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
			finded = false
		}
		aux = aux.Next
	}
	return finded
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

func (this *List)GetStores() []Tiendas {
	var tiendas []Tiendas
	aux := this.first
	if this.GetSize() == 0{
		return tiendas
	}else {
		for aux != nil{
			tiendas = append(tiendas, aux.Tienda)
			aux = aux.Next
		}

		return tiendas
	}

}

func (this *List)Graphic(cadenita *strings.Builder)  {
	graphic(this.first, cadenita, nil)
}

func graphic(back *Node, s *strings.Builder, now *Node){
	if back != nil{
		fmt.Fprintf(s, "node%p[shape=\"record\" label=\"%v| Ascii:%v| %v \" fillcolor=\"olivedrab1\"]; \n ",&(*back),back.Tienda.Nombre, back.Ascii,back.Tienda.Contacto)
		if now != nil{
			fmt.Fprintf(s, "node%p->node%p; \n", &(*now),&(*back))
			fmt.Fprintf(s, "node%p->node%p; \n", &(*back),&(*now))
		}
		graphic(back.Next, s , back)
	}

}


