import React from 'react'
import {Button, Header, Icon, Segment } from 'semantic-ui-react'
import '../css/CargarTiendas.css'


const SegmentExamplePlaceholderInline = () => (
    <Segment placeholder>
      <Header icon>
        <Icon name='search' />
        Choose your file json in your device !
      </Header>
      <Segment.Inline>
        <input type="file" id="files"></input>
      </Segment.Inline>
      <Segment.Inline>
          <Button primary className="botoncito">Cargar Productos</Button>
      </Segment.Inline>
    </Segment>
  )
function CargarProductos() {
    return (
        <div>
            <SegmentExamplePlaceholderInline/>
        </div>
    )
}

export default CargarProductos
