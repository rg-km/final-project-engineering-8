import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import Navbar from './Navigation';

const TeacherDetail = () => {
    const [detail, setDetail] = useState(null);
    const [loading, setLoading] = useState(false);
    const { id } = useParams();
    const [urlWA, setURLWA] = useState();
    const controller = new AbortController();

    const loadDetail = async () => {
        setLoading(true);
        try {
            const url = "https://api-dev-halloguru.herokuapp.com/v1/teacher/" + id;
            const { data } = await axios.get(url);

            setDetail(data.data);
            console.log("detail teacher : ", data);
            setURLWA("https://wa.me/" + detail.no_hp);
        } catch (error) {
            console.log(error);
        }
        setLoading(false);
    };

    useEffect(() => {
        loadDetail();
    }, []);

    return (

        <div className="card-guru-detail" >
            <Navbar />

            {!loading ? (
                <div className="col-md-6 mx-auto my-5">
                    <h3 className="mt-5 mb-3">DETAIL GURU</h3>
                    <table className="table">
                        <tbody>
                            <tr>
                                <th>Nama</th>
                                <td>{detail?.name}</td>
                            </tr>
                            <tr>
                                <th>Address</th>
                                <td>{detail?.address}</td>
                            </tr>
                            <tr>
                                <th>Description</th>
                                <td>{detail?.description} cm</td>
                            </tr>
                            <tr>
                                <th>Fee</th>
                                <td>{detail?.fee}</td>
                            </tr>
                            <tr>
                                <th>Rating</th>
                                <td>{detail?.rating}</td>
                            </tr>
                            <tr>
                                <th>Kategori</th>
                                <td>{detail?.teaching_category}</td>
                            </tr>
                            <tr>
                                <th>Tingkatan Mengajar</th>
                                <td>{detail?.teaching_level}</td>
                            </tr>
                            <tr>
                                <th>Mata Pelajaran</th>
                                <td>{detail?.teaching_subject}</td>
                            </tr>
                        </tbody>
                    </table>
                    <button className="btn btn-primary mt-3"> <a href={urlWA} target="_blank">Hubungi</a>

                    </button>
                </div>
            ) : (
                <h2>Loading...</h2>
            )
            }
        </div >

    );
};

export default TeacherDetail;
