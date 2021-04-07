import React, { useState } from 'react'
import { Modal, Icon, Image, Grid, Button } from 'semantic-ui-react'
import MenuLateral from './MenuCargarPedidos'
import Tablita from './TablaPedidos'
import ReactDOM from "react-dom"


const axios = require('axios')
function MostrarPedidos() {
    const [open, setOpen] = useState(false)
    const [obtener, setobtener] = useState(String)
    const [arPedidos, setarPedidos] = useState([])
    const encabezados = ["Cliente","Fecha", "Tienda", "Departamento", "Productos"]


    const obtenerMatriz = () => {
        const a = (document.getElementById("years")).value
        setobtener(a)
        const matriz = async function (a) {
            const data = await axios.get(`http://localhost:3000/obtenerMatriz/${a}`);
            var imagen = `http://localhost:3000/matriz/${a}`
            console.log(imagen)

            const Example = ({ imagen }) => <img src={imagen} alt="" style={{ maxWidth: "100%" }} />
            ReactDOM.render(<Example imagen={imagen} />, document.getElementById('matriz'))

        }
        matriz(a)

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

                        <Button secondary>Descargar Matriz</Button>
                        <Button secondary>Guardar Estructura Years</Button>
                        <Button secondary>Guardar Estructura de Meses</Button>
                        <Button secondary>Guardar Pedidos por Dia</Button>
                    </Button.Group>
                </Grid.Column>
                <Grid.Column width={10}>
                    <div id="matriz">
                    </div>
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
