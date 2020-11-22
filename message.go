package goarxml

import "sort"

const (
	BIG_ENDIAN = iota
	LITTLE_ENDIAN
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
}

type Message struct {
	Name    string   `json:"name"`
	Id      int32    `json:"id"`
	Vlan    string   `json:"vlan"`
	Length  int32    `json:"length"`
	Signals []Signal `json:"signals"`
}

func NewSignal(name string, endian int32, startbit int32, length int32, slope float64, intercept float64, max float64, min float64, unit string) Signal {
	return Signal{name, endian, startbit, length, slope, intercept, max, min, unit}
}

func (s Signal) String() string {
	return ToJson(s)
}

func NewMessage(name string, id int32, vlan string, length int32, signals []Signal) Message {
	return Message{name, id, vlan, length, signals}
}

func (m Message) String() string {
	return ToJson(m)
}

type ByStartbit []Signal

func (s ByStartbit) Len() int { return len(s) }

func (s ByStartbit) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByStartbit) Less(i, j int) bool { return s[i].StartBit < s[j].StartBit }

func (m Message) SortByStartbit() {
	sort.Sort(ByStartbit(m.Signals))
}

func SortByStartbit(msgs []Message) {
	for _, msg := range msgs {
		msg.SortByStartbit()
	}
}
