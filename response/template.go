package response

import (
//  "fmt"
//  "strings"
)

const (
  TEMPLATE_BUTTON_VISIBLE =        "VISIBLE"
  TEMPLATE_BUTTON_HIDDEN  =        "HIDDEN"
 
)

type Template interface {}

type TemplateCommon struct {
  Type                  string                   `json:"type"`
}

type BodyTemplate1 struct {
  TemplateCommon
  Token                 string                  `json:"token,omitempty"`
  BackButton            string                  `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage           `json:"backgroundImage,omitempty"`
  Title                 string                  `json:"title,omitempty"`
  TextContent           *TextContent            `json:"textContent,omitempty"`
}

type BodyTemplate2 struct {
  TemplateCommon
  Token                 string                  `json:"token,omitempty"`
  BackButton            string                  `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage           `json:"backgroundImage,omitempty"`
  Title                 string                  `json:"title,omitempty"`
  Image                 *DisplayImage           `json:"image,omitempty"`
  TextContent           *TextContent            `json:"textContent,omitempty"`
}

type BodyTemplate3 struct {
  TemplateCommon
  Token                 string                   `json:"token,omitempty"`
  BackButton            string                   `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage            `json:"backgroundImage,omitempty"`
  Title                 string                   `json:"title,omitempty"`
  Image                 *DisplayImage            `json:"image,omitempty"`
  TextContent           *TextContent             `json:"textContent,omitempty"`
}

type BodyTemplate6 struct {
  TemplateCommon
  Token                 string                   `json:"token,omitempty"`
  BackButton            string                   `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage `json:"backgroundImage,omitempty"`
  Title                 string `json:"title,omitempty"`
  Image                 *DisplayImage `json:"image,omitempty"`
  TextContent           *TextContent `json:"textContent,omitempty"`
}

type BodyTemplate7 struct {
  TemplateCommon
  Token                 string                   `json:"token,omitempty"`
  BackButton            string                   `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage `json:"backgroundImage,omitempty"`
  Title                 string `json:"title,omitempty"`
  Image                 *DisplayImage `json:"image,omitempty"`
  TextContent           *TextContent `json:"textContent,omitempty"`
}

type ListTemplate1 struct {
  TemplateCommon
  Token                 string                   `json:"token,omitempty"`
  BackButton            string                   `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage `json:"backgroundImage,omitempty"`
  Title                 string `json:"title,omitempty"`
  Image                 *DisplayImage `json:"image,omitempty"`
  TextContent           *TextContent `json:"textContent,omitempty"`
}
 
type ListTemplate2 struct {
  TemplateCommon
  Token                 string                   `json:"token,omitempty"`
  BackButton            string                   `json:"backButton,omitempty"`
  BackgroundImage       *DisplayImage `json:"backgroundImage,omitempty"`
  Title                 string `json:"title,omitempty"`
  Image                 *DisplayImage `json:"image,omitempty"`
  TextContent           *TextContent `json:"textContent,omitempty"`
}




type TemplateBuilder interface {
  Build() Directive
}

func NewTemplateBuilder(templatetype string) TemplateBuilder {
  switch templatetype {
    case "BodyTemplate1": return &BodyTemplate1{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
    case "BodyTemplate2": return &BodyTemplate2{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
    case "BodyTemplate3": return &BodyTemplate3{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
    case "BodyTemplate6": return &BodyTemplate6{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
    case "BodyTemplate7": return &BodyTemplate7{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
    case "ListTemplate1": return &ListTemplate1{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
    case "ListTemplate2": return &ListTemplate2{ TemplateCommon: TemplateCommon{Type: templatetype, }, }
  }
  return nil
}

// Main functions build

func (builder *BodyTemplate1)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

func (builder *BodyTemplate2)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

func (builder *BodyTemplate3)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

func (builder *BodyTemplate6)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

func (builder *BodyTemplate7)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

func (builder *ListTemplate1)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

func (builder *ListTemplate2)Build() Directive {
  tpl := &DirectiveRenderTemplate{DirectiveCommon: DirectiveCommon{ Type: "Display.RenderTemplate", }, Template: builder,}
  return tpl
}

// SPECIFIC INTERFACE FUNCTIONS FOR EACH TEMPLATES

// ============================================================================================
// BodyTemplate1

func (builder *BodyTemplate1)WithToken(token string) {
  builder.Token = token
}

// use TEMPLATE_BUTTON_VISIBLE and TEMPLATE_BUTTON_HIDDEN as parameter
func (builder *BodyTemplate1)WithBackButton(button string) {
  builder.BackButton = button
}

func (builder *BodyTemplate1)WithBackgroundImage(image string) *DisplayImage {
  builder.BackgroundImage = &DisplayImage{ Sources: &[]ImageSource{ ImageSource{ URL: image, }, }, }
  return builder.BackgroundImage
}

func (builder *BodyTemplate1)WithTitle(title string) {
  builder.Title = title
}

func (builder *BodyTemplate1)WithPrimaryText(text string) *TextContent {
  builder.TextContent = &TextContent{}
  builder.TextContent.WithPrimaryText(text)
  return builder.TextContent
}

// ============================================================================================
// BodyTemplate2

func (builder *BodyTemplate2)WithToken(token string) {
  builder.Token = token
}

// use TEMPLATE_BUTTON_VISIBLE and TEMPLATE_BUTTON_HIDDEN as parameter
func (builder *BodyTemplate2)WithBackButton(button string) {
  builder.BackButton = button
}

func (builder *BodyTemplate2)WithBackgroundImage(image string) *DisplayImage {
  builder.BackgroundImage = &DisplayImage{ Sources: &[]ImageSource{ ImageSource{ URL: image, }, }, }
  return builder.BackgroundImage
}

func (builder *BodyTemplate2)WithTitle(title string) {
  builder.Title = title
}

func (builder *BodyTemplate2)WithImage(image string) *DisplayImage {
  builder.Image = &DisplayImage{ Sources: &[]ImageSource{ ImageSource{ URL: image, }, }, }
  return builder.Image
}

func (builder *BodyTemplate2)WithPrimaryText(text string) *TextContent {
  builder.TextContent = &TextContent{}
  builder.TextContent.WithPrimaryText(text)
  return builder.TextContent
}

// ============================================================================================
// BodyTemplate3

func (builder *BodyTemplate3)WithToken(token string) {
  builder.Token = token
}

// use TEMPLATE_BUTTON_VISIBLE and TEMPLATE_BUTTON_HIDDEN as parameter
func (builder *BodyTemplate3)WithBackButton(button string) {
  builder.BackButton = button
}

func (builder *BodyTemplate3)WithBackgroundImage(image string) *DisplayImage {
  builder.BackgroundImage = &DisplayImage{ Sources: &[]ImageSource{ ImageSource{ URL: image, }, }, }
  return builder.BackgroundImage
}

func (builder *BodyTemplate3)WithTitle(title string) {
  builder.Title = title
}

func (builder *BodyTemplate3)WithImage(image string) *DisplayImage {
  builder.Image = &DisplayImage{ Sources: &[]ImageSource{ ImageSource{ URL: image, }, }, }
  return builder.Image
}

func (builder *BodyTemplate3)WithPrimaryText(text string) *TextContent {
  builder.TextContent = &TextContent{}
  builder.TextContent.WithPrimaryText(text)
  return builder.TextContent
}

// BodyTemplate6


// BodyTemplate7


// ListTemplate1



// ListTemplate2

