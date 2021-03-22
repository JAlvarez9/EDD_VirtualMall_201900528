import React from 'react'
import { Header, Icon, Segment } from 'semantic-ui-react'
import '../css/CargarTiendas.css'

const axios = require('axios')
const cargarTienditas = (event) => {
  console.log("asd")
  const json = event.target.files[0];
  axios.post('http://localhost:3000/cargartienda',
    json,
    { headers: { 'content-type': 'application/json' } }
  ).then(data => {
    alert('file uploaded')
    console.log(data)
  }).catch(e => {
    console.log('error')
    console.log(e)
  })


}


function CargarTiendas() {
  return (
    <Segment placeholder>
      <Header icon>
        <Icon name='search' />
        Choose your file json with stores in your device !
      </Header>
      <Segment.Inline>
        <input type="file"
          id='file'
          className='input-file'
          accept='.json'
          onChange={cargarTienditas}
        />
      </Segment.Inline>


    </Segment>
  )
}

export default CargarTiendas
