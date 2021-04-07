import React, { useState } from 'react'
import { Button, Form, Input, Radio, Segment, Grid } from 'semantic-ui-react'
import { useHistory } from 'react-router-dom'


const axios = require('axios')
function CrearUsu() {
    const histori = useHistory();
    const [name, setname] = useState("")
    const [dpi, setdpi] = useState("")
    const [pass, setpass] = useState("")
    const [email, setemail] = useState("")
    const [cuenta, setcuenta] = useState("")

    const handleChange = (event, {value}) => setcuenta({ value });

    const goInicioSesion = () =>{
        histori.push(`/inicio`)
    }
    const send = () =>{
        var usu = [
            dpi,
            name,
            email,
            pass,
            cuenta.value
        ]
        axios.post('http://localhost:3000/CreateUsu',
          usu,
          { headers: { 'content-type': 'application/json' } }
        ).then(data => {
          alert('El Usuario se creo exitosamente')
          console.log(data)
          histori.push(`/inicio`)
        }).catch(e => {
          alert('Porfavor llene todoas las casillas')
          console.log(e)
        })

    }

    return (
        <>
            <Grid centered>
                <Grid.Column style={{ maxWidth:700, marginTop:20 }}>
            <Form>
                <Form.Group widths='equal'>
                    <Form.Input
                        onChange={e => setname(e.target.value)}
                        name = "name"
                        label='Nombre'
                        placeholder='Nomnbre'
                    />
                    <Form.Input
                        onChange={e => setdpi(e.target.value)} 
                        name ="DPI"
                        label='DPI'
                        placeholder='DPI'
                    />
                    <Form.Input 
                    onChange={e => setpass(e.target.value)} 
                    fluid label='Password' 
                    placeholder='Password' 
                    type="password" />
                </Form.Group>
                <Form.Group inline>
                    <label>Type</label>
                    <Form.Field
                        control={Radio}
                        label='Admin'
                        value='Admin'
                        checked={cuenta.value === 'Admin'}
                        onChange={handleChange}
                    />
                    <Form.Field
                        control={Radio}
                        label='Usuario'
                        value='Usuario'
                        checked={cuenta.value === 'Usuario'}
                        onChange={handleChange}
                    />
                </Form.Group>
                <Form.Input
                    onChange={e => setemail(e.target.value)}
                    name = "email"
                    type="email"
                    label='Email'
                    placeholder='example@something.com'
                />
                <Form.Group widths="equal">
                <Form.Button onClick={send}>Submit</Form.Button>
                <Form.Button onClick={goInicioSesion}> Back Iniciar Sesión </Form.Button>
                </Form.Group>
            </Form>
            </Grid.Column>
            </Grid>
        </>
    )
}

export default CrearUsu
