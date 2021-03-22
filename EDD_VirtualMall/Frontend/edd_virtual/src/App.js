import React from 'react'
import {BrowserRouter as Router, Route} from 'react-router-dom'
import NavBar from "./components/Navbar";
import Header from "./components/Headerr"
import CargarTiendas from './components/CargarTiendas';
import CargarProductos from './components/CargarProductos';
import CargarPedidos from './components/CargarPedidos';
import MostrarPedidos from './components/MostrarPedidos';
import MostrarTiendas from './components/MostrarTiendas';
import PaginaPrincipal from './components/PaginaPrincipal';
import MostrarProductos from './components/MostrarProductos';
import Arbolito from './components/Arbolito'
import Carrito from './components/Carrito';



function App() {
  return (
    <>
    <Router>
        <Header />
        <NavBar/>
        <Route path="/principal" component={PaginaPrincipal}/>
        <Route path="/cargartienda" component={CargarTiendas}/>
        <Route path="/cargarproducto" component={CargarProductos}/>
        <Route path="/cargarpedidos" component={CargarPedidos}/>
        <Route path="/mostrarpedidos" component={MostrarPedidos}/>
        <Route path="/mostrartiendas" component={MostrarTiendas}/>
        <Route path="/mostrarproductos/:id" component={MostrarProductos}/>
        <Route path="/mostrararbol/:id" component={Arbolito}/>
        <Route path="/carrito" component={Carrito}/>
    </Router>
    </>
  );
}

export default App;