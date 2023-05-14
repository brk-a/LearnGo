package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToFetchRecord = "failed to fetch record"
)
var(
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		//create query
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	//fetch from dynamodb
	res, err := dynaClient.GetItem(input)
	if err!nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	//record found. umarshal it (translate it to JSON for front-end)
	item := new(User)
	err = dynamodbattribute.UnmarshalMap(res.Item, item)
	if err!=nil{
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{
		//create query
		TableName: aws.String(tableName),
	}

	//fetch from dynamodb
	res, err := dynaClient.Scan(input)
	if err!=nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	//records found. unmarshal them
	item := new([]User)
	err = dynamodbattribute.UnmarshalMap(res.Items[], item)
	return item, nil
}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() error {

}
