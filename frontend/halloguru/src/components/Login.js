import React from "react";
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './LoginRegister.css'
import image from '../images/login.svg'

function Navigation() {
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
                            />
                        </div>
                        <div className="mb-3">
                            <label>Password</label>
                            <input
                                type="password"
                                className="form-control"
                                placeholder="Masukkan password"
                            />
                        </div>
                        <div className="mb-3">
                            <div className="custom-control custom-checkbox">
                                <input
                                    type="checkbox"
                                    className="custom-control-input"
                                    id="customCheck1"
                                />
                                <label className="custom-control-label" htmlFor="customCheck1">
                                    Remember me
                                </label>
                            </div>
                        </div>
                        <div className="d-grid">
                            <button type="submit" className="btn btn-primary">
                                Submit
                            </button>
                        </div>
                        <p className="forgot-password text-right">
                            Belum memiliki akun? <a href="/register">Register</a>
                        </p>
                    </form>
                </div>

                <div className="login-right">
                    <img src={image} className="align-middle" />
                    <p className="text-justify align-middle">Tempat terpercaya mencari guru private terbaik</p>
                </div>
            </div>

        </div>
    );
}

export default Navigation;