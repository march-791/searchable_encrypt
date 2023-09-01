/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package trapdoor is construct
package trapdoor

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
	"gotest/algorithm/dualcode"
	"gotest/algorithm/keygen"
	"gotest/algorithm/matrix"
)

// Build trap door

/**
 * @Author peng
 * @Description // Build trap door
 * @Date 16:37 2022/6/28
 * @Param  cipher Query string
 * @return slice of VecDense
 **/

// GenTrapdoor Generate trap door
func GenTrapdoor(query []string, cipher *keygen.Cipher) []*mat.VecDense {
	if query == nil || cipher == nil {
		return nil
	}
	I := make([]*mat.VecDense, 2)
	var dualWordCode *mat.VecDense

	if cipher.Scope == keygen.LETTER {
		dualWordCode = dualcode.DualWordCodeSliceOnlyLetters(query, cipher.BaseUnit)
	} else if cipher.Scope == keygen.LETTERANDNURBER {
		dualWordCode = dualcode.DualWordCodeSliceLettersAndNumber(query, cipher.BaseUnit)
	} else if cipher.Scope == keygen.CHINESE {
		dualWordCode = dualcode.DualWordCodeSliceLettersAndNumber(query, cipher.BaseUnit)
	}
	if dualWordCode == nil {
		return nil
	}

	sk := cipher.SK
	rand.Seed(time.Now().UnixNano() + 3)
	r := rand.Float64() //nolint:gosec
	dimension := cipher.BaseUnit * cipher.BaseUnit
	// 构建B1 B2
	B1 := matrix.GenVecDense(dimension, 0)
	B2 := matrix.GenVecDense(dimension, 0)
	for j := 0; j < dimension; j++ {
		if sk.At(j, 0) == 1 {
			B1.SetVec(j, dualWordCode.AtVec(j))
			B2.SetVec(j, dualWordCode.AtVec(j))
		} else {
			B1.SetVec(j, 0.5*float64(dualWordCode.AtVec(j))+r)
			B2.SetVec(j, 0.5*float64(dualWordCode.AtVec(j))-r)
		}
	}

	I1 := mat.NewVecDense(dimension, nil)
	I2 := mat.NewVecDense(dimension, nil)

	// B1 B2 matrix * 正定矩阵M1 M2的逆
	tempMatrixM1 := mat.NewDense(dimension, dimension, nil)
	tempMatrixM2 := mat.NewDense(dimension, dimension, nil)
	err := tempMatrixM1.Inverse(cipher.M1)
	if err != nil {
		return nil
	}
	err = tempMatrixM2.Inverse(cipher.M2)
	if err != nil {
		return nil
	}
	I1.MulVec(tempMatrixM1, B1)
	I2.MulVec(tempMatrixM2, B2)
	I[0] = I1
	I[1] = I2
	return I
}
