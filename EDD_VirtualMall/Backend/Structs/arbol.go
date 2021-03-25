package Structs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type (
	NodeTree struct {
		Tienda       *string
		Departamento *string
		Calificacion *int
		Productos    Productos
		Factor       int
		Left         *NodeTree
		Right        *NodeTree
	}
	TreeAVL struct {
		Root *NodeTree
	}
)

func NewArbol() *TreeAVL {
	return &TreeAVL{nil}
}

func NewNodeTree(p Productos, tienda *string, depa *string, cali *int) *NodeTree {
	return &NodeTree{
		Tienda:       tienda,
		Departamento: depa,
		Calificacion: cali,
		Productos:    p,
		Factor:       0,
		Left:         nil,
		Right:        nil,
	}
}

func insertar(raiz *NodeTree, p Productos, hc *bool, tienda *string, depa *string, cali *int) *NodeTree {
	var n1 *NodeTree
	if raiz == nil {
		raiz = NewNodeTree(p, tienda, depa, cali)
		*hc = true
	} else if p.Codigo < raiz.Productos.Codigo {
		left := insertar(raiz.Left, p, hc, tienda, depa, cali)
		raiz.Left = left
		if *hc {
			switch raiz.Factor {
			case 1:
				raiz.Factor = 0
				*hc = false
				break
			case 0:
				raiz.Factor = -1
				break
			case -1:
				n1 = raiz.Left
				if n1.Factor == -1 {
					raiz = rotacionII(raiz, n1)
				} else {
					raiz = rotacionID(raiz, n1)
				}
				*hc = false
			}
		}
	} else if p.Codigo > raiz.Productos.Codigo {
		right := insertar(raiz.Right, p, hc, tienda, depa, cali)
		raiz.Right = right
		if *hc {
			switch raiz.Factor {
			case 1:
				n1 = raiz.Right
				if n1.Factor == 1 {
					raiz = rotacionDD(raiz, n1)
				} else {
					raiz = rotacionDI(raiz, n1)
				}
				*hc = false
				break
			case 0:
				raiz.Factor = 1
				break
			case -1:
				raiz.Factor = 0
				*hc = false
			}

		}
	}
	return raiz
}

func rotacionII(n *NodeTree, n1 *NodeTree) *NodeTree {
	n.Left = n1.Right
	n1.Right = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func rotacionDD(n *NodeTree, n1 *NodeTree) *NodeTree {
	n.Right = n1.Left
	n1.Left = n
	if n1.Factor == 1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func rotacionDI(n *NodeTree, n1 *NodeTree) *NodeTree {
	n2 := n1.Left
	n.Right = n2.Left
	n2.Left = n
	n1.Left = n2.Right
	n2.Right = n1
	if n2.Factor == 1 {
		n.Factor = -1
	} else {
		n.Factor = 0
	}
	if n2.Factor == -1 {
		n1.Factor = 1
	} else {
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func rotacionID(n *NodeTree, n1 *NodeTree) *NodeTree {
	n2 := n1.Right
	n.Left = n2.Right
	n2.Right = n
	n1.Right = n2.Left
	n2.Left = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	} else {
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	} else {
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func (this *TreeAVL) GetProducts() []Productos {
	var aux []Productos
	aux = PostOrden(this.Root, aux)

	return aux
}

func PostOrden(aux *NodeTree, arreglo []Productos) []Productos {
	if aux != nil {
		arreglo = PostOrden(aux.Left, arreglo)
		arreglo = PostOrden(aux.Right, arreglo)
		arreglo = append(arreglo, aux.Productos)
	}
	return arreglo
}

func (this *TreeAVL) Insert(p Productos, tienda *string, depa *string, cali *int) {
	b := false
	a := &b
	this.Root = insertar(this.Root, p, a, tienda, depa, cali)
}

func (this *TreeAVL) SearchPrduc(codigo int) *string {
	var product *string
	if this.Root != nil {
		product = busqueda(this.Root, codigo, product)
	}
	return product
}

func busqueda(root *NodeTree, codigo int, product *string) *string {
	if root != nil {
		if codigo == root.Productos.Codigo {
			aux := root.Productos.Nombre + "||" + strconv.Itoa(root.Productos.Codigo)
			return &aux
		}
		product = busqueda(root.Left, codigo, product)
		product = busqueda(root.Right, codigo, product)
	}
	return product
}

func (this *TreeAVL) SearchPrduc2(codigo int) *NodeTree {
	var product *NodeTree
	if this.Root != nil {
		product = busqueda2(this.Root, codigo, product)
	}
	return product
}

func busqueda2(root *NodeTree, codigo int, product *NodeTree) *NodeTree {
	if root != nil {
		if codigo == root.Productos.Codigo {
			return root
		}
		product = busqueda2(root.Left, codigo, product)
		product = busqueda2(root.Right, codigo, product)
	}
	return product
}

func (this *TreeAVL) Generate() {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if this.Root != nil {
		fmt.Fprintf(&cadena, "node%p[label=\"<f0> | <f1> code: %v|<f2> Name:%v |<f3> Mount:%v  | <f4>\" style = filled, fillcolor = darkolivegreen2];\n", &(*this.Root), this.Root.Productos.Codigo, this.Root.Productos.Nombre, this.Root.Productos.Cantidad)
		this.generate(&cadena, (this.Root), this.Root.Left, true)
		this.generate(&cadena, this.Root, this.Root.Right, false)
	}
	fmt.Fprintf(&cadena, "} \n")
	savedot(cadena.String(), this.Root.Productos.Contacto)

}

func (this *TreeAVL) generate(cadena *strings.Builder, padre *NodeTree, actual *NodeTree, izquierda bool) {
	if actual != nil {
		fmt.Fprintf(cadena, "node%p[label=\"<f0>|<f1> code: %v|<f2> Name:%s |<f3> Mount:%v  | <f4>\" style = filled, fillcolor = darkolivegreen2];\n", &(*actual), actual.Productos.Codigo, actual.Productos.Nombre, actual.Productos.Cantidad)
		if izquierda {
			fmt.Fprintf(cadena, "node%p:f0 -> node%p:f2 \n", &(*padre), &(*actual))
		} else {
			fmt.Fprintf(cadena, "node%p:f4 -> node%p:f2 \n", &(*padre), &(*actual))
		}
		this.generate(cadena, actual, actual.Left, true)
		this.generate(cadena, actual, actual.Right, false)
	}
}

func savedot(s string, i string) {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	nombre := string("Arbol" + i + ".png")
	nombre = strings.Replace(nombre, " ", "-", -1)

	_ = ioutil.WriteFile(path+"\\Dots\\arbol.dot", []byte(s), 0644)

	p := "dot -Tpng " + path + "\\Dots\\arbol.dot -o " + path + "\\Imagenes\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}
}
