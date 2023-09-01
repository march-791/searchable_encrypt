/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package trapdoor

import (
	"fmt"
	"testing"

	"gotest/algorithm/keygen"
)

func TestGenTrapdoor(t *testing.T) {
	strs := make([]string, 1)
	strs[0] = "and"
	I := GenTrapdoor(strs, keygen.GenCipher(keygen.LETTER))
	fmt.Println(I[0])
	fmt.Println(I[1])
}
