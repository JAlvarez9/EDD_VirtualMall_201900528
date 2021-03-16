import React from 'react'
import { Dropdown, Menu } from 'semantic-ui-react'

function MenuCargarPedidos() {
    return (
        <Menu vertical>
      <Menu.Item>Years</Menu.Item>
      <Dropdown text='2019' pointing className='link item'>
        <Dropdown.Menu>
          <Dropdown.Header>Months</Dropdown.Header>
          <Dropdown.Item> May </Dropdown.Item>
          <Dropdown.Item>April</Dropdown.Item>
          <Dropdown.Item>December</Dropdown.Item>
        </Dropdown.Menu>
      </Dropdown>
      <Dropdown text='1999' pointing className='link item'>
        <Dropdown.Menu>
          <Dropdown.Header>Months</Dropdown.Header>
          <Dropdown.Item> June </Dropdown.Item>
          <Dropdown.Item>January</Dropdown.Item>
          <Dropdown.Item>September</Dropdown.Item>
        </Dropdown.Menu>
      </Dropdown>
      
    </Menu>
    )
}

export default MenuCargarPedidos
