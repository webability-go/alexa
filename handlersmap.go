package alexa

import (
  "./request"
  "./response"
)

const (
  // Requests types
  LaunchRequest             = "LaunchRequest"
  SessionEndedRequest       = "SessionEndedRequest"
  IntentRequest             = "IntentRequest"

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
  ElementSelectedHandler    = "AMAZON.ElementSelectedHandler"
)

var HandlerTypeMap map[string]func(request.AlexaRequest) *response.AlexaResponse
var HandlerIntentMap map[string]func(request.AlexaRequest) *response.AlexaResponse

/* You should node edit this map file, but use the functions to assign handlers instead */

func BuildDefaultMap() {
  HandlerTypeMap = map[string]func(request.AlexaRequest) *response.AlexaResponse {
    // ==== Common types =====
    // ==== REQUESTS ====
    LaunchRequest:            DefaultLaunchHandler,
    SessionEndedRequest:      DefaultSessionEndedHandler,
    IntentRequest:            DefaultIntentTypeHandler,
  }

  HandlerIntentMap = map[string]func(request.AlexaRequest) *response.AlexaResponse {
    // ==== Common intents =====
    // ==== INTENTS ====
    CancelIntent:             DefaultIntentHandler,
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
    StopIntent:               DefaultIntentHandler,
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
    // ==== LOCAL DEVICE NOTIFICATION ====
    ElementSelectedHandler:   DefaultIntentHandler,

    // ==== Custom intents =====
    // Add all the intents you need with the functions
  }
}

func AddHandlerType(id string, fct func(request.AlexaRequest) *response.AlexaResponse) {
  HandlerTypeMap[id] = fct
}

func AddHandlersType(m map[string]func(request.AlexaRequest) *response.AlexaResponse) {
  for i, v := range(m) {
    HandlerTypeMap[i] = v
  }
}

func AddHandlerIntent(id string, fct func(request.AlexaRequest) *response.AlexaResponse) {
  HandlerIntentMap[id] = fct
}

func AddHandlersIntent(m map[string]func(request.AlexaRequest) *response.AlexaResponse) {
  for i, v := range(m) {
    HandlerIntentMap[i] = v
  }
}


