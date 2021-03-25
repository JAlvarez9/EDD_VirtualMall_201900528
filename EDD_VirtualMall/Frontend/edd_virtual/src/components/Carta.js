import React from 'react'
import '../css/Carta.css'
import { useHistory } from 'react-router-dom'
import { Button, Icon } from 'semantic-ui-react'



function Carta(props) {
    const histori = useHistory();
    const mandar = () => {
        histori.push(`/mostrarproductos/${props.id}`)
    }
    const mandararbolito = () => {
        histori.push(`/mostrararbol/${props.id}`)
    }
    return (
        <div className="column carta">
            <div className="ui card">

                <div className="image">
                    <img  src={props.logo} />
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
                        <Button basic color='green' floated='left' onClick={mandar}>
                            Mostrar Productos
                        </Button>
                    </div>
                    

                </div>
                <div className="extra content">
                    <span>Contacto: {props.contacto}</span>
                </div>
            </div>
        </div>
    )
}

export default Carta
