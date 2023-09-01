/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package util

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenRandMatrix 根据行列生成随机矩阵
// rand.Seed(time.Now().UnixNano())保证每次都不相同
// 矩阵中每个元素都是int 范围是[0,n)
// GenRandMatrix Gen Rand Matrix
func GenRandMatrix(column, row, scope int) *mat.Dense {
	rander := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	randMatrix := mat.NewDense(column, row, nil)
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			randMatrix.Set(i, j, float64(rander.Intn(scope))) //nolint:gosec
		}
	}
	return randMatrix
}

// 根据行列和指定值生成随机矩阵
// 矩阵中每个元素都是int
/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenMatrix gen Matrix
func GenMatrix(column, row, value int) *mat.Dense {
	if column <= 0 || row <= 0 || value <= 0 {
		return nil
	}
	randMatrix := mat.NewDense(column, row, nil)

	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			randMatrix.Set(i, j, float64(value))
		}
	}
	return randMatrix
}

// GenSymmetricPositiveDefiniteMatrix 生成正定矩阵
/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
// GenSymmetricPositiveDefiniteMatrix Gen SymmetricPositive Definite Matrix
func GenSymmetricPositiveDefiniteMatrix(column, row int) *mat.Dense {
	if column <= 0 || row <= 0 {
		return nil
	}
	randMatrix := mat.NewDense(column, row, nil)
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			if i == j {
				randMatrix.Set(i, j, 1.0)
			} else {
				randMatrix.Set(i, j, 0.0)
			}
		}
	}
	return randMatrix
}

//生成向量 n 维度 value为向量的值
/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenVecDense Gen Vec Dense
func GenVecDense(n, value int) *mat.VecDense {
	if n <= 0 || value <= 0 {
		return nil
	}
	randMatrix := mat.NewVecDense(n, nil)
	for i := 0; i < n; i++ {
		randMatrix.SetVec(i, float64(value))
	}
	return randMatrix
}
