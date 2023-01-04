import React, {useState} from "react";
import {Container, Button, Table, Accordion} from "react-bootstrap";
import AddWorkoutModal from "./AddWorkoutModal";

const Dashboard = ({workoutList, filter}) => {
    const [addWorkoutModal, setAddWorkoutModal] = useState(false);
    const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

    // have a useState for what the filter is: filter will have three options: current year, showing a list of the months and the workouts that were added for those months
    // on default, it will open to the current month. onClick, expand them to show all the workouts that were created for that month.
    // another filter will be to show other years, can choose 2022, 2021, 2020, etc. and then another filter will be to show all workouts that have been created
    // have the onClick expand be similar to that of the design of the NeetCode website, and can put the filters on the top maybe as well.

    let filterMenu;
    if (filter === "current-year") {
        filterMenu = (
            <div>
                <Container fluid className="d-flex align-items-center justify-content-center px-4" style={{width: "80%"}}>
                    <Accordion defaultActiveKey={(new Date().getMonth())} alwaysOpen style={{width: "100%"}}>
                        {Array.isArray(monthNames) && monthNames.map((monthName, monthIndex) => (
                            <Accordion.Item eventKey={monthIndex} key={monthIndex}>
                                <Accordion.Header>{monthName}</Accordion.Header>
                                <Accordion.Body bsPrefix="accordion-custom-body">
                                    <Table striped bordered hover>
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
                                </Accordion.Body>
                            </Accordion.Item>
                        ))}
                    </Accordion>
                    {/*<Table striped bordered hover>
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
                    </Table>*/}
                </Container>

                {addWorkoutModal && <AddWorkoutModal addWorkoutModal={addWorkoutModal} setAddWorkoutModal={setAddWorkoutModal}/>}
            </div>
        )
    } else if (filter === "all-workouts") {
        filterMenu = (
            <div>
                <Container fluid className="d-flex align-items-center justify-content-center px-4" style={{width: "80%"}}>
                    <Table striped bordered hover>
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
    }

    return (
        <div>
            <Container fluid className="d-flex align-items-center justify-content-center p-4">
                <Button variant="outline-primary" onClick={() => setAddWorkoutModal(true)} style={{width: "80%"}}>Add a Workout</Button>
            </Container>
            {filterMenu}
        </div>
    )
}

export default Dashboard;