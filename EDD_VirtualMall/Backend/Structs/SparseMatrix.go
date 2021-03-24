package Structs

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"strconv"
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
			X int
			Right, Left, Up, Down interface{}
		}
		NodeY struct {
			Departamento string
			Ascii int
			Y int
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
	if dia <= sup.(*NodeX).Dia {
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

func (this *SperseMatrix)getLastY(header *NodeX, ascci int)interface{} {
	if header.Down ==nil {
		return header
	}
	sup := header.Down
	if ascci <= sup.(*NodeMatrix).Ascii {
		return header
	}
	for sup.(*NodeMatrix).Down != nil {
		if ascci > sup.(*NodeMatrix).Ascii && ascci <= sup.(*NodeMatrix).Down.(*NodeMatrix).Ascii {
			return sup
		}
		sup = sup.(*NodeMatrix).Down
	}
	if ascci <= sup.(*NodeMatrix).Ascii{
		return sup.(*NodeMatrix).Up
	}
	return sup
}

func (this *SperseMatrix)getLastX(header *NodeY, dia int)interface{} {
	if header.Right ==nil {
		return header
	}
	sup := header.Right
	if dia <= sup.(*NodeMatrix).Dia  {
		return header
	}
	for sup.(*NodeMatrix).Right != nil {
		if dia > sup.(*NodeMatrix).Dia && dia <= sup.(*NodeMatrix).Right.(*NodeMatrix).Dia {
			return sup
		}
		sup = sup.(*NodeMatrix).Right
	}
	if dia <= sup.(*NodeMatrix).Dia{
		return sup.(*NodeMatrix).Left
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
	izquierda := this.getLastX(vertical.(*NodeY), nuevo.Dia)
	superior := this.getLastY(horizontal.(*NodeX), nuevo.Ascii)
	if reflect.TypeOf(izquierda).String() == "*Structs.NodeMatrix" {
		if izquierda.(*NodeMatrix).Right == nil {
			izquierda.(*NodeMatrix).Right = nuevo
			nuevo.Left = izquierda
		}else {
			temp := izquierda.(*NodeMatrix).Right
			izquierda.(*NodeMatrix).Right = nuevo
			nuevo.Left = izquierda
			temp.(*NodeMatrix).Left = nuevo
			nuevo.Right = temp
		}
	}else{
		if izquierda.(*NodeY).Right == nil  {
			izquierda.(*NodeY).Right = nuevo
			nuevo.Left = izquierda
		}else{
			temp := izquierda.(*NodeY).Right
			izquierda.(*NodeY).Right = nuevo
			nuevo.Left = izquierda
			temp.(*NodeMatrix).Left = nuevo
			nuevo.Right = temp
		}
		
	}

	/* SUPERIOR xd */
	if reflect.TypeOf(superior).String() == "*Structs.NodeMatrix" {
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

func (this *SperseMatrix) ObtenerNodito(headx int, heady int) *NodeMatrix  {
	var aux interface{} = this.HeadX
	for aux != nil{
		if aux.(*NodeX).Dia == headx{
			var sup = aux.(*NodeX).Down
			for sup != nil{
				if sup.(*NodeMatrix).Ascii == heady{
					return sup.(*NodeMatrix)
				}
				sup = sup.(*NodeMatrix).Down
			}

		}
		aux = aux.(*NodeX).Right
	}

	return nil
}

func (this *SperseMatrix) Imprimir() {
	var aux interface{} = this.HeadY
	for aux != nil {
		fmt.Print(aux.(*NodeY).Departamento, "***************")
		tmp := aux.(*NodeY).Right
		for tmp != nil {
			fmt.Printf("%v,%v------", tmp.(*NodeMatrix).Dia, tmp.(*NodeMatrix).Value.Departamento)
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
			fmt.Printf("%v,%v-------", tmp.(*NodeMatrix).Dia, tmp.(*NodeMatrix).Value.Departamento)
			tmp = tmp.(*NodeMatrix).Down
		}
		fmt.Println("")
		aux = aux.(*NodeX).Right
	}
}

func (this *SperseMatrix)Graphviz() string{
	var cadenita strings.Builder
	//var sizeX int
	//var sizeY int
	fmt.Fprintf(&cadenita, "digraph G{ \n")
	fmt.Fprintf(&cadenita, "node [shape=box] \n")
	fmt.Fprintf(&cadenita, " Mt[ label = \"Matrix\", width = 1.5, style = filled, fillcolor = firebrick1, group = 1 ]; \n")
	fmt.Fprintf(&cadenita, "e0[ shape = point, width = 0 ]; \n")
	fmt.Fprintf(&cadenita, "e1[ shape = point, width = 0 ]; \n")
	var aux interface{} = this.HeadY
	cont := 0
	for aux != nil  {

		fmt.Fprintf(&cadenita, "node%p [label = \"%s\"    width = 1.5 style = filled, fillcolor = bisque1, group = 1 ]; \n",aux.(*NodeY),aux.(*NodeY).Departamento)
		if aux.(*NodeY).Down != nil {
			fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux.(*NodeY),aux.(*NodeY).Down)
			fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux.(*NodeY).Down,aux.(*NodeY))
		}

		aux = aux.(*NodeY).Down
		//sizeY = cont
		cont++
	}

	aux = this.HeadX
	cont = 0
	for aux != nil  {

		fmt.Fprintf(&cadenita, "node%p [label = \"%d\"    width = 1.5 style = filled, fillcolor = bisque1, group = %v ]; \n",aux.(*NodeX),aux.(*NodeX).Dia,cont+2)
		if aux.(*NodeX).Right != nil {
			fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux.(*NodeX),aux.(*NodeX).Right)
			fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux.(*NodeX).Right,aux.(*NodeX))
		}

		aux = aux.(*NodeX).Right
		//sizeX = cont
		cont++
	}
	aux = this.HeadY
	fmt.Fprintf(&cadenita, "Mt -> node%p \n",aux.(*NodeY) )
	aux = this.HeadX
	fmt.Fprintf(&cadenita, "Mt -> node%p \n",aux.(*NodeX))

	aux = this.HeadX
	fmt.Fprintf(&cadenita, "{ rank = same; Mt;  ")

	for aux != nil {
		fmt.Fprintf(&cadenita, "node%p;", aux.(*NodeX))
		aux = aux.(*NodeX).Right

	}
	fmt.Fprintf(&cadenita, "} \n")
	aux = this.HeadY
	cont = 0
	var aux2 interface{}

	aux = this.HeadX
	cont =0
	for aux != nil{

		aux2 = aux.(*NodeX).Down
		fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux.(*NodeX),aux2.(*NodeMatrix))
		for aux2 != nil{
			fmt.Fprintf(&cadenita, "node%p [label = \"Nodito\" style = filled, fillcolor = darkolivegreen2\t width = 1.5, group = %v ]; \n",aux2.(*NodeMatrix),cont+2)
			if aux2.(*NodeMatrix).Down != nil {
				fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux2.(*NodeMatrix),aux2.(*NodeMatrix).Down)
				fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux2.(*NodeMatrix).Down,aux2.(*NodeMatrix))
			}
			aux2 = aux2.(*NodeMatrix).Down
		}
		aux = aux.(*NodeX).Right
		cont++
	}
	aux = this.HeadY
	for aux != nil{
		cont =0
		aux2 = aux.(*NodeY).Right
		fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux.(*NodeY),aux2.(*NodeMatrix))
		for aux2 != nil{

			if aux2.(*NodeMatrix).Right != nil {
				fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux2.(*NodeMatrix),aux2.(*NodeMatrix).Right)
				fmt.Fprintf(&cadenita, "node%p -> node%p; \n",aux2.(*NodeMatrix).Right,aux2.(*NodeMatrix))
			}
			aux2 = aux2.(*NodeMatrix).Right
			cont++
		}
		aux = aux.(*NodeY).Down
		cont++
	}
	aux = this.HeadY
	for aux != nil{
		aux2 = aux.(*NodeY).Right
		fmt.Fprintf(&cadenita, "{ rank = same; node%p; ",aux.(*NodeY))
		for aux2 != nil{
			fmt.Fprintf(&cadenita, "node%p;", aux2.(*NodeMatrix))
			aux2 = aux2.(*NodeMatrix).Right

		}
		fmt.Fprintf(&cadenita, "} \n")

		aux = aux.(*NodeY).Down

	}

	fmt.Fprintf(&cadenita, "} \n")



	string64 := savedotMatriz(cadenita.String(), strconv.Itoa(this.HeadX.Down.(*NodeMatrix).Year) +"-" +strconv.Itoa(this.HeadX.Down.(*NodeMatrix).Month) )
	return string64


}

func savedotMatriz(s string, i string) string {

	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("Matriz"+i+".png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\matriz.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\matriz.dot -o "+path+"\\Matrices\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

	imgFile, err := os.Open(path+"\\Matrices\\"+nombre) // DIreccion de la imagen

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return imgBase64Str
}


