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
  id := req.GetSessionId()
  dynamodb.Upsert(id, data)
  return nil
}

func LoadPersistentAttributes(req request.AlexaRequest) (interface{}, error) {
  return nil, nil
}
