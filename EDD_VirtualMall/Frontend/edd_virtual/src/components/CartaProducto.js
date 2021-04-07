import React, { useState } from 'react'
import '../css/Carta.css'
import { Icon, Button, Header, Image, Modal, Form, Input } from 'semantic-ui-react'





function CartaProducto(props) {
    let f = new Date();
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
    const [fecha] = useState(f.getDate() + "-" + (f.getMonth() + 1) + "-" + f.getFullYear())
    const [cliente, setcliente] = useState(window.sessionStorage.getItem("dpi"))
    const enviar = () => {


        var json = [
            nombre,
            descripcion,
            image,
            mount.toString(),
            price.toString(),
            id.toString(),
            tienda,
            departa,
            califi.toString(),
            cantidad,
            fecha,
            cliente
        ]
        var datos = localStorage.getItem("productos")
        if (datos == null || datos === undefined) {
            if( cantidad <= mount ){
                localStorage.setItem("productos", JSON.stringify([json]))
                alert("El producto se agrego al carrito :)")
                setOpen(false)
            }else{
                alert("La cantidad ingresada es mayor a la cantidad dentro de la tienda :( no se agregara al carrito")
            }
            
        } else {
            if(cantidad <= mount){
                datos = JSON.parse(datos)
            datos.push(json)
            console.log(datos)
            localStorage.setItem("productos", JSON.stringify(datos))
            alert("Su producto se agrego al carrito :)")
            setOpen(false)
            }else{
                alert("La cantidad ingresada es mayor a la cantidad dentro de la tienda :( no se agregara al carrito")
            }
            
        }
        
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
                                        <Input onChange={e => setcantidad(e.target.value)} icon="edit" iconPosition='left' placeholder='###' />
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
