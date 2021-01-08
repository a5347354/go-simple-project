package util

import (
	"fmt"
	"math"
	"strconv"
)

func RoundValue(value float64) int {
	returnValue, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", math.Trunc(value*math.Pow10(0)+0.5)*math.Pow10(-0)), 64)
	return int(returnValue)
}
