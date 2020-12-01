package lib

import (
	"math/rand"
	"strconv"
	"time"
)

type Generator struct {
	Operator string
	Range    int
}

func (gen *Generator) Init(operator string, limit int) {
	gen.Operator = operator
	gen.Range = limit
}

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}

func (gen *Generator) Generate(Operand int) string {
	rand.Seed(time.Now().UTC().UnixNano())
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
	case "-":
		first := rand.Intn(gen.Range)
		second := rand.Intn(gen.Range)
		if first < second {
			swap(&first, &second)
		}
		return strconv.Itoa(first) + gen.Operator + strconv.Itoa(second) + "="
	default:
		return ""
	}
}
