import React from 'react'
import { Table, Button } from 'semantic-ui-react'
import Fila from './Fila'



const axios = require('axios')
function Tablita(props) {
    
    const mandarPedidos = () => {
            
        
        axios.post('http://localhost:3000/carrito',
          props.data,
          { headers: { 'content-type': 'application/json' } }
        ).then(data => {
          alert('file uploaded')
          console.log(props.data)
          window.localStorage.clear();
        }).catch(e => {
          alert('error')
          console.log(e)
        })
      
      }

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
                    <Button positive floated='right' onClick={mandarPedidos}>Confirmar Compra</Button>
                    </Table.HeaderCell>
                </Table.Row>
            </Table.Footer>
        </Table>
    )
}

export default Tablita
