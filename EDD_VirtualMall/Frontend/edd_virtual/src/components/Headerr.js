import React from 'react'
import { Segment, Icon, Button } from "semantic-ui-react"

import { useHistory } from 'react-router-dom'



function Headerr() {
    const histori = useHistory();
    const mandar = () =>{
        histori.push(`/carrito`)
    }
    return (
        <>
            <Segment basic inverted textAlign="center">
                <h1>Virtual Mall</h1>
                <h3>Jose Fernando Alvarez Morales</h3>
            </Segment>
            <div textAlign="right">
                <Button animated='vertical' onClick={ mandar }>
                    <Button.Content hidden>Shop</Button.Content>
                    <Button.Content visible>
                        <Icon name='shop' />
                    </Button.Content>
                </Button>
            </div>

        </>

    )
}

export default Headerr
