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
            <p>PT Hallo Raya Indonesia, melakukan bisnis sebagai Halloguru, adalah sebuah perusahaan rintisan digital asal Indonesia untuk anak berkebutuhan khusus yang bergerak di bidang pendidikan nonformal.
Halloguru menawarkan platform pembelajaran berbasis kurikulum sekolah melalui video tutorial interaktif oleh guru dan atau mendatangi guru ke rumah secara langsung. Perusahaan ini didirikan pada Juni 2022.</p>
        </div>
    )
}