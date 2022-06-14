import React, { Component } from 'react'
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './LoginRegister.css'
import image from '../images/register.svg'

export default class Register extends Component {
    render() {
        return (
            <div className="auth-wrapper">

                <div className='register-wrapper d-flex'>
                    <div className="register-right">
                        <img src={image} className="align-middle" />
                        <p className="text-justify align-middle">Tempat terpercaya mencari guru private terbaik</p>
                    </div>
                    <div className="register-left">
                        <form>
                            <h3>Register</h3>
                            <div className="mb-3">
                                <label>Nama</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Masukkan Nama"
                                />
                            </div>
                            <div className="mb-3">
                                <label>Username</label>
                                <input type="text" className="form-control" placeholder="Masukkan Username" />
                            </div>
                            <div className="mb-3">
                                <label>Alamat</label>
                                <textarea
                                    type="text"
                                    className="form-control"
                                    placeholder="Masukkan Alamat"
                                />
                            </div>
                            <div className="mb-3">
                                <label>No HP</label>
                                <input type="text" className="form-control" placeholder="Masukkan No HP" />
                            </div>
                            <div className="mb-3">
                                <label>Password</label>
                                <input
                                    type="password"
                                    className="form-control"
                                    placeholder="Enter password"
                                />
                            </div>
                            <div className="d-grid">
                                <button type="submit" className="btn btn-primary">
                                    Sign Up
                                </button>
                            </div>
                            <p className="forgot-password text-right">
                                Sudah memiliki akun? <a href="/login">Login</a>
                            </p>
                        </form>
                    </div>
                </div>
            </div>
        )
    }
}