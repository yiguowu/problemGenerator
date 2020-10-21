package lib

import (
	"math/rand"
	"strconv"
)

type Generator struct {
	Operator string
	Range    int
}

func (gen *Generator) Init(operator string, limit int) {
	gen.Operator = operator
	gen.Range = limit
}
func (gen *Generator) Generate(Operand int) string {
	switch gen.Operator {
	case "+":
		ret := ""
		for i := 0; i < Operand; i++ {
			if ret == "" {
				ret = ret + strconv.Itoa(rand.Intn(gen.Range))
			} else {
				ret = ret + gen.Operator + strconv.Itoa(rand.Intn(gen.Range))
			}
		}
		return ret + "="
	default:
		return ""
	}
}
