import React from 'react'
import { Segment, Icon, Button, Grid } from "semantic-ui-react"

import { useHistory } from 'react-router-dom'



function Headerr() {
    const histori = useHistory();
    const mandar = () => {
        histori.push(`/carrito`)
    }
    const cerrar = () => {
        window.sessionStorage.setItem("user", "");
            window.sessionStorage.setItem("cuenta", "");
        histori.push(`/inicio`)
    }
    const name = window.sessionStorage.getItem("user")
    if (name === "" ) {
        return (
            <>
                <Segment basic inverted textAlign="center">
                    <h1>Virtual Mall</h1>
                    <h3>Jose Fernando Alvarez Morales</h3>
                </Segment>

            </>

        )
    } else {
        return (
            <>
                <Segment basic inverted textAlign="center">
                    <h1>Virtual Mall</h1>
                    <h3>Jose Fernando Alvarez Morales</h3>
                </Segment>
                <Grid columns={2} stackable textAlign='center'>
                    <Grid.Row >
                        <Grid.Column textAlign="left">
                            <Button animated='vertical' onClick={mandar}>
                                <Button.Content hidden>Shop</Button.Content>
                                <Button.Content visible>
                                    <Icon name='shop' />
                                </Button.Content>
                            </Button>
                        </Grid.Column>
                        <Grid.Column textAlign="right">
                            <Button negative onClick={cerrar}> Cerrar Sesión </Button>
                        </Grid.Column>
                    </Grid.Row>

                </Grid>

            </>

        )
    }

}

export default Headerr
