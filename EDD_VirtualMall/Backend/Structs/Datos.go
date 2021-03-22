package Structs

type (
	CodProducto struct {
		Codigo int `json:Codigo`
	}

	Pedidos struct {
		Fecha        string        `json:Fecha`
		Tienda       string        `json:Tienda`
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
		Departamento string `json:"Departamento,omitempty"`
		Tienda       string `json:"Tienda,omitempty"`
		Calificacion int    `json:Calificacion,omitempty`
		Contacto 	 string    `json:Contacto,omitempty`
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
)
