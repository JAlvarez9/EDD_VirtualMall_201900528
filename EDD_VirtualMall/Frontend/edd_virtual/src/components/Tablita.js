import React from 'react'
import { Table, Button } from 'semantic-ui-react'
import Fila from './Fila'


function Tablita(props) {

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
                        datas={dato}

                    />
                ))}
            </Table.Body>
            <Table.Footer fullWidth>
                <Table.Row>
                    <Table.HeaderCell colSpan='4'>
                    <Button positive floated='right'>Confirmar Compra</Button>
                    </Table.HeaderCell>
                </Table.Row>
            </Table.Footer>
        </Table>
    )
}

export default Tablita
