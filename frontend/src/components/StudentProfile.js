import React, { useEffect, useState } from "react";
import axios from "axios";
import Navbar from './Navigation';
import { useNavigate } from 'react-router-dom'
import CInputImage from "./CInputImage";
import Modal from "./Modal"
import Footer from './Footer'

const StudentProfile = () => {
    const [detail, setDetail] = useState(null);
    const [loading, setLoading] = useState(false);
    const [isUpdating, setIsUpdating] = useState(false)
    const [modal, setModal] = useState({ isShow: false, message: '', onHide: () => { } })
    const [isEditMode, setIsEditMode] = useState(false)
    const [base64Image, setBase64Image] = useState(null)
    const auth = JSON.parse(localStorage.getItem("TOKEN"));
    const navigate = useNavigate();

    const loadDetail = async () => {
        setLoading(true);
        try {
            const url = "https://api-dev-halloguru.herokuapp.com/v1/student/update";
            const { data } = await axios.get(url, { headers: { "Authorization": `Bearer ${auth.token}` } });

            setDetail(data.data);
            console.log("profile user : ", data);

        } catch (error) {
            console.log(error);
        }
        setLoading(false);
    };

    const deleteUser = async () => {
        setLoading(true);
        try {
            const url = "https://api-dev-halloguru.herokuapp.com/v1/student";
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
            nama: detail?.nama,
            alamat: detail?.alamat,
            no_hp: detail?.no_hp,
            profile_pict: base64Image || detail?.profile_pict
        }
        console.log('PAYLOAD REQUEST', payload)
        try {
            const url = `https://api-dev-halloguru.herokuapp.com/v1/student/${detail?.user_id}`;
            const res = await axios.put(url, payload, { headers: { "Authorization": `Bearer ${auth.token}` } });
            if (!res.data.status) throw new Error(res.data.message)

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
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('nama')} value={detail?.nama} /> : detail?.nama}</td>
                            </tr>
                            <tr>
                                <th>Username</th>
                                <td>{isEditMode ? <CInputForm disabled onChange={handleOnChange('username')} value={detail?.username} /> : detail?.username}</td>
                            </tr>
                            <tr>
                                <th>Address</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('alamat')} value={detail?.alamat} /> : detail?.alamat}</td>
                            </tr>
                            <tr>
                                <th>No HP</th>
                                <td>{isEditMode ? <CInputForm onChange={handleOnChange('no_hp')} value={detail?.no_hp} /> : detail?.no_hp}</td>
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
