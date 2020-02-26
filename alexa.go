// Copyright Philippe Thomassigny 2018-2020.
// Use of this source code is governed by a MIT licence.
// license that can be found in the LICENSE file.

// Package alexa is a framework ready-to-go to build Apps for Lambda servers for Amazon Alexa
//
package alexa

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/webability-go/alexa/dynamodb"
	"github.com/webability-go/alexa/request"
)

const VERSION = "0.4.0"

var DEVEL = false
var ERRORCAPTURE = true

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
func Start() error {
	if dynamodb.DDB_tablename != "" {
		err := dynamodb.StartDynamoTable()
		if err != nil {
			return err
		}
	}
	lambda.Start(DefaultHandler)
	return nil
}

func SetErrorCapture(status bool) {
	ERRORCAPTURE = status
}

/* ==========================================================================
   HIJACK THE SESSION UNMARSHAL ON THE REQUEST TO REPLACE ATTRIBUTES
   ==========================================================================*/

func SetSessionUnmarshalerHandler(unmarshaler func(data []byte, s *request.Session) error) {
	request.SessionUnmarshalerHandler = unmarshaler
}
