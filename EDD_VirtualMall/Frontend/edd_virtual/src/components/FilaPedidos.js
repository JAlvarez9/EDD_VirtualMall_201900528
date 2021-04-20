import React from 'react'
import { Header, Table, Modal, Button } from 'semantic-ui-react'


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
            <Table.Cell>
                <Modal
                    trigger={<Button>Caminito</Button>}
                    header='CAMINO CORTO '
                    content={props.camino}
                    actions={['Snooze', { key: 'done', content: 'Done', positive: true }]}
                />
            </Table.Cell>
        </Table.Row>
    )
}

export default FilaPedidos
