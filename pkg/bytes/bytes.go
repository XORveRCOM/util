package bytes

// 特定値で初期化したバイト列を生成
func ClearBytes(size int, init byte) []byte {
	bytes := make([]byte, size)
	if size > 0 {
		// all clear trick
		bytes[0] = init
		for i := 1; i < len(bytes); i *= 2 {
			copy(bytes[i:], bytes[:i])
		}
	}
	return bytes
}

// バイト列をビット転置して八行のバイト列として返す
func TransposeBits(src []byte) [][]byte {
	srcByteCount := len(src)
	// 結果のバイト数
	resultCount := (srcByteCount + 8 - 1) / 8
	// 八行のビット列
	ret := make([][]byte, 8)
	// ビット位置ごとに八行のビット列を抽出
	for bit := 0; bit < 8; bit++ {
		// チェックするビット
		bitMask := byte(0x01 << bit)
		// 結果配列
		bytes := ClearBytes(resultCount, 0)
		// 横断してチェックビットを収集
		for index := 0; index < len(src); index++ {
			if (src[index] & bitMask) != 0 {
				// チェックビットが立っていたので対応ビットに記録
				setMask := byte(0x01 << (index % 8))
				byteOffset := index / 8
				bytes[byteOffset] |= setMask
			}
		}
		ret[bit] = bytes
	}
	return ret
}
