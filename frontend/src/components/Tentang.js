import React, { useState, useEffect } from "react"
import axios from 'axios';
import { Link } from "react-router-dom";
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import '../style/Home.css';
import '../App.css';
import Navbar from './Navigation';

function Tentang() {
    return (
        <div className="About">
            <Navbar />
            <p>PT Hallo Raya Indonesia, melakukan bisnis sebagai Halloguru, adalah sebuah perusahaan rintisan digital asal Indonesia untuk anak berkebutuhan khusus yang bergerak di bidang pendidikan nonformal.
Halloguru menawarkan platform pembelajaran berbasis kurikulum sekolah melalui video tutorial interaktif oleh guru dan atau mendatangi guru ke rumah secara langsung. Perusahaan ini didirikan oleh Amalia Suherman, Dinda Wahyu Rahmadani, Fermi Naufal Akbar, Gilang Safera Putra, dan Shafa Salsabilla Buchori pada Juni 2022.</p>
        </div>
    )
}
export default Tentang;