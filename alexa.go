package alexa

import (
  "github.com/aws/aws-lambda-go/lambda"
)

var DEVEL = false

/*
  Alexa Lamdba Library Manager for GO
  This is an implementation for a Full functional lamdba function for Alexa. 
  Build the skill overloading the default event handlers and adding your own event handlers
  
  Log:
  2019-03-27: Phil, Creation
  2019-04-01: Phil, Add Dynamodb support
*/


func init() {
  BuildDefaultMap()
}

// Anything we need to make alexa work
func Start() {
  if ddb_tablename != "" {
    StartDynamoTable()
  }
  lambda.Start(DefaultHandler)
}

