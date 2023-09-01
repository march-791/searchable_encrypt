/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package search search
package search

import (
	"fmt"
	"sort"

	"gonum.org/v1/gonum/mat"
)

/**
 * @Author pengpeng
 * @Description
 * @Date 16:36 2022/6/28
 * @Param index trapdoor
 * @return Similarity of each article
 **/

type Index []*mat.VecDense

type IndexAndID struct {
	Id string
	Index
	Date string
}

type Result struct {
	Id    string
	Score float64
	Date  string
}

type Results []*Result

// Len 实现sort.Interface接口取元素数量方法
/**
 * @Author peng
 * @Description Len
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
func (s Results) Len() int {
	return len(s)
}

// Swap 实现sort.Interface接口交换元素方法
/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
func (s Results) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less 实现sort.Interface接口比较元素方法
/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/
func (s Results) Less(i, j int) bool {
	if s[i].Score != s[j].Score {
		return s[i].Score > s[j].Score
	}
	return s[i].Id < s[j].Id
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// Search search
func Search(indexs []*IndexAndID, trap Index) (results Results) {
	if indexs == nil || trap == nil {
		return nil
	}
	for _, value := range indexs {
		result := new(Result)
		result.Id = value.Id
		result.Score = mat.Dot(value.Index[0], trap[0]) + mat.Dot(value.Index[1], trap[1])
		result.Date = value.Date
		results = append(results, result)
	}
	sort.Sort(results)
	return results
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// SearchDemo Search Demo
func SearchDemo(indexes [][]*mat.VecDense, trap []*mat.VecDense) {
	if indexes == nil || trap == nil {
		return
	}
	result := make([]float64, len(indexes))
	for index, value := range indexes {
		I0 := value[0]
		I1 := value[1]
		result[index] = mat.Dot(I0, trap[0]) + mat.Dot(I1, trap[1])
	}
	//sort.Slice(result, func(i, j int) bool {
	//	if result[i] > result[j] {
	//		return true
	//	}
	//	return false
	//})
	fmt.Println(result)
}
