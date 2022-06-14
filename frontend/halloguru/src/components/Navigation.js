import React from "react";
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { Link } from 'react-router-dom'


function Navigation() {
    return (

        <Navbar collapseOnSelect expand="lg" >
            <Container>
                <Navbar.Brand href="/">HALLOGURU</Navbar.Brand>
                <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                <Navbar.Collapse id="responsive-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link href={'/home'}>Home</Nav.Link>
                        <Nav.Link href="#pricing">Pricing</Nav.Link>
                    </Nav>
                    <Nav>
                        <Nav.Link className="nav-link" href={'/login'}>
                            Login
                        </Nav.Link>
                        <Link className="nav-link" to={'/register'}>
                            Register
                        </Link>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default Navigation;