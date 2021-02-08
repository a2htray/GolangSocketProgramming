package operator

import (
	"regexp"
	"strconv"
)

var (
	Map = map[string]func(a float64, b float64) float64{
		"+": Add,
		"-": Sub,
		"*": Multiple,
		"/": Div,
	}
)

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Multiple(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) float64 {
	return a / b
}

func ParseExpression(expression string) (func(float64, float64) float64, float64, float64, error) {
	reg, err := regexp.Compile(`^(.*)([\+-\/*])(.*)$`)
	if err != nil {
		return nil,0, 0, err
	}

	matches := reg.FindAllStringSubmatch(expression, -1)
	a := matches[0][1]
	mapKey := matches[0][2]
	b := matches[0][3]

	aFloat64, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return nil,0, 0, err
	}

	bFloat64, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return nil,0, 0, err
	}

	return Map[mapKey], aFloat64, bFloat64, nil
}