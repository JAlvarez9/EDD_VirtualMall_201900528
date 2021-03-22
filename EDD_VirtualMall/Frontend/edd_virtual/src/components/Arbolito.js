import React, {useState} from 'react'
import { Image } from 'semantic-ui-react'

const axios = require('axios')
function Arbolito(props) {
    let id = props.match.params.id
    var contacto = id.split("$")
    function mostrar(){
        const data = axios.get(`http://localhost:3000/arbolito/${contacto[1]}`);  
        console.log(contacto[1])
        
        return "http://localhost:3000/arbolito/"+contacto[1]
        

    }

    const [imagen] = useState(mostrar())
    console.log(imagen)
    

    const data = axios.get(`http://localhost:3000/obtenerArbolito/${id}`);
    console.log(data)
    return (
        <Image src={ imagen }></Image>
    )


}

export default Arbolito
