package goarxml

import (
    "fmt"
    "github.com/antchfx/xmlquery"
    "testing"
)

func loadTestDoc() *xmlquery.Node {
    fileName := "/Users/jerrykim/workspaces/data/arxml/ARXML_BE_40.1.arxml"
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