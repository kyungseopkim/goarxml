package goarxml

type Network struct {
	Name 	string 		`json:"name"`
	Vlan	int32		`json:"vlan"`
	PduRef  []PduRef	`json:"pdu"`
}

func newNetwork(name string, vlan int32, pdu []PduRef) Network {
	n := Network{name, vlan, pdu}
	return n
}

func (network Network) String() string {
	return ToJson(network)
}
