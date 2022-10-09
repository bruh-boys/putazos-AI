package brain

import "math"

func relu(x float64) float64 {
	return math.Max(0.0, x)
}
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
func tanh(x float64) float64 {
	return math.Tanh(x)
}

var (
	mathFuncs = map[string]func(float64) float64{
		"relu":    relu,
		"sigmoid": sigmoid,
		"tanh":    tanh,
	}
)
