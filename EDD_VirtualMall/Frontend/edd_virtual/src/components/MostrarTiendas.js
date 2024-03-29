import React, { useEffect, useState } from 'react'
import Mosaico from './Mosaico'


const axios = require('axios')
function MostrarTiendas() {
    const [stores, setstores] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (stores.length === 0) {
                const data = await axios.get(`http://localhost:3000/obtenerTiendas`);
                if (data.status !== 204) {
                    console.log(data.data)
                    setstores(data.data)
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
            <div className="ImportList">
                <br></br>
                <Mosaico stores={stores} />
            </div>
        )
    }
}

export default MostrarTiendas
