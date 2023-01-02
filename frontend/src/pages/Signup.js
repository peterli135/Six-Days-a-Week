import React, {useState} from "react";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.css";
import {Container, Button, Form, FloatingLabel} from "react-bootstrap";
import {Link, Navigate} from "react-router-dom";

const Signup = () => {

    const [newAccount, setNewAccount ]= useState({"firstname": "", "lastname": "", "email": "", "password": "", "userid": ""})
    const [redirect, setRedirect] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()

        var url = "http://localhost:5000/api/signup"
        await axios.post(url, {
            "firstname": newAccount.firstname,
            "lastname": newAccount.lastname,
            "email": newAccount.email,
            "password": newAccount.password,
            "userid": newAccount.userid,
        }).then(response => {
            if (response.status === 200) {
                console.log(response)
                setRedirect(true)
            }
        })
    }

    return (
        <div>
            <Container fluid className="form-signin d-flex align-items-center justify-content-center">
                <Form onSubmit={handleSubmit} style={{minWidth: "400px"}}>

                    <h1 className="h3 mb-4 fw-normal">Create an account or <Link to="/login" className="h3 mb-3 fw-normal" style={{textDecoration: "none"}}>log in</Link></h1>

                    <Form.Group className="mb-3" controlId="formBasicFirstName">
                        <FloatingLabel controlId="floatingInput" label="First Name" className="mb-3 form-floating-label">
                            <Form.Control className="form-label" placeholder="First Name" 
                                onChange={e => setNewAccount({...newAccount, firstname: e.target.value})}
                            />
                        </FloatingLabel>
                    </Form.Group>

                    <Form.Group className="mb-3" controlId="formBasicLastName">
                        <FloatingLabel controlId="floatingInput" label="Last Name" className="mb-3 form-floating-label">
                            <Form.Control className="form-label" placeholder="Last Name" 
                                onChange={e => setNewAccount({...newAccount, lastname: e.target.value})}
                            />
                        </FloatingLabel>
                    </Form.Group>

                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <FloatingLabel controlId="floatingInput" label="Email Address" className="mb-3 form-floating-label">
                            <Form.Control className="form-label" type="email" placeholder="name@example.com" 
                                onChange={e => setNewAccount({...newAccount, email: e.target.value})}
                            />
                        </FloatingLabel>
                    </Form.Group>

                    <Form.Group className="mb-4" controlId="formBasicPassword">
                        <FloatingLabel controlId="floatingPassword" label="Password" className="mb-3 form-floating-label">
                            <Form.Control className="form-label" type="password" placeholder="Password" 
                                onChange={e => setNewAccount({...newAccount, password: e.target.value})}
                            />
                        </FloatingLabel>
                    </Form.Group>

                    <div className="d-grid gap-2">
                        <Button variant="primary" size="lg" type="submit">Sign Up</Button>
                    </div>
                </Form>
            </Container>
            
            {redirect && <Navigate to="/login" state={redirect} replace={true} /> }
        </div>
    )

}

export default Signup;