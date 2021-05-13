import React from 'react'
import { Comment} from 'semantic-ui-react'

function ComentarioTienda(props) {
    return (
        <Comment>
            
            <Comment.Content>
                <Comment.Author as='a'>Efe</Comment.Author>
                
                <Comment.Text>{props.comentario}</Comment.Text>
                
            </Comment.Content>
        </Comment>
    )
}

export default ComentarioTienda
