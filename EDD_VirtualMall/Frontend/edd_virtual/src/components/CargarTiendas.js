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
        <input type="file"
          id='file'
          className='input-file'
          accept='.json'
        />
      </Segment.Inline>
      <Segment.Inline>
          <Button className="botoncito" primary>Cargar Tiendas</Button>
      </Segment.Inline>
    </Segment>
  )

function CargarTiendas() {
    return (
        <SegmentExamplePlaceholderInline/>
    )
}

export default CargarTiendas
