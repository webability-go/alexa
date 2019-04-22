package alexa

import (
  "fmt"
  
  "github.com/webability-go/alexa/locale"
  "github.com/webability-go/alexa/request"
  "github.com/webability-go/alexa/response"
)

func DefaultHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

//  fmt.Println("HANDLER: ", req)

  loc := req.GetLocale()
  messages := locale.Get(loc)
  requesttype := req.GetRequestType()
  fct, ok := HandlerTypeMap[requesttype]
  if !ok {
    errormessage := messages.Get("RequestTypeNotFound") + requesttype
    fmt.Println(errormessage)
    return response.NewSimpleResponse(messages.Get("ERROR"), errormessage, true), nil
  }
  return fct(req), nil
}

func DefaultLaunchHandler(req request.AlexaRequest) *response.AlexaResponse {
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("Launch.Title"), messages.Get("Launch.Message"), false)
}

func DefaultSessionEndedHandler(req request.AlexaRequest) *response.AlexaResponse {
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("SessionEnded.Title"), messages.Get("SessionEnded.Message"), true)
}

func DefaultIntentTypeHandler(req request.AlexaRequest) *response.AlexaResponse {

  loc := req.GetLocale()
  messages := locale.Get(loc)
  var intentname = req.GetRequestIntentName()
  fct, ok := HandlerIntentMap[intentname]
  if !ok {
    errormessage := messages.Get("IntentNameNotFound") + intentname
    fmt.Println(errormessage)
    return response.NewSimpleResponse(messages.Get("ERROR"), errormessage, true)
  }
  return fct(req)
}

func DefaultIntentHandler(req request.AlexaRequest) *response.AlexaResponse {
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("Handler.Title"), messages.Get("Handler.Message") + req.GetRequestIntentName(), false)
}
