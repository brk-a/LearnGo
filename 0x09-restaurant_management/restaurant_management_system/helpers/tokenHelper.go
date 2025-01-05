package helpers

import (
	"context"
	"errors"
	"log"
	"os"
	"restaurant_management_system/database"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.RegisteredClaims
}

func (sd *SignedDetails) GetAudience() (jwt.ClaimStrings, error) {
	return sd.RegisteredClaims.Audience, nil
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(
	userEmail string,
	userFirstName string,
	userLastName string,
	userId string,
) (
	signedtoken string, signedRefreshToken string, err error,
) {
	claims := &SignedDetails{
		Email:      userEmail,
		First_name: userFirstName,
		Last_name:  userLastName,
		Uid:        userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{},
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Duration(1) * time.Hour)),
			Issuer:    "admin",
			Subject:   "",
			NotBefore: jwt.NewNumericDate(time.Now().Local().Add(time.Duration(1) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
			ID:        "sample_id_string",
		},
	}

	refeshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{},
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Duration(24) * time.Hour)),
			Issuer:    "admin",
			Subject:   "",
			NotBefore: jwt.NewNumericDate(time.Now().Local().Add(time.Duration(1) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
			ID:        "sample_id_string",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", errors.New("error while creating JWT")
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refeshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", errors.New("error while creating JWT")
	}

	return token, refreshToken, nil
}

func UpdateAllTokens(
	signedToken string,
	signedRefreshToken string,
	foundUserId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{
		Key: "$set", Value: bson.D{
			{Key: "updated_at", Value: Updated_at},
			{Key: "token", Value: signedToken},
			{Key: "refresh_token", Value: signedRefreshToken},
		},
	})

	upsert := true
	filter := bson.M{"user_id": foundUserId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		updateObj,
		&opt,
	)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}

	defer cancel()
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err!=nil{
		log.Fatal(err)
		msg = "error parsing token"
		return nil, msg
	}

	//is token invalid?
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Invalid token"
		return nil, msg
	}

	//is token expired?
	if claims.ExpiresAt.Time.Before(time.Now().Local()) {
		msg = "Token expired"
		return nil, msg
	}

	return claims, ""
}
