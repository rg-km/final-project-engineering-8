import React, { useState, useEffect } from "react"
import axios from 'axios';
import { Link } from "react-router-dom";
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import '../style/Home.css';
import '../App.css';
import Teachers from "./Teachers";
import Navbar from './Navigation';
import TeacherSection from "./TeacherSection";
import HomeAnimation from "../images/Kids-Studying.gif"

function Home() {

    const [cardGuru, setCardGuru] = useState([]);

    const getCardGuruData = async () => {
        try {
            const SW_API_URL = `https://api-dev-halloguru.herokuapp.com/v1/teachers`
            const list = await axios.get(SW_API_URL);
            setCardGuru(list.data.data)
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

            <div className="container mt-5">
                <div className="d-flex justify-content-around">
                    <div className="col-md-5">
                        <img src={HomeAnimation} alt="kids-studying-animation" className="animated-gif" />

                    </div>

                    <div className="caption col-md-5">
                        <h3>SOLUSI TEPAT MENCARI GURU TERBAIK</h3>
                        <p>Lorem ipsum is a placeholder text commonly used to demonstrate the visual form of a document or a typeface without relying on meaningful content.</p>
                        <div className="d-flex flex-wrap col-md-12 pt-2">
                            <div className="col-md-4 col-sm-6 col-12">
                                <div className="card card-fitur">
                                    <p>Menyalurkan tenaga pendidik untuk ABK</p>
                                </div>
                            </div>
                            <div className="col-md-4 col-sm-6 col-12">
                                <div className="card card-fitur">
                                    <p>Menyalurkan tenaga pendidik untuk umum</p>
                                </div>
                            </div>
                            <div className="col-md-4 col-sm-6 col-12">
                                <div className="card card-fitur">
                                    <p>Menghadirkan tenaga pendidik ke rumah</p>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
            </div>

            <TeacherSection />

            <div className="container my-5 py-5">
                <h3>CARI GURU TERBAIK UNTUKMU</h3>

                <div className="d-flex justify-content-center my-5">
                    <div className="col-md-5 mx-3">
                        <input
                            type="text"
                            className="form-control"
                            placeholder="Cari"
                        />
                    </div>
                    <button type="button" className="btn btn-primary">
                        Cari
                    </button>
                </div>
                <div className="d-flex flex-wrap col-md-12 wrap-teacher">
                    {cardGuru.map((item, index) => {
                        return <div className="col-md-4 col-sm-6 col-12" key={index}>
                            <Link to={`/teacher/${item.id}`}>
                                <Teachers
                                    name={item.name}
                                    description={item.description}
                                    rating={item.rating}
                                    teaching_category={item.teaching_category}
                                    teaching_level={item.teaching_level}
                                    teaching_subject={item.teaching_subject}
                                    profilePict={item.profile_pict}
                                />
                            </Link>
                        </div>
                    }
                    )}
                </div>
            </div>
        </div >
    );
}

export default Home;