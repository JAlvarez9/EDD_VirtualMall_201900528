import React, { useState } from 'react'
import { Menu, MenuItem } from 'semantic-ui-react'
import { Link } from 'react-router-dom'

const state = { activeItem: 'home' }


const opciones = ['Principal', 'Cargar Tiendas', 'Cargar Productos', 'Cargar Pedidos', 'Mostrar Pedidos', 'Mostrar Tiendas','Reportes']
const urls = ['/principal', '/cargartienda', '/cargarproducto', 'cargarpedidos', '/mostrarpedidos', '/mostrartiendas', '/reportes']


const opciones2 = ['Principal', 'Mostrar Tiendas']
const urls2 = ['/principal', '/mostrartiendas']

function Navbar() {
  
    const [cuent] = useState(window.sessionStorage.getItem("cuenta"))
    console.log(cuent)
    const { activeItem } = state

    if (cuent === "Admin") {
        return (
            <Menu pointing>
                {opciones.map((c, index) => (
                    <MenuItem
                        as={Link} to={urls[index]}
                        name={opciones[index]}
                        active={activeItem === c}
                        key={index}
                    />
                ))}
            </Menu>
        )
    } else if (cuent === "Usuario") {
        return (
            <Menu pointing>
                {opciones2.map((c, index) => (
                    <MenuItem
                        as={Link} to={urls2[index]}
                        name={opciones[index]}
                        active={activeItem === c}
                        key={index}
                    />
                ))}
            </Menu>
        )

    } else {
        return (
            <Menu pointing>

            </Menu>
        )

    }

}

export default Navbar
