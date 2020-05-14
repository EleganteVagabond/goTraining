package camelcase

func CamelWordCount(camelstr string) int {
	ret := 1
	strlen := len(camelstr)
	for ix := 0; ix < strlen; ix++ {
		if int(camelstr[ix])-91 < 0 {
			ret++
		}
	}
	return ret
}
