package lib

import "flag"

type ParamSet struct {
	Operator *string
	Count    *int
	Range    *int
	Operand  *int
}

func ParseParameter() ParamSet {
	var params ParamSet
	params.Operator = flag.String("o", "+", "Operator")
	params.Count = flag.Int("c", 20, "Count")
	params.Range = flag.Int("r", 20, "Range")
	params.Operand = flag.Int("a", 2, "Operand")
	flag.Parse()
	return params
}
