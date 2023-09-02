/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package search test
package search

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strconv"
	"strings"
	"testing"

	"gotest/algorithm/index"
	"gotest/algorithm/keygen"
	"gotest/algorithm/trapdoor"
	"gotest/tool"
)

func TestSearch(t *testing.T) {
	key := keygen.GenCipher(keygen.LETTER)                   // 密钥生成
	I := index.BuildIndexDemo("../../testdata/", 5, 10, key) // //索引构建
	strs := make([]string, 2)
	strs[0] = "you"
	strs[1] = "am"
	T := trapdoor.GenTrapdoor(strs, key) //陷门构建
	SearchDemo(I, T)                     //搜索
}

func TestIndexSearch(t *testing.T) {
	key := keygen.GenCipher(keygen.LETTER)
	I := index.BuildIndexDemo("../../testdata/", 5, 10, key)
	strs := make([]string, 2)
	strs[0] = "my"  //my
	strs[1] = "nor" //nor
	T := trapdoor.GenTrapdoor(strs, key)
	indexAndID := make([]*IndexAndID, 5)
	for index2, i := range I {
		indexAndID[index2] = &IndexAndID{
			Id:    strconv.Itoa(index2),
			Index: i,
		}
	}
	results := Search(indexAndID, T)
	date1, err1 := json.Marshal(results)
	if err1 != nil {
		return
	}
	fmt.Println(string(date1))
}

func TestSDKSearch(t *testing.T) {
	indexLen := 5
	keywords := make([]string, indexLen)
	keywords[0] = "he as it no of at in you am on"
	keywords[1] = "to am at he on she so an by do"
	keywords[2] = "of in he so do and as no or to"
	keywords[3] = "as in on one he at my nor of she"
	//keywords[4] = "as in on one he at my nor of she"

	indexAndID := make([]*IndexAndID, indexLen)

	key := keygen.GenCipher(keygen.LETTER)
	for i, value := range keywords {
		keyWord := strings.Split(value, " ")
		indexAndID[i] = &IndexAndID{
			Id:    strconv.Itoa(i),
			Index: index.BuildIndex(keyWord, key),
		}
		byteIndex, err := tool.Encode(indexAndID[i].Index)
		if err != nil {
			return
		}
		fmt.Println(i)
		fmt.Println("======")
		fmt.Println(hex.EncodeToString(byteIndex))
		fmt.Println("-------------")
		fmt.Println("index")
		IndexByte, err := hex.DecodeString(hex.EncodeToString(byteIndex))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(IndexByte)
		var I []*mat.VecDense
		err = tool.Decode(IndexByte, &I)
		fmt.Println(I[1].MarshalBinary())
	}

	strs := make([]string, 2)
	strs[0] = "youu" //my
	strs[1] = "amam" //nor
	T := trapdoor.GenTrapdoor(strs, key)
	byteTrapdoor, err := tool.Encode(T)
	if err != nil {
		return
	}
	fmt.Println("++++++++++")
	fmt.Println(hex.EncodeToString(byteTrapdoor))
	results := Search(indexAndID, T)
	date1, err1 := json.Marshal(results)
	if err1 != nil {
		return
	}
	fmt.Println(string(date1))
}
