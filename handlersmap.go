package alexa

import (
  "github.com/webability-go/alexa/request"
  "github.com/webability-go/alexa/response"
)

const (
  // Requests types
  LaunchRequest             = "LaunchRequest"
  SessionEndedRequest       = "SessionEndedRequest"
  IntentRequest             = "IntentRequest"
  Fallback                  = "Fallback"

  // Intents names
  CancelIntent              = "AMAZON.CancelIntent"
  HelpIntent                = "AMAZON.HelpIntent"
  LoopOffIntent             = "AMAZON.LoopOffIntent"
  LoopOnIntent              = "AMAZON.LoopOnIntent"
  NextIntent                = "AMAZON.NextIntent"
  NoIntent                  = "AMAZON.NoIntent"
  PauseIntent               = "AMAZON.PauseIntent"
  PreviousIntent            = "AMAZON.PreviousIntent"
  RepeatIntent              = "AMAZON.RepeatIntent"
  ResumeIntent              = "AMAZON.ResumeIntent"
  ShuffleOffIntent          = "AMAZON.ShuffleOffIntent"
  ShuffleOnIntent           = "AMAZON.ShuffleOnIntent"
  StartOverIntent           = "AMAZON.StartOverIntent"
  StopIntent                = "AMAZON.StopIntent"
  YesIntent                 = "AMAZON.YesIntent"
    // ==== DISPLAY HANDLERS INTENTS ====
  ScrollUpIntent            = "AMAZON.ScrollUpIntent"
  ScrollLeftIntent          = "AMAZON.ScrollLeftIntent"
  ScrollDownIntent          = "AMAZON.ScrollDownIntent"
  ScrollRightIntent         = "AMAZON.ScrollRightIntent"
  PageUpIntent              = "AMAZON.PageUpIntent"
  PageDownIntent            = "AMAZON.PageDownIntent"
  MoreIntent                = "AMAZON.MoreIntent"
  NavigateHomeIntent        = "AMAZON.NavigateHomeIntent"
  NavigateSettingsIntent    = "AMAZON.NavigateSettingsIntent"
    // ==== LOCAL DEVICE NOTIFICATION ====
  ElementSelectedHandler    = "Display.ElementSelected"
)

var HandlerTypeMap map[string]func(request.AlexaRequest) (*response.AlexaResponse, error)
var HandlerIntentMap map[string]func(request.AlexaRequest) (*response.AlexaResponse, error)

/* You should node edit this map file, but use the functions to assign handlers instead */

func BuildDefaultMap() {
  HandlerTypeMap = map[string]func(request.AlexaRequest) (*response.AlexaResponse, error) {
    // ==== Common types =====
    Fallback:                 DefaultFallbackHandler,
    // ==== REQUESTS ====
    LaunchRequest:            DefaultLaunchHandler,
    SessionEndedRequest:      DefaultSessionEndedHandler,
    IntentRequest:            DefaultIntentTypeHandler,
    // ==== LOCAL DEVICE NOTIFICATION ====
    ElementSelectedHandler:   DefaultFallbackHandler,
    
    // ==== Custom handlers =====
    // Add all the custom type handlers you need with the functions
  }

  HandlerIntentMap = map[string]func(request.AlexaRequest) (*response.AlexaResponse, error) {
    // ==== Common intents =====
    Fallback:                 DefaultFallbackHandler,
    // ==== INTENTS ====
    CancelIntent:             DefaultSessionEndedHandler,
    HelpIntent:               DefaultIntentHandler,
    LoopOffIntent:            DefaultIntentHandler,
    LoopOnIntent:             DefaultIntentHandler,
    NextIntent:               DefaultIntentHandler,
    NoIntent:                 DefaultIntentHandler,
    PauseIntent:              DefaultIntentHandler,
    PreviousIntent:           DefaultIntentHandler,
    RepeatIntent:             DefaultIntentHandler,
    ResumeIntent:             DefaultIntentHandler,
    ShuffleOffIntent:         DefaultIntentHandler,
    ShuffleOnIntent:          DefaultIntentHandler,
    StartOverIntent:          DefaultIntentHandler,
    StopIntent:               DefaultSessionEndedHandler,
    YesIntent:                DefaultIntentHandler,
    // ==== DISPLAY HANDLERS INTENTS ====
    ScrollUpIntent:           DefaultIntentHandler,
    ScrollLeftIntent:         DefaultIntentHandler,
    ScrollDownIntent:         DefaultIntentHandler,
    ScrollRightIntent:        DefaultIntentHandler,
    PageUpIntent:             DefaultIntentHandler,
    PageDownIntent:           DefaultIntentHandler,
    MoreIntent:               DefaultIntentHandler,
    NavigateHomeIntent:       DefaultIntentHandler,
    NavigateSettingsIntent:   DefaultIntentHandler,

    // ==== Custom intents =====
    // Add all the intents you need with the functions
  }
}

func AddHandlerType(id string, fct func(request.AlexaRequest) (*response.AlexaResponse, error) ) {
  HandlerTypeMap[id] = fct
}

func AddHandlersType(m map[string]func(request.AlexaRequest) (*response.AlexaResponse, error) ) {
  for i, v := range(m) {
    HandlerTypeMap[i] = v
  }
}

func AddHandlerIntent(id string, fct func(request.AlexaRequest) (*response.AlexaResponse, error) ) {
  HandlerIntentMap[id] = fct
}

func AddHandlersIntent(m map[string]func(request.AlexaRequest) (*response.AlexaResponse, error) ) {
  for i, v := range(m) {
    HandlerIntentMap[i] = v
  }
}


