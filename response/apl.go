package response

import (
  "os"
  "io/ioutil"
  "fmt"
  "encoding/json"
)

type APLBuilder struct {
  Datatype string
  Version string
  Document string
  DataSources *APLDataSources
}

type APLDocument map[string]interface {}
type APLDataSources map[string]interface {}

func NewAPLBuilder(datatype string, version string, document string, datasources *APLDataSources) *APLBuilder {
  return &APLBuilder{
    Datatype: datatype,
    Version: version,
    Document: document,
    DataSources: datasources,
  }
}



func (builder *APLBuilder) Build() Directive {
  
  jsonFile, err := os.Open(builder.Document)
  // if we os.Open returns an error then handle it
  if err != nil {
    fmt.Println(err)
  }
  byteValue, _ := ioutil.ReadAll(jsonFile)
  jsonFile.Close()

  var document map[string]interface{}
  json.Unmarshal([]byte(byteValue), &document)

  apl := &DirectiveAPL{
        Token: "abcdef",
        Document: document,
        Datasources: builder.DataSources,
      }
  apl.Type = builder.Datatype
  return apl
}

