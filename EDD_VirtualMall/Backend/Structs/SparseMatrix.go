package Structs

import (
	"fmt"
	"reflect"
)

type (
		NodeMatrix struct {
			StackPedidos *Stack
			Value Pedidos
			Year int
			Dia int
			Month int
			MonthString string
			Ascii int
			Right, Left, Up, Down interface{}
		}

		NodeX struct {
			Dia int
			Right, Left, Up, Down interface{}
		}
		NodeY struct {
			Departamento string
			Ascii int
			Right, Left, Up, Down interface{}
		}
		SperseMatrix struct {
			HeadX *NodeX
			HeadY *NodeY
		}
)

func (this *SperseMatrix)getX(dia int) interface{} {
	if this.HeadX == nil{
		return nil
	}
	var sup interface{} = this.HeadX
	for sup != nil{
		if sup.(*NodeX).Dia == dia {
			return sup
		}
		sup = sup.(*NodeX).Right
	}
	return nil
}

func (this *SperseMatrix)getY(ascii int) interface{}{
	if this.HeadY == nil{
		return nil
	}
	var sup interface{} = this.HeadY
	for sup != nil {
		if sup.(*NodeY).Ascii == ascii {
			return sup
		}
		sup = sup.(*NodeY).Down
	}

	return nil
}

func (this *SperseMatrix)createY(ascci int, depa string) *NodeY  {
	if this.HeadY == nil {
		nueva := &NodeY{
			Departamento: depa,
			Ascii:        ascci,
			Right:        nil,
			Left:         nil,
			Up:           nil,
			Down:         nil,
		}
		this.HeadY = nueva
		return nueva
	}
	var sup interface{} = this.HeadY
	if ascci <= sup.(*NodeY).Ascii {
		nueva := &NodeY{
			Departamento: depa,
			Ascii:        ascci,
			Right:        nil,
			Left:         nil,
			Up:           nil,
			Down:         nil,
		}
		nueva.Down = this.HeadY
		this.HeadY.Up = nueva
		this.HeadY = nueva
		return nueva
	}
	for sup.(*NodeY).Down != nil{
		if ascci > sup.(*NodeY).Ascii && ascci <= sup.(*NodeY).Down.(*NodeY).Ascii {
			nueva := &NodeY{
				Departamento: depa,
				Ascii:        ascci,
				Right:        nil,
				Left:         nil,
				Up:           nil,
				Down:         nil,
			}
			aux := sup.(*NodeY).Down
			aux.(*NodeY).Up = nueva
			nueva.Down = aux
			sup.(*NodeY).Down = nueva
			nueva.Up = sup
			return nueva
		}
		sup = sup.(*NodeY).Down
		
	}
	nueva := &NodeY{
		Departamento: depa,
		Ascii:        ascci,
		Right:        nil,
		Left:         nil,
		Up:           nil,
		Down:         nil,
	}
	sup.(*NodeY).Down = nueva
	nueva.Up = sup
	return nueva
}

func (this *SperseMatrix)createX(dia int) *NodeX {
	if this.HeadX == nil {
		nueva := &NodeX{
			Dia:   dia,
			Right: nil,
			Left:  nil,
			Up:    nil,
			Down:  nil,
		}
		this.HeadX = nueva
		return nueva
	}
	var sup interface{} = this.HeadX
	if dia > sup.(*NodeX).Dia {
		nueva := &NodeX{
			Dia:   dia,
			Right: nil,
			Left:  nil,
			Up:    nil,
			Down:  nil,
		}
		nueva.Right = this.HeadX
		this.HeadX.Left = nueva
		this.HeadX = nueva
		return nueva
	}
	for sup.(*NodeX).Right != nil {
		if dia > sup.(*NodeX).Dia && dia <= sup.(*NodeX).Right.(*NodeX).Dia {
			nueva := &NodeX{
				Dia:   dia,
				Right: nil,
				Left:  nil,
				Up:    nil,
				Down:  nil,
			}
			aux := sup.(*NodeX).Right
			aux.(*NodeX).Left = nueva
			nueva.Right = aux
			sup.(*NodeX).Right = nueva
			nueva.Left = sup
			return nueva
		}
		sup = sup.(*NodeX).Right
	}
	nueva := &NodeX{
		Dia:   dia,
		Right: nil,
		Left:  nil,
		Up:    nil,
		Down:  nil,
	}
	sup.(*NodeX).Right = nueva
	nueva.Left = sup
	return nueva

}

func (this *SperseMatrix)getLastY(header *NodeY, ascci int)interface{} {
	if header.Down ==nil {
		return header
	}
	sup := header.Down
	if ascci <= sup.(*NodeY).Ascii {
		return header
	}
	for sup.(*NodeY).Down != nil {
		if ascci > sup.(*NodeY).Ascii && ascci <= sup.(*NodeY).Ascii {
			return sup
		}
		sup = sup.(*NodeY).Down
	}
	if ascci <= sup.(*NodeY).Ascii{
		return sup.(*NodeY).Up
	}
	return sup
}

func (this *SperseMatrix)getLastX(header *NodeX, dia int)interface{} {
	if header.Right ==nil {
		return header
	}
	sup := header.Right
	if dia <= sup.(*NodeX).Dia  {
		return header
	}
	for sup.(*NodeX).Right != nil {
		if dia > sup.(*NodeX).Dia && dia <= sup.(*NodeX).Dia {
			return sup
		}
		sup = sup.(*NodeX).Right
	}
	if dia <= sup.(*NodeX).Dia{
		return sup.(*NodeX).Left
	}
	return sup
}

func (this *SperseMatrix)Add(nuevo *NodeMatrix) {
	vertical := this.getY(nuevo.Ascii)
	horizontal := this.getX(nuevo.Dia)

	if vertical == nil {
		vertical = this.createY(nuevo.Ascii,nuevo.Value.Departamento)
	}
	if horizontal == nil{
		horizontal = this.createX(nuevo.Dia)
	}
	superior := this.getLastX(horizontal.(*NodeX), nuevo.Dia)
	izquierda := this.getLastY(vertical.(*NodeY), nuevo.Ascii)
	if reflect.TypeOf(izquierda).String() == "*main.NodeMatrix" {
		if izquierda.(*NodeMatrix).Right == nil {
			izquierda.(*NodeMatrix).Right = nuevo
			nuevo.Left = izquierda
		}else {
			temp := izquierda.(*NodeMatrix).Right
			izquierda.(*NodeMatrix).Right = nuevo
			nuevo.Left = izquierda
			temp.(*NodeMatrix).Left = nuevo
			nuevo.Left = temp
		}
	}else{
		if izquierda.(*NodeY).Right == nil  {
			izquierda.(*NodeY).Right = nuevo
			nuevo.Left = izquierda
		}else{
			temp := izquierda.(*NodeY).Right
			izquierda.(*NodeY).Right = nuevo
			nuevo.Left = izquierda
			temp.(*NodeMatrix).Right = nuevo
			nuevo.Right = temp
		}
		
	}

	/* SUPERIOR xd */
	if reflect.TypeOf(superior).String() == "*main.NodeMatrix" {
		if superior.(*NodeMatrix).Down == nil{
			superior.(*NodeMatrix).Down = nuevo
			nuevo.Up = superior
		}else {
			temp := superior.(*NodeMatrix).Down
			superior.(*NodeMatrix).Down = nuevo
			nuevo.Up = superior
			temp.(*NodeMatrix).Up = nuevo
			nuevo.Down = temp
		}
	}else {
		if superior.(*NodeX).Down == nil {
			superior.(*NodeX).Down = nuevo
			nuevo.Up = superior
		}else {
			temp:= superior.(*NodeX).Down
			superior.(*NodeX).Down = nuevo
			nuevo.Up = superior
			temp.(*NodeMatrix).Up = nuevo
			nuevo.Down = temp
		}
	}
}

func (this *SperseMatrix) Imprimir() {
	var aux interface{} = this.HeadY
	for aux != nil {
		fmt.Print(aux.(*NodeY).Departamento, "***************")
		tmp := aux.(*NodeY).Right
		for tmp != nil {
			fmt.Printf("%v,%v------", tmp.(*NodeMatrix).Dia, tmp.(*NodeMatrix).Ascii)
			tmp = tmp.(*NodeMatrix).Right
		}
		fmt.Print("\n")
		aux = aux.(*NodeY).Down
	}
}

func (this *SperseMatrix) Imprimir2() {
	var aux interface{} = this.HeadX
	for aux != nil {
		fmt.Print(aux.(*NodeX).Dia, "*****************")
		tmp := aux.(*NodeX).Down
		for tmp != nil {
			fmt.Printf("%v,%v-------", tmp.(*NodeMatrix).Dia, tmp.(*NodeMatrix).Ascii)
			tmp = tmp.(*NodeMatrix).Down
		}
		fmt.Println("")
		aux = aux.(*NodeX).Right
	}
}


