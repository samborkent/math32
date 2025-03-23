package math32_test

import (
	"testing"

	. "github.com/samborkent/math32"
)

var round = []float32{
	5,
	8,
	Copysign(0, -1),
	-5,
	10,
	3,
	5,
	3,
	2,
	-9,
}

var vfroundSC = [][2]float32{
	{0, 0},
	{1.39067167e-39, 0}, // denormal
	{0.49999994, 0},     // 0.5-epsilon
	{0.5, 1},
	{0.50000001, 1}, // 0.5+epsilon
	{-1.5, -2},
	{-2.5, -3},
	{NaN(), NaN()},
	{Inf(1), Inf(1)},
	{225179949.5, 225179950}, // 1 bit fraction
	{225179950.5, 225179951},
	{450359995.5, 450359996}, // 1 bit fraction, rounding to 0 bit fraction
	{450359997, 450359997},   // large integer
}

func TestRound(t *testing.T) {
	for i := range len(vf) {
		if f := Round(vf[i]); !alike(round[i], f) {
			t.Errorf("Round(%g) = %g, want %g", vf[i], f, round[i])
		}
	}

	for i := range vfroundSC {
		if f := Round(vfroundSC[i][0]); !alike(vfroundSC[i][1], f) {
			t.Errorf("Round(%g) = %g, want %g", vfroundSC[i][0], f, vfroundSC[i][1])
		}
	}
}
