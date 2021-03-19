import React, { useEffect, useState } from 'react'
import MosaicoP from './MosaicoP'


const axios = require ('axios')
function MostrarProductos(props) {
    let id = props.match.params.id
    const [products, setproducts] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (products.length === 0) {
                const data = await axios.get(`http://localhost:3000/obtenerTiendas/${id}`);
                if (data.status !== 204) {
                    console.log(data.data)
                    setproducts(data.data)
                    setloading(true)
                }

            }
        }
        obtener()
    });
    if (loading === false) {
        return (
            <div className="ui segment carga">
                <div className="ui active dimmer">
                    <div className="ui text loader">Loading</div>
                </div>
                <p />
            </div>
        )
    } else {
        return (
            <div className="Productos">
                <br></br>
                <MosaicoP products={products} />
            </div>
        )
    }

}

export default MostrarProductos
