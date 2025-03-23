// Adapted from math/bits.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

const smallestNormal = 1.17549435e-38 // 2**-126

// normalize returns a normal number y and exponent exp
// satisfying x == y × 2**exp. It assumes x is finite and non-zero.
func normalize(x float32) (y float32, exp int) {
	if Abs(x) < smallestNormal {
		return x * (1 << mantissaLen), -mantissaLen
	}

	return x, 0
}
