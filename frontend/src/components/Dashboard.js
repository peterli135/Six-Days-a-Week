import React, {useState} from "react";
import {Container, Button, Table, Accordion} from "react-bootstrap";
import AddWorkoutModal from "./AddWorkoutModal";
import WorkoutModal from "./WorkoutModal";

const Dashboard = ({workoutList, filter}) => {
    const [addWorkoutModal, setAddWorkoutModal] = useState(false);
    const [workoutModal, setWorkoutModal] = useState({state: false, workoutData: {}});

    var monthListWorkouts = [{month: "1", data: []}, {month: "2", data: []}, {month: "3", data: []}, {month: "4", data: []}, {month: "5", data: []}, {month: "6", data: []},
                             {month: "7", data: []}, {month: "8", data: []}, {month: "9", data: []}, {month: "10", data: []}, {month: "11", data: []} , {month: "12", data: []}];

    // have a useState for what the filter is: filter will have three options: current year, showing a list of the months and the workouts that were added for those months
    // on default, it will open to the current month. onClick, expand them to show all the workouts that were created for that month.
    // another filter will be to show other years, can choose 2022, 2021, 2020, etc. and then another filter will be to show all workouts that have been created

    // puts workouts in the month that the workout was created in
    workoutList && workoutList.forEach((workout) => {
        var month = workout.date.split("/")[0]
        for (let j = 0; j < monthListWorkouts.length; j++) {
            if (monthListWorkouts[j].month === month) {
                monthListWorkouts[j].data.push(workout);
            }
        }
    })

    console.log(workoutModal);

    let filterMenu;
    if (filter === "current-year") {
        filterMenu = (
            <div>
                <Container fluid className="d-flex align-items-center justify-content-center px-4" style={{width: "80%"}}>
                    <Accordion defaultActiveKey={(new Date().getMonth())} alwaysOpen style={{width: "100%"}}>
                        {Array.isArray(monthListWorkouts) && monthListWorkouts.map((monthWorkouts, monthIndex) => (
                            <Accordion.Item eventKey={monthIndex} key={monthIndex}>
                                <Accordion.Header>{Intl.DateTimeFormat("en", { month: "long" }).format(new Date(monthWorkouts.month))}</Accordion.Header>
                                <Accordion.Body bsPrefix="accordion-custom-body">
                                    {(monthWorkouts.data.length > 0) 
                                        ? 
                                        <Table striped bordered hover>
                                            <thead>
                                                <tr>
                                                    <th>Date</th>
                                                    <th>Name</th>
                                                    <th>Number of Exercises</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                            {Array.isArray(monthWorkouts.data) && monthWorkouts.data.map((workout, index) => (
                                                <tr key={index} onClick={() => setWorkoutModal({...workoutModal, state: true, workoutData: workout})} style={{cursor: "pointer"}}>
                                                    <td>{workout.date}</td>
                                                    <td>{workout.name}</td>
                                                    <td>{workout.exercises.length}</td>
                                                </tr>
                                            ))}
                                            {/*<tr>
                                                <td colSpan={3}>
                                                <Button variant="outline-primary" onClick={() => setAddWorkoutModal(true)} style={{width: "100%"}}>Add a Workout</Button>
                                                </td>
                                            </tr>*/}
                                            </tbody>
                                        </Table>
                                        : 
                                        <Table striped bordered hover>
                                            <thead>
                                                <tr>
                                                    <th>No workouts created for this month.</th>
                                                </tr>
                                            </thead>
                                        </Table>
                                    }
                                </Accordion.Body>
                            </Accordion.Item>
                        ))}
                    </Accordion>
                </Container>

                {addWorkoutModal && <AddWorkoutModal addWorkoutModal={addWorkoutModal} setAddWorkoutModal={setAddWorkoutModal} />}
                {workoutModal.state && <WorkoutModal workoutModal={workoutModal} setWorkoutModal={setWorkoutModal} />}
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