package norm

import (
	"exercises/7-phnorm/norm/db"
	"regexp"
)

var regex = regexp.MustCompile("\\D")

// Normalize normalizes a phone number to include only number characters
func Normalize(num db.PhoneNumber) string {
	return regex.ReplaceAllString(num.Number, "")
}

func NormalizeAll(nums []db.PhoneNumber) []db.PhoneNumber {
	pnMap := make(map[string]bool, 0)
	for ix, num := range nums {
		normalized := Normalize(num)
		if pnMap[normalized] {
			nums[ix].Duplicate = true
		}
		pnMap[normalized] = true
		if normalized != num.Number {
			nums[ix].Number = Normalize(num)
			nums[ix].Modified = true
		}
	}
	return nums
}
