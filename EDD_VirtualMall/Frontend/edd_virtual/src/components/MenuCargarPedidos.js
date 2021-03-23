import React, { useEffect, useState } from 'react'
import { Menu } from 'semantic-ui-react'
import ListElement from './MonthsElements'

const axios = require('axios')
function MenuCargarPedidos() {
  const [years, setyears] = useState([])
  const [loading, setloading] = useState(false)
  useEffect(() => {
    async function obtener() {
      if (years.length === 0) {
        const data = await axios.get(`http://localhost:3000/obtenerYears`);
        if (data.status !== 204) {
          console.log(data.data)
          setyears(data.data)
          setloading(true)
        }

      }
    }
    obtener()
  });
  if (loading === false) {
    return (
      <div className="ui segment carga">
        <div className="ui active dimmer">
          <div className="ui text loader">Loading</div>
        </div>
        <p />
      </div>
    )
  } else {
    return (
      <Menu vertical>
        <Menu.Item>Years</Menu.Item>
        {years.map((c, index) => (
          <ListElement year={c.Year}
            months={c.Months}
          />
        ))}
      </Menu>
    )
  }
}

export default MenuCargarPedidos
