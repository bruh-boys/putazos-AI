package brain

import "math"

func Argmax(output []float32) (b int) {
	for i := 0; i < len(output); i++ {
		if output[i] > output[b] {
			b = i
		}
	}
	return
}

func relu(x float32) float32 {
	return float32(math.Max(0, float64(x)))

}
func sigmoid(x float32) float32 {

	return float32(1 / (1 + math.Exp(float64(x)*(-1))))

}

func tanh(x float32) float32 {
	return float32(math.Tanh(float64(x)))
}

func devSigmoid(x float32) float32 {
	return x * (1 - x)
}
func devRelu(x float32) float32 {
	out := 0.0
	if x > 0 {
		out = 1
	}
	return float32(out)
}
func devTanh(x float32) float32 {
	return 1 - (x * x)

}

var (
	mathFuncs = map[string]map[string]func(float32) float32{
		"sigmoid": {
			"derivative": devSigmoid,
			"activate":   sigmoid,
		},
		"relu": {
			"derivative": devRelu,
			"activate":   relu,
		},
		"tanh": {
			"derivative": devTanh,
			"activate":   tanh,
		},
	}
)
