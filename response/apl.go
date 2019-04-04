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
type APLDataSource map[string]interface {}

func NewAPLBuilder(datatype string, version string, document string, datasources *APLDataSources) *APLBuilder {
  return &APLBuilder{
    Datatype: datatype,
    Version: version,
    Document: document,
    DataSources: datasources,
  }
}

func NewAPLDataSources() *APLDataSources {
  return &APLDataSources{}
}

func (ds *APLDataSources)NewAPLDataSource(id string, objecttype string) *APLDataSource {
  src := &APLDataSource{}
  (*src)["type"] = objecttype
  (*src)["properties"] = make(map[string]interface{})
  (*ds)[id] = src
  return src
}

func (ds *APLDataSource)AddData(id string, val interface{}) {
  (*ds)["properties"].(map[string]interface{})[id] = val
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

