package goarxml

import (
	"sort"
	"strings"
)

const (
	BIG_ENDIAN = iota
	LITTLE_ENDIAN
)

const (
	NORMAL_MSG       = "normal"
	SEC_MSG          = "sec"
	MULTIPLEXING_MSG = "multiplex"
)

type Signal struct {
	Name      string  `json:"name"`
	Endian    int32   `json:"endian"`
	StartBit  int32   `json:"startBit"`
	Length    int32   `json:"length"`
	Slope     float64 `json:"slope"`
	Intercept float64 `json:"intercept"`
	Max       float64 `json:"max"`
	Min       float64 `json:"min"`
	Unit      string  `json:"unit"`
	IsSigned  bool    `json:"signed"`
	DataType  string  `json:"dataType"`
	Desc      string  `json:"description"`
}

type Message struct {
	Name    string   `json:"name"`
	Id      int32    `json:"id"`
	Vlan    string   `json:"vlan"`
	Length  int32    `json:"length"`
	Crc     bool     `json:"crc"`
	Type    string   `json:"type"`
	Signals []Signal `json:"signals"`
}

type MultiplexMessage struct {
	Name           string            `json:"name"`
	Id             int32             `json:"id"`
	Length         int32             `json:"length"`
	Type           string            `json:"type"`
	SelectorStart  int32             `json:"selectorStart"`
	SelectorLength int32             `json:"selectorLength"`
	SelectorEndian int32             `json:"selectorEndian"`
	Alternative    map[int32]Message `json:"alternative"`
}

func NewSignal(name string, endian int32, startbit int32, length int32, slope float64,
	intercept float64, max float64, min float64, unit string, signed bool, dataType string,
	desc string) Signal {
	return Signal{name, endian, startbit, length, slope,
		intercept, max, min, unit, signed, dataType, desc}
}

func (s Signal) String() string {
	return ToJson(s)
}

func NewMultiplexMessage(name string, id int32, length int32, msgType string,
	selectorStart int32, selectorLength int32, selectorEndian int32,
	alternative map[int32]Message) MultiplexMessage {
	return MultiplexMessage{name, id, length, msgType,
		selectorStart, selectorLength, selectorEndian,
		alternative}
}

func (m MultiplexMessage) String() string {
	return ToJson(m)
}

func NewMessage(name string, id int32, vlan string, length int32, crc bool, msgType string, signals []Signal) Message {
	return Message{name, id, vlan, length, crc, msgType, signals}
}

func (m Message) String() string {
	return ToJson(m)
}

type ByStartbit []Signal

func (s ByStartbit) Len() int { return len(s) }

func (s ByStartbit) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByStartbit) Less(i, j int) bool { return s[i].StartBit < s[j].StartBit }

func (s ByStartbit) IsCrc() bool {
	if len(s) > 0 && (strings.HasSuffix(s[0].Name, "CRC") ||
		strings.HasSuffix(s[0].Name, "CRC1") ||
		strings.HasSuffix(strings.ToUpper(s[0].Name), "CHECKSUM")) {
		return true
	}
	return false
}

func (m Message) SortByStartbit() {
	sort.Sort(ByStartbit(m.Signals))
}

func SortByStartbit(msgs []Message) {
	for _, msg := range msgs {
		msg.SortByStartbit()
	}
}

func Message2Lookup(msgs []Message) map[string]Message {
	ret := make(map[string]Message)
	for _, msg := range msgs {
		ret[msg.Name] = msg
	}
	return ret
}

func DetectEndian(text string) int {
	if text == "MOST-SIGNIFICANT-BYTE-LAST" {
		return LITTLE_ENDIAN
	}
	return BIG_ENDIAN
}
