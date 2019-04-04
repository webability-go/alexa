@UTF-8

# Alexa Ready to use SDK and Framework for GO

Alexa v0
=============================

The library is a full SDK with a framework, ready to deploy a lambda function on Amazon AWS to build a skill.
The SDK support all type of Alexa Requests, Attributes and Responses.
The Framework is multilanguage, supports SSML, Cards, Templates and APLs
The Framework comes also with all default intents pre-programmed (they answer the name of the intent) for basic operation.

First start:

Enter in your go environment and install the needed libraries:

```
go get github.com/aws/aws-sdk-go
go get github.com/aws/aws-lambda-go
go get github.com/webability-go/alexa
```

Creates a directory named "skill" into your go environment:

```
mkdir ¬/go/src/skill
```

creates a file called skill.go into your ¬/go/src/skill

```
package main

import "github.com/webability-go/alexa"

function main() {
  alexa.Start()
}
```

Then compile your file and zip it:

```
go build skill.go
zip -a skill.zip skill
```

Then import your zip file to your lambda function, with language GO 1.x
Name the Controler to the name of your executable (in this case, "skill")

( I will pass the "how to creates and compile an interaction model and link it with your lambda function", 
  there are enough tutorials of that already )

Once it is all linked, launch your skill in your test environment, and the skill works, magically.

It will tell the intents you say, for example if you invoke your skill with "my skill to test":

```
User: "Alexa, open my skill to test"
Alexa: "Alexa Skill Default Launch Handler."
User: "Help"
Alexa: "Alexa Skill Default Handler For HelpIntent"
User: "End skill"
Alexa: "Alexa Skill Default Handler For CancelIntent. Goodbye"
```

Refer to the full manual to implement your intents, use the SDK, framework and much more.


TO DO:
======
- Full Templates support (if works but not all the posibilities)
- Full APL support (if works but not all the posibilities)
- DynamoDB support to load/save attributes


Version Changes Control
=======================

V0.0.1 - 2019-04-04
-----------------------
- Framework working, SDK working
- AlexaRequest interpreter, Attributes interpreter and basic functions, AlexaResponse builder



# Reference Manual:
=======================






---
