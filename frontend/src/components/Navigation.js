import React, { useEffect } from "react";
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { useNavigate } from 'react-router-dom'


function Navigation() {
    const auth = localStorage.getItem("user-info");
    const navigate = useNavigate();
    const logout = () => {
        localStorage.removeItem("user-info");
        navigate('/');
    }

    return (

        <Navbar collapseOnSelect expand="lg" >
            <Container>
                <Navbar.Brand href="/">HALLOGURU</Navbar.Brand>
                <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link href={'/home'}>Home</Nav.Link>
                        <Nav.Link href="#tentang">Tentang</Nav.Link>
                    </Nav>
                    <Nav>
                        {!auth ?
                            <div className="d-flex">
                                <Nav.Link className="nav-link" href="/login">Login</Nav.Link>
                                <Nav.Link className="nav-link" href="/register">Register</Nav.Link>
                            </div>
                            :
                            <Nav.Link onClick={logout} href="/login">Logout</Nav.Link>
                        }

                        {/* <Nav.Link className="nav-link" href={'/login'}>
                            Login
                        </Nav.Link>
                        <Link className="nav-link" to={'/register'}>
                            Register
                        </Link> */}


                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default Navigation;