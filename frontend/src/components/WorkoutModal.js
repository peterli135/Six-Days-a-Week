import React, {useState, useEffect} from "react";
import axios from "axios";
import {Modal, Table} from "react-bootstrap";

const WorkoutModal = ({workoutModal, setWorkoutModal}) => {

    const [exerciseData, setExerciseData] = useState([])

    useEffect(() => {
        const getExercisesInWorkout = async () => {
        var url = "http://localhost:5000/api/user/workout/" + workoutModal.workoutData._id;
        await axios.get(url, { 
            withCredentials: true,
        }).then(response => {
            console.log(response);
            if (response.status === 200) {
                setExerciseData(response.data);
            }
        }).catch(response => {
            console.log(response);
            setExerciseData([]);
        })
        }
        getExercisesInWorkout();
    }, []);

    return (
        <Modal show={workoutModal} onHide={() => setWorkoutModal({state: false, workoutData: {}})}
            dialogClassName="modal-60w"
            aria-labelledby="contained-modal-title-vcenter"
            centered
        >
            <Modal.Header closeButton>
                <Modal.Title>Workout Name: {workoutModal.workoutData.name} - {workoutModal.workoutData.date}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>Exercise Name</th>
                            <th>Weight</th>
                            <th>Sets</th>
                            <th>Reps</th>
                        </tr>
                    </thead>
                    <tbody>
                    {Array.isArray(exerciseData) && exerciseData.map((exercise, index) => (
                        <tr key={index}>
                            <td>{exercise.name}</td>
                            <td>{exercise.weight}</td>
                            <td>{exercise.sets}</td>
                            <td>{exercise.reps}</td>
                        </tr>
                    ))}
                    {/*<tr>
                        <td colSpan={3}>
                        <Button variant="outline-primary" onClick={() => setAddWorkoutModal(true)} style={{width: "100%"}}>Add a Workout</Button>
                        </td>
                    </tr>*/}
                    </tbody>
                </Table>
            </Modal.Body>
        </Modal>
    )
}

export default WorkoutModal;