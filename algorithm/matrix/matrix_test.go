/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// package matrix test of matrix
package matrix

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

// TestGenRandMatrix Test generate RandMatrix
func TestGenRandMatrix(t *testing.T) {
	a := GenRandMatrix(5, 6, 9, 213124124925887)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

// TestGenVecDense Test generate VecDense
func TestGenVecDense(t *testing.T) {
	a := GenVecDense(5, 1)
	fmt.Println(a)
}

// TestGenRandK generate 676/2 location
func TestGenRandK(t *testing.T) {
	a := genRandK(676)
	fmt.Println(a)
}

// TestGenSkVec generate Vec
func TestGenSkVec(t *testing.T) {
	a := GenSkVec(676)
	fmt.Println(a)
}
