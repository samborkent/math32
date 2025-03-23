// Adapted from math/all_test.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32_test

import (
	"testing"

	. "github.com/samborkent/math32"
)

var modf = [][2]float32{
	{4, 0.979012},
	{7, 0.7388725},
	{Copysign(0, -1), -2.76880059e-01},
	{-5, -0.010603905},
	{9, 0.6362934},
	{2, 0.9263773},
	{5, 0.22908354},
	{2, 0.72793984},
	{1, 0.82530844},
	{-8, -0.68592453},
}

var vfmodfSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	Inf(1),
	NaN(),
}

var modfSC = [][2]float32{
	{Inf(-1), NaN()}, // [2]float64{Copysign(0, -1), Inf(-1)},
	{Copysign(0, -1), Copysign(0, -1)},
	{Inf(1), NaN()}, // [2]float64{0, Inf(1)},
	{NaN(), NaN()},
}

func TestModf(t *testing.T) {
	for i := range len(vf) {
		if f, g := Modf(vf[i]); !veryclose(modf[i][0], f) || !veryclose(modf[i][1], g) {
			t.Errorf("Modf(%g) = %g, %g, want %g, %g", vf[i], f, g, modf[i][0], modf[i][1])
		}
	}

	for i := range len(vfmodfSC) {
		if f, g := Modf(vfmodfSC[i]); !alike(modfSC[i][0], f) || !alike(modfSC[i][1], g) {
			t.Errorf("Modf(%g) = %g, %g, want %g, %g", vfmodfSC[i], f, g, modfSC[i][0], modfSC[i][1])
		}
	}
}
