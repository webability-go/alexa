@UTF-8

# Alexa Ready to use SDK and Framework for GO

Alexa v0.4
=============================

The library is a full SDK with a framework, ready to deploy a lambda function on Amazon AWS to build a skill.
The SDK support all type of Alexa Requests, Attributes and Responses.
The Framework is multilanguage, supports SSML, Cards, Templates and APLs
The Framework comes also with all default intents pre-programmed (they answer the name of the intent) for basic operation.

Full manual is below.

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

Refer to the full manual below to implement your intents, use the SDK, framework and much more.


TO DO:
======
Important:
- Full APL support (it works but not all the posibilities i.e. missing transformers)
- Get user account address and country still not totally working
- Implement Dynamic Entities for Customized Interactions

Not so important:
- Finish the implementation of Amazon API for user data (still missing todo lists and shopping lists)
- Verify beta intent request canfulfillintentrequest for english skills
- Verify special requests (game requests, playback requests, gadgets requests)
- Gadget controlers responses ( i.e. buttons )



# Reference Manual:
=======================

Define your own Handlers map:
=======================

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
  alexa.AddHandlersType(map[string]func(request.AlexaRequest) (*response.AlexaResponse, error) {
    alexa.LaunchRequest:                 yourLaunchHandler,
    alexa.SessionEndedRequest:           yourSessionEndedHandler,
    alexa.Fallback:                      yourFallbackHandler,
  })

  // Handlers intents:
  alexa.AddHandlersIntent(map[string]func(request.AlexaRequest) (*response.AlexaResponse, error) {
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

    // fallback
    alexa.Fallback:                      yourFallbackIntentHandler,

  })

  alexa.Start()
}

// ======================================================================
// EXAMPLE: LAUNCH HANDLER
// ======================================================================
func yourLaunchHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

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

  return resp, nil
}

// all the other defined handlers

```

Attributes:
======================

```
  const REGION = "us-east-1"
  const TABLENAME = "my_dynamo_table"

  // Before start:
  alexa.WithDynamoDbClient("latest", REGION)
  alexa.WithTableName(TABLENAME)
  alexa.WithAutoCreateTable(true)
  alexa.Start()
```

Request data:
======================

```

  reqtype        := Request.GetRequestType()            // string
  intentname     := Request.GetRequestIntentName()      // string
  sessionid      := Request.GetSessionId()              // string
  isnewsession   := Request.GetNewSession()             // bool
  userid         := Request.GetUserId()                 // string
  attributes     := Request.GetAttributes()             // object
  slots          := Request.GetSlots()                  // *map[string]Slot

  display        := Request.GetDisplay()                // object
  video          := Request.GetVideo()                  // object
  apl            := Request.GetAPL()                    // object
  hasdisplay     := Request.HasDisplay()                // bool
  hasvideo       := Request.HasVideo()                  // bool
  hasapl         := Request.HasAPL()                    // bool

  newSession     := Request.GetNewSession()             // bool
  locale         := Request.GetLocale()                 // string es_MX

```


Use attributes:
======================

```
  att := &MySkillAttributes{}    // if hijacked , (see example below)

  // load persistent attributes with dynamoDB
  err := alexa.LoadPersistentAttributes(req, att)
  if err != nil {
    fmt.Println(err)
  }

  // no persistent attributes ? load the request attributes
  att := Request.GetAttributes()

  ...

  // play with attributes
  att["Something"] = "Some data"
  // Create att.AddData, GetData AddString, GetString, GetBool, GetInt etc (or use xcore.DataSet)

  // Add the attributes to the response
  resp.AddAttributes(att)     // rename to SetAttributes ?   ADD should Ads something to a set of attributes.

  // set persistent attributes with dynamoDB
  err := alexa.SavePersistentAttributes(req, att)
  if err != nil {
    fmt.Println(err)
  }

```

Hijack default attribute with your own attribute structure
======================

```
func main() {

  alexa.SetSessionUnmarshalerHandler(AttributesHijack)   // Custom attributes
  err := alexa.Start()
  if err != nil {
    fmt.Println(err)
  }
}

type MySkillAttributes struct {
  Version             int            // attributes version
  Sessions            int            // Quantity of launched sessions
  First               time.Time      // First use of the skill
  Last                time.Time      // Last use of the skill
}

func AttributesHijack(data []byte, session *request.Session) error {

  type Alias request.Session
  aux := &struct {
    Attributes *MySkillAttributes `json:"attributes"`
    *Alias
  }{
    Alias: (*Alias)(session),
  }
  if err := json.Unmarshal(data, &aux); err != nil {
    return err
  }
  session.Attributes = aux.Attributes
  return nil
}


func yourIntentHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

  MyAttributes := req.GetAttributes().(*MySkillAttributes)    // is now a MySkillAttributes, not a default map[string]interface{}
  (*MyAttributes).Sessions ++

  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support SSML (mandatory)
  speech := response.NewSSMLBuilder()
  speech.Say("Hello " + givenname)
  resp.AddSpeech(speech);

  resp.AddAttributes(MyAttributes)

  return rest, nil
}


```


Locale:
======================

```

func yourIntentHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

  loc := req.GetLocale()

  // You can implement locale dependant translation table for your skill texts

  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support SSML (mandatory)
  speech := response.NewSSMLBuilder()
  speech.Say("Your locale is " + loc)
  resp.AddSpeech(speech);

  return rest, nil
}

```

Build a Speech / SSML
======================

Every function you call on a builder will ADDs the message to the output

```

  resp := response.NewResponse(false)   // false: launch does not close the skill

  // SSML build
  speech := response.NewSSMLBuilder()
  // adds a raw text (syntax is supposed to be "good")
  speech.Raw("<s>Welcome to Demo Skill</s> <say-as interpret-as="cardinal">12345</say-as>. <say-as interpret-as='spell-out'>hello</say-as>.")

  // Simple text
  speech.Say("Welcome to Demo Skill")

  // With a break
  speech.Break("0.5s")

  // sentence
  speech.Say("This is a sentence")
  speech.SetSentence()   // apply on previous "Say"

  // paragraph
  speech.Say("This is a paragraph")
  speech.SetParagraph()   // apply on previous "Say"


  speech.Say("This is a text with lots of effects")
  speech.AddEffect("whispered")    // apply on previous "Say"     values in Alexa developpers SSML Manuals
  speech.AddEmphasis("moderate")   // apply on previous "Say"     values in Alexa developpers SSML Manuals
  speech.AddLang("fr-FR")          // apply on previous "Say"     values in Alexa developpers SSML Manuals
  speech.AddVoice("Kendra")        // apply on previous "Say"     values in Alexa developpers SSML Manuals

  speech.Say("12345")
  speech.AddSayAs("spell-out")

  // Finally adds an audio sound
  speech.Audio("soundbank://soundlibrary/animals/amzn_sfx_bear_groan_roar_01")

  resp.AddSpeech(speech);

```


Build a Card
======================

```
  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support CARD
  card := response.NewCardBuilder( "Welcome", "Welcome to Demo Skill", "https://yourcdn.com/icon-1024.png", "https://yourcdn.com/icon-192.png" )

  resp.AddCard(card);

```

Build a Permission Card
======================

```
  resp := response.NewResponse(false)   // false: launch does not close the skill

  // permission CARD
  card := response.NewPermissionCardBuilder([]string{response.PERMISSION_EMAIL, response.PERMISSION_FIRSTNAME,})

  resp.AddCard(card);

```

Possible permissions:

```
const (
  PERMISSION_FULLNAME = "alexa::profile:name:read"
  PERMISSION_FIRSTNAME = "alexa::profile:given_name:read"
  PERMISSION_EMAIL = "alexa::profile:email:read"
  PERMISSION_MOBILE = "alexa::profile:mobile_number:read"
  PERMISSION_COUNTRY_AND_POSTAL_CODE = "read::alexa:device:all:address:country_and_postal_code"
  PERMISSION_ADDRESS = "read::alexa:device:all:address"
)
```




Build a Template
======================

```
  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support TEMPLATE
  template := response.NewTemplateBuilder("BodyTemplate3").(*response.BodyTemplate3)
  template.WithTitle("Example:")
  template.WithImage("https://yourcdn.com/icon-1024.png");
  template.WithPrimaryRichText("<div align='center'>Help.<br/>Start over.<br/>Close the skill.</div>");

  resp.AddTemplate(template);

```

Build a Template with a list
======================

```
  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support TEMPLATE
  template := response.NewTemplateBuilder("ListTemplate1").(*response.ListTemplate1)
  template.WithTitle("Select Something:")
  template.WithPrimaryRichText("<div align='center'>Help.<br/>Start over.<br/>Close the skill.</div>")
  template.AddListItem("Element 1", "https://yourcdn.com/icon1-1024.png", "The name of your item 1")
  template.AddListItem("Element 2", "https://yourcdn.com/icon2-1024.png", "The name of your item 2")
  template.AddListItem("Element 3", "https://yourcdn.com/icon3-1024.png", "The name of your item 3")
  template.AddListItem("Element 4", "https://yourcdn.com/icon4-1024.png", "The name of your item 4")
  template.AddListItem("Element 5", "https://yourcdn.com/icon5-1024.png", "The name of your item 5")

  resp.AddTemplate(template);

```


Build an APL
======================

```
  resp := response.NewResponse(false)   // false: launch does not close the skill

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

```


Launch a video
======================

```

  var resp *response.AlexaResponse

  video := &response.DirectiveVideoAppLaunch{}
  video.Type = "VideoApp.Launch"
  video.VideoItem.Source = "Your_Video_Source_MP4_or_M3U"
  video.VideoItem.Metadata = &response.VideoMetadata{ Title: "Title of the video", Subtitle: "Description of the video" }
  resp.AddVideo(video);

```


Consume Alexa/Amazon APIs
======================

Every API data is supposed to be authorized by the user of the skill

The timezone, distante and temperatureunit does not need authorization.

```

func yourIntentHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

  fullname, _ := alexa.GetAccountFullName(req)
  givenname, _ := alexa.GetAccountGivenName(req)
  email, _ := alexa.GetAccountEmail(req)
  number, _ := alexa.GetAccountMobileNumber(req)
  country_and_postalcode, _ := alexa.GetDeviceCountry(req)
  address, _ := alexa.GetDeviceAddress(req)
  timezone, _ := alexa.GetDeviceTimeZone(req)
  distunit, _ := alexa.GetDeviceDistanceUnit(req)
  tempunit, _ := alexa.GetDeviceTemperatureUnit(req)

  resp := response.NewResponse(false)   // false: launch does not close the skill

  // support SSML (mandatory)
  speech := response.NewSSMLBuilder()
  speech.Say("Hello " + givenname)
  resp.AddSpeech(speech);

  return rest, nil
}

```



Version Changes Control
=======================

v0.4.0 - 2020-02-25
-----------------------
- Licence added
- Few changes for publication
- Few bugs corrected

v0.3.3 - 2019-06-19
-----------------------
- Response.ssml enhanced to support some effects and raw text

v0.3.2 - 2019-05-10
-----------------------
- Upgrade documentation with all previous changes


v0.3.1 - 2019-05-08
-----------------------
- Alexa API Settings working (timezone, distance and temperature units)
- Alexa API Address implemented (not fully working yet)
- Response Permission cards added for address and country/postal code


v0.3.0 - 2019-05-06
-----------------------
- Alexa API working (get user account email, name, full name, mobile number implemented)
- Permission cards implemented
- Some minor bugs corrected


v0.2.0 - 2019-04-29
-----------------------
- Added Fallback Handlers
- Added error propagation on all the handlers to be more compliant with error management. If you catch the error and manage it, then you should return "nil" as error parameter.
- The error is captured and modified into the default main handler, and transformed to a voice error message. This can be deactivated with the SetErrorCapture(false) function


v0.1.0 - 2019-04-22
-----------------------
- DynamoDB implemented to manage attributes persistance: create table, LoadPersistentAttributes, SavePersistentAttributes
- Amazon API implemented (name, fullname, email, mobile number, address, country, timezone, distanceunit, temperatureunit)


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
