/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package index

import (
	"fmt"
	"testing"
)

// TestInitialKey test
func TestInitialKey(t *testing.T) {
	frequentWordSet := InitialKey("../../testdata/", 5, 10)
	for index, value := range frequentWordSet {
		fmt.Print(index)
		fmt.Println(value)
	}
}
