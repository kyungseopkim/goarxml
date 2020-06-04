package goarxml

type ComputeNum struct {
    V1    float32   `json:"v1"`
    V2    float32   `json:"v2"`
}

type CompuScale struct {
    Label       string          `json:"label"`
    Min         float32         `json:"min"`
    Max         float32         `json:"max"`
    Numerators  ComputeNum      `json:"numerators"`
    Denominator float32         `json:"denominator"`
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

func NewCompuScale(label string, min float32, max float32, numerators ComputeNum, denominator float32, constant string) CompuScale {
    return CompuScale{label, min, max, numerators, denominator, constant}
}

func NewCompuNum(v1 float32, v2 float32) ComputeNum {
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
