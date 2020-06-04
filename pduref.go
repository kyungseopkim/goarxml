package goarxml

type PduRef struct {
	Name 	string	`json:"name"`
	Ref 	string	`json:"ref"`
	Id 		int32	`json:"id"`
}

func newPduRef(name string, ref string, id int32) PduRef {
	return PduRef{name, ref, id}
}

func (pdu PduRef) String() string {
	return ToJson(pdu)
}

