import React from 'react'
import {  Menu } from 'semantic-ui-react'
import ListElement from './MonthsElements'

function MenuCargarPedidos(props) {
  console.log(props)
  return (
    <Menu vertical>
      <Menu.Item>Years</Menu.Item>
      {props.years.map((c, index) => (
        <ListElement year={c.Year}
          months={c.Months}
        />
      ))}
    </Menu>
  )
}

export default MenuCargarPedidos
