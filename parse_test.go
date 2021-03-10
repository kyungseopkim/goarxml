package goarxml

import (
    "fmt"
    "github.com/antchfx/xmlquery"
    "testing"
)

func loadTestDoc() *xmlquery.Node {
    fileName := "/Users/jerrykim/workspace/arxml/type/S21.6_WholeSys_Trial_20210303.arxml"
    doc, err := parseXml(fileName)
    if err != nil {
        panic(err)
    }
    return doc
}

func TestCompute(t *testing.T) {
    doc := loadTestDoc()
    fmt.Println(getDataTypes(doc))
}

func TestMessage_String(t *testing.T) {
    doc := loadTestDoc()
    vlan := getNetwork(doc)
    isignal := getISignal(doc)
    compu := getDataTypes(doc)

    fmt.Println(getMessage(doc, vlan, isignal, compu))
}

func TestGetNetwork(t *testing.T) {
    doc := loadTestDoc()
    vlan := getNetwork(doc)
    fmt.Println(vlan)
}

func TestGetDataType(t *testing.T) {
    doc := loadTestDoc()
    compus := getDataTypes(doc)
    compumap := getCompuMap(compus)
    value, ok := compumap["Headlight_signals"]
    if ok {
        fmt.Println(value)
    }
}