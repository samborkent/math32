// Adapted from math/ldexp.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

// Ldexp is the inverse of [Frexp].
// It returns frac × 2**exp.
//
// Special cases are:
//
//	Ldexp(±0, exp) = ±0
//	Ldexp(±Inf, exp) = ±Inf
//	Ldexp(NaN, exp) = NaN
func Ldexp(frac float32, exp int) float32 {
	// special cases
	switch {
	case frac == 0:
		return frac // correctly return -0
	case IsInf(frac, 0) || IsNaN(frac):
		return frac
	}

	frac, e := normalize(frac)
	exp += e
	x := math.Float32bits(frac)
	exp += int(int32(x>>shift)&mask - bias)

	if exp < -150 { // 2**(8-1) + (23-1)
		return Copysign(0, frac) // underflow
	}

	if exp > 127 { // overflow
		if frac < 0 {
			return Inf(-1)
		}

		return Inf(1)
	}

	var m float32 = 1

	if exp < -126 { // denormal
		exp += mantissaLen + 1
		m = 1.0 / (1 << (mantissaLen + 1)) // 2**-24
	}

	x &^= mask << shift
	x |= uint32(exp+bias) << shift

	return m * math.Float32frombits(x)
}
