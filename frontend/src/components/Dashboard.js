import React, {useState} from "react";
import axios from "axios";
import {Container, Button, Table} from "react-bootstrap";
import AddWorkoutModal from "./AddWorkoutModal";
import {Navigate} from "react-router-dom";

const Dashboard = ({workoutList}) => {
    const [addWorkoutModal, setAddWorkoutModal] = useState(false);

    // have a useState for what the filter is: filter will have three options: current year, showing a list of the months and the workouts that were added for those months
    // on default, it will open to the current month. onClick, expand them to show all the workouts that were created for that month.
    // another filter will be to show other years, can choose 2022, 2021, 2020, etc. and then another filter will be to show all workouts that have been created
    // have the onClick expand be similar to that of the design of the NeetCode website, and can put the filters on the top maybe as well.

    if (workoutList.length > 0) {
        return (
            <div>
                <Container fluid className="d-flex align-items-center justify-content-center p-4">
                    <Button variant="outline-primary" onClick={() => setAddWorkoutModal(true)} style={{width: "80%"}}>Add a Workout</Button>
                </Container>
                <Container fluid className="d-flex align-items-center justify-content-center px-4" style={{width: "80%"}}>
                    <Table striped>
                        <thead>
                            <tr>
                                <th>Date</th>
                                <th>Name</th>
                                <th>Number of Exercises</th>
                            </tr>
                        </thead>
                        <tbody>
                        {Array.isArray(workoutList) && workoutList.map((workout, index) => (
                            <tr key={index}>
                                <td>{workout.date}</td>
                                <td>{workout.name}</td>
                                <td>{workout.exercises.length}</td>
                            </tr>
                        ))}
                        </tbody>
                    </Table>
                </Container>

                {addWorkoutModal && <AddWorkoutModal addWorkoutModal={addWorkoutModal} setAddWorkoutModal={setAddWorkoutModal}/>}
            </div>
        )
    } else {
        return (
            <div>
                <Container fluid className="d-flex align-items-center justify-content-center">
                    <Button variant="outline-primary" onClick={() => setAddWorkoutModal(true)}>Add a Workout</Button>
                </Container>
            </div>
        )
    }
}

export default Dashboard;