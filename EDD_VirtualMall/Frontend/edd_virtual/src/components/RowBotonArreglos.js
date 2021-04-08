import React, { useState, useEffect } from 'react'
import { Grid } from 'semantic-ui-react'
import Boton from './ButtonArreglo'

const axios = require('axios')
function RowBotonArreglos() {

    const [botones, setbotones] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (botones.length === 0) {
                const data = await axios.get(`http://localhost:3000/botoncitos`);
                
                if (data.status !== 204) {
                    setbotones(data.data)
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
            <>
                {botones.map((dato, index) => (

                    <Grid.Row>
                        <Boton
                            name={dato}
                        />
                    </Grid.Row>
                ))}
            </>
        )
    }
}

export default RowBotonArreglos
