import React from 'react'
import { Grid, Modal, Image, Button } from 'semantic-ui-react'
import Fila from './RowBotonArreglos'
import FilaA from './RowArbolesUsu'
import Ejem from './RowBotonesMer'

function Reportes() {

    return (
        <Grid columns={4} divided>
            <Grid.Column textAlign="center">
                <Grid.Row><h1>Arreglo Linealizado</h1></Grid.Row>
                <Fila />
            </Grid.Column>
            <Grid.Column textAlign="center">
                <Grid.Row><h1>Arboles de Usuarios</h1></Grid.Row>
                <FilaA />
            </Grid.Column>
            <Grid.Column textAlign="center">
                <Grid.Row><h1>Grafo De Almacenamiento</h1></Grid.Row>
                <Modal
                    trigger={<Button>Show Graph</Button>}
                    header='Graph!'
                    content={<Image src="http://localhost:3000/grafita"></Image>}
                    actions={[ { key: 'done', content: 'Done', positive: true }]}
                />
            </Grid.Column>
            <Grid.Column textAlign="center">
                <Grid.Row><h1>Arboles de Merckle</h1></Grid.Row>
                <Ejem></Ejem>
            </Grid.Column>

        </Grid>
    )
}

export default Reportes
