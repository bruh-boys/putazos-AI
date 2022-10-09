package brain

import (
	"math/rand"
)

type Network struct {
	Bias            [][]float64
	Weights         [][][]float64
	ActivationFuncs []string
}

func NewNetwork(layers []int, activationFuncs []string) (net Network) {
	net.ActivationFuncs = activationFuncs
	net.Bias = make([][]float64, len(layers)-1)
	net.Weights = make([][][]float64, len(layers)-1)
	// this only works for defining the bias and the weights
	for l := 0; l < len(layers)-1; l++ {
		// its really simple and you should know why this look like this
		net.Weights[l] = make([][]float64, layers[l])
		net.Bias[l] = make([]float64, layers[l+1])
		// each neuron have a connection that have a weight
		for n := 0; n < layers[l]; n++ {
			net.Weights[l][n] = make([]float64, layers[l+1])
			for c := 0; c < layers[l+1]; c++ {
				net.Weights[l][n][c] = rand.Float64() - 0.5
			}
		}
		// the input doesnt have a bias , this is only for the hidden layers and the last layer
		for n := 0; n < layers[l+1]; n++ {
			net.Bias[l][n] = rand.Float64() - 0.5
		}

	}
	return
}

func (net Network) Foward(input []float64) []float64 {
	output := []float64{}
	for l := 0; l < len(net.Weights); l++ {
		output = make([]float64, len(net.Bias[l]))
		copy(output, net.Bias[l])

		// this is really basic stuff, is jsut
		// input*weights(matrix multiplication)
		for n := 0; n < len(net.Weights[l]); n++ {
			for c := 0; c < len(net.Weights[l][n]); c++ {
				output[c] += input[n] * net.Weights[l][n][c]
			}
		}
		// activation(output)
		for n := 0; n < len(output); n++ {
			output[n] = mathFuncs[net.ActivationFuncs[l]](output[n])
		}
		input = make([]float64, len(output))
		copy(input, output)
	}
	return output
}

func (net *Network) Mutate(netCopy Network) {
	for l := 0; l < len(net.Bias); l++ {
		for n := 0; n < len(net.Weights[l]); n++ {
			for c := 0; c < len(net.Weights[l][n]); c++ {

				net.Weights[l][n][c] = randomValue(netCopy.Weights[l][n][c])

			}
		}
		for n := 0; n < len(net.Bias[l]); n++ {

			net.Bias[l][n] = randomValue(netCopy.Bias[l][n])

		}
	}
}

func randomValue(val float64) float64 {
	// it only have a 10% of probability of having another value
	if rand.Float64() < 0.1 {
		return rand.Float64() - 0.5

	}
	// this doesnt change too much
	return val + rand.NormFloat64()
}
