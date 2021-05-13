package Structs

import (
	"math"
)

type (
	HashNode struct {
		hash 	int
		Value   string
	}

	HashTable struct {
		size int
		carga int
		porcentaje int
		porc_crecimiento int
		arreglo 	[]*HashNode
	}
)

func NewHashTable(size int, porcentaje int, porc_crecimiento int)*HashTable{
	arreglo := make([]*HashNode,size)
	return &HashTable{
		size:             size,
		carga:            0,
		porcentaje:       porcentaje,
		porc_crecimiento: porc_crecimiento,
		arreglo:          arreglo,
	}
}



func (this *HashTable)Insertar(valor string, nuevo int){
		new_node := HashNode{
			hash:  nuevo,
			Value: valor,
		}
		pos:= this.posicion(nuevo)
		this.arreglo[pos] = &new_node
		this.carga++
		if((this.carga*100)/this.size) > this.porcentaje {
			new_size:= this.size
			for{
				new_size++
				if((this.carga*100)/new_size) <= this.porc_crecimiento{
					break
				}
			}
			nuevo_array:= make([]*HashNode,new_size)
			antiguo:= this.arreglo
			this.arreglo = nuevo_array
			this.size = new_size
			aux:= 0
			for i := 0; i < len(antiguo); i++ {
				if antiguo[i] != nil{
					aux= this.posicion(antiguo[i].hash)
					nuevo_array[aux] = antiguo[i]

				}
			}
		}


}

func (this *HashTable)posicion(clave int) int{
	i,p:=0,0
	aux:= math.Mod(1,float64(clave) * .2520)


	p= int(float64(this.size) * aux)
	if p >= this.size{
		this.size = this.size + 8
		nuevo_array:= make([]*HashNode,this.size)
		copy(nuevo_array,this.arreglo)
		this.arreglo = nuevo_array
	}
	for this.arreglo[p] != nil && this.arreglo[p].hash != clave{
		i = int(math.Pow(float64(i),float64(2)))
		p = p + i
		if p >= this.size{
			p = p -this.size
		}

	}
	return p
}

func (this *HashTable)ArregloValores() []string {
	var aux []string
	for _, node := range this.arreglo {
		if node != nil{
			if node.Value != ""{
				aux = append(aux, node.Value)
			}
		}

	}
	return aux
}



