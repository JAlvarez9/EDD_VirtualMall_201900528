import React from 'react'
import Carta from './Carta'

function Mosaico(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.stores.map((c, index) => (
                    <Carta nombre={c.Nombre}
                        descripcion={c.Descripcion}
                        contacto = {c.Contacto}
                        logo={c.Logo}
                        calificacion = {c.Calificacion}
                        id={c.Key}
                        key={c.Key}
                    />
                ))}
            </div>
        </div>
    )
}

export default Mosaico
