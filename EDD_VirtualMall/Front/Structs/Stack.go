package Structs

type (

	NodeStack struct {
		Value CodProducto
		Next *NodeStack
		Prev *NodeStack
	}

	Stack struct {
		Top *NodeStack
		Size int
	}
	
)

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

func (this *Stack)Pop() *CodProducto{
	if this.Size == 0{
		aux := this.Top
		this.Top.Prev.Next = nil
		aux.Prev = this.Top
		this.Top.Next = nil
		return &aux.Value
	}
	return nil
}
