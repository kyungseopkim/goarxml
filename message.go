package goarxml

const  (
    BIG_ENDIAN = iota
    LITTLE_ENDIAN
)

type Signal struct {
    Name        string      `json:"name"`
    Endian      int32        `json:"endian"`
    StartBit    int32       `json:"start_bit"`
    Length      int32       `json:"length"`
    Slope       float32     `json:"slope"`
    Intercept   float32     `json:"intercept"`
    Max         float32     `json:"max"`
    Min         float32     `json:"min"`
}

type Message struct {
    Name    string      `json:"name"`
    Id      int32       `json:"id"`
    Vlan    string      `json:"vlan"`
    Length  int32       `json:"length"`
    Signals []Signal    `json:"signals"`
}

func NewSignal(name string, endian int32, startbit int32, length int32, slope float32, intercept float32, max float32, min float32) Signal {
    return Signal{name, endian, startbit, length, slope, intercept, max, min}
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

