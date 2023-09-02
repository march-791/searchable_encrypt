/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package tool

import (
	"bytes"
	"encoding/gob"
)

// Encode 用 gob 数据编码
func Encode(e interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(e)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode 用 gob 数据解码
func Decode(data []byte, e interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(e)
}
