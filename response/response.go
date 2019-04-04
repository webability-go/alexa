package response

import (
  "github.com/webability-go/alexa/attributes"
)

const (
  GENERATOR = "alexa-1.0.0/Webability/GO"
)

// MAIN ALEXA RESPONSE STRUCTURE
type AlexaResponse struct {
  Version               string                   `json:"version"`
  SessionAttributes     *attributes.Attributes   `json:"sessionAttributes,omitempty"`      // map[string]interface{}, controlled by attributes module. can be empty
  Response              Response                 `json:"response"`                         // response is mandatory (not a pointer *)
  UserAgent             string                   `json:"userAgent"`
}

// Response structures
type Response struct {
  OutputSpeech          *OutputSpeech            `json:"outputSpeech,omitempty"`
  Card                  *Card                    `json:"card,omitempty"`
  Reprompt              *Reprompt                `json:"reprompt,omitempty"`
  Directives            *[]Directive             `json:"directives,omitempty"`
  ShouldEndSession      bool                     `json:"shouldEndSession"`
}

// Reprompt object
type Reprompt struct {
  OutputSpeech          *OutputSpeech            `json:"outputSpeech,omitempty"`           // optional
}

// OutputSpeech object
type OutputSpeech struct {
  Type                  string                   `json:"type"`
  Text                  string                   `json:"text,omitempty"`
  SSML                  string                   `json:"ssml,omitempty"`
  PlayBehavior          string                   `json:"playBehavior,omitempty"`
}

// Card object
type Card struct {
  Type                  string                   `json:"type"`
  Title                 string                   `json:"title,omitempty"`
  Content               string                   `json:"content,omitempty"`
  Text                  string                   `json:"text,omitempty"`
  Image                 *CardImage               `json:"image,omitempty"`                 // optional
}

// Images for a card
type CardImage struct {
  SmallImageURL         string                   `json:"smallImageUrl,omitempty"`
  LargeImageURL         string                   `json:"largeImageUrl,omitempty"`
}

// Directives for audio, video, display, dialog
type Directive interface {}

type DirectiveCommon struct {
  Type                  string                   `json:"type"`
}

// Basic objects for directives
type DisplayImage struct {
  ContentDescription  string                   `json:"contentDescription"`
  Sources             *[]ImageSource           `json:"sources"`
}

type ImageSource struct {
  URL                   string                   `json:"url"`
  Size                  string                   `json:"size,omitempty"`
  WidthPixels           int                      `json:"widthPixels,omitempty"`
  HeightPixels          int                      `json:"heightPixels,omitempty"`
} 

type TextContent struct {
  PrimaryText           *TextField               `json:"primaryText"`
  SecondaryText         *TextField               `json:"secondaryText,omitempty"`
  TertiaryText          *TextField               `json:"tertiaryText,omitempty"`
}

type TextField struct {
  Type                  string                   `json:"type"`
  Text                  string                   `json:"text"`
}

// AUDIO PLAY DIRECTIVE
type DirectiveAudioPlay struct {
  DirectiveCommon
  PlayBehavior          string                   `json:"playBehavior"`
  AudioItem struct {
    Stream struct {
      URL               string                   `json:"url"`
      Token             string                   `json:"token"`
      expectedPreviousToken string               `json:"token,omitempty"`
      OffsetInMilliseconds int                   `json:"offsetInMilliseconds"`
    }                                            `json:"stream"`
    Metadata            *AudioMetaData           `json:"token,omitempty"`
  }                                              `json:"audioItem"`
}

type AudioMetaData struct {
  Title                 string                   `json:"title,omitempty"`
  Subtitle              string                   `json:"subtitle,omitempty"`
  Art                   *DisplayImage            `json:"art,omitempty"`
  BackgroundImage       *DisplayImage            `json:"backgroundImage,omitempty"`
}

type DirectiveAudioStop struct {
  DirectiveCommon
}

type DirectiveAudioClearQueue struct {
  DirectiveCommon
  ClearBehavior         string                   `json:"playBehavior"`
}


// DIALOG DIRECTIVE -- not implemented yet


// DISPLAY INTERFACE DIRECTIVE

// Render Templates
type DirectiveRenderTemplate struct {
  DirectiveCommon
  Template              Template                 `json:"template"`
}

// *** Render templates implemented into template.go

// APL Templates
type DirectiveAPL struct {
  DirectiveCommon
  Document              APLDocument              `json:"document"`
  Datasources           *APLDataSources          `json:"datasources,omitempty"`
  Token                 string                   `json:"token"`
}

// *** APL templates implemented into apl.go



// Basic Response creator
func NewResponse(close bool) *AlexaResponse {
  r := &AlexaResponse{
    Version: "1.0",
    Response: Response{
      ShouldEndSession: close,
    },
    UserAgent: GENERATOR,
  }
  return r
}

// Some Common responses pre-build
func NewSSMLResponse(text string, close bool) *AlexaResponse {
  r := &AlexaResponse{
    Version: "1.0",
    Response: Response{
      OutputSpeech: &OutputSpeech{
        Type: "SSML",
        SSML: text,
      },
      ShouldEndSession: close,
    },
    UserAgent: GENERATOR,
  }
  return r
}

func NewSimpleResponse(title string, text string, close bool) *AlexaResponse {
  r := &AlexaResponse{
    Version: "1.0",
    Response: Response{
      OutputSpeech: &OutputSpeech{
        Type: "PlainText",
        Text: text,
      },
      Card: &Card{
        Type:    "Simple",
        Title:   title,
        Content: text,
      },
      ShouldEndSession: close,
    },
    UserAgent: "kiwiask-1.0.0/GO/Kiwilimon",
  }
  return r
}



// Add things to the response
func (r *AlexaResponse)AddAttributes(attributes *attributes.Attributes) *AlexaResponse {
  r.SessionAttributes = attributes
  return r
}

func (r *AlexaResponse)AddSpeech(speech *SSMLBuilder) {
  if r.Response.OutputSpeech == nil {
    r.Response.OutputSpeech = &OutputSpeech{
        Type: "SSML",
        SSML: speech.Build(),
      }
  } else {
    r.Response.OutputSpeech.Text += speech.Build()
  }
}

func (r *AlexaResponse)AddCard(card *CardBuilder) {
  r.Response.Card = card.Build()
}

func (r *AlexaResponse)AddTemplate(template TemplateBuilder) {
  res := template.Build()
  if r.Response.Directives == nil {
    r.Response.Directives = &[]Directive{res}
  } else {
    *r.Response.Directives = append(*r.Response.Directives, res)
  }
}

func (r *AlexaResponse)AddAPL(apl *APLBuilder) {
  res := apl.Build()
  if r.Response.Directives == nil {
    r.Response.Directives = &[]Directive{res}
  } else {
    *r.Response.Directives = append(*r.Response.Directives, res)
  }
}

