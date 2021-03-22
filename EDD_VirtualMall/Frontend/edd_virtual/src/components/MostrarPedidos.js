import React, { useState, useEffect } from 'react'
import {  Image, Grid, Input, Button } from 'semantic-ui-react'
import MenuLateral from './MenuCargarPedidos'

const axios = require('axios')
function obtenerMatriz(fecha) {
    axios.get(`http://localhost:3000/obtenerMatriz/${fecha}`);

}
function MostrarPedidos() {
    const [years, setyears] = useState([])
    const [loading, setloading] = useState(false)
    const [obtener, setobtener] = useState("")

    console.log(obtener)


    useEffect(() => {
        async function obtener() {
            if (years.length === 0) {
                const data = await axios.get(`http://localhost:3000/obtenerYears`);
                if (data.status !== 204) {
                    console.log(data.data)
                    setyears(data.data)
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

            <Grid celled>
                <Grid.Row>
                    <Grid.Column width={3}>

                        <MenuLateral
                            years={years}
                        />

                    </Grid.Column>
                    <Grid.Column width={8}>
                        <Input placeholder='YYYY-MM' onChange={e => setobtener(e.target.value)} />
                        <Button secondary floated="left" onClick={obtenerMatriz(obtener)}> Generar Matriz </Button>
                    </Grid.Column>
                </Grid.Row>

                <Grid.Row>
                    <Grid.Column width={3}>
                        <Button.Group vertical>
                            <Button secondary>Generar Matriz</Button>
                            <Button secondary>Guardar Estructura Years</Button>
                            <Button secondary>Guardar Estructura de Meses</Button>
                            <Button secondary>Guardar Pedidos por Dia</Button>
                        </Button.Group>
                    </Grid.Column>
                    <Grid.Column width={10}>
                        <Image src={`http://localhost:3000/matriz/${obtener}`}></Image>
                    </Grid.Column>
                    <Grid.Column width={3}>
                        <Input placeholder="Departamento"></Input>
                        <Input placeholder="Dia"></Input>
                        <Button secondary> Observar Pedidos </Button>
                    </Grid.Column>
                </Grid.Row>
            </Grid>


        )
    }
}

export default MostrarPedidos
