package usecase

import (
	"math"

	"github.com/bramsedana/law-cots-1/entity"
)

func Arithmetic(equation entity.Equation) float64 {
	switch equation.Operation {
	case "add":
		return equation.First + equation.Second
	case "sub":
		return equation.First - equation.Second
	case "mul":
		return equation.First * equation.Second
	case "div":
		return equation.First / equation.Second
	case "pow":
		return math.Pow(equation.First, equation.Second)
	}
	return 0
}
