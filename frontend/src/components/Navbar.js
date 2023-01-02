import React from "react";
import {NavLink} from "react-router-dom";
import {Nav, Navbar, Container} from "react-bootstrap";
import axios from "axios";

const NavHome = ({setAccountDetails, setUserLoggedIn, userLoggedIn}) => {

    const Logout = async () => {
        var url = "http://localhost:5000/api/logout"
        await axios.post(url, {}, { 
            withCredentials: true,
        }).then(response => {
            if (response.status === 200) {
                console.log(response);
            }
        })
        setAccountDetails({firstname: "", lastname: "", email: "", userid: "",});
        setUserLoggedIn(false);
    }
    
    let activeClassName = " active-nav"

    let NavigationMenu;

    if (!userLoggedIn) {
        NavigationMenu = (
            <Nav className="justify-content-end">
                <NavLink to="/login" className={({ isActive }) => "me-2 px-4 btn btn-bd-primary" + (isActive ? activeClassName : "")}>Login</NavLink>
                <NavLink to="/signup" className={({ isActive }) => "me-2 px-4 btn btn-bd-primary" + (isActive ? activeClassName : "")}>Sign Up</NavLink>
            </Nav>
        )
    } else {
        NavigationMenu = (
            <Nav>
                <NavLink to="/login" onClick={Logout} className={({ isActive }) => "btn btn-bd-primary" + (isActive ? activeClassName : "")}>Logout</NavLink>
            </Nav>
        )
    }

    return (
        <Navbar collapseOnSelect expand="md" bg="dark" variant="dark" style={{minHeight: "80px"}}>
            <Container fluid className="px-4">
                <Navbar.Toggle/>
                <Navbar.Collapse className="justify-content-end">
                    <Nav className="me-auto">
                        <NavLink to="/" className={({ isActive }) => "px-4 btn btn-bd-primary" + (isActive ? activeClassName : "")}>Home</NavLink>
                    </Nav>
                    {NavigationMenu}
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}

export default NavHome;