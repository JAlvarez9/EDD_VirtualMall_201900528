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



func (this *HashTable)insertar(valor string, nuevo int){
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
	aux:= int(float64(clave) * .2520)
	aux2:= aux % 1

	p= this.size * aux2

	for this.arreglo[p] != nil && this.arreglo[p].hash != clave{
		i++
		p = int(math.Pow(float64(p), float64(2)))
		if p >= this.size{
			p = p -this.size
		}

	}
	return p
}

