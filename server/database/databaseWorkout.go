package database

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var exerciseCollection *mongo.Collection = OpenCollection(Client, "exercise")
var workoutDateCollection *mongo.Collection = OpenCollection(Client, "workoutdate")

func GetExerciseByName(exerciseName *string, weight *int, sets *int, reps *int) (models.Exercise, error) {
	var exerciseFound models.Exercise

	err := exerciseCollection.FindOne(context.TODO(), bson.M{"name": exerciseName, "weight": weight, "sets": sets, "reps": reps}).Decode(&exerciseFound)
	if err != nil {
		return exerciseFound, err
	}

	return exerciseFound, err
}

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

func InsertExercise(Exercise *models.Exercise) (interface{}, error) {
	exercise, err := exerciseCollection.InsertOne(context.TODO(), Exercise)
	if err != nil {
		return nil, err
	}

	return exercise.InsertedID, err
}

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

func InsertWorkoutDate(WorkoutDate *models.WorkoutDate) (interface{}, error) {
	workoutDate, err := workoutDateCollection.InsertOne(context.TODO(), WorkoutDate)
	if err != nil {
		return nil, err
	}

	return workoutDate.InsertedID, err
}

func GetWorkoutsByUser(userID primitive.ObjectID) ([]primitive.M, error) {
	var workouts []bson.M

	cursor, err := workoutDateCollection.Find(context.TODO(), bson.M{"UserID": userID})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &workouts); err != nil {
		return nil, err
	}

	return workouts, err
}

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
