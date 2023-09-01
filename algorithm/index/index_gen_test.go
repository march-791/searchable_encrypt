/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

// Package index test
package index

import (
	"fmt"
	"testing"

	"gotest/algorithm/keygen"
)

// TestBuildIndex test Index construction input filepath, Cipher
func TestBuildIndex(t *testing.T) {
	I := BuildIndexDemo("../../testdata/", 5, 10, keygen.GenCipher(keygen.LETTER))
	for _, value := range I {
		fmt.Println(value[0])
		fmt.Println(value[1])
		fmt.Println("============================================")
	}
}
