import React,{useState} from 'react'
import { Button, Image, Modal } from 'semantic-ui-react'


function ButtonArreglo(props) {

    const [open, setOpen] = React.useState(false)

    const [imagen] = useState("http://localhost:3000/arreglito/"+props.name)

    return (
        <Modal
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
            open={open}
            trigger={<Button style={{marginTop:10}}>{props.name}</Button>}
        >
            <Modal.Header>{props.name}</Modal.Header>
            <Modal.Content image>
                <Image size='big' src={imagen} />
            </Modal.Content>
            <Modal.Actions>
                <Button onClick={() => setOpen(false)}>Cancel</Button>
                <Button onClick={() => setOpen(false)} positive>
                    Ok
        </Button>
            </Modal.Actions>
        </Modal>
    )
}

export default ButtonArreglo
