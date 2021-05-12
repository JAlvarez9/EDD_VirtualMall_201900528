import React, { useState, useEffect } from 'react'
import { Grid } from 'semantic-ui-react'
import Boton from './ButtonMerckle'

const axios = require('axios')
function RowBotonesMerckle() {

    const [botones, setbotones] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (botones.length === 0) {
                const data = await axios.get(`http://localhost:3000/botoncitosmerckle`);
                
                if (data.status !== 204) {
                    setbotones(data.data)
                    console.log(data.data)
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
                            name={dato.Nombre}
                            graph={dato.Grafiquita}

                        />
                    </Grid.Row>
                ))}
            </>
        )
    }
}

export default RowBotonesMerckle
