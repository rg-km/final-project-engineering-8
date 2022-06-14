import React from "react";
import './Home.css';
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import Navbar from './Navigation';


function Home() {
    return (

        <div className='bg-white col-md-12 min-vh-100'>
            <Navbar />
            <h1>Selamat Datang di Home
            </h1>
        </div>
    );
}

export default Home;