package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bruh-boys/putazos-ai/backend/brain"
)

var letters = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " ", ",",
}
var la = map[byte]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
	' ': 26,
	',': 27,
}
var expect = []string{"clown timeline but never gona give you up", "just think in me", "never think in what is real"}

func Train(nn brain.NN) {
	for k := 0; k < 10000; k++ {
		bd, wd := [][][]float32{}, [][][][]float32{}

		for _, expected := range expect {
			input := make([]float32, len(letters))
			var b [][]float32 = nil

			input[la[byte(expected[0])]] = 1

			for i := 1; i < len(expected); i++ {

				layers, bef := nn.FeedFoward(input, b)
				exp := make([]float32, len(letters))
				exp[la[expected[i]]] = 1

				w, bi := nn.BackPropagation(layers, b, exp)
				bd = append(bd, bi)
				wd = append(wd, w)
				input = make([]float32, len(layers[len(layers)-1]))
				input[brain.Argmax(layers[len(layers)-1])] = 1

				b = bef

			}

			if k%100 == 0 {
				var b [][]float32 = nil

				input := make([]float32, len(letters))
				input[la[byte(expected[0])]] = 1

				fmt.Print(string(expected[0]))
				for i := 1; i < len(expected); i++ {
					out, bef := nn.Predict(input, b)
					input = make([]float32, len(out))
					input[brain.Argmax(out)] = 1

					fmt.Print(letters[brain.Argmax(out)])
					b = bef
				}
				fmt.Println("")
			}
		}
		for i := 0; i < len(bd); i++ {
			nn.UpdateWeightAndBias(float32(len(bd)), 0.02, wd[i], bd[i])
		}

	}
	nn.SaveModel("net-example.json")
}
func main() {
	nn := brain.OpenModel("net-example.json")

	input := make([]float32, len(letters))
	rand.Seed(time.Now().Unix())
	l := rand.Intn(len(letters))
	input[la[[]byte(letters[l])[0]]] = 1
	var b [][]float32 = nil

	fmt.Print(string(letters[l][0]))
	for i := 1; i < 50; i++ {
		out, bef := nn.Predict(input, b)
		input = make([]float32, len(out))
		input[brain.Argmax(out)] = 1

		fmt.Print(letters[brain.Argmax(out)])
		b = bef
	}
	fmt.Println("")

}
