package alexa

import (
  "fmt"
  "encoding/json"

  
  "github.com/webability-go/alexa/locale"
  "github.com/webability-go/alexa/request"
  "github.com/webability-go/alexa/response"
)

func DefaultHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

  if DEVEL {
    d, _ := json.Marshal(req)
    fmt.Println("HANDLER: ", string(d))
  }

  loc := req.GetLocale()
  messages := locale.Get(loc)
  requesttype := req.GetRequestType()
  fct, ok := HandlerTypeMap[requesttype]
  if !ok {
    fct, ok = HandlerTypeMap[Fallback]
    if !ok {
      errormessage := messages.Get("RequestTypeNotFound") + requesttype
      if DEVEL {
        fmt.Println(errormessage)
      }
      return response.NewSimpleResponse(messages.Get("ERROR"), errormessage, true), nil
    }
  }
  data, err := fct(req)
  if err != nil && ERRORCAPTURE {
    return response.NewSimpleResponse(messages.Get("ERROR"), fmt.Sprint(err), true), nil
  }
  return data, err
}

func DefaultFallbackHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("Fallback.Title"), messages.Get("Fallback.Message"), false), nil
}

func DefaultLaunchHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {
  
  // simulate an error
  return nil, fmt.Errorf("Esto es un error simulado de alexa")
  
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("Launch.Title"), messages.Get("Launch.Message"), false), nil
}

func DefaultSessionEndedHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("SessionEnded.Title"), messages.Get("SessionEnded.Message"), true), nil
}

func DefaultIntentTypeHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {

  loc := req.GetLocale()
  messages := locale.Get(loc)
  var intentname = req.GetRequestIntentName()
  fct, ok := HandlerIntentMap[intentname]
  if !ok {
    fct, ok = HandlerIntentMap[Fallback]
    if !ok {
      errormessage := messages.Get("IntentNameNotFound") + intentname
      fmt.Println(errormessage)
      return response.NewSimpleResponse(messages.Get("ERROR"), errormessage, true), nil
    }
  }
  return fct(req)
}

func DefaultIntentHandler(req request.AlexaRequest) (*response.AlexaResponse, error) {
  loc := req.GetLocale()
  messages := locale.Get(loc)
  return response.NewSimpleResponse(messages.Get("Handler.Title"), messages.Get("Handler.Message") + req.GetRequestIntentName(), false), nil
}
