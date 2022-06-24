import React, { useEffect, useState } from "react";
import axios from "axios";
import Navbar from './Navigation';
import { Modal } from 'react-bootstrap';
import { Button } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom'


const StudentProfile = () => {
    const [detail, setDetail] = useState(null);
    const [loading, setLoading] = useState(false);
    const auth = JSON.parse(localStorage.getItem("user-info"));
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

    useEffect(() => {
        loadDetail();
    }, []);

    return (

        <div className="card-guru-detail" >
            <Navbar />

            {!loading ? (
                <div className="col-md-6 mx-auto my-5">
                    <h3 className="mt-5 mb-3">My Profile</h3>
                    <img src={detail?.profile_pict} />
                    <table className="table">
                        <tbody>
                            <tr>
                                <th>Nama</th>
                                <td>{detail?.nama}</td>
                            </tr>
                            <tr>
                                <th>Username</th>
                                <td>{detail?.username}</td>
                            </tr>
                            <tr>
                                <th>Address</th>
                                <td>{detail?.alamat}</td>
                            </tr>
                            <tr>
                                <th>No HP</th>
                                <td>{detail?.no_hp}</td>
                            </tr>
                        </tbody>
                    </table>
                    <button className="btn btn-primary m-3">Update</button>
                    <button className="btn btn-primary m-3" onClick={() => deleteUser()}>Delete</button>
                </div>
            ) : (
                <h2>Loading...</h2>
            )
            }
        </div >

    );
};

export default StudentProfile;
