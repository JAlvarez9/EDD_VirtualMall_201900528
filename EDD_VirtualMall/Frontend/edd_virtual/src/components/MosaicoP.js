import React from 'react'

import Carta from './CartaProducto'

function MosaicoP(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.products.map((c, index) => (
                    <Carta name={c.Nombre}
                        descripcion={c.Descripcion}
                        image={c.Imagen}
                        mount = {c.Cantidad}
                        price = {c.Precio}
                        id={c.Codigo}
                        key={c.Codigo}
                        tienda={c.Tienda}
                        departa={c.Departamento}
                        califi={c.Calificacion}
                        conta={c.Contacto}
                    />
                ))}
            </div>
        </div>
    )
}

export default MosaicoP
