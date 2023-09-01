/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package dualcode Package decode dual function including chinese
package dualcode

import (
	"strings"

	"gonum.org/v1/gonum/mat"
	"gotest/algorithm/keygen"
	"gotest/algorithm/matrix"
)

// dual function
// input string
// input string include three case
// 1.only letters
// 2.letters and number
// 3.chinese
// output vector Composed of 0 and 1

// Main processes
// generated vector is composed of 0
// traverse every character of a word
// adjacent characters are mapped to a position in vector

// ASCIIS define border
const (
	ASCIIOFNUMZORE = 48
	ASCIIOFNINE    = 57

	ASCIIOFCAPITALA = 65
	ASCIIOFCAPITALZ = 90

	ASCIIOFSMALLA = 97
	ASCIIOFSMALLZ = 122
)

/**
 * @Author peng
 * @Description  dual function OnlyLetters
 * @Date 16:30 2022/6/28
 * @Param  string, m is the len of vector
 * @return vector, index is a slice made of location with 1 element
 **/

// DualWordCodeOnlyLetters dual function OnlyLetters
func DualWordCodeOnlyLetters(str string, m int, vector *mat.VecDense) (vector1 *mat.VecDense, index []int) {
	if len(str) == 0 || m <= 0 {
		return vector, nil
	}

	if !checkStrRange(str, keygen.LETTER) {
		return vector, nil
	}

	if vector == nil {
		vector = matrix.GenVecDense(m*m, 0)
	}

	// deal string
	str = strings.Trim(str, " ")
	str = strings.ToLower(str)
	strLen := len(str)
	index = make([]int, strLen)

	//process letter
	for i := 0; i < strLen; i++ {
		var currentIndex int
		if i == strLen-1 {
			currentIndex = int(str[i]-ASCIIOFSMALLA)*m + int(str[0]-ASCIIOFSMALLA)
		} else {
			currentIndex = int(str[i]-ASCIIOFSMALLA)*m + int(str[i+1]-ASCIIOFSMALLA)
		}

		if currentIndex < 0 || currentIndex > m*m {
			return vector, nil
		}

		index[i] = currentIndex
		vector.SetVec(currentIndex, 1)
	}
	return vector, index
}

/**
 * @Author peng
 * @Description  dual function OnlyLetters slice
 * @Date 16:30 2022/6/28
 * @Param  string slice, m is the len of vector
 * @return vector, index is a slice made of location with 1 element
 **/

// DualWordCodeSliceOnlyLetters dual function process OnlyLetters slice
func DualWordCodeSliceOnlyLetters(strs []string, m int) (vector *mat.VecDense) {
	if len(strs) == 0 || m <= 0 {
		return nil
	}

	vector = matrix.GenVecDense(m*m, 0)
	for _, str := range strs {
		DualWordCodeOnlyLetters(str, m, vector)
	}

	return vector
}

/**
 * @Author peng
 * @Description  dual function Letters And Number
 * @Date 16:30 2022/6/28
 * @Param  string, m is the len of vector
 * @return vector, index is a slice made of location with 1 element
 **/

// DualWordCodeLettersAndNumber dual function Letters And Number
func DualWordCodeLettersAndNumber(str string, m int, vector *mat.VecDense) (vector1 *mat.VecDense, index []int) {
	if len(str) == 0 || m <= 0 {
		return vector, nil
	}

	if !checkStrRange(str, keygen.LETTERANDNURBER) {
		return vector, nil
	}

	//preparation
	str = strings.Trim(str, " ")
	strLen := len(str)
	if vector == nil {
		vector = matrix.GenVecDense(m*m, 0)
	}
	index = make([]int, strLen)
	nums := make([]int, strLen)

	//map letter and num to [97,132]
	for i := 0; i < strLen; i++ {
		num := int(str[i])
		if num >= ASCIIOFNUMZORE && num <= ASCIIOFNINE {
			num = num + (ASCIIOFSMALLZ - ASCIIOFNUMZORE + 1)
			nums[i] = num
		} else if num >= ASCIIOFCAPITALA && num <= ASCIIOFCAPITALZ {
			num = num + (ASCIIOFSMALLA - ASCIIOFCAPITALA)
			nums[i] = num
		} else if num >= ASCIIOFSMALLA && num <= ASCIIOFSMALLZ {
			nums[i] = num
		} else {
			return nil, nil
		}
	}

	// encoded
	var currentIndex int
	for i := 0; i < strLen; i++ {
		if i == strLen-1 {
			currentIndex = (nums[i]-ASCIIOFSMALLA)*m + (nums[0] - ASCIIOFSMALLA)
		} else {
			currentIndex = (nums[i]-ASCIIOFSMALLA)*m + (nums[i+1] - ASCIIOFSMALLA)
		}
		if currentIndex < 0 || currentIndex > m*m {
			return vector, nil
		}
		index[i] = currentIndex
		vector.SetVec(currentIndex, 1)
	}
	return vector, index
}

/**
 * @Author peng
 * @Description  dual slice Letters And Number
 * @Date 16:30 2022/6/28
 * @Param  string, m is the len of vector
 * @return vector, index is a slice made of location with 1 element
 **/

// DualWordCodeSliceLettersAndNumber dual slice Letters And Number
func DualWordCodeSliceLettersAndNumber(strs []string, m int) (vector *mat.VecDense) {
	if strs == nil || m <= 0 {
		return nil
	}
	vector = matrix.GenVecDense(m*m, 0)
	for _, str := range strs {
		DualWordCodeLettersAndNumber(str, m, vector)
	}
	return vector
}

/**
 * @Author peng
 * @Description  dual function Chinese
 * @Date 16:30 2022/6/28
 * @Param  string, m is the len of vector
 * @return vector, index is a slice made of location with 1 element
 **/

// DualWordCodeChinese dual function Chinese
func DualWordCodeChinese(str string, m int, vector *mat.VecDense) (vector1 *mat.VecDense, index []int) {
	if len(str) == 0 || m <= 0 {
		return vector, nil
	}

	if !checkStrRange(str, keygen.CHINESE) {
		return vector, nil
	}

	str = strings.Trim(str, " ")
	strRune := []rune(str)
	strLen := len(strRune)
	if vector == nil {
		vector = matrix.GenVecDense(m*m, 0)
	}
	index = make([]int, strLen)

	for i, value := range strRune {
		index[i] = int(value) % (m * m)
		vector.SetVec(index[i], 1)
	}
	return vector, index
}

/**
 * @Author peng
 * @Description  dual slice function Chinese
 * @Date 16:30 2022/6/28
 * @Param  string, m is the len of vector
 * @return vector, index is a slice made of location with 1 element
 **/

// DualWordCodeSliceChinese dual slice Chinese
func DualWordCodeSliceChinese(strs []string, m int) (vector *mat.VecDense) {
	if strs == nil {
		return nil
	}
	vector = matrix.GenVecDense(m*m, 0)
	for _, str := range strs {
		DualWordCodeChinese(str, m, vector)
	}
	return vector
}

/*
*
* @Author peng

* @Description  check number and only number

* //TODO  check chinese range

* @Date 2022/11/5 15:44

* @Param

* @return
*
 */
func checkStrRange(str string, expectRange keygen.SEARCHSCOPE) bool {
	if len(str) == 0 {
		return false
	}

	str = strings.Trim(str, " ")

	if len(str) == 0 {
		return false
	}
	//TODO  check chinese range
	if expectRange == keygen.CHINESE {
		return true
	} else if expectRange == keygen.LETTER {
		for _, value := range str {
			num := int(value)
			// CAPITAL SMALL
			if !(num >= ASCIIOFCAPITALA && num <= ASCIIOFCAPITALZ) && !(num >= ASCIIOFSMALLA && num <= ASCIIOFSMALLZ) {
				return false
			}
		}
		return true
	} else if expectRange == keygen.LETTERANDNURBER {
		for _, value := range str {
			num := int(value)
			// CAPITAL SMALL
			if !(num >= ASCIIOFCAPITALA && num <= ASCIIOFCAPITALZ) &&
				!(num >= ASCIIOFSMALLA && num <= ASCIIOFSMALLZ) &&
				!(num >= ASCIIOFNUMZORE && num <= ASCIIOFNINE) {
				return false
			}
		}
		return true
	}
	return false
}
