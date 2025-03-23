// Adapted from math/all_test.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32_test

import . "github.com/samborkent/math32"

var vf = []float32{
	4.97901195e+00,
	7.73887245e+00,
	-2.76880059e-01,
	-5.01060369e+00,
	9.63629373e+00,
	2.92637726e+00,
	5.22908346e+00,
	2.72793992e+00,
	1.82530850e+00,
	-8.68592473e+00,
}

func alike(a, b float32) bool {
	switch {
	case IsNaN(a) && IsNaN(b):
		return true
	case a == b:
		return Signbit(a) == Signbit(b)
	}
	return false
}

func tolerance(a, b, e float32) bool {
	// Multiplying by e here can underflow denormal values to zero.
	// Check a==b so that at least if a and b are small and identical
	// we say they match.
	if a == b {
		return true
	}
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}

func veryclose(a, b float32) bool { return tolerance(a, b, 4e-10) }
