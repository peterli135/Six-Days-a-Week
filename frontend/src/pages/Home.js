import React, {useEffect, useState} from "react";
import Dashboard from "../components/Dashboard.js";
import axios from "axios";

const Home = ({accountDetails, userLoggedIn}) => {

    // TODO:
    // what should home page display? a log of your workouts that can be displayed in either list form or calendar form... for now, can just display list form
        // when clicking on a workout date, display a modal that shows a list of your exercises you did for that day.
    // have a modal that allows you to add a workout on the home page
    // then have another page that displays your progress... this could be in list form and graph form which displays how your strength is progressing for different exercises

    const [workouts, setWorkouts] = useState([]);

    useEffect(() => {
        const getUserWorkouts = async () => {
          var url = "http://localhost:5000/api/user/workouts"
          await axios.get(url, { 
            withCredentials: true,
          }).then(response => {
            console.log(response);
            if (response.status === 200) {
              setWorkouts(response.data);
            }
          }).catch(response => {
            console.log(response);
            setWorkouts([]);
          })
        }
        
        getUserWorkouts();
    }, []);

    if (userLoggedIn) {
        return (
            <Dashboard workoutList={workouts}/>
        )
    } else {
        // display here the default home screen if no user is logged in
        return (
            <div>No user logged in...</div>
        )
    }
}

export default Home;
