import React, { useState, useEffect } from 'react'
import '../css/Carta.css'
import { useHistory } from 'react-router-dom'
import { Button, Icon, Modal, Form, Image, Comment } from 'semantic-ui-react'
import Comentario from './ComentarioTienda'

const axios = require('axios')

function Carta(props) {
    const [comentario, setcomentario] = useState([])
    const [open, setOpen] = React.useState(false)
    const [come, setcome] = useState("")
    const [dpi] = useState(window.sessionStorage.getItem("dpi"))
    const histori = useHistory();
    const mandar = () => {
        console.logo(props.id)
        histori.push(`/mostrarproductos/${props.id}`)
    }
    const mandararbolito = () => {
        histori.push(`/mostrararbol/${props.id}`)
    }
    const comentacion = () => {
        var comenta = [
            come,
            dpi,
        ]
        console.log(comenta)
        axios.post(`http://localhost:3000/agregarComentTienda/${props.id}`,
            comenta,
            { headers: { 'content-type': 'application/json' } }
        ).then(data => {

        }).catch(e => {
            alert('Error :( ')
            console.log(e)
        })
    }
    useEffect(() => {
        async function obtener() {
            if (comentario.length === 0) {
                const data = await axios.get(`http://localhost:3000//obtenerArregloTiendas/${props.id}`);
                if (data.status !== 204) {
                    if(data.data === null){

                    }else{
                        setcomentario(data.data)
                    }
                    
                }

            }
        }
        obtener()
    });
    return (
        <div className="column carta">
            <div className="ui card">

                <div className="image">
                    <img src={props.logo} />
                </div>

                <div className="content">

                    <div className="header">{props.nombre}</div>

                    <div className="meta">
                        <p>Calificacion: {props.calificacion}</p>
                    </div>
                    <div className="meta">
                        <p>Departamento: {props.depa}</p>
                    </div>

                    <div className="description">{props.descripcion}</div>
                    <div>
                        <Button floated='right' onClick={mandararbolito}><Icon name='tree' /></Button>
                        <Modal
                            onClose={() => setOpen(false)}
                            onOpen={() => setOpen(true)}
                            open={open}
                            trigger={<Button floated='center' ><Icon name='file text' /></Button>}
                        >
                            <Modal.Header>Comentarios</Modal.Header>
                            <Modal.Content image>
                                <Image size='medium' src={props.logo} wrapped />
                                <Modal.Description>
                                    <Comment.Group>
                                        {comentario.map((dato, index) => (
                                            <Comentario
                                                comentario={dato}
                                            />
                                        ))}
                                        <Form reply>
                                            <Form.TextArea 
                                                onChange={e => setcome(e.target.value)}
                                            />
                                            <Button content='Add Reply' labelPosition='left' icon='edit' primary onClick={comentacion}/>
                                        </Form>
                                   
                                    </Comment.Group>
                                </Modal.Description>
                            </Modal.Content>
                        <Modal.Actions>

                            <Button
                                content="Yep, i finsih"
                                labelPosition='right'
                                icon='checkmark'
                                onClick={() => setOpen(false)}
                                positive
                            />
                        </Modal.Actions>
                        </Modal>

                    <Button basic color='green' floated='left' onClick={mandar}>
                        Mostrar Productos
                        </Button>
                </div>


            </div>
            <div className="extra content">
                <span>Contacto: {props.contacto}</span>
            </div>
        </div>
        </div >
    )
}

export default Carta
