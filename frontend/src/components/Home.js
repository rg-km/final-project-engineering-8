import React, { useState, useEffect } from "react"
import axios from 'axios';
import './Home.css';
import '../App.css';
import CardGuru from "./CardGuru";
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import Navbar from './Navigation';

function Home() {

    const [cardGuru, setCardGuru] = useState([]);

    const getCardGuruData = async () => {
        try {
            const SW_API_URL = `https://api-dev-halloguru.herokuapp.com/v1/teachers`
            const list = await axios.get(SW_API_URL, { withCredentials: true });
            setCardGuru(list.data.data)
            console.log(list.data.data, "dinda")
        } catch (error) {
            console.log('Error', error.message);
        }
    };

    useEffect(() => {
        getCardGuruData();
    }, []);

    return (

        <div className='bg-white col-md-12 min-vh-100'>
            <Navbar />
            <h1>Selamat Datang di Home
            </h1>
            <div className="container">

                <div className="d-flex flex-wrap">
                    <CardGuru
                        name="Dinda"
                        description="Saya memiliki keahlian dalam bidang matematika"
                        rating="4,5"
                        teaching_category="Anak Berkebutuhan Khusus"
                        teaching_level="SD"
                        teaching_subject="Matematika"
                    />
                    <CardGuru
                        name="Dinda"
                        description="Saya memiliki keahlian dalam bidang matematika"
                        rating="4,5"
                        teaching_category="Anak Berkebutuhan Khusus"
                        teaching_level="SD"
                        teaching_subject="Matematika"
                    />
                    <CardGuru
                        name="Dinda"
                        description="Saya memiliki keahlian dalam bidang matematika"
                        rating="4,5"
                        teaching_category="Anak Berkebutuhan Khusus"
                        teaching_level="SD"
                        teaching_subject="Matematika"
                    />
                    <CardGuru
                        name="Dinda"
                        description="Saya memiliki keahlian dalam bidang matematika"
                        rating="4,5"
                        teaching_category="Anak Berkebutuhan Khusus"
                        teaching_level="SD"
                        teaching_subject="Matematika"
                    />
                </div>
            </div>
        </div>
    );
}

export default Home;