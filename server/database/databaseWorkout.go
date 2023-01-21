package database

import (
	"context"
	"fmt"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var exerciseCollection *mongo.Collection = OpenCollection(Client, "exercise")
var workoutDateCollection *mongo.Collection = OpenCollection(Client, "workoutdate")

// function to get exercise from the name
func GetExerciseByName(exerciseName *string, weight *int, sets *int, reps *int) (models.Exercise, error) {
	var exerciseFound models.Exercise

	err := exerciseCollection.FindOne(context.TODO(), bson.M{"name": exerciseName, "weight": weight, "sets": sets, "reps": reps}).Decode(&exerciseFound)
	if err != nil {
		return exerciseFound, err
	}

	return exerciseFound, err
}

// function to get exercise by its id
func GetExerciseByID(exerciseID string) (models.Exercise, error) {
	var exerciseFound models.Exercise

	primitiveExerciseID, pErr := primitive.ObjectIDFromHex(exerciseID)
	if pErr != nil {
		return exerciseFound, pErr
	}

	err := exerciseCollection.FindOne(context.TODO(), bson.M{"_id": primitiveExerciseID}).Decode(&exerciseFound)
	if err != nil {
		return exerciseFound, err
	}

	return exerciseFound, err
}

// function to insert exercise into database
func InsertExercise(Exercise *models.Exercise) (interface{}, error) {
	exercise, err := exerciseCollection.InsertOne(context.TODO(), Exercise)
	if err != nil {
		return nil, err
	}

	return exercise.InsertedID, err
}

// function to insert multiple exercise into database
func InsertExercises(Exercises []models.Exercise) (interface{}, error) {
	var exerciseIDs []interface{}
	for _, exercise := range Exercises {
		exerciseFound, err := GetExerciseByName(exercise.Name, exercise.Weight, exercise.Sets, exercise.Reps)
		if err != nil {
			exercise.ID = primitive.NewObjectID()
			insertExercise, iErr := InsertExercise(&exercise)
			if iErr != nil {
				return exerciseIDs, iErr
			}
			exerciseIDs = append(exerciseIDs, insertExercise.(primitive.ObjectID))
		} else {
			exerciseIDs = append(exerciseIDs, exerciseFound.ID)
		}
	}
	return exerciseIDs, nil
}

// function to edit an exercise
func UpdateExercise(ExerciseID string, newExerciseData models.UpdateExercise) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(ExerciseID)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: newExerciseData}}

	result, err := exerciseCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

// function to insert a workout into database
func InsertWorkoutDate(WorkoutDate *models.WorkoutDate) (interface{}, error) {
	workoutDate, err := workoutDateCollection.InsertOne(context.TODO(), WorkoutDate)
	if err != nil {
		return nil, err
	}

	return workoutDate.InsertedID, err
}

// function to get workouts by a user's id
func GetWorkoutsByUser(userID primitive.ObjectID) ([]primitive.M, error) {
	cursor, err := workoutDateCollection.Find(context.TODO(), bson.M{"UserID": userID})
	if err != nil {
		return nil, err
	}

	var workouts []bson.M
	if err = cursor.All(context.TODO(), &workouts); err != nil {
		return nil, err
	}

	return workouts, err
}

// function to get a user's workouts in the current year
func GetCurrentYearWorkoutsByUser(userID primitive.ObjectID, currentYear string) ([]primitive.M, error) {
	dateIndex := mongo.IndexModel{Keys: bson.D{{Key: "date", Value: "text"}}}
	_, err := workoutDateCollection.Indexes().CreateOne(context.TODO(), dateIndex)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "UserID", Value: userID}},
				bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: currentYear}}}},
			},
		},
	}
	cursor, err := workoutDateCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var workouts []bson.M
	if err = cursor.All(context.TODO(), &workouts); err != nil {
		return nil, err
	}

	return workouts, err
}

// function to get a workout by its id
func GetWorkoutByID(workoutID string) (models.WorkoutDate, error) {
	var workoutFound models.WorkoutDate

	primitiveWorkoutID, pErr := primitive.ObjectIDFromHex(workoutID)
	if pErr != nil {
		return workoutFound, pErr
	}

	err := workoutDateCollection.FindOne(context.TODO(), bson.M{"_id": primitiveWorkoutID}).Decode(&workoutFound)
	if err != nil {
		return workoutFound, err
	}

	return workoutFound, err
}

// function to get all the exercises in a workout by the workoutID
func GetExercisesByWorkoutID(workout models.WorkoutDate) ([]models.Exercise, error) {
	var exercises []models.Exercise

	for _, exerciseID := range workout.Exercises {
		exerciseData, err := GetExerciseByID(exerciseID.Hex())
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, exerciseData)
	}

	return exercises, nil
}

// function to edit a workout (just update values for name/date and add/remove exercises)
func UpdateWorkout(WorkoutID string, newWorkoutData models.UpdateWorkoutDate) (interface{}, error) {
	// what to account for when updating workout?

	// 1: if adding/removing exercise,
	// first create/delete the exercise -- pipeline: first send API request to create/delete workout
	// then add/remove id to the workoutDate struct -- then send request to add/remove the id in the workoutDate
	// then update it using this function -- then update the workoutData after with this function...
	// 2: if updating name/date value, doesn't matter, just have those values in the struct
	id, _ := primitive.ObjectIDFromHex(WorkoutID)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: newWorkoutData}}

	result, err := exerciseCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
