/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package keygen Generate key
package keygen

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
	"gotest/algorithm/matrix"
)

type SEARCHSCOPE int

const (
	LETTER SEARCHSCOPE = iota
	LETTERANDNURBER
	CHINESE

	LETTERLENS       = 26
	LETTERANDNUMLENS = 36
	CHINESELENS      = 36
)

// Cipher include two matrix M1 M2 and SK (vector)
// M1 M2 The elements of the matrix(M1 M2) consist of 0 or 1
// SK  consist of 0 or 1
type Cipher struct {
	M1       *mat.Dense //matrix
	M2       *mat.Dense //matrix
	SK       *mat.Dense //vector
	BaseUnit int        //BaseUnit*BaseUnit is Dimension of vector
	R        float64    //random number[0,1)
	Scope    SEARCHSCOPE
}

/**
 * @Author peng
 * @Description  Generate key, baseUnit is a range of search scope
 * example
 * @Date 16:34 2022/6/28
 * @Param Dimension of vector
 * @return Cipher
 **/

// GenCipher based on baseUnit Generate key
func GenCipher(scope SEARCHSCOPE) *Cipher {
	var baseUnit int
	if scope == LETTER {
		baseUnit = LETTERLENS
	} else if scope == LETTERANDNURBER {
		baseUnit = LETTERANDNUMLENS
	} else if scope == CHINESE {
		baseUnit = CHINESELENS
	} else {
		scope = LETTER
		baseUnit = LETTERLENS
	}
	k := baseUnit * baseUnit
	currentTime := time.Now().UnixNano()
	m1 := matrix.GenRandMatrix(k, k, 2, currentTime)
	m2 := matrix.GenRandMatrix(k, k, 2, currentTime+1)
	sk := matrix.GenSkVec(k)
	rand.Seed(currentTime)
	r := rand.Float64() //nolint:gosec
	return &Cipher{
		M1:       m1,
		M2:       m2,
		SK:       sk,
		BaseUnit: baseUnit,
		R:        r,
		Scope:    scope,
	}
}

/**
 * @Author peng
 * @Description InitCipher
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// InitCipher Init Cipher
func InitCipher(baseUnit int, r float64, m1, m2, sk *mat.Dense, scope SEARCHSCOPE) *Cipher {
	if sk == nil || m1 == nil || m2 == nil || baseUnit <= 0 || r < 0 || r >= 1 || scope < LETTER || scope > CHINESE {
		return nil
	}
	if !(checkMatirx(m1, scope) || checkMatirx(m2, scope) || checkVector(sk, scope)) {
		return nil
	}
	return &Cipher{
		M1:       m1,
		M2:       m2,
		SK:       sk,
		BaseUnit: baseUnit,
		R:        r,
		Scope:    scope,
	}
}

func checkMatirx(matrix *mat.Dense, scope SEARCHSCOPE) bool {
	if matrix == nil {
		return false
	}
	row, col := matrix.Caps()
	if scope == LETTER {
		if row != LETTERLENS*LETTERLENS || col != LETTERLENS*LETTERLENS {
			return false
		}
	} else if scope == LETTERANDNURBER {
		if row != LETTERANDNUMLENS*LETTERANDNUMLENS || col != LETTERANDNUMLENS*LETTERANDNUMLENS {
			return false
		}
	} else if scope == CHINESE {
		if row != CHINESELENS*CHINESELENS || col != CHINESELENS*CHINESELENS {
			return false
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix.At(i, j) != 0 && matrix.At(i, j) != 1 {
				return false
			}
		}
	}
	return true
}

func checkVector(vector *mat.Dense, scope SEARCHSCOPE) bool {
	if vector == nil {
		return false
	}
	row, col := vector.Caps()
	if scope == LETTER {
		if row != LETTERANDNUMLENS*LETTERANDNUMLENS || col != 1 {
			return false
		}
	} else if scope == LETTERANDNURBER {
		if row != LETTERANDNUMLENS*LETTERANDNUMLENS || col != 1 {
			return false
		}
	} else if scope == CHINESE {
		if row != CHINESELENS*CHINESELENS || col != 1 {
			return false
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if vector.At(i, j) != 0 && vector.At(i, j) != 1 {
				return false
			}
		}
	}
	return true
}
