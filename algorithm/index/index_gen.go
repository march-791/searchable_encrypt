/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package index build index
package index

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gotest/algorithm/dualcode"
	"gotest/algorithm/keygen"
	"gotest/algorithm/matrix"
)

// according to cipher Index each document
// filepath File path, numberOfFile Number of documents
// topN Number of keywords
// cipher Key generated in the previous step

/**
 * @Author peng
 * @Description //according to cipher Index each document
 * @Date 16:31 2022/6/28
 * @Param filepath File path, numberOfFile Number of documents
 * @return index
 **/

// BuildIndexDemo buildIndexDemo for document
func BuildIndexDemo(filePath string, numberOfFile, topN int, cipher *keygen.Cipher) [][]*mat.VecDense {

	// dualWordCodes[article][keywords]
	dualWordCodes := make([]*mat.VecDense, numberOfFile)
	// topKeys[article][keywords]
	topKeys := InitialKey(filePath, numberOfFile, topN)
	for _, value := range topKeys {
		fmt.Println(value)
	}
	for indexPage, topKey := range topKeys { //
		if cipher != nil {
			if cipher.Scope == keygen.LETTER {
				dualWordCodes[indexPage] = dualcode.DualWordCodeSliceOnlyLetters(topKey, cipher.BaseUnit)
			} else if cipher.Scope == keygen.LETTERANDNURBER {
				dualWordCodes[indexPage] = dualcode.DualWordCodeSliceLettersAndNumber(topKey, cipher.BaseUnit)
			} else if cipher.Scope == keygen.CHINESE {
				dualWordCodes[indexPage] = dualcode.DualWordCodeSliceLettersAndNumber(topKey, cipher.BaseUnit)
			}
		}
	}

	// Extract private key
	sk := cipher.SK

	// Build B1 B2
	// Traverse SK and B
	// if s[i] = 1 then b[i] = b1[i] = b2[i]
	// if s[i]=0 then b1[i] = 1/2 * b[i] + r and b2[i] = 1/2 * b[i] - r
	I := make([][]*mat.VecDense, numberOfFile)
	for i := 0; i < numberOfFile; i++ {
		I1 := make([]*mat.VecDense, 2)
		B1 := matrix.GenVecDense(cipher.BaseUnit*cipher.BaseUnit, 0)
		B2 := matrix.GenVecDense(cipher.BaseUnit*cipher.BaseUnit, 0)
		for j := 0; j < cipher.BaseUnit*cipher.BaseUnit; j++ {
			if sk.At(j, 0) == 1 {
				B1.SetVec(j, dualWordCodes[i].AtVec(j))
				B2.SetVec(j, dualWordCodes[i].AtVec(j))
			} else {
				B1.SetVec(j, 0.5*float64(dualWordCodes[i].AtVec(j))+cipher.R)
				B2.SetVec(j, 0.5*float64(dualWordCodes[i].AtVec(j))-cipher.R)
			}
		}
		B1.MulVec(cipher.M1.T(), B1)
		B2.MulVec(cipher.M2.T(), B2)
		I1[0] = B1
		I1[1] = B2
		I[i] = I1
	}
	return I
}

/**
 * @Author peng
 * @Description Build Index
 * @Date 16:31 2022/6/28
 * @Param keys, cipher key
 * @return index
 **/

// BuildIndex Build Index
func BuildIndex(keys []string, cipher *keygen.Cipher) []*mat.VecDense {
	if keys == nil || cipher == nil {
		return nil
	}
	var dualWordCode *mat.VecDense
	if cipher != nil {
		if cipher.Scope == keygen.LETTER {
			dualWordCode = dualcode.DualWordCodeSliceOnlyLetters(keys, cipher.BaseUnit)
		} else if cipher.Scope == keygen.LETTERANDNURBER {
			dualWordCode = dualcode.DualWordCodeSliceLettersAndNumber(keys, cipher.BaseUnit)
		} else if cipher.Scope == keygen.CHINESE {
			dualWordCode = dualcode.DualWordCodeSliceLettersAndNumber(keys, cipher.BaseUnit)
		}
	}
	if dualWordCode == nil {
		return nil
	}
	I1 := make([]*mat.VecDense, 2)
	sk := cipher.SK

	//构建B1 B2 矩阵
	B1 := matrix.GenVecDense(cipher.BaseUnit*cipher.BaseUnit, 0)
	B2 := matrix.GenVecDense(cipher.BaseUnit*cipher.BaseUnit, 0)
	for j := 0; j < cipher.BaseUnit*cipher.BaseUnit; j++ {
		if sk.At(j, 0) == 1 {
			B1.SetVec(j, dualWordCode.AtVec(j))
			B2.SetVec(j, dualWordCode.AtVec(j))
		} else {
			B1.SetVec(j, 0.5*float64(dualWordCode.AtVec(j))+cipher.R)
			B2.SetVec(j, 0.5*float64(dualWordCode.AtVec(j))-cipher.R)
		}
	}
	// B1 B2 matrix * 正定矩阵M1 M2的转置
	B1.MulVec(cipher.M1.T(), B1)
	B2.MulVec(cipher.M2.T(), B2)
	I1[0] = B1
	I1[1] = B2
	return I1
}
