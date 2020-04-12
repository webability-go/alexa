package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	APLDATATYPE = "Alexa.Presentation.APL.RenderDocument"
	APLVERSION  = "1.0"
)

type APLBuilder struct {
	Datatype    string
	Version     string
	Document    string
	DataSources *APLDataSources
}

type APLDocument map[string]interface{}
type APLDataSources map[string]interface{}
type APLDataSource map[string]interface{}
type APLDataSourceList []APLDataSource

func NewAPLBuilder(document string, datasources *APLDataSources) *APLBuilder {
	return &APLBuilder{
		Datatype:    APLDATATYPE,
		Version:     APLVERSION,
		Document:    document,
		DataSources: datasources,
	}
}

func NewAPLDataSources() *APLDataSources {
	return &APLDataSources{}
}

func (ds *APLDataSources) NewDataSource(id string, objecttype string) *APLDataSource {
	src := &APLDataSource{}
	(*src)["type"] = objecttype
	(*ds)[id] = src
	return src
}

func (ds *APLDataSource) NewDataSet(id string) *APLDataSource {
	src := &APLDataSource{}
	(*ds)[id] = src
	return src
}

func (ds *APLDataSource) AddData(id string, val interface{}) {
	(*ds)[id] = val
}

func (ds *APLDataSource) NewDataList(id string) *APLDataSourceList {
	src := &APLDataSourceList{}
	(*ds)[id] = src
	return src
}

func (ds *APLDataSourceList) NewDataListItem() *APLDataSource {
	src := &APLDataSource{}
	*ds = append(*ds, *src)
	return src
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
		Token:       "abcdef",
		Document:    document,
		Datasources: builder.DataSources,
	}
	apl.Type = builder.Datatype
	return apl
}
