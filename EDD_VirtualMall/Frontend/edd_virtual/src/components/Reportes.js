import React from 'react'
import { Grid} from 'semantic-ui-react'
import Fila from './RowBotonArreglos'
import FilaA from './RowArbolesUsu'

function Reportes() {
   
    return (
        <Grid columns={2} divided>
            <Grid.Column textAlign="center">
                <Grid.Row><h1>Arreglo Linealizado</h1></Grid.Row>
                <Fila/>
            </Grid.Column>
            <Grid.Column textAlign="center">
                <Grid.Row><h1>Arboles de Usuarios</h1></Grid.Row>
                <FilaA/>
            </Grid.Column>

        </Grid>
    )
}

export default Reportes
