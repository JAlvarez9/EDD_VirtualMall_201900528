package main

import (
	"EDD_VirtualMall/Structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


var tiendas2 [] Structs.List
var pedidios Structs.ListYear
var cubix [][] Structs.NodeListas
var sizedep int
var sizeindex int
var departa [] string
var indice [] string
var cont int
var prueba Structs.Enlace

func example(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to my REST API of EDD, hopefully you enjoy it! :)")
	aux := int(06)
	fmt.Println(aux)


}

func cargaArchivos(w http.ResponseWriter, r *http.Request){
	var newDoc Structs.Enlace
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	for i, indice := range newDoc.Datos{
		for j, _ := range indice.Departamentos{
			sizedep = j+1
		}
		sizeindex = i+1
	}
	cubix = make([][] Structs.NodeListas, sizeindex)
	for i:= 0; i < len(newDoc.Datos[0].Departamentos); i++{
		departa = append(departa, newDoc.Datos[0].Departamentos[i].Nombre )
	}
	for i, datos := range newDoc.Datos{

		indice = append(indice, datos.Indice)

		sup := make([]Structs.NodeListas, sizedep)
		for j, departamentos := range datos.Departamentos{

			for _, tienda := range departamentos.Tiendas{

				aux2:= Structs.Node{
					tienda,
					departamentos.Nombre,
					datos.Indice,
					convertAscii(tienda.Nombre),
					nil,
					nil,
				}

				sup = putStore(aux2,sup,j)
			}
		}
		cubix[i] = sup

	}

	for i:= 0; i< sizedep; i++{
		for j:=0; j< sizeindex; j++{

			tiendas2 = append(tiendas2, cubix[j][i].Lista1)
			tiendas2 = append(tiendas2, cubix[j][i].Lista2)
			tiendas2 = append(tiendas2, cubix[j][i].Lista3)
			tiendas2 = append(tiendas2, cubix[j][i].Lista4)
			tiendas2 = append(tiendas2, cubix[j][i].Lista5)

		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente los archivos"}
	json.NewEncoder(w).Encode(error)
}

func CargarProductos(w http.ResponseWriter, r *http.Request){
	var newDoc Structs.EnlaceInventario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		error := Structs.JsonErrors{Mensaje: "Ha ocurrido un problema! :("}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqBody, &newDoc)
	var finded *Structs.Tiendas
	var position int

	for _, inven := range newDoc.Inventarios{
		sup := Structs.PedidosS{
			Departamento: inven.Departamento,
			Nombre:       inven.Tienda,
			Calificacion: inven.Calificacion,
		}
		position = searchingVectorS(&sup)
		finded = tiendas2[position].Search(&sup)
		if position >= 0 && finded.Nombre != ""{
			if finded.Arbolito == nil {
				finded.Arbolito = Structs.NewArbol()
				for _, product := range inven.Productos{
					product.Tienda = inven.Tienda
					product.Departamento = inven.Departamento
					product.Calificacion = finded.Calificacion
					product.Contacto = finded.Contacto
					finded.Arbolito.Insert(product, &inven.Departamento, &inven.Departamento, &inven.Calificacion)
				}

			}

		}else {

		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente los archivos"}
	json.NewEncoder(w).Encode(error)
}

func CargarPedidos(w http.ResponseWriter, r*http.Request){
	var newDoc Structs.EnlacePedidos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
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
			Value: *aux3,
			Year:  getYear(pedido.Fecha),
			Dia:   getDay(pedido.Fecha),
			Month: getMonth(pedido.Fecha),
			MonthString: getStringMonth(getMonth(pedido.Fecha)),
			Ascii: convertAscii(pedido.Departamento),
			Right: nil,
			Left:  nil,
			Up:    nil,
			Down:  nil,
		}

		pedidios.AddYear(&aux)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	error := Structs.JsonErrors{Mensaje: "Se han cargado correctamente los archivos"}
	json.NewEncoder(w).Encode(error)
	//fmt.Print("asd")

}

func Deletition(w http.ResponseWriter, r *http.Request){
	var newDoc Structs.PedidosE
	reqBody, err := ioutil.ReadAll(r.Body)
	var contain bool
	var position int
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)
	position = searchingVectorE(&newDoc)
	contain = tiendas2[position].Delete(&newDoc)
	if contain && position >= 0{
		error := Structs.JsonErrors{Mensaje: "The store was deleted succesfully"}
		json.NewEncoder(w).Encode(error)
	}else {
		error := Structs.JsonErrors{Mensaje: "We don´t find the store check your values"}
		json.NewEncoder(w).Encode(error)
	}
	var solo = len(tiendas2)
	fmt.Println(solo)

}

func Search(w http.ResponseWriter, r *http.Request){
	var newDoc Structs.PedidosS
	reqBody, err := ioutil.ReadAll(r.Body)
	var finded Structs.Tiendas
	var position int
	if err != nil{
		fmt.Fprintf(w,"Insert correct Values")
	}
	json.Unmarshal(reqBody, &newDoc)

	position = searchingVectorS(&newDoc)
	finded = *tiendas2[position].Search(&newDoc)

	if position >= 0 && finded.Nombre != ""{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(finded)
	}else {
		fmt.Fprintf(w,"We don't found that store please check your values")
	}
	//fmt.Print("asd")
}

func ShowList (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	PosVector, err :=strconv.Atoi(vars["id"])
	var stores []Structs.Tiendas
	if err != nil {
		fmt.Fprintf(w,"Invalid ID")
		return
	}
	var list = tiendas2[PosVector]
	stores = list.Show()
	if len(stores)==0{
		fmt.Fprintf(w,"In this list don't exist a store :(")
	}else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stores)
	}
}

func Graphviz(w http.ResponseWriter, r *http.Request){
	var cadenita strings.Builder
	var nodes []string
	cont := int(0)
	fmt.Fprintf(&cadenita, "digraph G{ \n")
	fmt.Fprintf(&cadenita, "rankdir= \"LR\" \n")
	fmt.Fprintf(&cadenita, "node[fontname=\"Arial\" style=\"filled\" shape=\"record\" color=\"blue\" fillcolor=\"mediumspringgreen\"]; \n")
	for i:= 0; i<= len(tiendas2);{
		if cont < 5{
			if i != len(tiendas2){
				if tiendas2[i].GetSize() == 0{
					fmt.Fprintf(&cadenita, "node%d[label=\"vacio | %d \"]; \n ", i,i)
					fmt.Fprintf(&cadenita, "node%dv[label=\" \", color=\"white\" fillcolor=\"white\"] \n",i)
					fmt.Fprintf(&cadenita, "node%d->node%dv; \n",i,i )
					aux := string("node"+strconv.Itoa(i))
					nodes = append(nodes, aux)
				}else{
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
					aux := string("node"+strconv.Itoa(i))
					nodes = append(nodes, aux)
					tiendas2[i].Graphic(&cadenita)
					fmt.Fprintf(&cadenita, "node%d->node%p; \n",i,&(*tiendas2[i].GetFirst()) )

				}
				cont++
			}
			i++
		}else{
			fmt.Fprintf(&cadenita, "{rank=\"same\" ;"+nodes[0]+" \n")
			for i:= 1; i<len(nodes); i++{
				fmt.Fprintf(&cadenita, ";"+nodes[i]+"\n")
			}
			fmt.Fprintf(&cadenita, " }\n")
			aux2:= string(" ")
			for i, node := range nodes{
				if i == 0{
					aux2 = node
				}else {
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

func SaveStuff(w http.ResponseWriter, r *http.Request){
	var stuffdatos []Structs.Datos
	var stuffdepa []Structs.Departamentos2

	for _, indi := range indice{
		for i:= 0; i < len(departa); i++{
			aux2 := []Structs.Tiendas {}
			aux:= Structs.Departamentos2{
				Nombre:  departa[i],
				Tiendas: aux2,
				Indice:  indi,
			}
			stuffdepa = append(stuffdepa, aux)

		}
	}



	for _, lista := range tiendas2{
		l := lista.GetStores()
		if lista.GetSize() > 0{
			for j, depa := range stuffdepa{
				if depa.Nombre == lista.GetFirst().Departamento && depa.Indice == lista.GetFirst().Indice{
					stuffdepa[j].Tiendas = append(stuffdepa[j].Tiendas, l...)
				}
			}
		}

	}
	for i:= 0; i < len(indice); i++{
		aux3:= Structs.Datos{
			Indice:        "",
			Departamentos: nil,
		}
		for _, depa := range stuffdepa{
			if indice[i] == depa.Indice{
				aux6:= Structs.Departamentos{
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
	f,_ := json.MarshalIndent(sending,""," ")
	_ = ioutil.WriteFile("NewStuff.json",f,0644)
	msg := Structs.JsonErrors{Mensaje: "The json file was created!"}
	json.NewEncoder(w).Encode(msg)


	/*var stuffsvaing []Structs.Datos

	 for i:= 0; i< len(indice); i++{
		 var stuffdepa []Structs.Departamentos
	 	aux := Structs.Datos{
			Indice:        indice[i],
			Departamentos: stuffdepa,
		}
	 	stuffsvaing = append(stuffsvaing, aux)
	 }

	 for i:= 0; i<len(indice); i++{
	 	for j:= 0; j<len(departa) ;j++{
			stores := []Structs.Tiendas{}
			aux2:= Structs.Departamentos{
				Nombre:  departa[j],
				Tiendas: stores,
			}
			stuffsvaing[i].Departamentos = append(stuffsvaing[i].Departamentos, aux2)
		 }

	 }

	for _, listas := range tiendas2{
		t := listas.GetStores()
		if listas.GetSize()> 0 {
			for i, indices := range stuffsvaing{
				if indices.Indice == listas.GetFirst().Indice{
					for j, depars := range indices.Departamentos{
						if depars.Nombre == listas.GetFirst().Departamento{
							stuffsvaing[i].Departamentos[j].Tiendas = append(stuffsvaing[i].Departamentos[j].Tiendas, t...)
							break
						}
					}
				}
			}
		}
	}
	var sending Structs.Enlace
	sending.Datos = stuffsvaing
	f,_ := json.MarshalIndent(sending,""," ")
	_ = ioutil.WriteFile("NewStuff.json",f,0644)
	*/

}

func searchingVectorE(pedido *Structs.PedidosE) int{
	var indicefound, depafound bool
	var first, second, result int

	for i, s := range indice{
		if s[0] == pedido.Nombre[0]{
			first = i
			indicefound = true
		}
	}
	for i, s := range departa{
		if s == pedido.Categoria{
			second = i
			depafound = true
		}
	}

	if !indicefound {
		return -1
	}

	if !depafound{
		return -1
	}

	f := second - 0
	s := f * len(indice) + first
	result = s * 5 + pedido.Calificacion-1
	return result
	/*[i][j][w]
	primero = j-0
	segundo = primero * cantidad filas + i
	tercero = segundo*5 + w*/

}

func searchingVectorS(pedido *Structs.PedidosS) int{
	var indicefound, depafound bool
	var first, second, result int

	for i, s := range indice{
		if s[0] == pedido.Nombre[0]{
			first = i
			indicefound = true
		}
	}
	for i, s := range departa{
		if s == pedido.Departamento{
			second = i
			depafound = true
		}
	}

	if !indicefound {
		return -1
	}

	if !depafound{
		return -1
	}

	f := second - 0
	s := f * len(indice) + first
	result = s * 5 + pedido.Calificacion-1
	return result
	/*[i][j][w]
	primero = j-0
	segundo = primero * cantidad filas + i
	tercero = segundo*5 + w*/

}

func putStore(aux2 Structs.Node, sup []Structs.NodeListas, depa int,) []Structs.NodeListas{

	switch aux2.Tienda.Calificacion {
	case 1:
		sup[depa].Lista1.SortedInsert(&aux2)
		break
	case 2:sup[depa].Lista2.SortedInsert(&aux2)
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
	if(len(tiendas2) != 0){
		for i:= 0; i < len(tiendas2) ; i++{
			aux := tiendas2[i].GetStores()
			for _, tienda := range aux {
				tienda.Key = strconv.Itoa(i)+"$"+tienda.Contacto
				ShopList.Tiendas = append(ShopList.Tiendas, tienda)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		arreglo := ShopList.Tiendas
		json.NewEncoder(w).Encode(arreglo)
	}else {
		err := Structs.JsonErrors{Mensaje: "No hay tiendas cargadas"}
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(err)
	}

}

func getProducts(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	elements := strings.Split(vars["id"],"$")
	vectorPos, _ := strconv.Atoi(elements[0])
	ShopList := tiendas2[vectorPos]
	selectShop := ShopList.GetShop(elements[1])
	products := selectShop.Arbolito.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	arreglo := products
	json.NewEncoder(w).Encode(arreglo)


}

func getArbolito(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	elements := strings.Split(vars["id"],"$")
	vectorPos, _ := strconv.Atoi(elements[0])
	ShopList := tiendas2[vectorPos]
	selectShop := ShopList.GetShop(elements[1])
	selectShop.Arbolito.Generate()
}

func getMatrix(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	elements := strings.Split(vars["id"],"-")
	year, err := strconv.Atoi(elements[0])
	month, err2 := strconv.Atoi(elements[1])
	if err != nil && err2 != nil {
		fmt.Println("error")
	}
	monthUsed := pedidios.SearchYear(year)
	sperceMatrix := monthUsed.Monts.SearchMonth(month)

	sperceMatrix.Matrix.Graphviz()

}

func getPedidos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ :=vars["id"]
	is := strings.Split(i,"$")
	var datos []string

	var finded *Structs.Tiendas
	var position int
	datos = strings.Split(is[0],"-")
	aux,_ := strconv.Atoi(datos[0])
	monthUsed:= pedidios.SearchYear(aux)
	aux,_ = strconv.Atoi(datos[1])
	sperceMatrix := monthUsed.Monts.SearchMonth(aux)
	aux2,_ := strconv.Atoi(is[1])
	aux3 := strings.Replace(is[2],"_"," ",-1)

	pedidos := sperceMatrix.Matrix.ObtenerNodito(aux2,convertAscii(aux3))
	arregloPedidos := pedidos.StackPedidos.ArregloPedidos()
	var pedidosEnviar []Structs.ShowingPedidos
	for _, pedi := range arregloPedidos{
		sup := Structs.PedidosS{
			Departamento: pedi.Departamento,
			Nombre:       pedi.Tienda,
			Calificacion: pedi.Calificacion,
		}
		position = searchingVectorS(&sup)
		finded = tiendas2[position].Search(&sup)
		var sup3 [] string
		sup2 := Structs.ShowingPedidos{
			Fecha:        pedi.Fecha,
			Tiendas:      pedi.Tienda,
			Departamento: pedi.Departamento,
			Producto:     sup3,
		}
		for _, codProd := range pedi.Productos{
			aux4 := finded.Arbolito.SearchPrduc(codProd.Codigo)
			sup2.Producto = append(sup2.Producto, *aux4)
		}
		pedidosEnviar = append(pedidosEnviar, sup2)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedidosEnviar)
}

func saveDot(s string,i int){
	nombre := string("lista"+strconv.Itoa(i)+".pdf")
	f, err := os.Create("lista.dot")
	if err != nil{
		fmt.Println("There was an error!")
	}
	l, err := f.WriteString(s)
	if err != nil{
		fmt.Println("There was an error!")
		f.Close()
		return
	}
	fmt.Println(l,"Created Succesfully")
	p := "dot -Tpdf lista.dot -o " + nombre

	args := strings.Split(p, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("A ocurrido un error", err)

	}
	fmt.Printf("%s\n", b)
}

func convertAscii(s string)int{
	ascii:= int(0)
	runes := []rune(s)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	for i := 0; i < len(result);i++ {
		ascii = ascii + result[i]
	}

	return ascii
}

func getDay(s string)int{
	aux := strings.Split(s,"-")
	i, err := strconv.Atoi(aux[0])
	if err != nil {
		fmt.Println("Na")
	}
	return i
}

func getYear(s string)int{
	aux := strings.Split(s,"-")
	i, err := strconv.Atoi(aux[2])
	if err != nil {
		fmt.Println("Na")
	}
	return i
}

func getMonth(s string)int{
	aux := strings.Split(s,"-")
	i, err := strconv.Atoi(aux[1])
	if err != nil {
		fmt.Println("Na")
	}
	return i
}

func getStringMonth(s int)string{
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
	i, _ :=vars["id"]
	http.ServeFile(w,r,"Imagenes/Arbol"+i+".png")

}

func mostrarImagenMatriz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ :=vars["id"]
	http.ServeFile(w,r,"Matrices/Matriz"+i+".png")

}

func getYears(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	arreglo := pedidios.GetYeats()
	json.NewEncoder(w).Encode(arreglo)

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", example).Methods("GET")
	router.HandleFunc("/cargartienda", cargaArchivos).Methods("POST")
	router.HandleFunc("/Eliminar", Deletition).Methods("DELETE")
	router.HandleFunc("/TiendaEspecifica", Search).Methods("POST")
	router.HandleFunc("/id/{id}", ShowList).Methods("GET")
	router.HandleFunc("/getArreglo", Graphviz).Methods("GET")
	router.HandleFunc("/guardar", SaveStuff).Methods("GET")
	router.HandleFunc("/obtenerTiendas", getShops).Methods("GET")
	router.HandleFunc("/obtenerTiendas/{id}", getProducts).Methods("GET")
	router.HandleFunc("/obtenerArbolito/{id}", getArbolito).Methods("GET")
	router.HandleFunc("/obtenerMatriz/{id}", getMatrix).Methods("GET")
	router.HandleFunc("/obtenerYears", getYears).Methods("GET")
	router.HandleFunc("/obtenerPedidos/{id}", getPedidos).Methods("GET")
	router.HandleFunc("/cargarproductos", CargarProductos).Methods("POST")
	router.HandleFunc("/cargarpedidos", CargarPedidos).Methods("POST")
	router.HandleFunc("/arbolito/{id}", mostrarImagenArbolito)
	router.HandleFunc("/matriz/{id}", mostrarImagenMatriz)


	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type","Authorization"})
	methods := handlers.AllowedMethods([]string{"GET","PUT","POST","DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
}
