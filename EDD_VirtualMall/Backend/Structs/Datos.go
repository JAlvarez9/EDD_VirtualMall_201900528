package Structs

type (
	CodProducto struct {
		Codigo int `json:Codigo`
	}

	Pedidos struct {
		Fecha        string        `json:Fecha`
		Cliente      int           `json:Cliente`
		Tienda       string         `json:Tienda`
		Departamento string        `json:Departamento`
		Calificacion int           `json:Calificacion`
		Productos    []CodProducto `json:Productos`
	}

	EnlacePedidos struct {
		Pedidos []Pedidos `json:Pedidos`
	}

	Productos struct {
		Nombre       string `json:Nombre`
		Codigo       int    `json:Codigo`
		Descripcion  string `json:Descripcion`
		Precio       float64`json:Precio`
		Cantidad     int    `json:Cantidad`
		Imagen       string `json:Imagen`
		Almacenamiento string `json:Almacenamiento`
		Departamento string `json:Departamento,omitempty`
		Tienda       string `json:Tienda,omitempty`
		Calificacion int    `json:Calificacion,omitempty`
		Contacto 	 string    `json:Contacto,omitempty`
		Tablita 	HashTable  `json:Tablita,omitempty`
	}
	Inventario struct {
		Tienda       string      `json:Tienda`
		Departamento string      `json:Departamento`
		Calificacion int         `json:Calificacion`
		Productos    []Productos `json:Productos`
	}

	EnlaceInventario struct {
		Inventarios []Inventario `json:Inventarios`
	}
	Tiendas struct {
		Nombre       string   `json:Nombre`
		Descripcion  string   `json:Descripcion`
		Contacto     string   `json:Contacto`
		Calificacion int      `json:Calificacion`
		Logo         string   `json:Logo`
		Arbolito     *TreeAVL `json:"Arbol,omitempty"`
		Key          string   `json:"Key,omitempty"`
		Tablita 	HashTable  `json:Tablita,omitempty`

	}

	JsonErrors struct {
		Mensaje string `json:Mensajer`
	}

	Departamentos struct {
		Nombre  string    `json:Nombre`
		Tiendas []Tiendas `json:Tiendas`
	}

	Departamentos2 struct {
		Nombre  string    `json:Nombre`
		Tiendas []Tiendas `json:Tiendas`
		Indice  string
	}

	Datos struct {
		Indice        string          `json:Indice`
		Departamentos []Departamentos `json:Departamentos`
	}

	Enlace struct {
		Datos []Datos `json:Datos`
	}

	Shops struct {
		Tiendas []Tiendas
	}

	Anios struct {
		Year string
		Months []string
	}

	Producto2 struct {
		Nombre string
		Codigo int
	}

	ShowingPedidos struct {
		Fecha string
		Tiendas string
		Departamento string
		Cliente int
		Producto [] string
		CaminmoCorto string
		CodProductos []CodProducto
		Calificacion int
	}


	ArregloPedidos struct {
		Pedidos []ShowingPedidos
	}

	Carrito struct {
		Nombre string
		Descripcion string
		Image string
		Mount int
		Price float64
		Id int
		Tienda string
		Departamento string
		Calificacion int
		Cantidad int
		Fecha string
		Cliente int
	}

	ValidarPedidos struct {
		Tienda string
		Departamento string
		Calificacion int
		Cliente int
		Producto   CodProducto
		Productos   []CodProducto
	}

	Usuarios struct {
		DPI int 	`json:DPI`
		Nombre string  `json:Nombre`
		Correo string	`json:Correo`
		Password string  `json:Password`
		Cuenta string	 `json:Cuenta`
	}

	EnlaceUsuarios struct {
		Usuarios []Usuarios `json:Usuarios`
	}

	InicioSesion struct {
		DPI string
		Password string
	}

	UsuariosEncrit struct {
		DPI string
		DPIN int
		Nombre string
		Correo string
		Password string
		Cuenta string
	}

	Enlaces struct {
		Nombre string `json:Nombre`
		Distancia float64 `json:Distancia`
	}

	Vertices struct {
		Nombre string `json:Nombre`
		Enlaces []Enlaces `json:Enlaces`
	}

	EnlaceGrafos struct {
		Nodos []Vertices `json:Nodos`
		PosicionInicialRobot string `json:PosicionInicialRobot`
		Entrega string `json:Entrega`
	}

	Grafo struct {
		Inicio string
		Final string
		Nodos *Stack3
	}

	CaminosCortos struct {
		Inicia string
		Termina string
		Recorrido []string
		Peso int
	}

	CaminosProductos struct {
		Almacenamiento string
		Productos []Productos
	}

	TransaccionPedidos struct {
		Id string
		Sha string
		Sha2 string
		Dpi int
		Fecha string
		Monto float64
	}

	TransaccionTiendas struct {
		Id string
		Sha string
		Sha2 string
		Accion string
		Nombre string
		Departamento string
		Descripcion string
		Contacto string
		Calificacion int
		Logo string
	}

	TransaccionProductos struct {
		Id string
		Sha string
		Sha2 string
		Accion string
		Tienda string
		Departamento string
		Calificacion int
		Nombre string
		Codigo int
		Descripcion string
		Precio float64
		Cantidad int
		Imagen string
		Almacenamiento string
	}

	TransaccionUsuarios struct {
		Id string
		Sha string
		Sha2 string
		Accion string
		DPI int
		Nombre string
		Correo string
		Pass string
		Cuenta string
	}

	Comentarios struct {
		Comen string
	}

	Bloqusitos struct {
		Indice int
		Fecha string
		Data string
		Nonce int
		PreviousHash string
		Hash string
	}

	ArbolitioMerckle struct {
		Nombre string
		Grafiquita string
	}

	Bloquito struct {
		Indice string
		Fecha string
		Data ArreglosBloc
		Nonce string
		PreviousHash string
		Hash string
	}

	ArreglosBloc struct {
		Tienditas []TransaccionTiendas
		Productitos []TransaccionProductos
		Usuaritos []TransaccionUsuarios
		Pediditos []TransaccionPedidos
	}


)
