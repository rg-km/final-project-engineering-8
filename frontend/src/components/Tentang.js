import React, { useState, useEffect } from "react"
import axios from 'axios';
import { Link } from "react-router-dom";
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import '../style/Home.css';
import '../App.css';
import Navbar from './Navigation';
import Footer from './Footer';

function Tentang() {
    return (
        <div className="About">
            <Navbar />
            <div className="container my-5">
                <h3>Tentang Kami</h3>
                <p>PT Hallo Raya Indonesia, melakukan bisnis sebagai Halloguru, adalah sebuah perusahaan rintisan digital asal Indonesia untuk anak berkebutuhan khusus yang bergerak di bidang pendidikan nonformal.
                    Halloguru menawarkan platform pembelajaran berbasis kurikulum sekolah melalui video tutorial interaktif oleh guru dan atau mendatangi guru ke rumah secara langsung. Perusahaan ini didirikan oleh Amalia Suherman, Dinda Wahyu Rahmadani, Fermi Naufal Akbar, Gilang Safera Putra, dan Shafa Salsabilla Buchori pada Juni 2022.</p>
            </div>
            <div className="container my-5">
                <h3>Team Kami</h3>
                <div className="d-flex flex-wrap">
                    <div className="col-md-4">
                        <div className="card p-3">
                            <p>FERMI NAUFAL AKBAR</p>
                            <p>Backend</p>
                        </div>
                    </div>
                    <div className="col-md-4">
                        <div className="card p-3">
                            <p>Gilang Safera Putra</p>
                            <p>Backend</p>
                        </div>
                    </div>
                    <div className="col-md-4">
                        <div className="card p-3">
                            <p>Shafa Salsabilla Buchori</p>
                            <p>Backend</p>
                        </div>
                    </div>
                    <div className="col-md-4">
                        <div className="card p-3">
                            <p>Dinda Wahyu Rahmadani</p>
                            <p>Frontend</p>

                        </div>
                    </div>
                    <div className="col-md-4">
                        <div className="card p-3">
                            <p>Amalia Suherman</p>
                            <p>Frontend</p>
                        </div>
                    </div>
                </div>
            </div>

            <Footer />
        </div>
    )
}
export default Tentang;