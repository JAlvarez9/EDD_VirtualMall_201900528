package Structs

import (
	"fmt"
	"strconv"
)

type (
	/* Stack1 */
	NodeStack struct {
		Value Pedidos
		Next *NodeStack
		Prev *NodeStack
	}

	Stack struct {
		Top *NodeStack
		Size int
	}

	/* Stack2 */

	NodeStack2 struct {
		Pedido ValidarPedidos
		Next *NodeStack2
		Prev *NodeStack2
	}

	Stack2 struct {
		Top *NodeStack2
		Size int
	}

	/* Stack3 */

	NodeStack3 struct {
		Nodo Vertices
		Next *NodeStack3
		Prev *NodeStack3
	}

	Stack3 struct {
		Top *NodeStack3
		Size int
	}

	/* Stack4 */
	NodeStack4 struct {
		Nodo CaminosCortos
		Next *NodeStack4
		Prev *NodeStack4
	}

	Stack4 struct {
		Top *NodeStack4
		Size int
	}

	/* Stack5 */

	NodeStack5 struct {
		Value CaminosProductos
		Next *NodeStack5
		Prev *NodeStack5
	}

	Stack5 struct {
		Top *NodeStack5
		Size int
	}

)

/*  Stack de Producto dentro Pedidos   */

func NewStack() *Stack {
	return &Stack{
		Top:  nil,
		Size: 0,
	}
}

func (this *Stack)Push(stack *NodeStack) {
	aux := this.Top
	if aux == nil {
		this.Top = stack
	}else {
		aux.Next = stack
		stack.Prev = this.Top
		this.Top = stack

	}
	this.Size++

}

func (this *Stack)Pop() *Pedidos{
	if this.Size == 0{
		aux := this.Top
		this.Top.Prev.Next = nil
		aux.Prev = this.Top
		this.Top.Next = nil
		return &aux.Value
	}
	return nil
}

func (this *Stack)First() *Pedidos{
	return &this.Top.Value
}

func (this *Stack)ArregloPedidos() []Pedidos {
	aux:= this.Top
	var sup []Pedidos
	for aux != nil{
		sup = append(sup,aux.Value)

		aux = aux.Next
	}
	return sup

}

/*  Stack de Validar Pedidos   */

func NewStack2() *Stack2 {
	return &Stack2{
		Top:  nil,
		Size: 0,
	}
}

func (this *Stack2)Push2(stack *NodeStack2) {
	aux := this.Top
	if aux == nil {
		this.Top = stack
		this.Size++
	}else {
		aux.Next = stack
		stack.Prev = this.Top
		this.Top = stack
		this.Size++
	}
}

func (this *Stack2)VerificarExsite(stack *NodeStack2)bool {
	aux := this.Top
	for aux != nil{
		if stack.Pedido.Tienda == aux.Pedido.Tienda && stack.Pedido.Departamento == aux.Pedido.Departamento && stack.Pedido.Calificacion == aux.Pedido.Calificacion {
			aux.Pedido.Productos = append(aux.Pedido.Productos, stack.Pedido.Producto )
			return true
		}
		aux = aux.Prev
	}
	return false
}

func (this *Stack2)ArregloVPedidos() *[]ValidarPedidos {
	aux := this.Top
	var sup []ValidarPedidos
	for aux!= nil{
		sup = append(sup,aux.Pedido)
		aux = aux.Prev
	}

	return &sup
}

/*  Stack de Nodos de los Grafos  */

func NewStack3() *Stack3 {
	return &Stack3{
		Top:  nil,
		Size: 0,
	}
}

func (this *Stack3)Push3(stack *NodeStack3) {
	aux := this.Top
	if aux == nil {
		this.Top = stack
	}else {
		aux.Next = stack
		stack.Prev = this.Top
		this.Top = stack

	}
	this.Size++

}

func (this *Stack3)Pop3() *Vertices{
	if this.Size == 0{
		aux := this.Top
		this.Top.Prev.Next = nil
		aux.Prev = this.Top
		this.Top.Next = nil
		return &aux.Nodo
	}
	return nil
}

func (this *Stack3)ArregloVGrafo() *[]Vertices {
	aux := this.Top
	var sup []Vertices
	for aux!= nil{
		sup = append(sup,aux.Nodo)
		aux = aux.Prev
	}

	return &sup
}

func (this *Stack3)ArregloDobleD()*[][]string  {
	matriz := make([][]string, this.Size+1)
	aux := this.Top
	sup := make([]string, this.Size+1)
	for i:= 1; i < this.Size+1;i++{
		sup[i]= aux.Nodo.Nombre
		aux = aux.Prev
	}
	matriz[0] = sup
	aux = this.Top
	for i := 1; i < this.Size+1; i++ {
		sup3 := make([]string, this.Size+1)
		sup3[0] = aux.Nodo.Nombre
		matriz[i] = sup3
		aux = aux.Prev
	}

	var aux3 []Vertices

	for aux = this.Top; aux != nil; aux=aux.Prev{
		for _, a := range aux.Nodo.Enlaces {
			for aux2:= this.Top; aux2 != nil; aux2=aux2.Prev{
				if aux2.Nodo.Nombre == a.Nombre{
					sup3 := Enlaces{
						Nombre:    aux.Nodo.Nombre,
						Distancia: a.Distancia,
					}
					var sup4 []Enlaces
					sup4 = append(sup4, sup3)
					sup2 :=Vertices{
						Nombre:    a.Nombre,
						Enlaces: sup4,
					}
					aux3 = append(aux3, sup2)
				}

			}
		}

	}

	for aux = this.Top; aux != nil; aux=aux.Prev{

		for _, i2 := range aux3 {
			if aux.Nodo.Nombre == i2.Nombre{
				aux.Nodo.Enlaces = append(aux.Nodo.Enlaces, i2.Enlaces...)
			}
		}

	}

	for aux = this.Top; aux != nil; aux = aux.Prev{
		for i := 1; i < this.Size+1; i++{
			if aux.Nodo.Nombre == matriz[i][0]{
				for j:=1; j< this.Size+1; j++{
					for _, i2 := range aux.Nodo.Enlaces{
						if i2.Nombre == matriz[0][j]{

							matriz[i][j] = FloatTostring(i2.Distancia)

						}
					}
				}
			}

		}
	}

	for i := 1; i < this.Size+1; i++ {
		for j := 1; j < this.Size+1; j++ {
			if matriz[i][0] == matriz[0][j]{
				matriz[i][j] = "0"
			}else if matriz[i][j] == ""{
				matriz[i][j] = "999999"
			}
		}
	}
	/*
	for i := 1; i < this.Size+1; i++ {
		for j := 1; j < this.Size+1; j++ {
			fmt.Printf("| %s |", matriz[i][j])
		}
		fmt.Print("\n")
	}

	fmt.Print("asdf")
	*/
	return &matriz
}

/*  Stack de Caminos mas Cortos*/

func NewStac4k() *Stack4 {
	return &Stack4{
		Top:  nil,
		Size: 0,
	}
}

func (this *Stack4)Push4(stack *NodeStack4) {
	aux := this.Top
	if aux == nil {
		this.Top = stack
	}else {
		aux.Next = stack
		stack.Prev = this.Top
		this.Top = stack

	}
	this.Size++

}

func (this *Stack4)Pop4() *CaminosCortos{
	if this.Size == 0{
		aux := this.Top
		this.Top.Prev.Next = nil
		aux.Prev = this.Top
		this.Top.Next = nil
		return &aux.Nodo
	}
	return nil
}

func (this *Stack4)ArregloStack4() *[]CaminosCortos{
	aux := this.Top
	var sup []CaminosCortos
	for aux!= nil{
		sup = append(sup,aux.Nodo)
		aux = aux.Prev
	}

	return &sup
}

/*  Stack de Caminos mas Cortos*/

func NewStack5() *Stack5 {
	return &Stack5{
		Top:  nil,
		Size: 0,
	}
}

func (this *Stack5)Push5(stack *NodeStack5) {
	aux := this.Top
	if aux == nil {
		this.Top = stack
	}else {
		aux.Next = stack
		stack.Prev = this.Top
		this.Top = stack

	}
	this.Size++

}

func (this *Stack5)VerificarExsite5(stack *Productos)bool {
	aux := this.Top
	for aux != nil{
		if stack.Almacenamiento == aux.Value.Almacenamiento {
			aux.Value.Productos = append(aux.Value.Productos, *stack)
			return true
		}
		aux = aux.Prev
	}
	return false
}

func (this *Stack5)ArregloCaminosProductos() *[]CaminosProductos {
	aux := this.Top
	var sup []CaminosProductos
	for aux!= nil{
		sup = append(sup,aux.Value)
		aux = aux.Prev
	}

	return &sup
}


/* Funciones Extras */

func StringToint(cadena string) int {
	numero, _ := strconv.Atoi(cadena)
	return numero
}

func IntTostring(numero int) string {
	cadena := strconv.Itoa(numero)
	return cadena
}

func FloatTostring(f float64) string {
	s := fmt.Sprint(f)
	return s
}

func StringTofloat(s string) float64 {
	f, _ := strconv.ParseFloat(s,64)
	return f
}