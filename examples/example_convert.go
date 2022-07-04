package examples

import (
	"github.com/tobycroft/Calc"
	"testing"
)

func convert_to_string(T *testing.B) string {
	str := ""
	str += Calc.Any2String(int64(1))
	str += Calc.Any2String(int(1))
	str += Calc.Any2String(float64(1.314))
	str += Calc.Any2String(1.314)
	str += Calc.Any2String("string")
	return str
}

func convert_to_int64() int64 {
	number := int64(0)
	number += Calc.Any2Int64("5")
	t, err := Calc.Any2Int64_2("5")
	if err != nil {
		return 0 //errors
	}
	number += t
	return number
}

func convert_to_float64() float64 {
	number := float64(0)
	number += Calc.Any2Float64("5.5")
	t, err := Calc.Any2Float64_2("4.5")
	if err != nil {
		return 0 //errors
	}
	number += t
	return number
}
