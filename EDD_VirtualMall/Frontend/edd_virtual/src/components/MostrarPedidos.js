import React, { useState } from 'react'
import { Modal, Icon, Image, Grid, Button } from 'semantic-ui-react'
import MenuLateral from './MenuCargarPedidos'
import Tablita from './TablaPedidos'

const axios = require('axios')
function MostrarPedidos() {
    //var obtener = ""
    const [open, setOpen] = useState(false)
    const [obtener, setobtener] = useState("")
    const [arPedidos, setarPedidos] = useState([])
    const encabezados = ["Fecha", "Tienda", "Departamento", "Productos"]
    const obtenerMatriz = () => {

        setobtener((document.getElementById("years")).value)
        console.log(obtener)
        axios.get(`http://localhost:3000/obtenerMatriz/${obtener}`);

    }

    const pedidos = () => {
        const name = (document.getElementById("depa")).value
        const dia = (document.getElementById("dia")).value

        let aux2 = obtener + "$" + dia + "$" + name
        const id = aux2

        const asdf = async function () {
            const data = await axios.get(`http://localhost:3000/obtenerPedidos/${id}`);
            setarPedidos(data.data)
        }
        asdf()


    }
    return (

        <Grid celled>
            <Grid.Row>
                <Grid.Column width={3}>

                    <MenuLateral />

                </Grid.Column>
                <Grid.Column width={8}>

                    <div class="ui input">
                        <input id="years" type="text" placeholder="YYYY-MM" />
                    </div>
                    <Button secondary floated="left" onClick={obtenerMatriz}> Generar Matriz </Button>
                </Grid.Column>
            </Grid.Row>

            <Grid.Row>
                <Grid.Column width={3}>
                    <Button.Group vertical>

                        <button  download="matriz.png" href={`http://localhost:3000/matriz/${obtener}`} class="ui secondary button">
                           Descargar Matriz
                        </button>
                        <Button secondary>Guardar Estructura Years</Button>
                        <Button secondary>Guardar Estructura de Meses</Button>
                        <Button secondary>Guardar Pedidos por Dia</Button>
                    </Button.Group>
                </Grid.Column>
                <Grid.Column width={10}>
                    <Image src={`http://localhost:3000/matriz/${obtener}`}></Image>
                </Grid.Column>
                <Grid.Column width={3}>
                    <div class="ui input">
                        <input id="depa" type="text" placeholder="Departamento" />
                    </div>
                    <div class="ui input">
                        <input id="dia" type="text" placeholder="Dia" />
                    </div>
                    <Modal
                        onClose={() => setOpen(false)}
                        onOpen={() => setOpen(true)}
                        open={open}
                        trigger={<Button onClick={pedidos}>Observar Pedidos</Button>}
                    >
                        <Modal.Header>Pedidos</Modal.Header>
                        <Modal.Content >
                            <Tablita
                                enca={encabezados}
                                data={arPedidos}
                            />

                        </Modal.Content>
                        <Modal.Actions>
                            <Button negative onClick={() => setOpen(false)}>
                                <Icon name="trash" /> Cerrar
                            </Button>

                        </Modal.Actions>
                    </Modal>

                </Grid.Column>
            </Grid.Row>
        </Grid>


    )

}

export default MostrarPedidos
