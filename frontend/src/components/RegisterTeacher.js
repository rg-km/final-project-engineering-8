import React, { useState } from 'react'
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import '../style/LoginRegister.css'
import '../App.css'
import image from '../images/register.svg'
import { useNavigate } from 'react-router-dom'

export default function Register() {

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [nama, setNama] = useState("")
    const [alamat, setAlamat] = useState("");
    const [no_hp, setNo_hp] = useState("");

    const [namaErr, setNamaErr] = useState(false);
    const [userErr, setUserErr] = useState(false);
    const [alamatErr, setAlamatErr] = useState(false);
    const [noHPErr, setNoHPErr] = useState(false);
    const [passErr, setPassErr] = useState(false);
    const navigate = useNavigate();

    async function collectData() {
        //validation
        let aman = true;

        if (nama.length === 0) {
            setNamaErr(true);
            aman = false;
        } else {
            setNamaErr(false);
        }

        if (username.length <= 3) {
            setUserErr(true);
            aman = false;
        } else {
            setUserErr(false);
        }

        if (alamat.length === 0) {
            setAlamatErr(true);
            aman = false;
        } else {
            setAlamatErr(false);
        }

        if (no_hp.length === 0) {
            setNoHPErr(true);
            aman = false;
        } else {
            setNoHPErr(false);
        }

        if (password.length <= 7) {
            setPassErr(true)
            aman = false;
        } else {
            setPassErr(false)
        }

        console.log("aman ", aman)
        if (aman) {
            let item = { username, password, nama, alamat, no_hp };
            console.warn(item)

            let result = await fetch('https://api-dev-halloguru.herokuapp.com/register/teacher', {
                method: 'POST',
                body: JSON.stringify(item),
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json"
                }
            });
            result = await result.json();
            //console.warn("result", result);
            //localStorage.setItem("user-info", JSON.stringify(result));
            if (result.data) {

                navigate('/login');
            }
        }

    }

    return (
        <div className="auth-wrapper" >

            <div className='register-wrapper d-flex'>
                <div className="register-right">
                    <img src={image} className="align-middle" />
                    <p className="text-justify align-middle">Tempat terpercaya mencari guru private terbaik</p>
                </div>
                <div className="register-left">
                    <form>
                        <h3>Register Teacher</h3>
                        <div className='d-flex'>
                            <div className="mb-3 jarakKanan">
                                <label>Nama</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Masukkan Nama"
                                    value={nama}
                                    onChange={(e) => setNama(e.target.value)}
                                />
                                {namaErr ? <span className="warning">Nama tidak boleh kosong</span> : ""}

                            </div>
                            <div className="mb-3">
                                <label>Username</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Minimal 4 karakter"
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                />
                                {userErr ? <span className="warning">Minimal 4 karakter</span> : ""}
                            </div>
                        </div>
                        <div className="mb-3">
                            <label>Alamat</label>
                            <textarea
                                type="text"
                                className="form-control"
                                placeholder="Masukkan Alamat"
                                value={alamat}
                                onChange={(e) => setAlamat(e.target.value)}
                            />
                            {alamatErr ? <span className="warning">Alamat tidak boleh kosong</span> : ""}

                        </div>
                        <div className="mb-3">
                            <label>No HP</label>
                            <input
                                type="number"
                                className="form-control"
                                placeholder="Masukkan No HP"
                                value={no_hp}
                                onChange={(e) => setNo_hp(e.target.value)}
                            />
                            {noHPErr ? <span className="warning">No HP tidak boleh kosong</span> : ""}

                        </div>
                        <div className="d-flex">
                            <div className="mb-3 jarakKanan">
                                <label for="exampleFormControlSelect1">Jenjang Didik</label>
                                <select class="form-control" id="exampleFormControlSelect1">
                                    <option>SD</option>
                                    <option>SMP</option>
                                    <option>SMA</option>
                                    <option>Semua Jenjang</option>
                                </select>
                            </div>
                            <div className="mb-3">
                                <label for="exampleFormControlSelect1">Mata Pelajaran Didik</label>
                                <select class="form-control" id="exampleFormControlSelect1">
                                    <option>Matematika</option>
                                    <option>Bahasa Inggris</option>
                                    <option>Kimia</option>
                                    <option>Biologi</option>
                                    <option>Fisika</option>
                                    <option>Semua Mata Pelajaran</option>
                                </select>
                            </div>
                        </div>
                        <div className="mb-3">
                            <label for="exampleFormControlSelect1">Kategori Didik</label>
                            <select class="form-control" id="exampleFormControlSelect1">
                                <option>ABK</option>
                                <option>NON ABK</option>
                                <option>Semua Kategori</option>
                            </select>
                        </div>
                        <div className="mb-3">
                            <label>Password</label>
                            <input
                                type="password"
                                className="form-control"
                                placeholder="Password minimal 8 karakter"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            {passErr ? <span className="warning">Password minimal 8 karakter</span> : ""}
                        </div>

                        <div className="d-grid">
                            <button type="button" onClick={collectData} className="btn btn-primary">
                                Sign Up
                            </button>
                        </div>

                    </form>
                </div>
            </div >
        </div >
    )

}