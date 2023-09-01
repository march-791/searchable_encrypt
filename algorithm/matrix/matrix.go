/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package matrix provides matrix and vector operation
package matrix

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

/**
 * @Author peng
 * @Description Generate random matrix
 * @Date 16:35 2022/6/28
 * @Param column,row,scope of element,Random number seed
 * @return Random matrix
 **/

// GenRandMatrix Generate random matrix Based on rows and columns
func GenRandMatrix(column, row, scope int, seed int64) *mat.Dense {
	if column <= 0 || row <= 0 || scope <= 0 {
		return nil
	}
	rander := rand.New(rand.NewSource(seed)) //nolint:gosec
	randMatrix := mat.NewDense(column, row, nil)
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			randMatrix.Set(i, j, float64(rander.Intn(scope)))
		}
	}
	return randMatrix
}

// 生成SK 保证sk中0-1的个数相当
// k代表向量的维度

/**
 * @Author peng
 * @Description GenSkVec
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenSkVec Gen vector according to k, divide half-and-half 0 and 1
func GenSkVec(k int) *mat.Dense {
	if k <= 0 {
		return nil
	}
	randMatrix := mat.NewDense(k, 1, nil)
	a := genRandK(k)
	for i := 0; i < k; i++ {
		if isExist(a, i) {
			randMatrix.Set(i, 0, 1.0)
		} else {
			randMatrix.Set(i, 0, 0.0)
		}
	}
	return randMatrix
}

/**
 * @Author peng
 * @Description GenVecDense
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

//GenVecDense Gen vector according to k and value
func GenVecDense(n, value int) *mat.VecDense {
	if n <= 0 {
		return nil
	}
	randMatrix := mat.NewVecDense(n, nil)
	for i := 0; i < n; i++ {
		randMatrix.SetVec(i, float64(value))
	}
	return randMatrix
}

/**
 * @Author peng
 * @Description gen RandK
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// genRandK according to len Select half of len location random
func genRandK(len int) []int {
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = i
	}
	for i := 0; i < len; i++ {
		j := rand.Intn(len) //nolint:gosec
		if i == j {
			continue
		}
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}
	return arr[:len/2]
}

/**
 * @Author peng
 * @Description isExist
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// isExist Determine whether the item exists
func isExist(arr []int, item int) bool {
	for _, value := range arr {
		if item == value {
			return true
		}
	}
	return false
}
