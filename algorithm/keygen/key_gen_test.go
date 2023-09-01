/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package keygen

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestCipher_GenKeyLetter(t *testing.T) {

	key := GenCipher(LETTER)
	fa := mat.Formatted(key.M1, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\nM1 = %#v\n\n", fa)
	fb := mat.Formatted(key.M2, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\nM2 = %#v\n\n", fb)
	fc := mat.Formatted(key.SK, mat.Prefix("    "), mat.FormatPython())
	fmt.Printf("layout syntax:\nSK = %#v\n\n", fc)

	row1, col1 := key.M1.Caps()
	if row1 != LETTERLENS*LETTERLENS || col1 != LETTERLENS*LETTERLENS {
		t.Fail()
	}
	row2, col2 := key.M2.Caps()
	if row2 != LETTERLENS*LETTERLENS || col2 != LETTERLENS*LETTERLENS {
		t.Fail()
	}
	row3, col3 := key.SK.Caps()
	if row3 != LETTERLENS*LETTERLENS || col3 != 1 {
		t.Fail()
	}
	for i := 0; i < LETTERLENS; i++ {
		for j := 0; j < LETTERLENS; j++ {
			if key.M1.At(i, j) != 0 && key.M1.At(i, j) != 1 {
				t.Fail()
			}
			if key.M2.At(i, j) != 0 && key.M2.At(i, j) != 1 {
				t.Fail()
			}
		}
		if key.SK.At(i, 0) != 0 && key.SK.At(i, 0) != 1 {
			t.Fail()
		}
	}

	b := key.BaseUnit != LETTERLENS || key.R < 0.0 || key.R >= 1.0 || key.Scope != LETTER
	if b {
		t.Fail()
	}
}

func TestKeyInit(t *testing.T) {
	keyLetter := GenCipher(LETTER)
	keyLetterInit := InitCipher(keyLetter.BaseUnit, keyLetter.R, keyLetter.M1, keyLetter.M2, keyLetter.SK, keyLetter.Scope)
	if keyLetter.R != keyLetterInit.R || keyLetter.Scope != keyLetterInit.Scope || keyLetter.BaseUnit != keyLetterInit.BaseUnit {
		t.Fail()
	}
	if keyLetter.M1 != keyLetterInit.M1 || keyLetter.M2 != keyLetterInit.M2 || keyLetter.SK != keyLetterInit.SK {
		t.Fail()
	}

	keyNum := GenCipher(LETTERANDNURBER)
	keyNumInit := InitCipher(keyNum.BaseUnit, keyNum.R, keyNum.M1, keyNum.M2, keyNum.SK, keyNum.Scope)
	if keyNum.R != keyNumInit.R || keyNum.Scope != keyNumInit.Scope || keyNum.BaseUnit != keyNumInit.BaseUnit {
		t.Fail()
	}
	if keyNum.M1 != keyNumInit.M1 || keyNum.M2 != keyNumInit.M2 || keyNum.SK != keyNumInit.SK {
		t.Fail()
	}

	keyChinese := GenCipher(CHINESE)
	keyChineseInit := InitCipher(keyChinese.BaseUnit, keyChinese.R, keyChinese.M1, keyChinese.M2, keyChinese.SK, keyChinese.Scope)
	if keyChinese.R != keyChineseInit.R || keyChinese.Scope != keyChineseInit.Scope || keyChinese.BaseUnit != keyChineseInit.BaseUnit {
		t.Fail()
	}
	if keyChinese.M1 != keyChineseInit.M1 || keyChinese.M2 != keyChineseInit.M2 || keyChinese.SK != keyChineseInit.SK {
		t.Fail()
	}
}

func TestCipher_GenKeyNumber(t *testing.T) {
	key2 := GenCipher(LETTERANDNURBER)
	row21, col21 := key2.M1.Caps()
	if row21 != LETTERANDNUMLENS*LETTERANDNUMLENS || col21 != LETTERANDNUMLENS*LETTERANDNUMLENS {
		t.Fail()
	}
	row22, col22 := key2.M2.Caps()
	if row22 != LETTERANDNUMLENS*LETTERANDNUMLENS || col22 != LETTERANDNUMLENS*LETTERANDNUMLENS {
		t.Fail()
	}
	row23, col23 := key2.SK.Caps()
	if row23 != LETTERANDNUMLENS*LETTERANDNUMLENS || col23 != 1 {
		t.Fail()
	}
	for i := 0; i < LETTERANDNUMLENS; i++ {
		for j := 0; j < LETTERANDNUMLENS; j++ {
			if key2.M1.At(i, j) != 0 && key2.M1.At(i, j) != 1 {
				t.Fail()
			}
			if key2.M2.At(i, j) != 0 && key2.M2.At(i, j) != 1 {
				t.Fail()
			}
		}
		if key2.SK.At(i, 0) != 0 && key2.SK.At(i, 0) != 1 {
			t.Fail()
		}
	}

	b2 := key2.BaseUnit != LETTERANDNUMLENS || key2.R < 0.0 || key2.R >= 1.0 || key2.Scope != LETTERANDNURBER
	if b2 {
		t.Fail()
	}
}

func TestCipher_GenKeyChinese(t *testing.T) {
	key3 := GenCipher(CHINESE)
	row31, col31 := key3.M1.Caps()
	if row31 != CHINESELENS*CHINESELENS || col31 != CHINESELENS*CHINESELENS {
		t.Fail()
	}
	row32, col32 := key3.M2.Caps()
	if row32 != CHINESELENS*CHINESELENS || col32 != CHINESELENS*CHINESELENS {
		t.Fail()
	}
	row33, col33 := key3.SK.Caps()
	if row33 != CHINESELENS*CHINESELENS || col33 != 1 {
		t.Fail()
	}
	for i := 0; i < CHINESELENS; i++ {
		for j := 0; j < CHINESELENS; j++ {
			if key3.M1.At(i, j) != 0 && key3.M1.At(i, j) != 1 {
				t.Fail()
			}
			if key3.M2.At(i, j) != 0 && key3.M2.At(i, j) != 1 {
				t.Fail()
			}
		}
		if key3.SK.At(i, 0) != 0 && key3.SK.At(i, 0) != 1 {
			t.Fail()
		}
	}

	b3 := key3.BaseUnit != CHINESELENS || key3.R < 0.0 || key3.R >= 1.0 || key3.Scope != CHINESE
	if b3 {
		t.Fail()
	}

}
