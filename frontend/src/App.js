import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import "./scss/custom.scss";
import Login from "./pages/Login";
import NavHome from "./components/Navbar";
import Home from "./pages/Home";
import Signup from "./pages/Signup";
import axios from "axios";
import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {

  const [accountDetails, setAccountDetails] = useState({"firstname": "", "lastname": "", "email": "", "userid": ""})
  const [userLoggedIn, setUserLoggedIn] = useState(false);

  useEffect(() => {
    const getUser = async () => {
      var url = "http://localhost:5000/api/user"
      await axios.get(url, { 
        withCredentials: true,
      }).then(response => {
        if (response.status === 200 && !userLoggedIn) {
          setAccountDetails({
            firstname: response.data.firstname,
            lastname: response.data.lastname,
            email: response.data.email,
            userid: response.data.userid,
          });
          setUserLoggedIn(true);
        }
      }).catch(response => {
        console.log(response);
        setAccountDetails({firstname: "", lastname: "", email: "", userid: "",})
        setUserLoggedIn(false);
      })
    }

    getUser();
  }, [userLoggedIn]);

  return (
    <div className="App">
      <BrowserRouter>
        <NavHome setAccountDetails={setAccountDetails} setUserLoggedIn={setUserLoggedIn} userLoggedIn={userLoggedIn} />

        <main>
            <Routes>
              <Route path="/" exact element={<Home accountDetails={accountDetails} userLoggedIn={userLoggedIn} />}/>
              <Route path="/login" element={<Login setAccountDetails={setAccountDetails} setUserLoggedIn={setUserLoggedIn} />}/>
              <Route path="/signup" element={<Signup />}/>
            </Routes>


        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;