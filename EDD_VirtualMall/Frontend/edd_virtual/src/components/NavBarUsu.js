import React from 'react'
import { Menu, MenuItem } from 'semantic-ui-react'
import {Link} from 'react-router-dom'

const state = { activeItem: 'home' }

const opciones = ['Principal','Mostrar Tiendas']
const urls = ['/principal','/mostrartiendas']

function HeaderUsu() {
    const { activeItem } = state
    return (
        <Menu pointing>
        {opciones.map((c,index)=>(
            <MenuItem
                as={Link} to={urls[index]}
                name={opciones[index]}
                active={activeItem === c}
                key={index}
            />
        ))}
      </Menu>
    )
}

export default HeaderUsu