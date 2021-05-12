import React from 'react'
import { Button, Divider, Grid, Header, Icon, Search, Segment, Input } from 'semantic-ui-react'
import { useHistory } from 'react-router-dom'

const axios = require('axios')


function InicioSesion() {
    const histori = useHistory();
    const setUsu = () => {
        if ((document.getElementById("use")).value === "admin" && (document.getElementById("pas")).value) {
            alert('Se Inicio sesión')
            window.sessionStorage.setItem("user", "admin");
            window.sessionStorage.setItem("cuenta", "Admin");
            window.sessionStorage.setItem("dpi", "123456789");
            histori.push('/principal')
        } else {
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
                                <div class="ui input" style={{ marginRight: 10 }}>
                                    <input id="use" type="text" placeholder="User" />
                                </div>
                                <div class="ui input">
                                    <input id="pas" type="password" placeholder="Password" />
                                </div>
                            </div>
                            <div style={{ marginTop: 10 }}>
                                <Button primary onClick={setUsu}> Log In </Button>
                            </div>
                        </Grid.Column>

                        <Grid.Column>
                            <Header icon>
                                <Icon name='users' />
            Add New
          </Header>
                            <Button primary onClick={goCreate}>Create Single User</Button>
                            
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
            </Segment>
        </>
    )
}

export default InicioSesion
