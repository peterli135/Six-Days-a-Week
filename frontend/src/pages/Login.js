import React, {useState} from "react";
import axios from "axios";
import {Container, Button, Form, FloatingLabel} from "react-bootstrap";
import {Navigate} from "react-router-dom";

const Login = ({setAccountDetails, setUserLoggedIn}) => {
    const [accountLogin, setAccountLogin] = useState({"email": "", "password": ""})
    const [redirect, setRedirect] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()

        var url = "http://localhost:5000/api/login"
        await axios.post(url, {
            "email": accountLogin.email,
            "password": accountLogin.password,
        }, { 
            withCredentials: true,
        }).then(response => {
            if (response.status === 200) {
                console.log(response);
                setAccountDetails({
                    firstname: response.data.object.firstname,
                    lastname: response.data.object.lastname,
                    email: response.data.object.email,
                    userid: response.data.object.userid,
                });
                setUserLoggedIn(true);
                setRedirect(true);
            }
        })
    }

    return (
        <div>
            <Container fluid className="form-signin d-flex align-items-center justify-content-center">
                <Form onSubmit={handleSubmit} style={{minWidth: "400px"}}>
                    <h1 className="h3 mb-4 fw-normal">Please Sign In</h1>

                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <FloatingLabel controlId="floatingInput" label="Email Address" className="mb-3 form-floating-label">
                            <Form.Control className="form-label" type="email" placeholder="name@example.com" 
                                onChange = {e => setAccountLogin({...accountLogin, email: e.target.value})}
                            />
                        </FloatingLabel>
                    </Form.Group>

                    <Form.Group className="mb-4" controlId="formBasicPassword">
                        <FloatingLabel controlId="floatingPassword" label="Password" className="mb-3 form-floating-label">
                            <Form.Control className="form-label" type="password" placeholder="Password" 
                                onChange = {e => setAccountLogin({...accountLogin, password: e.target.value})}
                            />
                        </FloatingLabel>
                    </Form.Group>

                    <div className="d-grid gap-2">
                        <Button variant="primary" size="lg" type="submit">Sign In</Button>
                    </div>

                </Form>
            </Container>

            {redirect && <Navigate to="/" state={redirect} replace={true} /> }
        </div>
    )
}

export default Login;