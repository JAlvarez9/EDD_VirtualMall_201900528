import React, { useState } from 'react'
import '../css/Carta.css'
import { Icon, Button, Header, Image, Modal, Form, Input } from 'semantic-ui-react'





function CartaProducto(props) {
    const [open, setOpen] = useState(false)
    const [nombre] = useState(props.name)
    const [descripcion] = useState(props.descripcion)
    const [image] = useState(props.image)
    const [mount] = useState(props.mount)
    const [price] = useState(props.price)
    const [id] = useState(props.id)
    const [tienda] = useState(props.tienda)
    const [departa] = useState(props.departa)
    const [califi] = useState(props.califi)
    const [cantidad, setcantidad] = useState(0)

    const enviar = () => {
        var json = [
            nombre,
            descripcion,
            image,
            mount,
            price,
            id,
            tienda,
            departa,
            califi,
            cantidad
        ]
        var datos = localStorage.getItem("productos")
        if (datos == null || datos === undefined) {
            localStorage.setItem("productos", JSON.stringify([json]))
        } else {
            datos = JSON.parse(datos)
            datos.push(json)
            console.log(datos)
            localStorage.setItem("productos", JSON.stringify(datos))
        }
        alert(JSON.stringify(json))
        setOpen(false)
    }

    return (
        <div className="column carta">
            <div className="ui card">

                <div className="image">
                    <Image size='medium' src={props.image} wrapped />
                </div>

                <div className="content">

                    <div className="header">{props.name}</div>

                    <div className="meta">
                        <s>Precio: <Icon disabled name='dollar sign' /> {props.price}</s>
                    </div>

                    <div className="description">{props.descripcion}</div>
                    <Modal
                        onClose={() => setOpen(false)}
                        onOpen={() => setOpen(true)}
                        open={open}
                        trigger={<Button>Comprar</Button>}
                    >
                        <Modal.Header>{props.name}</Modal.Header>
                        <Modal.Content image>
                            <Image size='medium' src={props.image} wrapped />
                            <Modal.Description>
                                <Header><Icon disabled name='dollar sign' />{props.price}</Header>
                                <p>
                                    {props.descripcion}
                                </p>
                                <p>Cantidad disponible {props.mount}</p>
                                <Form>
                                    <Form.Field inline>
                                        <label>Cantidad deseada</label>
                                        <Input onChange={e => setcantidad(e.target.value) } icon="edit" iconPosition='left' placeholder='###' />
                                    </Form.Field>
                                </Form>
                            </Modal.Description>

                        </Modal.Content>
                        <Modal.Actions>
                            <Button color='black' onClick={() => setOpen(false)}>
                                Nope, no quiero comprar
                            </Button>
                            <Button
                                content="Yep, agregar al carrito"
                                labelPosition='right'
                                icon='checkmark'
                                onClick={enviar}
                                positive
                            />
                        </Modal.Actions>
                    </Modal>


                </div>
                <div className="extra content">
                    <span>Cantidad: {props.mount}</span>
                </div>

            </div>
        </div>
    )
}

export default CartaProducto
