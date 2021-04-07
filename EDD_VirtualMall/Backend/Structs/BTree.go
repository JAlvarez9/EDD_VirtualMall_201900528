package Structs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type (

	BTNode struct {
		Max int
		NodeF *BTNode
		Keys []*Key
	}

	Key struct {
	Value UsuariosEncrit
	Left *BTNode
	Right *BTNode
	}

	BTree struct {
		k int
		root *BTNode
		
	}
)

/*  Funciones del valor dentro del Nodo Arbol B   */

func NewKey(value UsuariosEncrit)*Key {
	k:=Key{
		Value: value,
		Left:  nil,
		Right: nil,
	} 
	return &k
}

/*  Funciones del Nodo del Arbol B   */

func NewBTNode(max int) *BTNode  {
	keys:= make([]*Key, max)
	n:= BTNode{
		Max:   max,
		NodeF: nil,
		Keys:  keys,
	}
	return &n
}

func (this *BTNode)Put(i int, llave *Key)  {
	this.Keys[i] = llave
}

/*  Funciones del Arbol B   */

func NewBTree(nivel int) *BTree  {

	a := BTree{
		k:    nivel,
		root: nil,
	}
	rootnode := NewBTNode(nivel)
	a.root = rootnode
	return &a
}

func (this *BTree)Insert(newKey *Key)  {

	if this.root.Keys[0] == nil {
		this.root.Put(0,newKey)
	}else if this.root.Keys[0].Left == nil{
		insertPlace:= -1
		node := this.root
		insertPlace = this.placeNode(node, newKey)
		if insertPlace != -1{
			if insertPlace == node.Max-1{
				middle := node.Max/2
				centralKey := node.Keys[middle]
				derecho := NewBTNode(this.k)
				izquierdo := NewBTNode(this.k)
				leftIndex := 0
				rightIndex := 0
				for j := 0; j<node.Max; j++{
					if node.Keys[j].Value.DPIN < centralKey.Value.DPIN{
						izquierdo.Put(leftIndex, node.Keys[j])
						leftIndex++
						node.Put(j,nil)
					}else if node.Keys[j].Value.DPIN >centralKey.Value.DPIN{
						derecho.Put(rightIndex,node.Keys[j])
						rightIndex++
						node.Put(j,nil)
					}
				}
				node.Put(middle,nil)
				this.root = node
				this.root.Put(0,centralKey)
				izquierdo.NodeF = this.root
				derecho.NodeF = this.root
				centralKey.Left = izquierdo
				centralKey.Right = derecho
			}
		}

	}else if this.root.Keys[0].Left != nil{
		node := this.root
		for node.Keys[0].Left != nil{
			loop := 0
			for i:=0; i<node.Max ;i,loop = i+1, loop+1{
					if node.Keys[i]!= nil{
						if node.Keys[i].Value.DPIN > newKey.Value.DPIN{
							node = node.Keys[i].Left
							break
						}

					}else {
						node = node.Keys[i-1].Right
						break
					}
			}
			if loop == node.Max{
				node = node.Keys[loop-1].Right
			}
		}
		indiceColocado := this.placeNode(node,newKey)
		if indiceColocado == node.Max-1{
			for node.NodeF != nil{
				middleIndex := node.Max/2
				centralKey := node.Keys[middleIndex]
				izquierdo := NewBTNode(this.k)
				derecho :=  NewBTNode(this.k)
				leftIndex, rightIndex := 0,0
				for i:= 0; i< node.Max;i++{
					if node.Keys[i].Value.DPIN < centralKey.Value.DPIN{
						izquierdo.Put(leftIndex,node.Keys[i])
						leftIndex++
						node.Put(i,nil)
					}else if node.Keys[i].Value.DPIN > centralKey.Value.DPIN{
						derecho.Put(rightIndex, node.Keys[i])
						rightIndex++
						node.Put(i,nil)
					}

				}
				node.Put(middleIndex, nil)
				centralKey.Left=izquierdo
				centralKey.Right = derecho
				node = node.NodeF
				izquierdo.NodeF = node
				derecho.NodeF = node
				for j:= 0;j< izquierdo.Max;j++{
					if izquierdo.Keys[j] != nil{
						if izquierdo.Keys[j].Left!= nil{
							izquierdo.Keys[j].Left.NodeF = izquierdo
						}
						if izquierdo.Keys[j].Right!= nil{
							izquierdo.Keys[j].Right.NodeF = izquierdo
						}
					}

				}
				for j:= 0; j< derecho.Max;j++{
					if derecho.Keys[j] != nil{
						if derecho.Keys[j].Left != nil{
							derecho.Keys[j].Left.NodeF = derecho
						}
						if derecho.Keys[j].Right != nil{
							derecho.Keys[j].Right.NodeF = derecho
						}
					}
				}
				placed := this.placeNode(node, centralKey)
				if placed == node.Max-1{
					if node.NodeF == nil{
						indiceCentralRoot := node.Max/2
						centralkeyRoot := node.Keys[indiceCentralRoot]
						izquierdoRoot := NewBTNode(this.k)
						derechoRoot := NewBTNode(this.k)
						rightIndexRoot , leftIndexRoot := 0,0
						for j:= 0; j< node.Max ; j++{
							if node.Keys[j].Value.DPIN < centralkeyRoot.Value.DPIN{
								izquierdoRoot.Put(leftIndexRoot, node.Keys[j])
								leftIndexRoot++
								node.Put(j,nil)
							}else if node.Keys[j].Value.DPIN > centralkeyRoot.Value.DPIN{
								derechoRoot.Put(rightIndexRoot,node.Keys[j])
								rightIndexRoot++
								node.Put(j,nil)
							}
						}
						node.Put(indiceCentralRoot,nil)
						node.Put(0,centralkeyRoot)
						for i:= 0; i<this.k ; i++{
							if izquierdoRoot.Keys[i] != nil{
								izquierdoRoot.Keys[i].Left.NodeF = izquierdoRoot
								izquierdoRoot.Keys[i].Right.NodeF = izquierdoRoot
							}
						}
						for i:= 0; i<this.k ; i++{
							if derechoRoot.Keys[i] != nil{
								derechoRoot.Keys[i].Left.NodeF = derechoRoot
								derechoRoot.Keys[i].Right.NodeF = derechoRoot
							}
						}
						centralkeyRoot.Left = izquierdoRoot
						centralkeyRoot.Right = derechoRoot
						izquierdoRoot.NodeF = node
						derechoRoot.NodeF = node
						this.root = node
					}
					continue
				}else {
					break
				}
			}
		}
	}

}

func (this *BTree)placeNode(node *BTNode, newKey *Key) int {
	index:= -1
	for i:= 0; i < node.Max ;i++{
		if node.Keys[i] == nil{
			placed := false
			for j := i-1; j>= 0 ;j-- {
				if node.Keys[j].Value.DPIN > newKey.Value.DPIN {
					node.Put(j+1,node.Keys[j])
				}else {
					node.Put(j+1,newKey)
					node.Keys[j].Right = newKey.Left
					if j+2 < this.k && node.Keys[j+2] != nil{
						node.Keys[j+2].Left = newKey.Right
					}
					placed = true
					break
				}

			}
			if placed == false{
				node.Put(0,newKey)
				node.Keys[1].Left = newKey.Right
			}
			index = i
			break
		}

	}
	return index
}

func (this *BTree)FindNode(usu *InicioSesion) *UsuariosEncrit{
	var sup *UsuariosEncrit
	if this.root != nil{
		sup = busquedaBTree(this.root, usu, sup)
	}

	return sup
}

func busquedaBTree(root *BTNode, usu *InicioSesion, product *UsuariosEncrit) *UsuariosEncrit {
	if root != nil {
		for _, key := range root.Keys{
			if key!=nil{
				if key.Value.DPI == usu.DPI && key.Value.Password == usu.Password{
					product = &key.Value
					return product
				}
				product = busquedaBTree(key.Right, usu, product)
				product = busquedaBTree(key.Left, usu, product)
			}

		}
	}
	return product
}

func (this *BTree)Graph() {
	cadena := strings.Builder{}
	fmt.Fprintf(&cadena, "digraph BTree{ \n node[shape=record] \n")
	m := make(map[string]*BTNode)
	gra(this.root, &cadena,m,nil,0)
	fmt.Fprintf(&cadena,"}")
	saveBtree(cadena.String())
}

func gra(actual *BTNode, cadena *strings.Builder, arr map[string]*BTNode, father *BTNode,pos int){
	if actual == nil{
		return
	}
	j:= 0
	contiene := arr[fmt.Sprint(&(*actual))]
	 if contiene != nil {
	 	arr[fmt.Sprint(&(*actual))] = nil
		 return
	 } else {
	 	arr[fmt.Sprint(&(*actual))] = actual
	 }
	 fmt.Fprintf(cadena, "node%p[label=\"",&(*actual))
	enlace := true
	for i:=0; i < actual.Max; i++ {
		if actual.Keys[i] == nil{
			return
		}else {
			if enlace{
				if i != actual.Max-1{
					fmt.Fprintf(cadena, "<f%d>|", j)
				}else{
					fmt.Fprintf(cadena,"<f%d>",j)
					break
				}
				enlace = false
				i--
				j++
			}else {
				fmt.Fprintf(cadena, "{<f%d>DPI:%s|Nombre: %s|Correo: %s|Password:%x|Cuenta: %s}|",j,actual.Keys[i].Value.DPI,actual.Keys[i].Value.Nombre,actual.Keys[i].Value.Correo,actual.Keys[i].Value.Password,actual.Keys[i].Value.Cuenta)
				j++
				enlace = true
				if i< actual.Max-1{
					if actual.Keys[i+1] == nil{
						fmt.Fprintf(cadena, "<f%d>",j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cadena,"\"] \n")
	ji := 0
	for i:= 0; i< actual.Max; i++{
		if actual.Keys[i] == nil{
			break
		}
		gra(actual.Keys[i].Left,cadena,arr,actual,ji)
		ji++
		ji++
		gra(actual.Keys[i].Right,cadena,arr,actual,ji)
		ji++
		ji--
	}
	if father != nil{
		fmt.Fprintf(cadena, "node%p:f%d -> node%p \n", &(*father),pos,&(*actual))
	}
}

func saveBtree(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("BTree.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\btree.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\btree.dot -o "+path+"\\Btree\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}
	
}

func (this *BTree)GraphBTreeES() {
	cadena := strings.Builder{}
	fmt.Fprintf(&cadena, "digraph BTreeES{ \n node[shape=record] \n")
	m := make(map[string]*BTNode)
	gra(this.root, &cadena,m,nil,0)
	fmt.Fprintf(&cadena,"}")
	saveBtreeES(cadena.String())
}

func saveBtreeES(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("BTreeES.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\btree.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\btree.dot -o "+path+"\\Btree\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

}

func (this *BTree)GraphBTreeE() {
	cadena := strings.Builder{}
	fmt.Fprintf(&cadena, "digraph BTreeE{ \n node[shape=record] \n")
	m := make(map[string]*BTNode)
	gra(this.root, &cadena,m,nil,0)
	fmt.Fprintf(&cadena,"}")
	saveBtreeE(cadena.String())
}

func saveBtreeE(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("BTreeE.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\btree.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\btree.dot -o "+path+"\\Btree\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

}

