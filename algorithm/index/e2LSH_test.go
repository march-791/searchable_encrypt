/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package index

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"gonum.org/v1/gonum/mat"
)

func TestGenPara(t *testing.T) {
	rander := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	a := GenPara(10, 100, rander)
	fmt.Println(a.A)
	fmt.Println(a.B)
}

func TestGenE2LSHFamily(t *testing.T) {
	a := GenE2LSHFamily(4, 5, 100)
	for index, value := range a {
		fmt.Print(index)
		fmt.Print(value.A)
		fmt.Println(value.B)
	}
}

func TestGenHashVals(t *testing.T) {
	ab := GenE2LSHFamily(5, 5, 10)
	for _, value := range ab {
		var sum float64 = 0
		for i := 0; i < 5; i++ {
			sum += value.A.AtVec(i)
		}
		fmt.Println(sum + value.B)
	}
	fmt.Println()
	v := mat.NewVecDense(5, nil)
	for i := 0; i < 5; i++ {
		v.SetVec(i, 1)
	}
	sliceHash := GenHashVals(ab, v, 1)
	for _, value := range sliceHash {
		fmt.Println(value)
	}
}

func TestH2(t *testing.T) {
	ab := GenE2LSHFamily(5, 5, 10)
	v := mat.NewVecDense(5, nil)
	for i := 0; i < 5; i++ {
		v.SetVec(i, 1)
	}
	sliceHash := GenHashVals(ab, v, 1)
	fpRand := make([]float64, 5)
	rander := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	for i := 0; i < 5; i++ {
		fpRand[i] = float64(rander.Intn(20) - 10)
		fmt.Println(fpRand[i])
		fmt.Println(sliceHash[i])
		fmt.Println(fpRand[i] * sliceHash[i])
	}
	var sum float64
	for i := 0; i < 5; i++ {
		sum = sum + (fpRand[i] * sliceHash[i])
	}
	c := math.Pow(2, 32) - 5
	result := int(sum) % int(c)
	result2 := H2(sliceHash, fpRand, 5, int(c))
	fmt.Println(result)
	fmt.Println(result2)
}

func TestE2LSH(t *testing.T) {
	//n := 5          // n是向量的维度
	//k := 20         // hash表的维度
	//l := 5          // 有l个 [h1, ...hk]这样的向量组  [[h1, ...hk], [h1, ..hk], ..., [h1, ...hk]]
	//tableSize := 20 // hash表的维度
	//r := 1          // r是ab中随机数b的范围[0,r)
	//
	////构造dataSet
	//fileName := "../test/test_data.csv"
	//fs, err := os.Open(fileName)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer fs.Close()
	//
	//fs1, _ := os.Open(fileName)
	//r1 := csv.NewReader(fs1)
	//content, err := r1.ReadAll()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//dataSet := make([]*mat.VecDense, len(content))
	//for index, row := range content {
	//	node := dataSet[index]
	//	if node == nil {
	//		dataSet[index] = mat.NewVecDense(len(content[index]), nil)
	//		for i := 0; i < len(content[index]); i++ {
	//			float64, _ := strconv.ParseFloat(row[i], 64)
	//			dataSet[index].SetVec(i, float64)
	//		}
	//	}
	//}
	//
	////验证数据正确性
	//for _, value := range dataSet {
	//	for i := 0 ;i < value.Len(); i++ {
	//		fmt.Print(value.AtVec(i))
	//		fmt.Print(",")
	//	}
	//	fmt.Println()
	//}
	//
	//// 构造查询*mat.VecDense
	//query := mat.NewVecDense(n, nil)
	//query.SetVec(0, -2.7769)
	//query.SetVec(1, -5.6967)
	//query.SetVec(2, 5.9179)
	//query.SetVec(3, 0.37671)
	//query.SetVec(4, 1)

	//	hashTable, hashFuncGroups, fpRand, _ := E2LSH(dataSet, k, l, r, tableSize)

	//C := math.Pow(2, 32) - 5
	//mapSet := util.NewSet()
	//for _, hashFuncGroup := range hashFuncGroups {
	//	queryFp := H2(GenHashVals(hashFuncGroup, query, r), fpRand, k, int(C))
	//	queryIndex := int(math.Abs(float64(queryFp % tableSize)))
	//	if hashTable[queryIndex][queryFp] != nil {
	//		for _, value := range hashTable[queryIndex][queryFp] {
	//			mapSet.Add(value)
	//		}
	//	}
	//}
	//fmt.Println()
	//indexs := mapSet.GetItems()
	//rander := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec
	//randDis := 0.0
	//dis := 0.0
	//for _, value := range indexs {
	//	intValue, ok := value.(int)
	//
	//	if ok {
	//		randDis += euclideanDistance(dataSet[rander.Intn(1372)], query)
	//		dis += euclideanDistance(dataSet[intValue], query)
	//		fmt.Println(euclideanDistance(dataSet[intValue], query))
	//	}
	//}
	//fmt.Println(randDis)
	//fmt.Println(dis)
}

func euclideanDistance(n1, n2 *mat.VecDense) float64 {
	result := 0.0
	for i := 0; i < n1.Len() && i < n2.Len(); i++ {
		result += math.Pow(n1.AtVec(i)-n2.AtVec(i), 2)
	}
	return math.Sqrt(result)
}
