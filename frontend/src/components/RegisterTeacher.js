import React, { useState } from 'react'
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import '../style/LoginRegister.css'
import '../App.css'
import image from '../images/register.svg'
import { useNavigate } from 'react-router-dom'
import Modal from "./Modal";


export default function Register() {

    //item
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [nama, setNama] = useState("")
    const [alamat, setAlamat] = useState("");
    const [no_hp, setNo_hp] = useState("");
    const [deskripsi, setDeskripsi] = useState("");
    const [fee, setFee] = useState("");
    const [jenjang_id, setJenjang_id] = useState(1);
    const [pelajaran_id, setPelajaran] = useState(1);
    const [kategori_id, setKategori] = useState(1);

    //modal
    const [modalShow, setModalShow] = useState(false);
    const [modalMessage, setModalMessage] = useState("Register Berhasil, silahkan Login sebagai guru untuk melanjutkan.");

    //error
    const [namaErr, setNamaErr] = useState(false);
    const [userErr, setUserErr] = useState(false);
    const [alamatErr, setAlamatErr] = useState(false);
    const [noHPErr, setNoHPErr] = useState(false);
    const [passErr, setPassErr] = useState(false);
    const [deskripsiErr, setDeskripsiErr] = useState(false);
    const [biayaErr, setBiayaErr] = useState(false);
    const navigate = useNavigate();

    const rupiah = (number) => {
        return new Intl.NumberFormat("id-ID", {
            style: "currency",
            currency: "IDR"
        }).format(number);
    }


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

        if (aman) {
            let biaya = "Rp. " + rupiah(fee);
            let item = { username, password, nama, alamat, no_hp, deskripsi, biaya, jenjang_id, pelajaran_id, kategori_id };
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

            if (result.data) {
                setModalShow(true);
                setUsername("");
                setNama("");
                setAlamat("")
                setPassword("")
                setNo_hp("")
                setFee("")
                setDeskripsi("")
            }
        }

    }

    return (
        <div className='container p-5'>
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
                            <div className="mb-3">
                                <label>deskripsi</label>
                                <textarea
                                    type="text"
                                    className="form-control"
                                    placeholder="Masukkan Deskripsi Singkat"
                                    value={deskripsi}
                                    onChange={(e) => setDeskripsi(e.target.value)}
                                />
                                {deskripsiErr ? <span className="warning">Alamat tidak boleh kosong</span> : ""}

                            </div>
                            <div className="mb-3">
                                <label>Biaya</label>
                                <input
                                    type="number"
                                    className="form-control"
                                    placeholder="Masukkan biaya mengajar"
                                    value={fee}
                                    onChange={(e) => setFee(e.target.value)}
                                />
                                {biayaErr ? <span className="warning">No HP tidak boleh kosong</span> : ""}

                            </div>
                            <div className="d-flex">
                                <div className="mb-3 jarakKanan">
                                    <label>Jenjang Didik</label>
                                    <select
                                        className="form-control"
                                        id="exampleFormControlSelect1"
                                        value={jenjang_id}
                                        onChange={(e) => setJenjang_id(parseInt(e.target.value))}>
                                        <option value="1">SD</option>
                                        <option value="2">SMP</option>
                                        <option value="3">SMA</option>
                                        <option value="4">Semua Jenjang</option>
                                    </select>
                                </div>
                                <div className="mb-3">
                                    <label>Mata Pelajaran Didik</label>
                                    <select
                                        className="form-control"
                                        id="exampleFormControlSelect2"
                                        value={pelajaran_id}
                                        onChange={(e) => setPelajaran(parseInt(e.target.value))}>
                                        <option value="1">Matematika</option>
                                        <option value="2">Bahasa Inggris</option>
                                        <option value="3">Kimia</option>
                                        <option value="4">Biologi</option>
                                        <option value="5">Fisika</option>
                                        <option value="6">Semua Mata Pelajaran</option>
                                    </select>
                                </div>
                            </div>
                            <div className="mb-3">
                                <label>Kategori Didik</label>
                                <select
                                    className="form-control"
                                    id="exampleFormControlSelect3"
                                    value={kategori_id}
                                    onChange={(e) => setKategori(parseInt(e.target.value))}>
                                    <option value="1">ABK</option>
                                    <option value="2">NON ABK</option>
                                    <option value="3">Semua Kategori</option>
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
            <Modal
                show={modalShow}
                onHide={() => setModalShow(false)}
                message={modalMessage}
            />
        </div>
    )

}