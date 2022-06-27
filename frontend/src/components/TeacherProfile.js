import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { useNavigate } from 'react-router-dom'
import CInputImage from "./CInputImage";
import axios from "axios";
import Navbar from './Navigation';
import Modal from "./Modal"
import Footer from './Footer'
import BackButton from "./BackButton";

const StudentProfile = () => {
    const [detail, setDetail] = useState(null);
    const [loading, setLoading] = useState(false);
    const [token, setToken] = useState('');
    const { id } = useParams();
    const [isUpdating, setIsUpdating] = useState(false)
    const [modal, setModal] = useState({ isShow: false, message: '', onHide: () => { } })
    const [isEditMode, setIsEditMode] = useState(false)
    const [base64Image, setBase64Image] = useState(null)
    const [jenjang_id, setJenjang_id] = useState(1);
    const [pelajaran_id, setPelajaran] = useState(1);
    const [kategori_id, setKategori] = useState(1);

    const auth = JSON.parse(localStorage.getItem("TOKEN"));
    const navigate = useNavigate();

    const loadDetail = async () => {
        setLoading(true);
        try {
            const url = "https://api-dev-halloguru.herokuapp.com/v1/teacher/" + id;
            const { data } = await axios.get(url, { headers: { "Authorization": `Bearer ${auth.token}` } });

            setDetail(data.data);
            console.log("profile teacher : ", data.data);

        } catch (error) {
            console.log(error);
        }
        setLoading(false);
    };

    const deleteUser = async () => {
        setLoading(true);
        try {
            const url = "https://api-dev-halloguru.herokuapp.com/v1/teacher/" + id;
            await axios.delete(url, { headers: { "Authorization": `Bearer ${auth.token}` } });
            localStorage.removeItem("user-info");
            navigate('/home');

        } catch (error) {
            console.log(error);
        }
        setLoading(false);
    }

    const handleOnChange = (type) => (e) => {
        e.preventDefault()
        setDetail(prevState => {
            return {
                ...prevState,
                [type]: e.target.value
            }
        })
    }

    const handleUpdateStudent = async () => {
        setIsUpdating(true);
        const payload = {
            name: detail?.name,
            address: detail?.address,
            no_hp: detail?.no_hp,
            description: detail?.description,
            fee: detail?.fee,
            teaching_category: kategori_id,
            teaching_level: jenjang_id,
            teaching_subject: pelajaran_id,
            profile_pict: base64Image || detail?.profile_pict
        }
        console.log('PAYLOAD REQUEST', payload)
        try {
            const url = `https://api-dev-halloguru.herokuapp.com/v1/teacher/${id}`;
            const res = await axios.put(url, payload, { headers: { "Authorization": `Bearer ${auth.token}` } });
            if (!res.data.status) throw new Error(res.data.message)

            const result = {
                id: res.data.data.id,
                name: res.data.data.name,
                role: "guru",
                profile_pict: res.data.data.profile_pict,
            }
            localStorage.setItem("user-info", JSON.stringify(result));
            setModal({
                isShow: true,
                message: "Update success",
                onHide: () => { navigate("/home") }
            })

        } catch (error) {
            console.log(error);
            setModal({
                isShow: true,
                message: "Failed Update, ERROR:" + error.message,
                onHide: () => { setModal(prev => ({ ...prev, isShow: false })) }
            })
        } finally {
            setIsUpdating(false);
        }
    }

    useEffect(() => {
        loadDetail();
    }, []);

    const buttonClass = isUpdating ? "btn btn-primary m-3 disabled" : "btn btn-primary m-3"

    return (

        <div className="card-guru-detail" >
            <Navbar />

            {!loading ? (
                <div className="col-md-6 mx-auto my-5">
                    <h3 className="mt-5 mb-3">My Profile</h3>
                    {/* <img src={detail?.profile_pict} /> */}
                    <CInputImage value={detail?.profile_pict} onFileChange={(value) => setBase64Image(value)} isShowUploadButton={isEditMode} resetDefaultState={!isEditMode} />
                    <table className="table">
                        <tbody>
                            <tr>
                                <th>Nama</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('name')} value={detail?.name} /> : detail?.name}</td>
                            </tr>
                            <tr>
                                <th>Address</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('address')} value={detail?.address} /> : detail?.address}</td>
                            </tr>
                            <tr>
                                <th>No HP</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('no_hp')} value={detail?.no_hp} /> : detail?.no_hp}</td>
                            </tr>
                            <tr>
                                <th>Description</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('description')} value={detail?.description} /> : detail?.description}</td>
                            </tr>
                            <tr>
                                <th>Fee</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('fee')} value={detail?.fee} /> : detail?.fee}</td>
                            </tr>
                            <tr>
                                <th>Rating</th>
                                <td>{isEditMode ? <CInputForm disabled onChange={handleOnChange('rating')} value={detail?.rating} /> : detail?.rating}</td>
                            </tr>
                            <tr>
                                <th>Kategori</th>
                                <td>
                                    {isEditMode ?
                                        <select
                                            className="form-control"
                                            id="exampleFormControlSelect1"
                                            value={kategori_id}
                                            onChange={(e) => setKategori(parseInt(e.target.value))}>
                                            <option value="1">ABK</option>
                                            <option value="2">NON ABK</option>
                                            <option value="3">Semua Kategori</option>
                                        </select>
                                        :
                                        detail?.teaching_category
                                    }

                                </td>
                            </tr>
                            <tr>
                                <th>Tingkatan Mengajar</th>
                                <td>
                                    {isEditMode ?
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
                                        :
                                        detail?.teaching_level
                                    }
                                </td>
                            </tr>
                            <tr>
                                <th>Mata Pelajaran</th>
                                <td>
                                    {isEditMode ?
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
                                        :
                                        detail?.teaching_subject
                                    }
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    {
                        isEditMode &&
                        <button className={buttonClass} onClick={handleUpdateStudent}>{isUpdating ? "Updating Profile..." : "Save Changes"}</button>
                    }
                    <button className={buttonClass} onClick={() => setIsEditMode(isEdit => !isEdit)}>{isEditMode ? "Cancel" : "Update"}</button>
                    <button className={buttonClass} onClick={() => deleteUser()}>Delete</button>
                </div>
            ) : (
                <h2>Loading...</h2>
            )
            }
            <Modal
                show={modal.isShow}
                onHide={modal.onHide}
                message={modal.message}
            />
            <Footer />
        </div >

    );
};

const CInputForm = ({ onChange, value, disabled }) => {
    return (
        <div>
            <input
                type="text"
                className="form-control"
                value={value}
                onChange={onChange}
                disabled={disabled}
            />
            {!value ? <span className="warning">Field harus diisi</span> : ""}
        </div>
    )
}

export default StudentProfile;
