package handlers

import(
	"github.com/brk-a/LearnGo/0x01_go_aws_serverless/pkg/user"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed, sorry"

type ErrorBody struct{
	ErrorMessage *string `json: "error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
){
	email := req.QueryStringParameters["email"]
	
	// query param has 1 email
	if len(email)>0{
		res, err := user.FetchUser(email, tableName, dynaClient)
		if err!=nil{
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, res)
	}

	//query param has more than 1 email
	res, err := user.FetchUsers(tableName, dynaClient)
	if err!nil{
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, res)

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error
){

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error
){

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error
){

}

func UnhandledMethod(*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrErrorMethodNotAllowed)
}
