/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package dualcode test package dual
package dualcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gotest/algorithm/keygen"
)

// TestDualWordCodeOnlyLetters Test word coding
func TestDualWordCodeOnlyLetters(t *testing.T) {
	a := assert.New(t)

	vector, index := DualWordCodeOnlyLetters(" AAZZaazz ", keygen.LETTERLENS, nil)

	a.NotNil(vector)
	a.NotNil(index)

	//[0 25 675 650 0 25 675 650]
	a.Equal(8, len(index))
	a.Equal(0, index[0])
	a.Equal(keygen.LETTERLENS-1, index[1])
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS-1, index[2])
	a.Equal(keygen.LETTERLENS*(keygen.LETTERLENS-1), index[3])
	a.Equal(0, index[4])
	a.Equal(keygen.LETTERLENS-1, index[5])
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS-1, index[6])
	a.Equal(keygen.LETTERLENS*(keygen.LETTERLENS-1), index[7])

	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS, vector.Len())
	a.Equal(float64(1), vector.AtVec(0))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS-1))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*(keygen.LETTERLENS-1)))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERLENS-1))

	vector1, index1 := DualWordCodeOnlyLetters(" A AZZaazz123 ", keygen.LETTERLENS, nil)

	a.Nil(vector1)
	a.Nil(index1)

	vector2, index2 := DualWordCodeOnlyLetters("  ", keygen.LETTERLENS, nil)

	a.Nil(vector2)
	a.Nil(index2)

	vector3, index3 := DualWordCodeOnlyLetters(" A A中文", keygen.LETTERLENS, nil)

	a.Nil(vector3)
	a.Nil(index3)

	vector4, index4 := DualWordCodeOnlyLetters(" A12A", keygen.LETTERLENS, nil)

	a.Nil(vector4)
	a.Nil(index4)

	vector5, index5 := DualWordCodeOnlyLetters(" AA12", keygen.LETTERLENS, nil)

	a.Nil(vector5)
	a.Nil(index5)
}

// TestDualWordCodeSliceOnlyLetters Test word slice coding
func TestDualWordCodeSliceOnlyLetters(t *testing.T) {
	a := assert.New(t)

	vector, index := DualWordCodeOnlyLetters(" AAZZaazz ", keygen.LETTERLENS, nil)

	a.NotNil(vector)
	a.NotNil(index)

	//[0 25 675 650 0 25 675 650]
	a.Equal(8, len(index))
	a.Equal(0, index[0])
	a.Equal(keygen.LETTERLENS-1, index[1])
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS-1, index[2])
	a.Equal(keygen.LETTERLENS*(keygen.LETTERLENS-1), index[3])
	a.Equal(0, index[4])
	a.Equal(keygen.LETTERLENS-1, index[5])
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS-1, index[6])
	a.Equal(keygen.LETTERLENS*(keygen.LETTERLENS-1), index[7])

	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS, vector.Len())
	a.Equal(float64(1), vector.AtVec(0))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS-1))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*(keygen.LETTERLENS-1)))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERLENS-1))

	vector1, index1 := DualWordCodeOnlyLetters(" A AZZaazz123 ", keygen.LETTERLENS, nil)

	a.Nil(vector1)
	a.Nil(index1)

	vector2, index2 := DualWordCodeOnlyLetters("  ", keygen.LETTERLENS, nil)

	a.Nil(vector2)
	a.Nil(index2)

	vector3, index3 := DualWordCodeOnlyLetters(" A A中文", keygen.LETTERLENS, nil)

	a.Nil(vector3)
	a.Nil(index3)

	vector4, index4 := DualWordCodeOnlyLetters(" A12A", keygen.LETTERLENS, nil)

	a.Nil(vector4)
	a.Nil(index4)

	vector5, index5 := DualWordCodeOnlyLetters(" AA12", keygen.LETTERLENS, nil)

	a.Nil(vector5)
	a.Nil(index5)

	strs := make([]string, 2)
	strs[0] = " AAZZ "
	strs[1] = " aazz "
	vector6 := DualWordCodeSliceOnlyLetters(strs, keygen.LETTERLENS)

	a.NotNil(vector6)
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS, vector6.Len())

	a.Equal(float64(1), vector.AtVec(0))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS-1))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*(keygen.LETTERLENS-1)))
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERLENS-1))
	a.Equal(vector, vector6)

	strs2 := make([]string, 2)
	strs2[0] = " aa "
	strs2[1] = " zz "
	vector7 := DualWordCodeSliceOnlyLetters(strs2, keygen.LETTERLENS)

	a.NotNil(vector7)
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS, vector7.Len())

	a.Equal(float64(1), vector7.AtVec(0))
	a.Equal(float64(1), vector7.AtVec(keygen.LETTERLENS*keygen.LETTERLENS-1))

	strs3 := make([]string, 2)
	strs3[0] = " aaa "
	strs3[1] = " 1234 "
	vector8 := DualWordCodeSliceOnlyLetters(strs3, keygen.LETTERLENS)

	a.NotNil(vector8)
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS, vector8.Len())

	a.Equal(float64(1), vector8.AtVec(0))
	a.Equal(float64(0), vector8.AtVec(keygen.LETTERLENS*keygen.LETTERLENS-1))

	strs4 := make([]string, 2)
	strs4[0] = " 中文 "
	strs4[1] = " 123 "
	vector9 := DualWordCodeSliceOnlyLetters(strs4, keygen.LETTERLENS)

	a.NotNil(vector9)
	a.Equal(keygen.LETTERLENS*keygen.LETTERLENS, vector9.Len())

	a.Equal(float64(0), vector9.AtVec(0))
	a.Equal(float64(0), vector9.AtVec(keygen.LETTERLENS*keygen.LETTERLENS-1))

}

// DualWordCodeLettersAndNumber Test word slice coding
func TestDualWordCodeLettersAndNumber(t *testing.T) {

	a := assert.New(t)

	vector, index := DualWordCodeLettersAndNumber(" AAZZaazz00990 ", keygen.LETTERANDNUMLENS, nil)

	a.NotNil(vector)
	a.NotNil(index)

	a.Equal(13, len(index))

	//[0 25 925 900 0 25 925 926 962 971 1295 1286 936]
	a.Equal(0, index[0])                                                                              //aa 0
	a.Equal(keygen.LETTERLENS-1, index[1])                                                            //az 25
	a.Equal((keygen.LETTERLENS-1)*keygen.LETTERANDNUMLENS+keygen.LETTERLENS-1, index[2])              //zz 925
	a.Equal(keygen.LETTERANDNUMLENS*(keygen.LETTERLENS-1), index[3])                                  //za 900
	a.Equal(0, index[4])                                                                              //aa 0
	a.Equal(keygen.LETTERLENS-1, index[5])                                                            //az 25
	a.Equal((keygen.LETTERLENS-1)*keygen.LETTERANDNUMLENS+keygen.LETTERLENS-1, index[6])              //zz 925
	a.Equal((keygen.LETTERLENS-1)*keygen.LETTERANDNUMLENS+keygen.LETTERLENS, index[7])                //z0 926
	a.Equal(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS, index[8])                    //00 962
	a.Equal(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERANDNUMLENS-1, index[9])            //09 971
	a.Equal(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERANDNUMLENS-1, index[10]) //99 1295
	a.Equal(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERLENS, index[11])         //90 1286
	a.Equal(keygen.LETTERLENS*keygen.LETTERANDNUMLENS, index[12])                                     //0a 936

	a.Equal(keygen.LETTERANDNUMLENS*keygen.LETTERANDNUMLENS, vector.Len())
	a.Equal(float64(1), vector.AtVec(0))                                                                             //0
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS-1))                                                           //25
	a.Equal(float64(1), vector.AtVec((keygen.LETTERLENS-1)*keygen.LETTERANDNUMLENS+keygen.LETTERLENS-1))             //925
	a.Equal(float64(1), vector.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERLENS-1)))                                 //900
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS))                   //926
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS))                   //962
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERANDNUMLENS-1))           //971
	a.Equal(float64(1), vector.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERANDNUMLENS-1)) //1295
	a.Equal(float64(1), vector.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERLENS))         //1286
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS))                                     //936
}

// TestDualWordCodeSliceLettersAndNumber
func TestDualWordCodeSliceLettersAndNumber(t *testing.T) {
	a := assert.New(t)
	strs := make([]string, 1)
	strs[0] = " AAZZaazz009900 "
	vector := DualWordCodeSliceLettersAndNumber(strs, keygen.LETTERANDNUMLENS)

	a.Equal(keygen.LETTERANDNUMLENS*keygen.LETTERANDNUMLENS, vector.Len())
	a.Equal(float64(1), vector.AtVec(0))                                                                             //0
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS-1))                                                           //25
	a.Equal(float64(1), vector.AtVec((keygen.LETTERLENS-1)*keygen.LETTERANDNUMLENS+keygen.LETTERLENS-1))             //925
	a.Equal(float64(1), vector.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERLENS-1)))                                 //900
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS))                   //926
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS))                   //962
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERANDNUMLENS-1))           //971
	a.Equal(float64(1), vector.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERANDNUMLENS-1)) //1295
	a.Equal(float64(1), vector.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERLENS))         //1286
	a.Equal(float64(1), vector.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS))                                     //936

	strs1 := make([]string, 2)
	strs1[0] = " AAZZaazz00990 "
	strs1[1] = " AAZZaazz0099001 "
	vector1 := DualWordCodeSliceLettersAndNumber(strs, keygen.LETTERANDNUMLENS)

	a.Equal(keygen.LETTERANDNUMLENS*keygen.LETTERANDNUMLENS, vector.Len())
	a.Equal(float64(1), vector1.AtVec(0))                                                                             //0
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERLENS-1))                                                           //25
	a.Equal(float64(1), vector1.AtVec((keygen.LETTERLENS-1)*keygen.LETTERANDNUMLENS+keygen.LETTERLENS-1))             //925
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERLENS-1)))                                 //900
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS))                   //926
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERLENS))                   //962
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS+keygen.LETTERANDNUMLENS-1))           //971
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERANDNUMLENS-1)) //1295
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERANDNUMLENS*(keygen.LETTERANDNUMLENS-1)+keygen.LETTERLENS))         //1286
	a.Equal(float64(1), vector1.AtVec(keygen.LETTERLENS*keygen.LETTERANDNUMLENS))                                     //936

}

// TestDualWordCodeChinese
func TestDualWordCodeChinese(t *testing.T) {
	a := assert.New(t)

	//string to rune
	vector, index := DualWordCodeChinese("彭鹏世界", keygen.CHINESELENS, nil)

	a.NotNil(vector)
	a.NotNil(index)

	//[1101 351 550 220] % 1296
	a.Equal(keygen.CHINESELENS*keygen.CHINESELENS, vector.Len())
	a.Equal(1101, index[0]) //彭 1101
	a.Equal(351, index[1])  //鹏 351
	a.Equal(550, index[2])  //世 550
	a.Equal(220, index[3])  //界 220

	a.Equal(keygen.CHINESELENS*keygen.CHINESELENS, vector.Len())
	a.Equal(float64(1), vector.AtVec(1101)) //彭 1101
	a.Equal(float64(1), vector.AtVec(351))  //鹏 351
	a.Equal(float64(1), vector.AtVec(550))  //世 550
	a.Equal(float64(1), vector.AtVec(220))  //界 220
}

// TestDualWordCodeSliceChinese
func TestDualWordCodeSliceChinese(t *testing.T) {
	a := assert.New(t)

	strs := make([]string, 2)
	strs[0] = "彭鹏"
	strs[1] = "世界"
	vector := DualWordCodeSliceChinese(strs, keygen.CHINESELENS)

	a.Equal(keygen.CHINESELENS*keygen.CHINESELENS, vector.Len())
	a.Equal(float64(1), vector.AtVec(1101)) //彭 1101
	a.Equal(float64(1), vector.AtVec(351))  //鹏 351
	a.Equal(float64(1), vector.AtVec(550))  //世 550
	a.Equal(float64(1), vector.AtVec(220))  //界 220
}

/**
* @Author peng

* @Description

* //

* @Date 2022/11/7 10:25

* @Param

* @return
**/

// TestCheckStrRange Test str range
func TestCheckStrRange(t *testing.T) {
	a := assert.New(t)

	onlyletter := "abAZ"
	letterAndNum := "09AZaz"
	chinese := "测试中文"
	chineseAndNum := "测试中文09"
	chineseAndNumLetter := "测试中文09za"

	//
	a.True(checkStrRange(onlyletter, keygen.LETTER))
	a.True(checkStrRange(onlyletter, keygen.LETTERANDNURBER))
	//a.False(checkStrRange(onlyletter, keygen.CHINESE))//todo chinese logic

	a.False(checkStrRange(letterAndNum, keygen.LETTER))
	a.True(checkStrRange(letterAndNum, keygen.LETTERANDNURBER))
	//a.False(checkStrRange(letterAndNum, keygen.CHINESE))

	//a.False(checkStrRange(chinese, keygen.LETTER))
	//a.False(checkStrRange(chinese, keygen.LETTERANDNURBER))
	a.True(checkStrRange(chinese, keygen.CHINESE))

	a.False(checkStrRange(chineseAndNum, keygen.LETTER))
	a.False(checkStrRange(chineseAndNum, keygen.LETTERANDNURBER))
	//a.False(checkStrRange(chineseAndNum, keygen.CHINESE))

	a.False(checkStrRange(chineseAndNumLetter, keygen.LETTER))
	a.False(checkStrRange(chineseAndNumLetter, keygen.LETTERANDNURBER))
	//a.False(checkStrRange(chineseAndNumLetter, keygen.CHINESE))
}
