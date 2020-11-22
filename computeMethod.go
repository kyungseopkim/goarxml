package goarxml

type ComputeNum struct {
    V1    float64   `json:"v1"`
    V2    float64   `json:"v2"`
}

type CompuScale struct {
    Label       string          `json:"label"`
    Min         float64         `json:"min"`
    Max         float64         `json:"max"`
    Numerators  ComputeNum      `json:"numerators"`
    Denominator float64         `json:"denominator"`
    Constant    string          `json:"const"`
}

type ComputeMethod struct {
    Name 		string 			`json:"name"`
    Category	string			`json:"category"`
    Unit 		string			`json:"unit"`
    Scale 		[]CompuScale 	`json:"scale"`
}

func NewComputeMethod(name string, category string, unit string, scale []CompuScale) ComputeMethod {
    return ComputeMethod{name, category, unit, scale }
}

func NewCompuScale(label string, min float64, max float64, numerators ComputeNum, denominator float64, constant string) CompuScale {
    return CompuScale{label, min, max, numerators, denominator, constant}
}

func NewCompuNum(v1 float64, v2 float64) ComputeNum {
   return ComputeNum{v1, v2}
}

func (cn ComputeNum) String() string  {
    return ToJson(cn)
}

func (cs CompuScale) String() string {
    return ToJson(cs)
}

func (cm ComputeMethod) String() string {
    return ToJson(cm)
}
