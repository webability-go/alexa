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
- Full Request Implementation
- Full APL support (it works but not all the posibilities i.e. transformers)
- Finish the implementation of Amazon API for user data


Version Changes Control
=======================

v0.1.0 - 2019-04-22
-----------------------
- DynamoDB implemented to manage attributes persistance: create table, LoadPersistentAttributes, SavePersistentAttributes


v0.0.9 - 2019-04-19
-----------------------
- Removed ./attributes and code. Attributes is just an interface{} by default and can be used to store anything.
- The hijack function works now correctly to build an app own attributes structure


v0.0.8 - 2019-04-16
-----------------------
- ShouldEndSession parameter is now a *bool since it must be omited when the directive is a video launch. Code adjusted


v0.0.7 - 2019-04-15
-----------------------
- Video Directive implemented in response
- Bug corrected on ResolutionsPerAuthority object in the request (missing an "s")


v0.0.6 - 2019-04-12
-----------------------
- Some bugs corrected on NewTextResponse on the text analyse to decode to correct type.


v0.0.5 - 2019-04-11
-----------------------
- NewTextResponse implemented. NewSSMLResponse removed. Please use NewTextResponse, the text can be a string or an SSMLBuilder, the system reacts intelligently to it and build the correct text string.


v0.0.4 - 2019-04-10
-----------------------
- Default CancelIntent and StopIntent now ends the skill (changed in default intents map)
- APL builder enhanced, to build better the datasources, functions added: NewDataSet, NewDataList, NewDataListItem
- APL builder: "properties" subset removed. If a properties subset si needed, please use NewDataSet("properties")
- Attributes handler can be hijacked to unmarshal attributes into a custom structure instead of a map[string]interface{} structure
- Request Attributes and Response Attributes are now an interface to be able to be overloaded with a custom attributes structure


v0.0.3 - 2019-04-09
-----------------------
- Added HasVideo, HasDisplay and HasAPL in request implementation
- Request, Intent and Slots ordered and cleaned
- Added Request.GetSlots function
- Added attributes basic access functions (Set, Get, GetInt, GetBool, GetString, GetFloat)


v0.0.2 - 2019-04-08
-----------------------
- Full implementation of BodyTemplate1, BodyTemplate2, BodyTemplate3 (WithToken, WithTitle, WithBackButton, WithImage, WithBackgroundImage, WithPrimaryText) 
- Functions added to control DisplayImage and TextContent objects (WithSize, WithPrimaryText, WithSecondaryText, WithTertiaryText, AddSource)
- Full implementation of ListTemplate1, ListTemplate2 and ListItem (WithToken, WithTitle, WithBackButton, WithImage, WithBackgroundImage, AddListItem) 


V0.0.1 - 2019-04-04
-----------------------
- Framework working, SDK working
- AlexaRequest interpreter, Attributes interpreter and basic functions, AlexaResponse builder



# Reference Manual:
=======================

Define Handlers map:
======================

```
package main

import (
  ...

  "github.com/webability-go/alexa/request"
  "github.com/webability-go/alexa/response"
)

// Build the handlers map befor calling the start
// The full supported handlers are in handlersmap.go library for reference

func main()
{
  // Handlers types:
  alexa.AddHandlersType(map[string]func(request.AlexaRequest) *response.AlexaResponse {
    alexa.LaunchRequest: yourLaunchHandler,
    alexa.SessionEndedRequest: yourSessionEndedHandler,
  })

  // Handlers intents:
  alexa.AddHandlersIntent(map[string]func(request.AlexaRequest) *response.AlexaResponse {
    // native intents
    alexa.CancelIntent:                  yourCancelIntentHandler,
    alexa.StopIntent:                    yourCancelIntentHandler,
    alexa.HelpIntent:                    yourHelpIntentHandler,
    alexa.NextIntent:                    yourNextIntentHandler,
    alexa.PreviousIntent:                yourPreviousIntentHandler,
    alexa.RepeatIntent:                  yourRepeatIntentHandler,
    alexa.StartOverIntent:               yourStartOverIntentHandler,
    alexa.MoreIntent:                    yourMoreIntentHandler,
    alexa.ElementSelectedHandler:        yourElementSelectedHandler,
    
    // custom intents
    "yourOwnIntent":                     yourOwnIntentHandler,
    "anotherCurtomIntent":               yourAnotherCustomIntentHandler,
    "navigationIntent":                  yourNavigationIntentHandler,
  })

  alexa.Start()
}

// ======================================================================
// EXAMPLE: LAUNCH HANDLER
// ======================================================================
func yourLaunchHandler(req request.AlexaRequest) *response.AlexaResponse {

  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support SSML (mandatory)
  speech := response.NewSSMLBuilder()
  speech.Say("Welcome to Demo Skill")
  resp.AddSpeech(speech);

  // support CARD
  card := response.NewCardBuilder( "Welcome", "Welcome to Demo Skill", "https://yourcdn.com/icon-1024.png", "https://yourcdn.com/icon-192.png" )
  resp.AddCard(card);
  
  // support TEMPLATE
  template := response.NewTemplateBuilder("BodyTemplate3").(*response.BodyTemplate3)
  template.WithTitle("Example:")
  template.WithImage("https://yourcdn.com/icon-1024.png");
  template.WithPrimaryRichText("<div align='center'>Help.<br/>Start over.<br/>Close the skill.</div>");
  resp.AddTemplate(template);
  
  // support APL
  aplsources := response.NewAPLDataSources()
  apldata := aplsources.NewAPLDataSource("welcomedata", "object")
  apldata.AddData("logo", "https://yourcdn.com/icon-192.png")
  apldata.AddData("image", "https://yourcdn.com/icon-1024.png")
  apldata.AddData("maintitle", "Welcome")
  apldata.AddData("titleshort", "Examples:")
  apldata.AddData("title", "Examples of what you can say:")
  apldata.AddData("subtitle", "Search something, Make an action like that:")
  apldata.AddData("primaryText", "Help.<br/>Start over.<br/>Close the skill.<br/>Say something intelligent.")
  
  apl := response.NewAPLBuilder( "Alexa.Presentation.APL.RenderDocument", "1.0", "./application/apl/yourapl.json", aplsources )
  resp.AddAPL(apl);
  
  return resp
}

// all the other defined handlers

```

Attributes:
======================

```
  // Before start:
  alexa.WithDynamoDbClient("latest", REGION)
  alexa.WithTableName(KIWITABLA)
  alexa.WithAutoCreateTable(true)
  alexa.Start()
```

Request data: ( pass the IsNil interface to alexa code with new functions HasDisplay, HasVideo, HasAPL )

```
  display := Request.GetDisplay()          // object
  video := Request.GetVideo()              // object
  apl := Request.GetAPL()                  // object

  newSession := Request.GetNewSession()    // bool
  locale := Request.GetLocale()            // string es_MX
```
  

Use attributes:
======================

```
  att := Request.GetAttributes()
  ...
  att["Something"] = "Some data"

  resp.AddAttributes(att)     // rename to SetAttributes ?   ADD should Ads something to a set of attributes.
   // Create att.AddData, GetData AddString, GetString, GetBool, GetInt etc (or use xcore.DataSet)
   
```

Hijack default attribute with your own attribute structure
======================



Locale:
======================

```

```

Build a Speech / SSML
======================

Build a Card
======================

Build a Template
======================

Build an APL
======================


---
