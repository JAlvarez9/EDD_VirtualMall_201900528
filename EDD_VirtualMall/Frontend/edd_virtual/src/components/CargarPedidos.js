import React from 'react'
import {Button, Header, Icon, Segment } from 'semantic-ui-react'
import '../css/CargarTiendas.css'

const axios = require('axios')
const cargarTienditas = (event) => {
  const json = event.target.files[0];
  axios.post('http://localhost:3000/cargarpedidos',
    json,
    { headers: { 'content-type': 'application/json' } }
  ).then(data => {
    alert('file uploaded')
    console.log(data)
  }).catch(e => {
    alert('error')
    console.log(e)
  })

}


const SegmentExamplePlaceholderInline = () => (
    <Segment placeholder>
      <Header icon>
        <Icon name='search' />
        Choose your file json with pedidos in your device !
      </Header>
      <Segment.Inline>
        <input type="file" 
        id="files"
        accept='.json'
        onChange={cargarTienditas}
        ></input>
      </Segment.Inline>
      
    </Segment>
  )
function CargarPedidos() {
    return (
        <div>
            <SegmentExamplePlaceholderInline/>
        </div>
    )
}

export default CargarPedidos