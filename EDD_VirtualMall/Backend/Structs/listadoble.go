package Structs

import (
	"fmt"
	"strings"
)

type (
	Node struct {
		Tienda Tiendas
		Departamento string
		Indice string
		Ascii int
		Next *Node
		Back *Node
	}
	NodeMonth struct {
		Matrix SperseMatrix
		Mont int
		MontjString string
		Next *NodeMonth
		Back *NodeMonth
	}
	NodeYear struct {
		Year int
		Monts *ListMonth
		Next *NodeYear
		Back *NodeYear
	}
	ListMonth struct {
		first *NodeMonth
		size int

	}
	ListYear struct {
		first *NodeYear
		size int
	}

	NodeListas struct {
		Lista1 List
		Lista2 List
		Lista3 List
		Lista4 List
		Lista5 List
	}
	List struct {
			first *Node
		size int
	}
	PedidosE struct {
		Categoria string `json:Categoria`
		Nombre string `json: Nombre`
		Calificacion int	`json:Calificacion`
	}
	PedidosS struct {
		Departamento string `json:Departamento`
		Nombre string `json: Nombre`
		Calificacion int	`json:Calificacion`
	}

)
/* Lista de Tiendas */
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

func (this *List)Search(ped *PedidosS) *Tiendas {
	var aux = this.first
	for aux != nil{
		if aux.Departamento == ped.Departamento && aux.Tienda.Nombre == ped.Nombre && aux.Tienda.Calificacion == ped.Calificacion{
			return &aux.Tienda
		}
		aux = aux.Next
	}
	nulo := Tiendas{
		Nombre:       "",
		Descripcion:  "",
		Contacto:     "",
		Calificacion: 0,
	}
	return &nulo
}

func (this *List)GetShop(s string) *Tiendas {
	var aux = this.first
	for aux != nil{
		if aux.Tienda.Contacto == s{
			return &aux.Tienda
		}
		aux = aux.Next
	}

	return nil
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

/* ------------- Lista de Años -------------- */

func (this *ListYear)SearchYear(year int) *NodeYear {
	aux := this.first
	for aux != nil{
		if aux.Year == year {
			return aux
		}
		aux = aux.Next
	}

	return nil
}

func (this *ListYear)AddYear(nuevo *NodeMatrix) {
	aux := this.SearchYear(nuevo.Year)
	if aux != nil{
		aux.Monts.AddMonth(nuevo)
	}else {

		sup2 := NodeYear{
			Year:  nuevo.Year,
			Monts: &ListMonth{
				first: nil,
				size:  0,
			},
			Next:  nil,
			Back:  nil,
		}
		this.SortedInsertYear(&sup2)
		aux = this.SearchYear(nuevo.Year)
		aux.Monts.AddMonth(nuevo)
	}
}

func (this *ListYear)SortedInsertYear(nuevo *NodeYear) {
	var aux *NodeYear
	if this.first == nil {
		this.first = nuevo

	}else if this.first.Year >= nuevo.Year{
		nuevo.Next = this.first

		nuevo.Next.Back = nuevo
		this.first = nuevo
	}else {
		aux = this.first
		for aux.Next != nil && aux.Next.Year < nuevo.Year{
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


/* --------------- Lista de Meses  --------------- */

func (this *ListMonth)SearchMonth(month int) *NodeMonth {
	aux := this.first
	for aux != nil {
		if aux.Mont == month {
			return aux
		}
		aux = aux.Next
	}
	return nil

}

func (this *ListMonth)AddMonth(nuevo *NodeMatrix) {
	aux := this.SearchMonth(nuevo.Month)
	if aux != nil {
		nuevoNM := nuevo
		aux.Matrix.Add(nuevoNM)
	}else {
			sup:=NodeMonth{
				Matrix:      SperseMatrix{
					HeadX: nil,
					HeadY: nil,
				},
				Mont:        nuevo.Month,
				MontjString: nuevo.MonthString,
				Next:        nil,
				Back:        nil,
			}
			this.SortedInsertMonth(&sup)
		aux = this.SearchMonth(nuevo.Month)
			nuevoNM := nuevo
			aux.Matrix.Add(nuevoNM)

	}

}

func (this *ListMonth)SortedInsertMonth(nuevo *NodeMonth) {
	var aux *NodeMonth
	if this.first == nil {
		this.first = nuevo

	}else if this.first.Mont >= nuevo.Mont{
		nuevo.Next = this.first

		nuevo.Next.Back = nuevo
		this.first = nuevo
	}else {
		aux = this.first
		for aux.Next != nil && aux.Next.Mont < nuevo.Mont{
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

