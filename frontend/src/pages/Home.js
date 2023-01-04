import React, {useEffect, useState} from "react";
import Dashboard from "../components/Dashboard.js";
import {Carousel, Tabs, Tab} from "react-bootstrap";
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
        console.log("hello")
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
            <Tab eventKey="longer-tab" title="Loooonger Tab">
    
            </Tab>
          </Tabs>

        )
    } else {
        // display here the default home screen if no user is logged in
        return (
          <Carousel>
            <Carousel.Item>
              <img
                className="d-block w-100"
                src="holder.js/800x400?text=First slide&bg=373940"
                alt="First slide"
              />
              <Carousel.Caption>
                <h3>First slide label</h3>
                <p>Nulla vitae elit libero, a pharetra augue mollis interdum.</p>
              </Carousel.Caption>
            </Carousel.Item>
            <Carousel.Item>
              <img
                className="d-block w-100"
                src="holder.js/800x400?text=Second slide&bg=282c34"
                alt="Second slide"
              />

              <Carousel.Caption>
                <h3>Second slide label</h3>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
              </Carousel.Caption>
            </Carousel.Item>
            <Carousel.Item>
              <img
                className="d-block w-100"
                src="holder.js/800x400?text=Third slide&bg=20232a"
                alt="Third slide"
              />

              <Carousel.Caption>
                <h3>Third slide label</h3>
                <p>
                  Praesent commodo cursus magna, vel scelerisque nisl consectetur.
                </p>
              </Carousel.Caption>
            </Carousel.Item>
          </Carousel>
        )
    }
}

export default Home;
