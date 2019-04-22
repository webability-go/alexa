package alexa

import (
  "github.com/webability-go/alexa/request"
  "github.com/webability-go/alexa/dynamodb"
)

func WithDynamoDbClient(api string, reg string) {
  dynamodb.DDB_apiVersion = api
  dynamodb.DDB_region = reg
}

func WithTableName(name string) {
  dynamodb.DDB_tablename = name
}

func WithAutoCreateTable(auto bool) {
  dynamodb.DDB_autocreate = auto
}

func SavePersistentAttributes(req request.AlexaRequest, data interface{}) error {
  id := req.GetUserId()
  return dynamodb.Upsert(id, data)
}

func LoadPersistentAttributes(req request.AlexaRequest, container interface{}) error {
  id := req.GetUserId()
  return dynamodb.Select(id, container)
}
