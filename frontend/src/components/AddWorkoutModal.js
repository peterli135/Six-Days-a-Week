import React, {useState} from "react";
import axios from "axios";
import {Modal, Button, Form, FloatingLabel, InputGroup, CloseButton} from "react-bootstrap";

const AddWorkoutModal = ({addWorkoutModal, setAddWorkoutModal}) => {

    const [formValues, setFormValues] = useState([{"name": "", "weight": Number, "sets": Number, "reps": Number}])
    const [workout, setWorkout] = useState({"name": "", "date": ""})

    const handleChange = (index, e) => {
        let newFormValues = [...formValues];
        newFormValues[index][e.target.name] = e.target.value;
        if (e.target.name !== "name") {
            newFormValues[index][e.target.name] = e.target.valueAsNumber;
        }
        setFormValues(newFormValues);
    }

    const addExerciseField = () => {
        setFormValues([...formValues, {"name": "", "weight": Number, "sets": Number, "reps": Number}])
    }

    const removeExerciseField = (index) => {
        let newFormValues = [...formValues];
        newFormValues.splice(index, 1);
        setFormValues(newFormValues)
    }

    const handleSubmit = async (e) => {
        e.preventDefault()

        var url = "http://localhost:5000/api/exercises"
        await axios.post(url, formValues, { 
            withCredentials: true,
        }).then(response => {
            console.log(response);
            if (response.status === 200) {
                handleSubmitWorkout(e, response.data);
            }
        })
    }

    const handleSubmitWorkout = async (e, exerciseIDs) => {
        e.preventDefault()

        var url = "http://localhost:5000/api/workoutdate"
        await axios.post(url, {
            "name": workout.name,
            "date": formatDateMMDDYY(workout.date),
            "exercises": exerciseIDs
        }, { 
            withCredentials: true,
        }).then(response => {
            console.log(response);
            setAddWorkoutModal(false);
        })
    }

    // Formats Date to MM/DD/YYYY
    const formatDateMMDDYY = (date) => {
        const dateObj = new Date(date + "T00:00:00");
        return new Intl.DateTimeFormat("en-US").format(dateObj);
    }

    return (
        <Modal show={addWorkoutModal} onHide={() => setAddWorkoutModal(false)}
            dialogClassName="modal-60w"
            aria-labelledby="contained-modal-title-vcenter"
            centered
        >
            <Form onSubmit={handleSubmit} style={{minWidth: "350px"}}>
                <Modal.Header closeButton>
                    <Modal.Title>Add a Workout</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <InputGroup className="mb-3">
                    <InputGroup.Text>Workout Split Name and Date</InputGroup.Text>
                    <FloatingLabel controlId="floatingInput" label="Workout Name" className="form-floating-label">
                        <Form.Control className="form-modal" placeholder="Workout Name"
                            onChange = {e => setWorkout({...workout, name: e.target.value})}
                        />
                    </FloatingLabel>
                    <FloatingLabel controlId="floatingInput" label="Date" className="form-floating-label">
                        <Form.Control className="form-modal" type="date" placeholder="Date"
                            onChange = {e => setWorkout({...workout, date: e.target.value})}
                        />
                    </FloatingLabel>
                    </InputGroup>
                    <h1 className="h5 fw-normal">Add your exercises here:</h1>
                    {formValues.map((element, index) => (
                        <InputGroup className="mb-2" key={index}>
                            <InputGroup.Text>Exercise {index + 1}</InputGroup.Text>
                            <FloatingLabel controlId="floatingInput" label="Name" className="form-floating-label">
                                <Form.Control className="form-modal" placeholder="Name" name="name"
                                    value={element.name || ""} onChange={e => handleChange(index, e)}
                                />
                            </FloatingLabel>
                            <FloatingLabel controlId="floatingInput" label="Weight" className="form-floating-label">
                                <Form.Control className="form-modal" type="number" step="5" placeholder="Weight" name="weight"
                                    onChange={e => handleChange(index, e)}
                                />
                            </FloatingLabel>
                            <FloatingLabel controlId="floatingInput" label="Sets" className="form-floating-label">
                                <Form.Control className="form-modal" type="number" placeholder="Sets" name="sets"
                                    onChange={e => handleChange(index, e)}
                                />
                            </FloatingLabel>
                            <FloatingLabel controlId="floatingInput" label="Reps" className="form-floating-label">
                                <Form.Control className="form-modal" type="number" placeholder="Reps" name="reps"
                                    onChange={e => handleChange(index, e)}
                                />
                            </FloatingLabel>
                            <InputGroup.Text><CloseButton onClick={() => removeExerciseField(index)}/></InputGroup.Text>
                        </InputGroup>
                    ))}
                    <div className="d-grid gap-2 mt-3">
                        <Button variant="primary" size="lg" onClick={() => addExerciseField()}>Click here to add another exercise</Button>
                    </div>
                </Modal.Body>
                <Modal.Footer className="d-grid gap-2">
                    <Button variant="primary" size="lg" type="submit">
                        Add Workout
                    </Button>
                </Modal.Footer>
            </Form>
        </Modal>
    )
}

export default AddWorkoutModal;