package main

import (
	"EDD_VirtualMall/Structs"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fernet/fernet-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tiendas2 []Structs.List
var pedidios Structs.ListYear
var botoncitos []string
var Btree *Structs.BTree
var BtreeE *Structs.BTree
var BtreeES *Structs.BTree
var cubix [][]Structs.NodeListas
var sizedep int
var sizeindex int
var departa []string
var indice []string
var cont int
var prueba Structs.Enlace
var grafito Structs.Grafo

func example(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API of EDD, hopefully you enjoy it! :)")

}

func cargaArchivos(w http.ResponseWriter, r *http.Request) {
	var newDoc Structs.Enlace
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	for i, indice := range newDoc.Datos {
		for j, _ := range indice.Departamentos {
			sizedep = j + 1
		}
		sizeindex = i + 1
	}
	cubix = make([][]Structs.NodeListas, sizeindex)
	for i := 0; i < len(newDoc.Datos[0].Departamentos); i++ {
		departa = append(departa, newDoc.Datos[0].Departamentos[i].Nombre)
	}
	for i, datos := range newDoc.Datos {

		indice = append(indice, datos.Indice)

		sup := make([]Structs.NodeListas, sizedep)
		for j, departamentos := range datos.Departamentos {

			for _, tienda := range departamentos.Tiendas {

				aux2 := Structs.Node{
					tienda,
					departamentos.Nombre,
					datos.Indice,
					convertAscii(tienda.Nombre),
					nil,
					nil,
				}

				sup = putStore(aux2, sup, j)
			}
		}
		cubix[i] = sup

	}

	for i := 0; i < sizedep; i++ {
		for j := 0; j < sizeindex; j++ {

			tiendas2 = append(tiendas2, cubix[j][i].Lista1)
			tiendas2 = append(tiendas2, cubix[j][i].Lista2)
			tiendas2 = append(tiendas2, cubix[j][i].Lista3)
			tiendas2 = append(tiendas2, cubix[j][i].Lista4)
			tiendas2 = append(tiendas2, cubix[j][i].Lista5)

		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente las tiendas"}
	json.NewEncoder(w).Encode(error)
}

func CargarProductos(w http.ResponseWriter, r *http.Request) {
	var newDoc Structs.EnlaceInventario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	var finded *Structs.Tiendas
	var position int

	for _, inven := range newDoc.Inventarios {
		sup := Structs.PedidosS{
			Departamento: inven.Departamento,
			Nombre:       inven.Tienda,
			Calificacion: inven.Calificacion,
		}
		position = searchingVectorS(&sup)
		finded = tiendas2[position].Search(&sup)
		if position >= 0 && finded.Nombre != "" {
			if finded.Arbolito == nil {
				finded.Arbolito = Structs.NewArbol()
				for _, product := range inven.Productos {
					product.Tienda = inven.Tienda
					product.Departamento = inven.Departamento
					product.Calificacion = finded.Calificacion
					product.Contacto = finded.Contacto
					finded.Arbolito.Insert(product, &inven.Departamento, &inven.Departamento, &inven.Calificacion)
				}

			}

		} else {

		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente los archivos"}
	json.NewEncoder(w).Encode(error)
}

func CargarPedidos(w http.ResponseWriter, r *http.Request) {
	var newDoc Structs.EnlacePedidos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)

	for _, pedido := range newDoc.Pedidos {
		//fmt.Println(pedido.Fecha)
		supp := Structs.Stack{
			Top:  nil,
			Size: 0,
		}
		aux2 := Structs.NodeStack{
			Value: pedido,
			Next:  nil,
			Prev:  nil,
		}

		supp.Push(&aux2)
		aux3 := supp.First()
		aux := Structs.NodeMatrix{
			StackPedidos: &supp,
			Value:        *aux3,
			Year:         getYear(pedido.Fecha),
			Dia:          getDay(pedido.Fecha),
			Month:        getMonth(pedido.Fecha),
			MonthString:  getStringMonth(getMonth(pedido.Fecha)),
			Ascii:        convertAscii(pedido.Departamento),
			Right:        nil,
			Left:         nil,
			Up:           nil,
			Down:         nil,
		}

		pedidios.AddYear(&aux)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente los archivos"}
	json.NewEncoder(w).Encode(error)
	fmt.Print("asd")

}

func CargarUsuarios(w http.ResponseWriter, r *http.Request)  {

	var newDoc Structs.EnlaceUsuarios
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	mk:= "cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4="
	Btree = Structs.NewBTree(5)
	BtreeE = Structs.NewBTree(5)
	BtreeES = Structs.NewBTree(5)
	for _, usu := range newDoc.Usuarios{
		sup := Structs.UsuariosEncrit{
			DPI:      strconv.Itoa(usu.DPI),
			DPIN:	  usu.DPI,
			Nombre:   usu.Nombre,
			Correo:   usu.Correo,
			Password: EncryptPass(usu.Password),
			Cuenta:   usu.Cuenta,
		}
		aux := Structs.NewKey(sup)
		Btree.Insert(aux)
		sup2 := Structs.UsuariosEncrit{
			DPI:      BtreeCodificado(strconv.Itoa(usu.DPI),mk),
			DPIN:	  usu.DPI,
			Nombre:   usu.Nombre,
			Correo:   BtreeCodificado(usu.Correo,mk),
			Password: EncryptPass(usu.Password),
			Cuenta:   usu.Cuenta,
		}
		aux2:= Structs.NewKey(sup2)
		BtreeES.Insert(aux2)
		sup3 := Structs.UsuariosEncrit{
			DPI:      BtreeCodificado(strconv.Itoa(usu.DPI),mk),
			DPIN:	  usu.DPI,
			Nombre:   BtreeCodificado(usu.Nombre,mk),
			Correo:   BtreeCodificado(usu.Correo,mk),
			Password: EncryptPass(usu.Password),
			Cuenta:   BtreeCodificado(usu.Cuenta,mk),
		}
		aux3:= Structs.NewKey(sup3)
		BtreeE.Insert(aux3)
	}
	Btree.Graph()
	BtreeE.GraphBTreeE()
	BtreeES.GraphBTreeES()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente los Usuarios"}
	json.NewEncoder(w).Encode(error)

}

func CargarGrafo(w http.ResponseWriter, r *http.Request)  {
	var newDoc Structs.EnlaceGrafos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	sup := Structs.NewStack3()
	for _, enlace := range newDoc.Nodos{
		aux:= Structs.NodeStack3{
			Nodo: enlace,
			Next: nil,
			Prev: nil,
		}
		sup.Push3(&aux)
	}
	grafito.Final = newDoc.Entrega
	grafito.Inicio = newDoc.PosicionInicialRobot
	grafito.Nodos = sup
	GraphvizGrafo()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	mensa := Structs.JsonErrors{Mensaje: "Se Creo exitosamente el Grafo"}
	json.NewEncoder(w).Encode(mensa)

}

func CreateUsu(w http.ResponseWriter, r *http.Request)  {
	var newUsu []string
	reqBody, err := ioutil.ReadAll(r.Body)
	mk:= "cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4="
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newUsu)
	
	aux := Structs.UsuariosEncrit{
		DPI:      newUsu[0],
		DPIN:     stringToint(newUsu[0]),
		Nombre:   newUsu[1],
		Correo:   newUsu[2],
		Password: EncryptPass(newUsu[3]),
		Cuenta:   newUsu[4],
	}
	sup := Structs.NewKey(aux)
	Btree.Insert(sup)
	aux2 := Structs.UsuariosEncrit{
		DPI:      BtreeCodificado(newUsu[0],mk),
		DPIN:     stringToint(newUsu[0]),
		Nombre:   newUsu[1],
		Correo:   BtreeCodificado(newUsu[2],mk),
		Password: EncryptPass(newUsu[3]),
		Cuenta:   newUsu[4],
	}
	sup2 := Structs.NewKey(aux2)
	BtreeES.Insert(sup2)
	aux3 := Structs.UsuariosEncrit{
		DPI:      BtreeCodificado(newUsu[0],mk),
		DPIN:     stringToint(newUsu[0]),
		Nombre:   BtreeCodificado(newUsu[1],mk),
		Correo:   BtreeCodificado(newUsu[2],mk),
		Password: EncryptPass(newUsu[3]),
		Cuenta:   BtreeCodificado(newUsu[4],mk),
	}
	sup3:= Structs.NewKey(aux3)
	BtreeE.Insert(sup3)
	Btree.Graph()
	BtreeE.GraphBTreeE()
	BtreeES.GraphBTreeES()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se Creo exitosamente el Usuario"}
	json.NewEncoder(w).Encode(error)
}

func EncryptPass(pass string) string {
	  contr := sha256.Sum256([]byte(pass))
	  aux := string(contr[:])
	  return aux
}

func BtreeCodificado(aux string, mk string) string {
	k := fernet.MustDecodeKeys(mk)
	tok, err := fernet.EncryptAndSign([]byte(aux), k[0])
	if err != nil {
		panic(err)
	}
	aux2 := string(tok[:])
	return aux2
}

func Deletition(w http.ResponseWriter, r *http.Request) {
	var newDoc Structs.PedidosE
	reqBody, err := ioutil.ReadAll(r.Body)
	var contain bool
	var position int
	if err != nil {
		fmt.Fprintf(w, "Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)
	position = searchingVectorE(&newDoc)
	contain = tiendas2[position].Delete(&newDoc)
	if contain && position >= 0 {
		error := Structs.JsonErrors{Mensaje: "The store was deleted succesfully"}
		json.NewEncoder(w).Encode(error)
	} else {
		error := Structs.JsonErrors{Mensaje: "We don´t find the store check your values"}
		json.NewEncoder(w).Encode(error)
	}
	var solo = len(tiendas2)
	fmt.Println(solo)

}

func Search(w http.ResponseWriter, r *http.Request) {
	var newDoc Structs.PedidosS
	reqBody, err := ioutil.ReadAll(r.Body)
	var finded Structs.Tiendas
	var position int
	if err != nil {
		fmt.Fprintf(w, "Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)

	position = searchingVectorS(&newDoc)
	finded = *tiendas2[position].Search(&newDoc)

	if position >= 0 && finded.Nombre != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(finded)
	} else {
		fmt.Fprintf(w, "We don't found that store please check your values")
	}
	//fmt.Print("asd")
}

func ShowList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PosVector, err := strconv.Atoi(vars["id"])
	var stores []Structs.Tiendas
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	var list = tiendas2[PosVector]
	stores = list.Show()
	if len(stores) == 0 {
		fmt.Fprintf(w, "In this list don't exist a store :(")
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stores)
	}
}

func Graphviz(w http.ResponseWriter, r *http.Request) {
	var cadenita strings.Builder
	var nodes []string
	cont := int(0)
	fmt.Fprintf(&cadenita, "digraph G{ \n")
	fmt.Fprintf(&cadenita, "rankdir= \"LR\" \n")
	fmt.Fprintf(&cadenita, "node[fontname=\"Arial\" style=\"filled\" shape=\"record\" color=\"blue\" fillcolor=\"mediumspringgreen\"]; \n")
	for i := 0; i <= len(tiendas2); {
		if cont < 5 {
			if i != len(tiendas2) {
				if tiendas2[i].GetSize() == 0 {
					fmt.Fprintf(&cadenita, "node%d[label=\"vacio | %d \"]; \n ", i, i)
					fmt.Fprintf(&cadenita, "node%dv[label=\" \", color=\"white\" fillcolor=\"white\"] \n", i)
					fmt.Fprintf(&cadenita, "node%d->node%dv; \n", i, i)
					aux := string("node" + strconv.Itoa(i))
					nodes = append(nodes, aux)
				} else {
					fmt.Fprintf(&cadenita, "node%d[style=\"filled\" color=\"blue\" fillcolor=\"mediumspringgreen\" label=< \n", i)
					fmt.Fprintf(&cadenita, "<TABLE BORDER=\"0\" ALIGN=\"LEFT\"> \n")
					fmt.Fprintf(&cadenita, "<TR> \n")
					fmt.Fprintf(&cadenita, "<TD >Indice %v</TD> \n", tiendas2[i].GetFirst().Indice)
					fmt.Fprintf(&cadenita, "<TD BORDER=\"1\"> No. %d </TD> \n", i)
					fmt.Fprintf(&cadenita, "</TR> \n")
					fmt.Fprintf(&cadenita, "<TR> \n")
					fmt.Fprintf(&cadenita, "<TD BORDER=\"1\">%v</TD> \n", tiendas2[i].GetFirst().Departamento)
					fmt.Fprintf(&cadenita, "<TD> Cal. %d </TD> \n", tiendas2[i].GetFirst().Tienda.Calificacion)
					fmt.Fprintf(&cadenita, "</TR> \n")
					fmt.Fprintf(&cadenita, "</TABLE> \n")
					fmt.Fprintf(&cadenita, ">, ]; \n")
					aux := string("node" + strconv.Itoa(i))
					nodes = append(nodes, aux)
					tiendas2[i].Graphic(&cadenita)
					fmt.Fprintf(&cadenita, "node%d->node%p; \n", i, &(*tiendas2[i].GetFirst()))

				}
				cont++
			}
			i++
		} else {
			fmt.Fprintf(&cadenita, "{rank=\"same\" ;"+nodes[0]+" \n")
			for i := 1; i < len(nodes); i++ {
				fmt.Fprintf(&cadenita, ";"+nodes[i]+"\n")
			}
			fmt.Fprintf(&cadenita, " }\n")
			aux2 := string(" ")
			for i, node := range nodes {
				if i == 0 {
					aux2 = node
				} else {
					fmt.Fprintf(&cadenita, "%v -> %v \n", aux2, node)
					aux2 = node
				}
			}
			aux2 = " "
			fmt.Print(aux2)
			fmt.Fprintf(&cadenita, "} \n")
			saveDot(cadenita.String(), i)
			cadenita.Reset()
			fmt.Fprintf(&cadenita, "digraph G{ \n")
			fmt.Fprintf(&cadenita, "rankdir= \"LR\" \n")
			fmt.Fprintf(&cadenita, "node[fontname=\"Arial\" style=\"filled\" shape=\"box\" color=\"blue\" fillcolor=\"mediumspringgreen\"]; \n")
			cont = 0
			nodes = nodes[:0]
		}

	}
	mens := Structs.JsonErrors{Mensaje: "The graphic was created!"}
	json.NewEncoder(w).Encode(mens)

}

func SaveStuff(w http.ResponseWriter, r *http.Request) {
	var stuffdatos []Structs.Datos
	var stuffdepa []Structs.Departamentos2

	for _, indi := range indice {
		for i := 0; i < len(departa); i++ {
			aux2 := []Structs.Tiendas{}
			aux := Structs.Departamentos2{
				Nombre:  departa[i],
				Tiendas: aux2,
				Indice:  indi,
			}
			stuffdepa = append(stuffdepa, aux)

		}
	}

	for _, lista := range tiendas2 {
		l := lista.GetStores()
		if lista.GetSize() > 0 {
			for j, depa := range stuffdepa {
				if depa.Nombre == lista.GetFirst().Departamento && depa.Indice == lista.GetFirst().Indice {
					stuffdepa[j].Tiendas = append(stuffdepa[j].Tiendas, l...)
				}
			}
		}

	}
	for i := 0; i < len(indice); i++ {
		aux3 := Structs.Datos{
			Indice:        "",
			Departamentos: nil,
		}
		for _, depa := range stuffdepa {
			if indice[i] == depa.Indice {
				aux6 := Structs.Departamentos{
					Nombre:  depa.Nombre,
					Tiendas: depa.Tiendas,
				}
				aux3.Indice = indice[i]
				aux3.Departamentos = append(aux3.Departamentos, aux6)

			}
		}
		stuffdatos = append(stuffdatos, aux3)
	}

	var sending Structs.Enlace
	sending.Datos = stuffdatos
	f, _ := json.MarshalIndent(sending, "", " ")
	_ = ioutil.WriteFile("NewStuff.json", f, 0644)
	msg := Structs.JsonErrors{Mensaje: "The json file was created!"}
	json.NewEncoder(w).Encode(msg)

}

func GraphvizGrafo()  {
	var s strings.Builder
	fmt.Fprintf(&s, "digraph Grafito{ \n")
	aux := grafito.Nodos.ArregloVGrafo()
	fmt.Fprintf(&s, "inicio -> %s \n", grafito.Inicio)
	for _, v := range *aux{
		for _, enla := range v.Enlaces{
			fmt.Fprintf(&s, "%s -> %s [label=\" %v \" dir=both] \n",v.Nombre, enla.Nombre,enla.Distancia)
		}
	}
	fmt.Fprintf(&s, "%s -> Final \n", grafito.Final)
	fmt.Fprintf(&s, "}")

	saveDotGrafo(s.String())


}

func CaminoMasCorto(w http.ResponseWriter, r *http.Request)  {

}

func searchingVectorE(pedido *Structs.PedidosE) int {
	var indicefound, depafound bool
	var first, second, result int

	for i, s := range indice {
		if s[0] == pedido.Nombre[0] {
			first = i
			indicefound = true
		}
	}
	for i, s := range departa {
		if s == pedido.Categoria {
			second = i
			depafound = true
		}
	}

	if !indicefound {
		return -1
	}

	if !depafound {
		return -1
	}

	f := second - 0
	s := f*len(indice) + first
	result = s*5 + pedido.Calificacion - 1
	return result

}

func searchingVectorS(pedido *Structs.PedidosS) int {
	var indicefound, depafound bool
	var first, second, result int
	aux := strings.ToUpper(pedido.Nombre)
	for i, s := range indice {
		if s[0] == aux[0] {
			first = i
			indicefound = true
		}
	}
	for i, s := range departa {
		if s == pedido.Departamento {
			second = i
			depafound = true
		}
	}

	if !indicefound {
		return -1
	}

	if !depafound {
		return -1
	}

	f := second - 0
	s := f*len(indice) + first
	result = s*5 + pedido.Calificacion - 1
	return result

}

func putStore(aux2 Structs.Node, sup []Structs.NodeListas, depa int) []Structs.NodeListas {

	switch aux2.Tienda.Calificacion {
	case 1:
		sup[depa].Lista1.SortedInsert(&aux2)
		break
	case 2:
		sup[depa].Lista2.SortedInsert(&aux2)
		break
	case 3:
		sup[depa].Lista3.SortedInsert(&aux2)
		break
	case 4:
		sup[depa].Lista4.SortedInsert(&aux2)
		break
	case 5:
		sup[depa].Lista5.SortedInsert(&aux2)
		break
	}
	return sup
}

func getShops(w http.ResponseWriter, r *http.Request) {
	var ShopList Structs.Shops
	if len(tiendas2) != 0 {
		for i := 0; i < len(tiendas2); i++ {
			aux := tiendas2[i].GetStores()
			for _, tienda := range aux {
				tienda.Key = strconv.Itoa(i) + "$" + tienda.Contacto
				ShopList.Tiendas = append(ShopList.Tiendas, tienda)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		arreglo := ShopList.Tiendas
		json.NewEncoder(w).Encode(arreglo)
	} else {
		err := Structs.JsonErrors{Mensaje: "No hay tiendas cargadas"}
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(err)
	}

}

func getProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	elements := strings.Split(vars["id"], "$")
	vectorPos, _ := strconv.Atoi(elements[0])
	ShopList := tiendas2[vectorPos]
	selectShop := ShopList.GetShop(elements[1])
	products := selectShop.Arbolito.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	arreglo := products
	json.NewEncoder(w).Encode(arreglo)

}

func getUsuario(w http.ResponseWriter, r *http.Request)  {
	var newDoc []string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	aux2:= Structs.InicioSesion{
		DPI:   newDoc[0],
		Password: EncryptPass(newDoc[1]),
	}
	aux:= Btree.FindNode(&aux2)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aux)
}

func getArbolito(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	elements := strings.Split(vars["id"], "$")
	vectorPos, _ := strconv.Atoi(elements[0])
	ShopList := tiendas2[vectorPos]
	selectShop := ShopList.GetShop(elements[1])
	selectShop.Arbolito.Generate()
}

func getMatrix(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	elements := strings.Split(vars["id"], "-")
	year, err := strconv.Atoi(elements[0])
	month, err2 := strconv.Atoi(elements[1])
	if err != nil && err2 != nil {
		fmt.Println("error")
	}
	monthUsed := pedidios.SearchYear(year)
	sperceMatrix := monthUsed.Monts.SearchMonth(month)
	sperceMatrix.Matrix.Graphviz()
	w.WriteHeader(http.StatusOK)
}

func getPedidos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := vars["id"]
	is := strings.Split(i, "$")
	var datos []string

	var finded *Structs.Tiendas
	var position int
	datos = strings.Split(is[0], "-")
	aux, _ := strconv.Atoi(datos[0])
	monthUsed := pedidios.SearchYear(aux)
	aux, _ = strconv.Atoi(datos[1])
	sperceMatrix := monthUsed.Monts.SearchMonth(aux)
	aux2, _ := strconv.Atoi(is[1])
	aux3 := strings.Replace(is[2], "_", " ", -1)

	pedidos := sperceMatrix.Matrix.ObtenerNodito(aux2, convertAscii(aux3))
	arregloPedidos := pedidos.StackPedidos.ArregloPedidos()
	var pedidosEnviar []Structs.ShowingPedidos
	for _, pedi := range arregloPedidos {
		sup := Structs.PedidosS{
			Departamento: pedi.Departamento,
			Nombre:       pedi.Tienda,
			Calificacion: pedi.Calificacion,
		}
		position = searchingVectorS(&sup)
		finded = tiendas2[position].Search(&sup)
		var sup3 []string
		sup2 := Structs.ShowingPedidos{
			Fecha:        pedi.Fecha,
			Tiendas:      pedi.Tienda,
			Departamento: pedi.Departamento,
			Cliente: pedi.Cliente,
			Producto:     sup3,
			CaminmoCorto: "",
		}
		for _, codProd := range pedi.Productos {
			aux4 := finded.Arbolito.SearchPrduc(codProd.Codigo)
			sup2.Producto = append(sup2.Producto, *aux4)
		}
		pedidosEnviar = append(pedidosEnviar, sup2)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedidosEnviar)
}

func saveDot(s string, i int) {

	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("lista" + strconv.Itoa(i) + ".png")
	botoncitos = append(botoncitos, nombre)

	_ = ioutil.WriteFile(path+"\\Dots\\arreglo.dot",[]byte(s),0644)

	p := "dot -Tpng " + path +"\\Dots\\arreglo.dot -o "+path+"\\Arreglo\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}
}
func saveDotGrafo(s string) {

	path, err := os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	nombre := string("Grafo.png")
	botoncitos = append(botoncitos, nombre)

	_ = ioutil.WriteFile(path+"\\Dots\\grafo.dot",[]byte(s),0644)

	p := "circo -Tpng " + path +"\\Dots\\grafo.dot -o "+path+"\\Grafo\\" + nombre
	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)
		fmt.Printf("%s\n", b)
	}
}

func convertAscii(s string) int {
	ascii := int(0)
	runes := []rune(s)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	for i := 0; i < len(result); i++ {
		ascii = ascii + result[i]
	}

	return ascii
}

func getDay(s string) int {
	aux := strings.Split(s, "-")
	i, err := strconv.Atoi(aux[0])
	if err != nil {
		fmt.Println("Na")
	}
	return i
}

func getYear(s string) int {
	aux := strings.Split(s, "-")
	i, err := strconv.Atoi(aux[2])
	if err != nil {
		fmt.Println("Na")
	}
	return i
}

func getMonth(s string) int {
	aux := strings.Split(s, "-")
	i, err := strconv.Atoi(aux[1])
	if err != nil {
		fmt.Println("Na")
	}
	return i
}

func getStringMonth(s int) string {
	switch s {
	case 1:
		return "Enero"
		break
	case 2:
		return "Febrero"
		break
	case 3:
		return "Marzo"
		break
	case 4:
		return "Abril"
		break
	case 5:
		return "Mayo"
		break
	case 6:
		return "Junio"
		break
	case 7:
		return "Julio"
		break
	case 8:
		return "Agosto"
		break
	case 9:
		return "Septiembre"
		break
	case 10:
		return "Octubre"
		break
	case 11:
		return "Noviembre"
		break
	case 12:
		return "Diciembre"
		break
	}
	return "0"
}

func mostrarImagenArbolito(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := vars["id"]
	http.ServeFile(w, r, "Imagenes/Arbol"+i+".png")

}

func mostrarImagenMatriz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := vars["id"]
	http.ServeFile(w, r, "Matrices/Matriz"+i+".png")

}

func mostrarArregloL(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	i, _ := vars["id"]
	http.ServeFile(w, r, "Arreglo/"+i)

}

func mostrarBTrees(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	i, _ := vars["id"]
	http.ServeFile(w, r, "Btree/"+i)

}

func getYears(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	arreglo := pedidios.GetYeats()
	json.NewEncoder(w).Encode(arreglo)

}

func getBotones(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(botoncitos)
}

func getBTrees(w http.ResponseWriter, r *http.Request)  {
	b := [3]string{"BTree.png","BTreeE.png","BTreeES.png"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func carritoPedidos(w http.ResponseWriter, r *http.Request) {
	var newDoc [][]string
	var pedidos []Structs.Carrito
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	for i := 0; i < len(newDoc); i++ {
		mount, _ := strconv.Atoi(newDoc[i][3])
		precio, _ := strconv.ParseFloat(newDoc[i][4], 64)
		idi, _ := strconv.Atoi(newDoc[i][5])
		cali, _ := strconv.Atoi(newDoc[i][8])
		canti, _ := strconv.Atoi(newDoc[i][9])
		aux := Structs.Carrito{
			Nombre:       newDoc[i][0],
			Descripcion:  newDoc[i][1],
			Image:        newDoc[i][2],
			Mount:        mount,
			Price:        precio,
			Id:           idi,
			Tienda:       newDoc[i][6],
			Departamento: newDoc[i][7],
			Calificacion: cali,
			Cantidad:     canti,
			Fecha:        newDoc[i][10],
			Cliente:      stringToint(newDoc[i][11]),
		}
		pedidos = append(pedidos, aux)
	}
	retirarArbol(&pedidos)
	agregarPedido(&pedidos)

}

func retirarArbol(pedidos *[]Structs.Carrito) {
	var finded *Structs.Tiendas
	var position int
	var hojita *Structs.NodeTree
	for _, carrito := range *pedidos {
		sup := Structs.PedidosS{
			Departamento: carrito.Departamento,
			Nombre:       carrito.Tienda,
			Calificacion: carrito.Calificacion,
		}
		position = searchingVectorS(&sup)
		finded = tiendas2[position].Search(&sup)
		hojita = finded.Arbolito.SearchPrduc2(carrito.Id)
		hojita.Productos.Cantidad = hojita.Productos.Cantidad - carrito.Cantidad
	}
}

func agregarPedido(pedidos *[]Structs.Carrito) {
	stack := Structs.NewStack2()

	for _, carrito := range *pedidos {
		var sup2 []Structs.CodProducto
		sup2 = append(sup2, Structs.CodProducto{Codigo: carrito.Id})
		sup := Structs.ValidarPedidos{
			Tienda:       carrito.Tienda,
			Departamento: carrito.Departamento,
			Calificacion: carrito.Calificacion,
			Cliente:       carrito.Cliente,
			Producto:     Structs.CodProducto{Codigo: carrito.Id},
			Productos:    sup2,
		}
		aux := Structs.NodeStack2{
			Pedido: sup,
			Next:   nil,
			Prev:   nil,
		}
		if stack.VerificarExsite(&aux) {

		} else {
			stack.Push2(&aux)
		}
	}
	aux := stack.ArregloVPedidos()
	var aux2 []Structs.Pedidos
	for _, p := range *aux {
		sup := Structs.Pedidos{
			Fecha:        "",
			Cliente:       p.Cliente,
			Tienda:       p.Tienda,
			Departamento: p.Departamento,
			Calificacion: p.Calificacion,
			Productos:    nil,
		}
		aux2 = append(aux2, sup)
	}
	for i, p := range aux2 {
		var aux3 []Structs.CodProducto
		for _, r := range *pedidos {
			if p.Calificacion == r.Calificacion && p.Departamento == r.Departamento && p.Calificacion == r.Calificacion && p.Cliente == r.Cliente {
				aux4 := Structs.CodProducto{Codigo: r.Id}
				aux3 = append(aux3, aux4)
			}
			aux2[i].Fecha = r.Fecha
		}
		aux2[i].Productos = aux3
	}

	for _, pedido := range aux2 {
		supp := Structs.Stack{
			Top:  nil,
			Size: 0,
		}
		aux4 := Structs.NodeStack{
			Value: pedido,
			Next:  nil,
			Prev:  nil,
		}

		supp.Push(&aux4)
		aux3 := supp.First()
		aux5 := Structs.NodeMatrix{
			StackPedidos: &supp,
			Value:        *aux3,
			Year:         getYear(pedido.Fecha),
			Dia:          getDay(pedido.Fecha),
			Month:        getMonth(pedido.Fecha),
			MonthString:  getStringMonth(getMonth(pedido.Fecha)),
			Ascii:        convertAscii(pedido.Departamento),
			Right:        nil,
			Left:         nil,
			Up:           nil,
			Down:         nil,
		}

		pedidios.AddYear(&aux5)

	}

}

func stringToint(cadena string) int {
	numero, _ := strconv.Atoi(cadena)
	return numero
}

func intTostring(numero int) string {
	cadena := strconv.Itoa(numero)
	return cadena
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", example).Methods("GET")
	router.HandleFunc("/cargartienda", cargaArchivos).Methods("POST")
	router.HandleFunc("/cargarusuarios", CargarUsuarios).Methods("POST")
	router.HandleFunc("/cargarGrafo", CargarGrafo).Methods("POST")
	router.HandleFunc("/Eliminar", Deletition).Methods("DELETE")
	router.HandleFunc("/TiendaEspecifica", Search).Methods("POST")
	router.HandleFunc("/CreateUsu", CreateUsu).Methods("POST")
	router.HandleFunc("/id/{id}", ShowList).Methods("GET")
	router.HandleFunc("/getArreglo", Graphviz).Methods("GET")
	router.HandleFunc("/guardar", SaveStuff).Methods("GET")
	router.HandleFunc("/obtenerTiendas", getShops).Methods("GET")
	router.HandleFunc("/obtenerTiendas/{id}", getProducts).Methods("GET")
	router.HandleFunc("/obtenerArbolito/{id}", getArbolito).Methods("GET")
	router.HandleFunc("/obtenerMatriz/{id}", getMatrix).Methods("GET")
	router.HandleFunc("/obtenerYears", getYears).Methods("GET")
	router.HandleFunc("/obtenerUsu", getUsuario).Methods("POST")
	router.HandleFunc("/obtenerPedidos/{id}", getPedidos).Methods("GET")
	router.HandleFunc("/botoncitos", getBotones).Methods("GET")
	router.HandleFunc("/btrees", getBTrees).Methods("GET")
	router.HandleFunc("/cargarproductos", CargarProductos).Methods("POST")
	router.HandleFunc("/carrito", carritoPedidos).Methods("POST")
	router.HandleFunc("/cargarpedidos", CargarPedidos).Methods("POST")
	router.HandleFunc("/arbolito/{id}", mostrarImagenArbolito)
	router.HandleFunc("/matriz/{id}", mostrarImagenMatriz)
	router.HandleFunc("/arreglito/{id}", mostrarArregloL)
	router.HandleFunc("/arbolitosb/{id}", mostrarBTrees)


	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
}
