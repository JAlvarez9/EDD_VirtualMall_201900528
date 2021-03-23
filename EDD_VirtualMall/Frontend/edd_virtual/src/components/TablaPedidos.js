import React from 'react'
import Fila from './FilaPedidos'
import { Table } from 'semantic-ui-react'

function TablaPedidos(props) {
    return (
        <Table celled padded>
            <Table.Header>
                <Table.Row>
                    {props.enca.map((encabezado, index) => (
                        <Table.HeaderCell textAlign="center">{encabezado}</Table.HeaderCell>

                    ))}
                </Table.Row>
            </Table.Header>
            <Table.Body>
                {props.data.map((dato, index) => (

                    <Fila
                        index={index}
                        fecha = {dato.Fecha}
                        tienda = {dato.Tiendas}
                        depa = {dato.Departamento}
                        producs = {dato.Producto}


                    />
                ))}
            </Table.Body>
            
        </Table>
    )
}

export default TablaPedidos
