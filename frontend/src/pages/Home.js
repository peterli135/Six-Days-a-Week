import React, {useEffect, useState} from "react";
import Dashboard from "../components/Dashboard.js";
import {Container, Carousel, Tabs, Tab} from "react-bootstrap";
import axios from "axios";

const Home = ({accountDetails, userLoggedIn}) => {

    // TODO:
    // what should home page display? a log of your workouts that can be displayed in either list form or calendar form... for now, can just display list form
        // when clicking on a workout date, display a modal that shows a list of your exercises you did for that day.
    // have a modal that allows you to add a workout on the home page
    // then have another page that displays your progress... this could be in list form and graph form which displays how your strength is progressing for different exercises

    const [filter, setFilter] = useState("current-year");
    const [workouts, setWorkouts] = useState([]);

    useEffect(() => {
      if (filter === "current-year") {
        const getCurrentYearUserWorkouts = async () => {
          var url = "http://localhost:5000/api/user/workouts/currentyear"
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
        getCurrentYearUserWorkouts();
      } else if (filter === "all-workouts") {
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
      }
    }, [filter]);
    console.log(filter);

    if (userLoggedIn) {
        return (
          <Tabs
            id="fill-tab-example"
            activeKey={filter}
            onSelect={(k) => setFilter(k)}
            className="mt-3"
            fill
          >
            <Tab eventKey="all-workouts" title="All Workouts">
              <Dashboard workoutList={workouts} filter={filter}/>
            </Tab>
            <Tab eventKey="current-year" title="Current Year">
              <Dashboard workoutList={workouts} filter={filter}/>
            </Tab>
            <Tab eventKey="longer-tab" title="Graph View">
    
            </Tab>
          </Tabs>

        )
    } else {
        // display here the default home screen if no user is logged in
        return (
          <Container fluid className="d-flex align-items-center justify-content-center px-4 pt-4" style={{width: "80%"}}>
            <Carousel>
              <Carousel.Item>
                <img
                  className="d-block w-100"
                  src="/home-view-current-year.png"
                  alt="First slide"
                />
                <Carousel.Caption>
                  <h3>View your current year's created workouts</h3>
                </Carousel.Caption>
              </Carousel.Item>
              <Carousel.Item>
                <img
                  className="d-block w-100"
                  src="/home-all-view.png"
                  alt="Second slide"
                />
                <Carousel.Caption>
                  <h3>View all your created workouts</h3>
                </Carousel.Caption>
              </Carousel.Item>
              <Carousel.Item>
                <img
                  className="d-block w-100"
                  src="/add-workout-modal.png"
                  alt="Third slide"
                />
                <Carousel.Caption>
                  <h3>Log your workouts</h3>
                </Carousel.Caption>
              </Carousel.Item>
              <Carousel.Item>
                <img
                  className="d-block w-100"
                  src="/workout-exercise-list-view.png"
                  alt="Fourth slide"
                />
                <Carousel.Caption>
                  <h3>View your exercises in your workouts</h3>
                </Carousel.Caption>
              </Carousel.Item>
            </Carousel>
          </Container>
        )
    }
}

export default Home;
