import React, { useEffect, useState } from 'react'
import { Segment } from 'semantic-ui-react'
import { useHistory } from 'react-router-dom'
import Tabla from './Tablita'




function Carrito() {
    const histori = useHistory();
    const verifyRoute = () => {
        if (!window.sessionStorage.getItem("cuenta")) {
            return histori.push('/inicio');
        }
    }; 
    verifyRoute()
    const encabezados = ["Usuario", "Producto", "Precio Unitario", "Cantidad", "Total"]
    const [listado, setlistado] = useState([])
    const [cliente] = useState(window.sessionStorage.getItem("dpi"))
    useEffect(() => {
        let data = localStorage.getItem('productos')
        if (data != null) {
            setlistado(JSON.parse(data))
        }
    }, [])

    return (
        <Segment inverted color='grey'>
            <h2></h2>
            <Tabla data={listado}
                enca={encabezados}
            />
        </Segment>
    )
}

export default Carrito

