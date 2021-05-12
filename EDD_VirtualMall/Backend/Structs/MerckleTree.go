package Structs

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
)

type (
	NodoMerckle struct{
		valor TransaccionPedidos
		right *NodoMerckle
		left *NodoMerckle
	}

	MerckleTree struct {
		root *NodoMerckle
	}

	NodoMerckleTiendas struct{
		valor TransaccionTiendas
		right *NodoMerckleTiendas
		left *NodoMerckleTiendas
	}

	MerckleTreeTiendas struct {
		root *NodoMerckleTiendas
	}
	NodoMerckleProductos struct{
		valor TransaccionProductos
		right *NodoMerckleProductos
		left *NodoMerckleProductos
	}

	MerckleTreeProductos struct {
		root *NodoMerckleProductos
	}
	NodoMerckleUsuarios struct{
		valor TransaccionUsuarios
		right *NodoMerckleUsuarios
		left *NodoMerckleUsuarios
	}

	MerckleTreeUsuarios struct {
		root *NodoMerckleUsuarios
	}

)

/*	----------------  Arbol De Pedidos	------------------	*/

func(this *NodoMerckle)suma()string {
	if this.right != nil && this.left != nil{
		return this.right.valor.Id + this.left.valor.Id
	}
	return "-1"
}

func NewnodoMerckle(valor TransaccionPedidos, right *NodoMerckle,left *NodoMerckle) *NodoMerckle{
	return &NodoMerckle{
		valor: valor,
		right: right,
		left:  left,
	}
}

func NewArbolMerckle() *MerckleTree{
	return &MerckleTree{}
}

func (this *MerckleTree)Insertar(valor TransaccionPedidos) {
	n := NewnodoMerckle(valor,nil,nil)
	if this.root == nil{
		lista:= list.New()
		lista.PushBack(n)
		aux := TransaccionPedidos{
			Id:    Sha256("-1"),
			Sha: "-1",
			Sha2: "",
			Dpi:   0,
			Fecha: "",
			Monto: 0,
		}
		lista.PushBack(NewnodoMerckle(aux, nil,nil))
		this.contruirarbol(lista)
	}else{
		lista:= this.ObtenerLista()
		lista.PushBack(n)
		this.contruirarbol(lista)

	}
}

func (this *MerckleTree)ObtenerLista()*list.List  {
	lista := list.New()
	obtenerLista(lista, this.root.left)
	obtenerLista(lista, this.root.right)

	return lista
}

func obtenerLista(lista *list.List, actual *NodoMerckle)  {
	if actual != nil{
		obtenerLista(lista, actual.left)
		if actual.right == nil && actual.valor.Sha != "-1"{
			lista.PushBack(actual)
		}
		obtenerLista(lista, actual.right)
	}
}

func (this *MerckleTree)contruirarbol(lista *list.List){
	size:= float64(lista.Len())
	cant := 1
	for (size/2) > 1{
		cant++
		size=size/2
	}
	nodostotales:= math.Pow(2,float64(cant))
	for lista.Len() < int(nodostotales){
		aux := TransaccionPedidos{
			Id:    Sha256("-1"),
			Sha: "-1",
			Sha2: "",
			Dpi:   0,
			Fecha: "",
			Monto: 0,
		}
		lista.PushBack(NewnodoMerckle(aux,nil,nil))

	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()

		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoMerckle)
		nodo2 := segundo.Value.(*NodoMerckle)

		//aux2 := nodo1.valor.Id + "\n" + nodo2.valor.Id
		aux := TransaccionPedidos{
			Id:    Sha256(nodo1.valor.Id + nodo2.valor.Id),
			Sha: nodo1.valor.Id,
			Sha2: nodo2.valor.Id,
			Dpi:   0,
			Fecha: "",
			Monto: 0,
		}
		nuevo :=  NewnodoMerckle(aux, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.root = lista.Front().Value.(*NodoMerckle)
}

func (this *MerckleTree) GrafiquitaPedidos() {
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph{ \n")
	fmt.Fprint(&cadena, "node[shape=\"record\"]; \n")
	if(this.root != nil){
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1> %x | %x | %x}| <f2> \"fillcolor=\"olivedrab1\"]; \n", &(*this.root),this.root.valor.Id,this.root.valor.Sha,this.root.valor.Sha2)
		this.genereaPedidos(&cadena, (this.root), this.root.left, true)
		this.genereaPedidos(&cadena, this.root, this.root.right, false)
	}
	fmt.Fprintf(&cadena, "} \n")

	saveMercklePedidos(cadena.String())

}

func (this *MerckleTree) genereaPedidos(cadena *strings.Builder, padre *NodoMerckle, actual *NodoMerckle, izquierda bool)  {
	if actual != nil{
		if actual.valor.Fecha != "" {
			fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x| id:%v| fecha:%v }| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id,actual.valor.Dpi,actual.valor.Fecha)
			if izquierda {
				fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
			}else{
				fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
			}

		}else {
			if actual.valor.Sha == "-1"{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %v | %v}| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}else{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %x | %x}| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}


		}
		this.genereaPedidos(cadena, actual, actual.left, true)
		this.genereaPedidos(cadena, actual, actual.right, false)

	}


}

func saveMercklePedidos(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("MercklePedidos.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\merckle.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\merckle.dot -o "+path+"\\Merckle\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

}


/*  ---------------------	  Arbol de Producto   -------------------   */

func(this *NodoMerckleProductos)sumaProductos()string {
	if this.right != nil && this.left != nil{
		return this.right.valor.Id+ "\n" + this.left.valor.Id
	}
	return "-1"
}

func NewnodoMerckleProductos(valor TransaccionProductos, right *NodoMerckleProductos,left *NodoMerckleProductos) *NodoMerckleProductos{
	return &NodoMerckleProductos{
		valor: valor,
		right: right,
		left:  left,
	}
}

func NewArbolMerckleProductos() *MerckleTreeProductos{
	return &MerckleTreeProductos{}
}

func (this *MerckleTreeProductos)InsertarProductos(valor TransaccionProductos) {
	n := NewnodoMerckleProductos(valor,nil,nil)
	if this.root == nil{
		lista:= list.New()
		lista.PushBack(n)
		aux := TransaccionProductos{
			Id:             Sha256("-1"),
			Sha:            "-1",
			Sha2: 			"",
			Accion:         "",
			Tienda:         "",
			Departamento:   "",
			Calificacion:   0,
			Nombre:         "",
			Codigo:         0,
			Descripcion:    "",
			Precio:         0,
			Cantidad:       0,
			Imagen:         "",
			Almacenamiento: "",
		}
		lista.PushBack(NewnodoMerckleProductos(aux, nil,nil))
		this.contruirarbolProductos(lista)
	}else{
		lista:= this.ObtenerListaProductos()
		lista.PushBack(n)
		this.contruirarbolProductos(lista)

	}
}

func (this *MerckleTreeProductos)ObtenerListaProductos()*list.List  {
	lista := list.New()
	obtenerListaProductos(lista, this.root.left)
	obtenerListaProductos(lista, this.root.right)

	return lista
}

func obtenerListaProductos(lista *list.List, actual *NodoMerckleProductos)  {
	if actual != nil{
		obtenerListaProductos(lista, actual.left)
		if actual.right == nil && actual.valor.Sha != "-1"{
			lista.PushBack(actual)
		}
		obtenerListaProductos(lista, actual.right)
	}
}

func (this *MerckleTreeProductos)contruirarbolProductos(lista *list.List){
	size:= float64(lista.Len())
	cant := 1
	for (size/2) > 1{
		cant++
		size=size/2
	}
	nodostotales:= math.Pow(2,float64(cant))
	for lista.Len() < int(nodostotales){
		aux := TransaccionProductos{
			Id:             Sha256("-1"),
			Sha:            "-1",
			Sha2: 			"",
			Accion:         "",
			Tienda:         "",
			Departamento:   "",
			Calificacion:   0,
			Nombre:         "",
			Codigo:         0,
			Descripcion:    "",
			Precio:         0,
			Cantidad:       0,
			Imagen:         "",
			Almacenamiento: "",
		}
		lista.PushBack(NewnodoMerckleProductos(aux,nil,nil))

	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()

		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoMerckleProductos)
		nodo2 := segundo.Value.(*NodoMerckleProductos)



		aux := TransaccionProductos{
			Id:             Sha256(nodo1.valor.Id+nodo2.valor.Id),
			Sha:            nodo1.valor.Id,
			Sha2: 			nodo2.valor.Id,
			Accion:         "",
			Tienda:         "",
			Departamento:   "",
			Calificacion:   0,
			Nombre:         "",
			Codigo:         0,
			Descripcion:    "",
			Precio:         0,
			Cantidad:       0,
			Imagen:         "",
			Almacenamiento: "",
		}
		nuevo :=  NewnodoMerckleProductos(aux, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.root = lista.Front().Value.(*NodoMerckleProductos)
}

func (this *MerckleTreeProductos) GrafiquitaProductos() {
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph{ \n")
	fmt.Fprint(&cadena, "node[shape=\"record\"]; \n")
	if(this.root != nil){
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1> %x | %x | %x}| <f2> \" fillcolor=\"olivedrab1\"]; \n", &(*this.root),this.root.valor.Id,this.root.valor.Sha,this.root.valor.Sha2)
		this.genereaProductos(&cadena, (this.root), this.root.left, true)
		this.genereaProductos(&cadena, this.root, this.root.right, false)
	}
	fmt.Fprintf(&cadena, "} \n")
	saveMerckleProductos(cadena.String())

}

func (this *MerckleTreeProductos) genereaProductos(cadena *strings.Builder, padre *NodoMerckleProductos, actual *NodoMerckleProductos, izquierda bool)  {
	if actual != nil{
		if actual.valor.Accion != "" {
			fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x| %v }| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual),	actual.valor.Id,actual.valor.Accion)
			if izquierda {
				fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
			}else{
				fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
			}

		}else {
			if actual.valor.Sha == "-1"{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %v | %v}| <f2> \" fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}else{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %x | %x}| <f2> \" fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}


		}
		this.genereaProductos(cadena, actual, actual.left, true)
		this.genereaProductos(cadena, actual, actual.right, false)

	}


}
func saveMerckleProductos(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("MerckleProductos.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\merckle.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\merckle.dot -o "+path+"\\Merckle\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

}

/*  ----------------- 	 Arbol de Tiendas  ----------------------- 	*/

func(this *NodoMerckleTiendas)sumaTiendas()string {
	if this.right != nil && this.left != nil{
		return this.right.valor.Id+ "\n" + this.left.valor.Id
	}
	return "-1"
}

func NewnodoMerckleTiendas(valor TransaccionTiendas, right *NodoMerckleTiendas,left *NodoMerckleTiendas) *NodoMerckleTiendas{
	return &NodoMerckleTiendas{
		valor: valor,
		right: right,
		left:  left,
	}
}

func NewArbolMerckleTiendas() *MerckleTreeTiendas{
	return &MerckleTreeTiendas{}
}

func (this *MerckleTreeTiendas)InsertarUsuario(valor TransaccionTiendas) {
	n := NewnodoMerckleTiendas(valor,nil,nil)
	if this.root == nil{
		lista:= list.New()
		lista.PushBack(n)
		aux := TransaccionTiendas{
			Id:           Sha256("-1"),
			Sha:          "-1",
			Sha2:         "",
			Accion:       "",
			Nombre:       "",
			Departamento: "",
			Descripcion:  "",
			Contacto:     "",
			Calificacion: 0,
			Logo:         "",
		}
		lista.PushBack(NewnodoMerckleTiendas(aux, nil,nil))
		this.contruirarbolTienda(lista)
	}else{
		lista:= this.ObtenerListaTienda()
		lista.PushBack(n)
		this.contruirarbolTienda(lista)

	}
}

func (this *MerckleTreeTiendas)ObtenerListaTienda()*list.List  {
	lista := list.New()
	obtenerListaTiendas(lista, this.root.left)
	obtenerListaTiendas(lista, this.root.right)

	return lista
}

func obtenerListaTiendas(lista *list.List, actual *NodoMerckleTiendas)  {
	if actual != nil{
		obtenerListaTiendas(lista, actual.left)
		if actual.right == nil && actual.valor.Sha != "-1"{
			lista.PushBack(actual)
		}
		obtenerListaTiendas(lista, actual.right)
	}
}

func (this *MerckleTreeTiendas)contruirarbolTienda(lista *list.List){
	size:= float64(lista.Len())
	cant := 1
	for (size/2) > 1{
		cant++
		size=size/2
	}
	nodostotales:= math.Pow(2,float64(cant))
	for lista.Len() < int(nodostotales){
		aux := TransaccionTiendas{
			Id:           Sha256("-1"),
			Sha:          "-1",
			Sha2:         "",
			Accion:       "",
			Nombre:       "",
			Departamento: "",
			Descripcion:  "",
			Contacto:     "",
			Calificacion: 0,
			Logo:         "",
		}
		lista.PushBack(NewnodoMerckleTiendas(aux,nil,nil))

	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()

		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoMerckleTiendas)
		nodo2 := segundo.Value.(*NodoMerckleTiendas)

		//aux2 := nodo1.valor.Id + "\n" + nodo2.valor.Id
		aux := TransaccionTiendas{
			Id:           Sha256(nodo1.valor.Id+nodo2.valor.Id),
			Sha:          nodo1.valor.Id,
			Sha2:         nodo2.valor.Id,
			Accion:       "",
			Nombre:       "",
			Departamento: "",
			Descripcion:  "",
			Contacto:     "",
			Calificacion: 0,
			Logo:         "",
		}
		nuevo :=  NewnodoMerckleTiendas(aux, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.root = lista.Front().Value.(*NodoMerckleTiendas)
}

func (this *MerckleTreeTiendas) GrafiquitaUsuarios() {
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph{ \n")
	fmt.Fprint(&cadena, "node[shape=\"record\"]; \n")
	if(this.root != nil){
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1> %x | %x | %x}| <f2> \"fillcolor=\"olivedrab1\"]; \n", &(*this.root),this.root.valor.Id,this.root.valor.Sha,this.root.valor.Sha2)
		this.genereaTiendas(&cadena, (this.root), this.root.left, true)
		this.genereaTiendas(&cadena, this.root, this.root.right, false)
	}
	fmt.Fprintf(&cadena, "} \n")
	saveMerckleTiendas(cadena.String())

}

func (this *MerckleTreeTiendas) genereaTiendas(cadena *strings.Builder, padre *NodoMerckleTiendas, actual *NodoMerckleTiendas, izquierda bool)  {
	if actual != nil{
		if actual.valor.Accion != "" {
			fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x| %v }| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual),
				actual.valor.Id,actual.valor.Accion)
			if izquierda {
				fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
			}else{
				fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
			}

		}else {
			if(actual.valor.Sha == "-1"){
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %v | %v}| <f2> \" fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1\n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}else{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %x | %x}| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}


		}
		this.genereaTiendas(cadena, actual, actual.left, true)
		this.genereaTiendas(cadena, actual, actual.right, false)

	}


}

func saveMerckleTiendas(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("MerckleTiendas.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\merckle.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\merckle.dot -o "+path+"\\Merckle\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

}



/*  --------------------	 Arbol de Usuarios	 ------------------------  */

func(this *NodoMerckleUsuarios)sumaUsuarios()string {
	if this.right != nil && this.left != nil{
		return this.right.valor.Id+ "\n" + this.left.valor.Id
	}
	return "-1"
}

func NewnodoMerckleUsuarios(valor TransaccionUsuarios, right *NodoMerckleUsuarios,left *NodoMerckleUsuarios) *NodoMerckleUsuarios{
	return &NodoMerckleUsuarios{
		valor: valor,
		right: right,
		left:  left,
	}
}

func NewArbolMerckleUsuarios() *MerckleTreeUsuarios{
	return &MerckleTreeUsuarios{}
}

func (this *MerckleTreeUsuarios)InsertarUsu(valor TransaccionUsuarios) {
	n := NewnodoMerckleUsuarios(valor,nil,nil)
	if this.root == nil{
		lista:= list.New()
		lista.PushBack(n)
		aux := TransaccionUsuarios{
			Id:     Sha256("-1"),
			Sha:    "-1",
			Sha2:   "",
			Accion: "",
			DPI:    0,
			Nombre: "",
			Correo: "",
			Pass:   "",
			Cuenta: "",
		}
		lista.PushBack(NewnodoMerckleUsuarios(aux, nil,nil))
		this.contruirarbolUsu(lista)
	}else{
		lista:= this.ObtenerListaUsu()
		lista.PushBack(n)
		this.contruirarbolUsu(lista)

	}
}

func (this *MerckleTreeUsuarios)ObtenerListaUsu()*list.List  {
	lista := list.New()
	obtenerListaUsu(lista, this.root.left)
	obtenerListaUsu(lista, this.root.right)

	return lista
}

func obtenerListaUsu(lista *list.List, actual *NodoMerckleUsuarios)  {
	if actual != nil{
		obtenerListaUsu(lista, actual.left)
		if actual.right == nil && actual.valor.Sha != "-1"{
			lista.PushBack(actual)
		}
		obtenerListaUsu(lista, actual.right)
	}
}

func (this *MerckleTreeUsuarios)contruirarbolUsu(lista *list.List){
	size:= float64(lista.Len())
	cant := 1
	for (size/2) > 1{
		cant++
		size=size/2
	}
	nodostotales:= math.Pow(2,float64(cant))
	for lista.Len() < int(nodostotales){
		aux := TransaccionUsuarios{
			Id:     Sha256("-1"),
			Sha:    "-1",
			Sha2:   "",
			Accion: "",
			DPI:    0,
			Nombre: "",
			Correo: "",
			Pass:   "",
			Cuenta: "",
		}
		lista.PushBack(NewnodoMerckleUsuarios(aux,nil,nil))

	}
	for lista.Len() > 1{
		primero := lista.Front()
		segundo := primero.Next()

		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoMerckleUsuarios)
		nodo2 := segundo.Value.(*NodoMerckleUsuarios)

		//aux2 := nodo1.valor.Id + "\n" + nodo2.valor.Id
		aux := TransaccionUsuarios{
			Id:     Sha256(nodo1.valor.Id+nodo2.valor.Id),
			Sha:    nodo1.valor.Id,
			Sha2:   nodo2.valor.Id,
			Accion: "",
			DPI:    0,
			Nombre: "",
			Correo: "",
			Pass:   "",
			Cuenta: "",
		}
		nuevo :=  NewnodoMerckleUsuarios(aux, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	this.root = lista.Front().Value.(*NodoMerckleUsuarios)
}

func (this *MerckleTreeUsuarios) GrafiquitaPedidos() {
	var cadena strings.Builder
	fmt.Fprint(&cadena, "digraph{ \n")
	fmt.Fprint(&cadena, "node[shape=\"record\"]; \n")
	if this.root != nil{
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1> %x | %x | %x}| <f2> \" fillcolor=\"olivedrab1\"]; \n", &(*this.root),this.root.valor.Id,this.root.valor.Sha,this.root.valor.Sha2)
		this.genereaUsu(&cadena, (this.root), this.root.left, true)
		this.genereaUsu(&cadena, this.root, this.root.right, false)
	}
	fmt.Fprintf(&cadena, "} \n")
	saveMerckleUsuarios(cadena.String())

}

func (this *MerckleTreeUsuarios) genereaUsu(cadena *strings.Builder, padre *NodoMerckleUsuarios, actual *NodoMerckleUsuarios, izquierda bool)  {
	if actual != nil{
		if actual.valor.Accion != "" {
			fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x| %v }| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual),
				actual.valor.Id,actual.valor.Accion)
			if izquierda {
				fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
			}else{
				fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
			}

		}else {
			if actual.valor.Sha == "-1"{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %v | %v}| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}else{
				fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1> id:%x | %x | %x}| <f2> \"fillcolor=\"olivedrab1\"] \n", &(*actual), actual.valor.Id, actual.valor.Sha, actual.valor.Sha2)
				if izquierda {
					fmt.Fprintf(cadena,"node%p:f0 -> node%p:f1 \n", &(*padre), &(*actual) )
				}else{
					fmt.Fprintf(cadena,"node%p:f2 -> node%p:f1 \n", &(*padre), &(*actual) )
				}
			}


		}
		this.genereaUsu(cadena, actual, actual.left, true)
		this.genereaUsu(cadena, actual, actual.right, false)

	}


}

func saveMerckleUsuarios(s string)  {
	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("MerckleUsuarios.png")
	nombre = strings.Replace(nombre," ","-",-1)

	_ = ioutil.WriteFile(path+"\\Dots\\merckle.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\merckle.dot -o "+path+"\\Merckle\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}

}



/*	----------- Graficar Pedidos---------------	*/



/*  SHA 256  */

func Sha256(pass string) string {
	contr := sha256.Sum256([]byte(pass))
	aux := string(contr[:])
	return aux
}