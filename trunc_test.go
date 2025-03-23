// Adapted from math/all_test.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32_test

import (
	"testing"

	. "github.com/samborkent/math32"
)

var trunc = []float32{
	4.00000000e+00,
	7.00000000e+00,
	Copysign(0, -1),
	-5.00000000e+00,
	9.00000000e+00,
	2.00000000e+00,
	5.00000000e+00,
	2.00000000e+00,
	1.00000000e+00,
	-8.00000000e+00,
}

var truncSC = append(ceilBaseSC,
	1<<23-1,
	1<<23-1,
	1<<23,
	-1<<23,
	-1<<23+1,
	-1<<23+1,
	1<<24,
	-1<<24,
)

func TestTrunc(t *testing.T) {
	t.Parallel()

	for i := range len(vf) {
		if f := Trunc(vf[i]); !alike(trunc[i], f) {
			t.Errorf("Trunc(%g) = %g, want %g", vf[i], f, trunc[i])
		}
	}

	for i := range len(vfceilSC) {
		if f := Trunc(vfceilSC[i]); !alike(truncSC[i], f) {
			t.Errorf("Trunc(%g) = %g, want %g", vfceilSC[i], f, truncSC[i])
		}
	}
}
