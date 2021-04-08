import React,{useState, useEffect} from 'react'
import {Grid} from 'semantic-ui-react'
import Boton from './ButtonArbolUsu'

const axios = require('axios')
function RowArbolesUsu() {

    const [arbolitos, setarbolitos] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (arbolitos.length === 0) {
                const data = await axios.get(`http://localhost:3000/btrees`);
                if (data.status !== 204) {
                    setarbolitos(data.data)
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
                {arbolitos.map((dato, index) => (

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

export default RowArbolesUsu
