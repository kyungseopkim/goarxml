package goarxml

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	var arxml = "/Users/jerrykim/local/resources/arxml/S22.4.aspen.arxml"

	var msg = Parse(arxml)

	fmt.Println(msg)
}
