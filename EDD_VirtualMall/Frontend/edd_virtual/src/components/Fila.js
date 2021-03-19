import React, { useState } from 'react'
import { Header, Table } from 'semantic-ui-react'


function Fila(props) {
    const [totala] = useState(props.datas[4] * props.datas[9])
    
    
    return (
        <Table.Row>
            <Table.Cell>
                <Header as='h2' textAlign='center'>
                    {props.datas[0]}
                </Header>
            </Table.Cell>
            <Table.Cell singleLine>${props.datas[4]}</Table.Cell>
            <Table.Cell textAlign='right'>
                {props.datas[9]}
            </Table.Cell >
            <Table.Cell>
                ${totala}
            </Table.Cell>
        </Table.Row>
    )
}

export default Fila
