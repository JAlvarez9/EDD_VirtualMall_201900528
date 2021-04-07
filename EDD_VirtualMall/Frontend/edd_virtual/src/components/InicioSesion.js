import React from 'react'
import { Button, Divider, Grid, Header, Icon, Search, Segment, Input } from 'semantic-ui-react'
import { useHistory } from 'react-router-dom'


const axios = require('axios')
const cargarUsuarios = (event) => {
    const json = event.target.files[0];
    console.log(json)
    axios.post('http://localhost:3000/cargarusuarios',
        json,
        { headers: { 'content-type': 'application/json' } }
    ).then(data => {
        alert('file uploaded')
        console.log(data)
    }).catch(e => {
        console.log('error')
        console.log(e)
    })


}



function InicioSesion() {
    const histori = useHistory();
    const setUsu = () => {

        var json = [
            (document.getElementById("use")).value,
            (document.getElementById("pas")).value
        ]
        axios.post('http://localhost:3000/obtenerUsu',
            json,
            { headers: { 'content-type': 'application/json' } }
        ).then(data => {
            alert('Se Inicio sesión')
            window.sessionStorage.setItem("user", data.data.Nombre);
            window.sessionStorage.setItem("cuenta", data.data.Cuenta);
            window.sessionStorage.setItem("dpi", data.data.DPI);
            histori.push('/principal')
        }).catch(e => {
            alert('Verifique los datos')
            console.log(e)
        })
    }

    const goCreate = () => {
        histori.push(`/form`)
    }

    return (
        <>

            <Segment placeholder>
                <Grid columns={2} stackable textAlign='center'>
                    <Divider vertical>Or</Divider>

                    <Grid.Row verticalAlign='middle'>
                        <Grid.Column>
                            <Header icon>
                                <Icon name='user' />
            Start Sesion
          </Header>
                            <div textAlign="center">
                                <div class="ui input" style={{marginRight:10}}>
                                    <input id="use" type="text" placeholder="User" />
                                </div>
                                <div class="ui input">
                                    <input id="pas" type="password" placeholder="Password" />
                                </div>
                            </div>
                            <div style={{marginTop:10}}>
                                <Button  primary onClick={setUsu}> Log In </Button>
                            </div>
                        </Grid.Column>

                        <Grid.Column>
                            <Header icon>
                                <Icon name='users' />
            Add New
          </Header>
                            <Button primary onClick={goCreate}>Create Single User</Button>
                            <Button style={{marginTop:10}} as="label" htmlFor="file" type="button">
                                Massive Charge
        </Button>
                            <input type="file" id="file" hidden onChange={cargarUsuarios} />
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
            </Segment>
        </>
    )
}

export default InicioSesion
