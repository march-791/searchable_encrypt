/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/
package util

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

func TestGenRandMatrix(t *testing.T) {
	a := GenRandMatrix(5, 6, 9)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
func TestGenMatrix(t *testing.T) {
	a := GenMatrix(5, 4, 1)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
func TestGenSymmetricPositiveDefiniteMatrix(t *testing.T) {
	a := GenSymmetricPositiveDefiniteMatrix(5, 5)
	fmt.Println(a)
	fa := mat.Formatted(a, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\na = %#v\n\n", fa)
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
func TestGenVecDense(t *testing.T) {
	a := GenVecDense(5, 1)
	fmt.Println(a)
}
