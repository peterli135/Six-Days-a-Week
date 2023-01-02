package database

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = OpenCollection(Client, "user")

// function that gets number of accounts associated with the email. should return 0 if there are no accounts using that email.
func GetUserCountByEmail(email string) (int64, error) {
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": email})
	if err != nil {
		return count, err
	}
	return count, nil
}

// function that gets the user associated with the email.
func GetUserByEmail(ctx context.Context, email *string) (models.User, error) {
	var userFound models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&userFound)

	if err != nil {
		return userFound, err
	}

	return userFound, err
}

// function that finds user by their userid (claims.Issuer)
func GetUserByUserID(issueClaims string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"userid": issueClaims}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func InsertUser(User *models.User) (interface{}, error) {
	user, err := userCollection.InsertOne(context.TODO(), User)
	if err != nil {
		return nil, err
	}
	return user.InsertedID, err
}

// UpdateAllTokens function renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, userID string) (models.User, error) {
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refreshtoken", Value: signedRefreshToken})

	upsert := true
	filter := bson.M{"userid": userID}
	opt := options.FindOneAndUpdateOptions{
		Upsert: &upsert,
	}
	opt.SetReturnDocument(options.After)
	var updatedUser models.User

	/*_, err := userCollection.UpdateOne(
		context.TODO(),
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)*/
	err := userCollection.FindOneAndUpdate(
		context.TODO(),
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	).Decode(&updatedUser)

	return updatedUser, err

}
