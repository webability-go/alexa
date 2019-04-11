package response

import (
  "strings"

  "github.com/webability-go/alexa/attributes"
)

const (
  GENERATOR             =                        "alexa-1.0.0/Webability/GO"
  
  IMAGESOURCE_X_SMALL   =                        "X_SMALL"
  IMAGESOURCE_SMALL     =                        "SMALL"
  IMAGESOURCE_MEDIUM    =                        "MEDIUM"
  IMAGESOURCE_LARGE     =                        "LARGE"
  IMAGESOURCE_X_LARGE   =                        "X_LARGE"
)


// MAIN ALEXA RESPONSE STRUCTURE
type AlexaResponse struct {
  Version               string                   `json:"version"`
  SessionAttributes     attributes.AttributesDef `json:"sessionAttributes,omitempty"`      // map[string]interface{}, controlled by attributes module. can be empty
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

// Structured elements functionality
func analyzeText(text string) string {
  
  // if text contains < > tags, it's a rich text, if not, a normal text
  posinf := strings.Index(text, "<")
  possup := strings.Index(text, "<")
  if posinf != -1 && possup != -1 {
    return "RichText"
  }
  return "PlainText"
}

func (tc *TextContent)WithPrimaryText(text string) {
  texttype := analyzeText(text)
  tc.PrimaryText = &TextField{ Type: texttype, Text: text, }
}

func (tc *TextContent)WithSecondaryText(text string) {
  texttype := analyzeText(text)
  tc.SecondaryText = &TextField{ Type: texttype, Text: text, }
}

func (tc *TextContent)WithTertiaryText(text string) {
  texttype := analyzeText(text)
  tc.TertiaryText = &TextField{ Type: texttype, Text: text, }
}

func (di *DisplayImage)AddSource(url string) *ImageSource {
  src := &ImageSource{ URL: url, }
  if di.Sources == nil {
    di.Sources = &[]ImageSource{ *src }
  } else {
    *di.Sources = append(*di.Sources, *src)
  }
  return src
}

func (is *ImageSource)WithURL(url string) {
  is.URL = url
}

// size is IMAGESOURCE_X_SMALL to X_LARGE
// if width or height are 0, they are ignored
func (is *ImageSource)WithSize(size string, width int, height int) {
  is.Size = size
  is.WidthPixels = width
  is.HeightPixels = height
}











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
func NewTextResponse(text interface{}, close bool) *AlexaResponse {

  ntype := ""
  ntext := ""
  switch text.(type) {
    case *SSMLBuilder:
      ntype = "SSML"
      ntext = text.(*SSMLBuilder).Build()
    case string:
      ntype = analyzeText(text.(string))
      ntext = text.(string)
      if ntype == "SSML" {
        if ntext[0:7] != "<speak>" {
          ntext = "<speak>" + ntext
        }
        if ntext[len(ntext)-8:len(ntext)] != "</speak>" {
          ntext += "</speak>"
        }
      }
  }
// be intelligent:
// if text is string and no tags, simple text
// if text is string and some tags: ssml text, check if <speak> is here and adds it if not
// if text is ssml speech object, build it and inject it as ssml text

  r := &AlexaResponse{
    Version: "1.0",
    Response: Response{
      OutputSpeech: &OutputSpeech{
        Type: ntype,
      },
      ShouldEndSession: close,
    },
    UserAgent: GENERATOR,
  }
  if ntype == "SSML" {
    r.Response.OutputSpeech.SSML = ntext
  } else {
    r.Response.OutputSpeech.Text = ntext
  }
  return r
}


// disappear this function wher NewTextResponse is done
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
  
// be intelligent: text as interface{}
// if text is string and no tags, simple text
// if text is string and some tags: ssml text, check if <speak> is here and adds it if not
// if text is ssml speech object, build it and inject it as ssml text

// gets text for card from original text stripping tags and etc (some painter function to transform it)
  
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
    UserAgent: GENERATOR,
  }
  return r
}



// Add things to the response
func (r *AlexaResponse)AddAttributes(attributes attributes.AttributesDef) *AlexaResponse {
  r.SessionAttributes = attributes
  return r
}

func (r *AlexaResponse)AddSpeech(speech *SSMLBuilder) {
  
// be intelligent: text as interface{}
// if text is string and no tags, simple text
// if text is string and some tags: ssml text, check if <speak> is here and adds it if not
// if text is ssml speech object, build it and inject it as ssml text
  
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

