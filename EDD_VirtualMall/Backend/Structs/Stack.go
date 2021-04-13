package Structs

type (

	NodeStack struct {
		Value Pedidos
		Next *NodeStack
		Prev *NodeStack
	}

	Stack struct {
		Top *NodeStack
		Size int
	}

	NodeStack2 struct {
		Pedido ValidarPedidos
		Next *NodeStack2
		Prev *NodeStack2
	}

	Stack2 struct {
		Top *NodeStack2
		Size int
	}

	NodeStack3 struct {
		Nodo Vertices
		Next *NodeStack3
		Prev *NodeStack3
	}

	Stack3 struct {
		Top *NodeStack3
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

func (this *Stack3)ArregloDobleD()*[][]int  {
	return nil
}
