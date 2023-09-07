// /*/*
// Copyright (C) BABEC. All rights reserved.
// Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
// All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0
//
// */
//
// // Package search test
package search

//
//import (
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"gonum.org/v1/gonum/mat"
//	"strconv"
//	"strings"
//	"testing"
//
//	"gotest/algorithm/index"
//	"gotest/algorithm/keygen"
//	"gotest/algorithm/trapdoor"
//	"gotest/tool"
//)
//
//func TestSearch(t *testing.T) {
//	key := keygen.GenCipher(keygen.LETTER)                   // 密钥生成
//	I := index.BuildIndexDemo("../../testdata/", 5, 10, key) // //索引构建
//	strs := make([]string, 2)
//	strs[0] = "you"
//	strs[1] = "am"
//	T := trapdoor.GenTrapdoor(strs, key) //陷门构建
//	SearchDemo(I, T)                     //搜索
//}
//
//func TestIndexSearch(t *testing.T) {
//	key := keygen.GenCipher(keygen.LETTER)
//	I := index.BuildIndexDemo("../../testdata/", 5, 10, key)
//	strs := make([]string, 2)
//	strs[0] = "my"  //my
//	strs[1] = "nor" //nor
//	T := trapdoor.GenTrapdoor(strs, key)
//	indexAndID := make([]*IndexAndID, 5)
//	for index2, i := range I {
//		indexAndID[index2] = &IndexAndID{
//			Id:    strconv.Itoa(index2),
//			Index: i,
//		}
//	}
//	results := Search(indexAndID, T)
//	date1, err1 := json.Marshal(results)
//	if err1 != nil {
//		return
//	}
//	fmt.Println(string(date1))
//}
//
//func TestSDKSearch(t *testing.T) {
//	indexLen := 5
//	keywords := make([]string, indexLen)
//	keywords[0] = "he as it no of at in you am on"
//	keywords[1] = "to am at he on she so an by do"
//	keywords[2] = "of in he so do and as no or to"
//	keywords[3] = "as in on one he at my nor of she"
//	//keywords[4] = "as in on one he at my nor of she"
//
//	indexAndID := make([]*IndexAndID, indexLen)
//
//	key := keygen.GenCipher(keygen.LETTER)
//	for i, value := range keywords {
//		keyWord := strings.Split(value, " ")
//		indexAndID[i] = &IndexAndID{
//			Id:    strconv.Itoa(i),
//			Index: index.BuildIndex(keyWord, key),
//		}
//		byteIndex, err := tool.Encode(indexAndID[i].Index)
//		if err != nil {
//			return
//		}
//		fmt.Println(i)
//		fmt.Println("======")
//		fmt.Println(hex.EncodeToString(byteIndex))
//		fmt.Println("-------------")
//		fmt.Println("index")
//		IndexByte, err := hex.DecodeString(hex.EncodeToString(byteIndex))
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		fmt.Println(IndexByte)
//		var I []*mat.VecDense
//		err = tool.Decode(IndexByte, &I)
//		fmt.Println(I[1].MarshalBinary())
//	}
//
//	strs := make([]string, 2)
//	strs[0] = "youu" //my
//	strs[1] = "amam" //nor
//	T := trapdoor.GenTrapdoor(strs, key)
//	byteTrapdoor, err := tool.Encode(T)
//	if err != nil {
//		return
//	}
//	fmt.Println("++++++++++")
//	fmt.Println(hex.EncodeToString(byteTrapdoor))
//	results := Search(indexAndID, T)
//	date1, err1 := json.Marshal(results)
//	if err1 != nil {
//		return
//	}
//	fmt.Println(string(date1))
//	resultStr := "16ff8102010107526573756c747301ff820001ff800000327f030102ff8000010401024964010c00010553636f7265010800010444617465010c00010846696c654e616d65010c000000ffedff82000101013001f857bbc948caa0604001ffd1687474703a2f2f3132332e35362e3138352e3130363a383038302f646f776e6c6f61643f746f6b656e3d65794a68624763694f694a49557a49314e694973496e523563434936496b705856434a392e65794a515958526f496a6f69625856736157356e5846787465575a706247567a584677694c434a476157786c546d46745a534936496a457564486830496977695a586877496a6f784e6a6b7a4f4467304f44637966512e725a5657365a6756334d755851586159744259515f42654b575f7342656857796c4a2d4e67352d565a304d0105312e74787400"
//	resultBytes2, err := hex.DecodeString(resultStr)
//	var results2 Results
//	err = tool.Decode(resultBytes2, &results2)
//	fmt.Println(results2[0])
//}
//*/
