import React from "react";
import 'bootstrap/dist/css/bootstrap.min.css'
import '../style/LoginRegister.css'
import image from '../images/login.svg'
import { useNavigate } from 'react-router-dom'
import Modal from "./Modal";
import useStore from "../store/User";

function Login() {

    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [userErr, setUserErr] = React.useState(false);
    const [passErr, setPassErr] = React.useState(false);
    const [modalShow, setModalShow] = React.useState(false);
    const [modalMessage, setModalMessage] = React.useState("Username/Pasword Salah");
    const { token, id, name, role, setToken, setId, setName, setRole } = useStore()
    const navigate = useNavigate();

    const handleLogin = async () => {
        //validation
        let aman = true;

        if (username.length === 0) {
            setUserErr(true);
            aman = false;
        } else {
            setUserErr(false);
        }

        if (password.length === 0) {
            setPassErr(true);
            aman = false;
        } else {
            setPassErr(false);
        }

        if (aman) {

            let result = await fetch(' https://api-dev-halloguru.herokuapp.com/login', {
                method: 'POST',
                body: JSON.stringify({ username, password }),
                headers: {
                    'Content-Type': 'application/json',
                }
            });
            result = await result.json();
            console.log(result);
            //zustand
            setToken(result.token)
            setId(result.id)
            setName(result.name)
            setRole(result.role)
            const test = { token: result.token }

            if (result.status === true) {
                localStorage.setItem("user-info", JSON.stringify(result));
                localStorage.setItem("TOKEN", JSON.stringify(test));
                navigate('/')
            } else {
                setModalShow(true)
            }
        }
    }

    return (
        <div className="auth-wrapper">
            <div className="login-wrapper d-flex">

                <div className="login-left ">
                    <form>
                        <h3>LOGIN</h3>
                        <div className="mb-3">
                            <label>Username</label>
                            <input
                                type="text"
                                className="form-control"
                                placeholder="Masukkan username"
                                onChange={(e) => setUsername(e.target.value)}
                            />
                            {userErr ? <span className="warning">Username tidak boleh kosong</span> : ""}
                        </div>
                        <div className="mb-3">
                            <label>Password</label>
                            <input
                                type="password"
                                className="form-control"
                                placeholder="Masukkan password"
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            {passErr ? <span className="warning">Password tidak boleh kosong</span> : ""}

                        </div>

                        <div className="d-grid">
                            <button onClick={handleLogin} type="button" className="btn btn-primary">
                                Submit
                            </button>
                        </div>
                        <p className="forgot-password text-right mt-3">

                            Belum memiliki akun?
                        </p>
                        <p className="forgot-password text-right">
                            <a href="/register">Register Murid </a> | <a href="/register/teacher">Register Guru</a>
                        </p>
                        <p className="forgot-password text-right">


                        </p>
                    </form>
                </div>

                <div className="login-right">
                    <img src={image} className="align-middle" alt="hallo guru" />
                    <p className="text-justify align-middle">Tempat terpercaya mencari guru private terbaik</p>
                </div>
            </div>

            <Modal
                show={modalShow}
                onHide={() => setModalShow(false)}
                message={modalMessage}
            />
        </div >
    );
}

export default Login;