/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package index

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

// AB AB是向量a和b对应的数据结构
type AB struct {
	A *mat.VecDense // 向量A
	B float64       // [0,r]中的随机浮点数
}

// 根据n,r生成一组向量a,和随机数b
// n 为最终生成向量a的维数,向量a中的每一个元素服从标准正态分布(N(0,1))
// b 随机数为[0,r]中的浮点数

/**
 * @Author peng
 * @Description Gen Para
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenPara Gen Para
func GenPara(n, r int, randerSeed *rand.Rand) *AB {
	a := mat.NewVecDense(n, nil)
	for i := 0; i < n; i++ {
		a.SetVec(i, randerSeed.NormFloat64())
	}
	b := randerSeed.Float64() * float64(r) // [0.0,1.0) * r
	return &AB{
		a,
		b,
	}
}

// 生成k组由n,r 组成的A b
// n是的维度 r是随机数b的范围 k是一共有多少组A b

/**
 * @Author peng
 * @Description Gen Para
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenE2LSHFamily Gen E2LSH Family
func GenE2LSHFamily(n, k, r int) []*AB {
	rander := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	sliceLSH := make([]*AB, k)
	for i := 0; i < k; i++ {
		sliceLSH[i] = GenPara(n, r, rander)
	}
	return sliceLSH
}

/**
 * @Author peng
 * @Description GenHashVals
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// GenHashVals 根据传入的向量组计算hash值公式是 向量内积a*v + b(随机数) / r
// GenHashVals Gen Hash Vals
func GenHashVals(ab []*AB, v *mat.VecDense, r int) []float64 {
	sliceHash := make([]float64, len(ab))
	for index, value := range ab {
		hashVal := (mat.Dot(value.A, v) + value.B) / float64(r)
		sliceHash[index] = hashVal
	}
	return sliceHash
}

/**
 * @Author peng
 * @Description h2
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// H2 将k组 向量内积a*v + b(随机数) / r 求和 然后与c求模
// H2
func H2(hashvals, fpRand []float64, k, c int) int {
	var sum float64 = 0
	for i := 0; i < k && i < len(fpRand) && i < len(hashvals); i++ {
		sum = (hashvals[i] * fpRand[i]) + sum
	}
	return int(sum) % c
}

// dataSet相当于数据集
// k是向量的维度
// 有l个 [h1, ...hk]这样的向量组  [[h1, ...hk], [h1, ..hk], ..., [h1, ...hk]]
// r是ab中随机数b的范围[0,r)
// tableSize是哈希桶的大小
/**
 * @Author peng
 * @Description Gen E2LSH
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// E2LSH E2 LSH
func E2LSH(dataSet []*mat.VecDense, k, l, r, tableSize int) ([]map[int][]int, [][]*AB, []float64, [][][]float64) {
	hashTable := make([]map[int][]int, tableSize)
	n := dataSet[0].Len()                                     // n是向量的维度
	m := len(dataSet)                                         // m是数据集中向量个数
	c := math.Pow(2, 32) - 5                                  // c是大素数
	fpRand := make([]float64, k)                              // k向量的维度
	rander := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	for i := 0; i < k; i++ {
		fpRand[i] = float64(rander.Intn(20) - 10)
	}
	var hashFuncs [][]*AB
	hashCollect := make([][][]float64, l)
	for i := 0; i < l; i++ { //有l个向量组 每组有k个Ab组成
		if hashCollect[i] == nil {
			hashCollect[i] = make([][]float64, m)
		}
		e2LSHFamily := GenE2LSHFamily(n, k, r)
		hashFuncs = append(hashFuncs, e2LSHFamily)
		for dataIndex := 0; dataIndex < m; dataIndex++ { // m是数据集中向量个数
			hashVals := GenHashVals(e2LSHFamily, dataSet[dataIndex], r) // 向量内积a*v + b(随机数) / r 有k组
			hashCollect[i] = append(hashCollect[i], hashVals)
			fp := H2(hashVals, fpRand, k, int(c))
			index := int(math.Abs(float64(fp % tableSize)))

			if hashTable[index] == nil {
				hashTable[index] = make(map[int][]int)
				if hashTable[index][fp] == nil {
					hashTable[index][fp] = make([]int, 0)
				}
			}
			hashTable[index][fp] = append(hashTable[index][fp], dataIndex)
		}

	}
	return hashTable, hashFuncs, fpRand, hashCollect
}
