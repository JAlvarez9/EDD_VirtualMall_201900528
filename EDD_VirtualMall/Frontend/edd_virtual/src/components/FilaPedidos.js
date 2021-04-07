import React from 'react'
import { Header, Table } from 'semantic-ui-react'


function FilaPedidos(props) {
    return (
        <Table.Row>
            <Table.Cell>
                <Header>
                    {props.cliente}
                </Header>
            </Table.Cell>
            <Table.Cell>
                <Header>
                    {props.fecha}
                </Header>
            </Table.Cell>
            <Table.Cell singleLine>{props.tienda}</Table.Cell>
            <Table.Cell textAlign='right'>
                {props.depa}
            </Table.Cell >
            <Table.Cell>
                {props.producs.map((dato, index) => (
                    <p>{dato}</p>
                    
                ))}
            </Table.Cell>
        </Table.Row>
    )
}

export default FilaPedidos
