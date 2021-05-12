import React from 'react'

import { Button, Image, Modal } from 'semantic-ui-react'
import { Graphviz } from 'graphviz-react';

function ButtonMerckle(props) {
    const optionsgp = { fit: true, lenght: 900, width: 850, zoom: true }
    return (
        <Modal
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
            open={open}
            trigger={<Button style={{marginTop:10}}>{props.name}</Button>}
        >
            <Modal.Header>{props.name}</Modal.Header>
            <Modal.Content image>
               <Graphviz options={optionsgp}  dot={props.graph}  ></Graphviz>
            </Modal.Content>
            <Modal.Actions>
                <Button onClick={() => setOpen(false)} positive>
                    Ok
        </Button>
            </Modal.Actions>
        </Modal>
    )
}

export default ButtonMerckle
