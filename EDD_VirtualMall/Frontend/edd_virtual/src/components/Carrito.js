import React, { useEffect, useState } from 'react'
import { Segment } from 'semantic-ui-react'
import Tabla from './Tablita'




function Carrito() {
    const encabezados = ["Producto", "Precio Unitario", "Cantidad", "Total"]
    const [listado, setlistado] = useState([])
    useEffect(() => {
        let data = localStorage.getItem('productos')
        if (data != null) {
            setlistado(JSON.parse(data))
        }
    }, [])

    return (
        <Segment inverted color='grey'>
            <Tabla data = {listado}
                enca = {encabezados}
            />
        </Segment>
    )
}

export default Carrito

