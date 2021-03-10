package goarxml

type ISignal struct {
    Name 		string 		`json:"name"`
    Length		int32		`json:"length"`
    Desc		string		`json:"desc"`
    Ref 		string 		`json:"ref"`
    Init 		float64		`json:"init"`
    IsSigned	bool		`json:"isSigned"`
    DataType    string      `json:"dataType"`
}

func NewISignal(name string, length int32, desc string,
    ref string, init float64, isSigned bool, dataType string) ISignal {
    return ISignal{name, length, desc,
        ref, init, isSigned, dataType}
}

func (isignal ISignal) String() string {
    return ToJson(isignal)
}
