/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package index  index Pretreatment
package index

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

/**
 * @Author peng
 * @Description Select topN keyWords
 * @Date 16:32 2022/6/28
 * @Param  filepath N
 * @return slice of keywords
 **/

/**
 * @Author peng
 * @Description InitialKey
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

// InitialKey Select topN keyWords for every paper
func InitialKey(path string, numberOfFile, topN int) [][]string {
	frequentWordSet := make([][]string, numberOfFile)
	for i := 1; i < numberOfFile+1; i++ {
		fileName := path + strconv.Itoa(i) + ".txt"
		buf, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		statisticTimes := make(map[string]int)
		wordsLength := strings.Fields(string(buf))
		for _, word := range wordsLength {
			wordL := strings.ToLower(word)
			_, ok := statisticTimes[wordL]
			if ok {
				statisticTimes[wordL] = statisticTimes[wordL] + 1
			} else {
				statisticTimes[wordL] = 1
			}
		}

		uniqueWords := make([]string, topN)
		for w := range statisticTimes {
			uniqueWords = append(uniqueWords, w)
		}
		sort.Slice(uniqueWords, func(i, j int) bool {
			s, t := uniqueWords[i], uniqueWords[j]
			return statisticTimes[s] > statisticTimes[t] || statisticTimes[s] == statisticTimes[t] && s < t
		})
		frequentWordSet[i-1] = uniqueWords[:topN]
	}
	return frequentWordSet
}
