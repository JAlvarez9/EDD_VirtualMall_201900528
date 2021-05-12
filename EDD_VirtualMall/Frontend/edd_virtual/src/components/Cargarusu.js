import React from 'react'
import { Header, Icon, Segment } from 'semantic-ui-react'
import { useHistory } from 'react-router-dom'
import '../css/CargarTiendas.css'

const axios = require('axios')
const cargarUsuarios = (event) => {
    const json = event.target.files[0];
    var id = (document.getElementById("master")).value
    console.log(id)
    if (id === "") {
        id = "$"
    }
    console.log(json)
    axios.post(`http://localhost:3000/cargarusuarios/${id}`,
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

function Cargarusu() {
  const histori = useHistory();
  const verifyRoute = () => {
    if (!window.sessionStorage.getItem("cuenta")) {
      return histori.push('/inicio');
    }
  };
  verifyRoute()
    return (
        <Segment placeholder>
          <Header icon>
            <Icon name='search' />
            Choose your file json with Users in your device !
          </Header>
          <Segment.Inline>
            <input type="file"
              id='file'
              className='input-file'
              accept='.json'
              onChange={cargarUsuarios}
            />
          </Segment.Inline>
    
    
        </Segment>
      )
}

export default Cargarusu
