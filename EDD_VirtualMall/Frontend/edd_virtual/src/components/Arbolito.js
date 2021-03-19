import React, {useState, useEffect} from 'react'
import { Image } from 'semantic-ui-react'

const axios = require ('axios')
function Arbolito(props) {
    let id = props.match.params.id
    const [products, setproducts] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (products.length === 0) {
                const data = await axios.get(`http://localhost:3000/obtenerTiendas/${id}`);
                if (data.status !== 204) {
                    console.log(data.data)
                    setproducts(data.data)
                    setloading(true)
                }

            }
        }
        obtener()
    });
    return (
        <Image src='/images/wireframe/image.png' fluid />
    )
}

export default Arbolito
