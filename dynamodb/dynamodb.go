package dynamodb

import (
//  "fmt"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var DDB_apiVersion = ""
var DDB_region = ""
var DDB_tablename = ""
var DDB_autocreate = false
var DDB_session *session.Session
var DDB_client *dynamodb.DynamoDB

// Dynamo for alexa params structure
type dynamoalexa struct {
  id string
  attributes interface{}
}

func OpenSession() {
  DDB_session = session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
    Config: aws.Config{Region: aws.String(DDB_region)},
      }))
  // Create DynamoDB client
  DDB_client = dynamodb.New(DDB_session)
}

func CreateTable() error {
  input := &dynamodb.CreateTableInput{
    AttributeDefinitions: []*dynamodb.AttributeDefinition{
        {
            AttributeName: aws.String("id"),
            AttributeType: aws.String("S"),
        },
        {
            AttributeName: aws.String("attributes"),
            AttributeType: aws.String("S"),
        },
    },
    KeySchema: []*dynamodb.KeySchemaElement{
        {
            AttributeName: aws.String("id"),
            KeyType:       aws.String("HASH"),
        },
        {
            AttributeName: aws.String("attributes"),
            KeyType:       aws.String("RANGE"),
        },
    },
    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits:  aws.Int64(5),
        WriteCapacityUnits: aws.Int64(5),
    },
    TableName: aws.String(DDB_tablename),
  }

  _, err := DDB_client.CreateTable(input)
  if err != nil {
    return err
  }
  return nil
}

func VerifyTable() error {
  // Get the list of tables
  result, err := DDB_client.ListTables(&dynamodb.ListTablesInput{})
  if err != nil {
    return err
  }
  found := false
  for _, n := range result.TableNames {
    if *n == DDB_tablename {
      found = true
    }
  }
  if !found {
    // creates table
    err = CreateTable()
    if err != nil {
      return err
    }
  }
  return nil
}

func StartDynamoTable() error {
  if DDB_client == nil {
    OpenSession()
  }
  // Verify if the table exists on Dynamo
  // No: autocreate = false::: set flag not available, nothing to do
  if DDB_autocreate {
    // No: autocreate = true::: create it
    err := VerifyTable()
    if err != nil {
      return err
    }
  }
  return nil
}

func Select(id string) (interface{}, error) {
  
  da := &dynamoalexa{id:  id,}

  key, err := dynamodbattribute.MarshalMap(da)
  if err != nil {
    return nil, err
  }

  input := &dynamodb.GetItemInput{
    Key:       key,
    TableName: aws.String(DDB_tablename),
  }

  result, err := DDB_client.GetItem(input)
  if err != nil {
    return nil, err
  }

  attributes := &dynamoalexa{}
  err = dynamodbattribute.UnmarshalMap(result.Item, attributes)
  if err != nil {
    return nil, err
  }

  return attributes.attributes, nil
}

func Upsert(id string, data interface{}) error {
  
  da := &dynamoalexa{ id: id, attributes: data,}
  
  av, err := dynamodbattribute.MarshalMap(da)
  if err != nil {
    return err
  }

  input := &dynamodb.PutItemInput{
    Item:      av,
    TableName: aws.String(DDB_tablename),
  }

  _, err = DDB_client.PutItem(input)
  if err != nil {
    return err
  }
  return nil
}

