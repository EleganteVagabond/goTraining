package cipher

func CaesarCipher(raw string, shift int32) string {
	if shift > 26 {
		shift = shift % 26
	}
	ret := make([]byte, len(raw))
	for ix := 0; ix < len(ret); ix++ {
		rawrune := rune(raw[ix])
		char := rotate(rawrune, 'A', 'Z', shift)
		if char == rawrune {
			char = rotate(rawrune, 'a', 'z', shift)
		}
		ret[ix] = byte(char)
	}
	return string(ret)
}

func rotate(char rune, min rune, max rune, shift int32) rune {
	if char >= min && char <= max {
		char += shift
		if char > max {
			char = char - (max - min + 1)
		}
	}
	return char
}
