package response

import (
//  "fmt"
)


type Template interface {}

type TemplateCommon struct {
  Type    string `json:"type"`
}

type BodyTemplate1 struct {
  TemplateCommon
  Token    string `json:"token,omitempty"`
  BackButton    string `json:"backButton,omitempty"`
  BackgroundImage   *DisplayImage
  Title          string `json:"title,omitempty"`
  TextContent       *TextContent `json:"title,omitempty"`
}

type BodyTemplate2 struct {
  TemplateCommon
  Image   struct {
    Sources []ImageSource `json:"sources,omitempty"`
  } `json:"image,omitempty"`
  Title          string `json:"title,omitempty"`
  TextContent    struct {
  } `json:"textContent,omitempty"`
  BackButton       string `json:"backButton,omitempty"`
}

type BodyTemplate3 struct {
  TemplateCommon
  Token                 string `json:"token,omitempty"`
  BackButton            string `json:"backButton,omitempty"`
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
    case "BodyTemplate1": return &BodyTemplate1{}
    case "BodyTemplate2": return &BodyTemplate2{}
    case "BodyTemplate3": return &BodyTemplate3{}
  }
  return nil
}

func (builder *BodyTemplate1)Build() Directive {
  template := new(BodyTemplate1)
  tpl := &DirectiveRenderTemplate{Template: *template,}
  return tpl
}

func (builder *BodyTemplate2)Build() Directive {
  template := new(BodyTemplate1)
  tpl := &DirectiveRenderTemplate{Template: *template,}
  return tpl
}

func (builder *BodyTemplate3)Build() Directive {
  builder.Type = "BodyTemplate3"
  tpl := &DirectiveRenderTemplate{Template: builder,}
  tpl.Type = "Display.RenderTemplate"
  return tpl
}

// SPECIFIC FUNCTIONS 
func (builder *BodyTemplate3)WithTitle(title string) {
  builder.Title = title
}

func (builder *BodyTemplate3)WithImage(image string) {
  builder.Image = &DisplayImage{ }
  builder.Image.Sources = &[]ImageSource{ ImageSource{ URL: image, }, }
  
}

func (builder *BodyTemplate3)WithPrimaryRichText(text string) {
  builder.TextContent = &TextContent{ PrimaryText: &TextField{ Type: "RichText", Text: text, }, }
}

