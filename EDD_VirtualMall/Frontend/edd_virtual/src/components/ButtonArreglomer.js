import React,{useState} from 'react'
import { Button, Image, Modal } from 'semantic-ui-react'
import { Graphviz } from 'graphviz-react';

function ButtonArreglomer(props) {

    const [open, setOpen] = React.useState(false)
    const [grafi, setgrafi] = useState("")

    const conversion = () => {
        let asdf = props.graph.replace('\"', '"')
        console.log(asdf)
        setgrafi =asdf
    }

    conversion()

    return (
        <Modal
            onClose={() => setOpen(false)}
            onOpen={() => setOpen(true)}
            open={open}
            trigger={<Button style={{marginTop:10}}>{props.name}</Button>}
        >
            <Modal.Header>{props.name}</Modal.Header>
            <Modal.Content image>
                <Graphviz dot={grafi}/>
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

export default ButtonArreglomer
