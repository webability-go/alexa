package alexa

import (
  "fmt"

//  "github.com/aws/aws-sdk-go/aws"
//  "github.com/aws/aws-sdk-go/aws/session"
//  "github.com/aws/aws-sdk-go/service/dynamodb"
//  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var ddb_apiVersion = ""
var ddb_region = ""
var ddb_tablename = ""
var ddb_autocreate = false

// Dynamo for alexa params structure
type dynamoalexa struct {
  id string
  attributes string
}

func WithDynamoDbClient(api string, reg string) {
  ddb_apiVersion = api
  ddb_region = reg
}

func WithTableName(name string) {
  ddb_tablename = name
}

func WithAutoCreateTable(auto bool) {
  ddb_autocreate = auto
}

func StartDynamoTable() {
  fmt.Println("Start Dynamo Table: ", ddb_tablename)
  
  // Verify if the table exists on Dynamo
  // No: autocreate = true::: create it
  // No: autocreate = false::: set flag not available
  // Yes: OK
}