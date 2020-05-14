package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the morganAndString function below.
func morganAndString(a string, b string) string {
	alen, blen := len(a), len(b)
	totalLen := alen + blen
	ret := make([]rune, totalLen)
	aix, bix, rix := 0, 0, 0
	for rix < totalLen && aix < alen && bix < blen {
		ca := a[aix]
		cb := b[bix]
		// fmt.Println("comparing", string(ca), "at ix", aix, "to", string(cb), "at ix", bix)
		if ca == cb && aix+1 != alen && bix+1 != blen {
			// peek ahead
			// find the next minimum
			aix2, bix2 := aix, bix
			acont, bcont := true, true
			caix, cbix := aix, bix
			ca2, cb2 := a[aix2], b[bix2]
			for ca2 == cb2 && aix2+1 < alen && bix2+1 < blen {
				aix2++
				bix2++
				ca2, cb2 = a[aix2], b[bix2]
				// if the same char as the match and all the chars have matched
				if ca2 == ca && acont {
					caix = aix2
				} else {
					acont = false
				}
				if cb2 == cb && bcont {
					cbix = bix2
				} else {
					bcont = false
				}
			}
			// fmt.Println("Inner loop done", aix, bix, aix2, bix2, caix, cbix)
			// copy all of the duplicate values at once
			if rix == 7911 {
				fmt.Println("SGW1", aix, alen, bix, blen, string(ca), string(cb), ca < cb || (ca == cb && bix+1 == blen))
				fmt.Println(a[aix:])
				fmt.Println(b[bix:])
			}
			if ca2 < cb2 || (ca2 == cb2 && bix2+1 == blen) {
				if rix == 7911 {
					fmt.Println("a..", a[aix:caix+1])
				}
				for aix <= caix {
					ret[rix] = rune(a[aix])
					rix++
					aix++
				}
			} else {
				if rix == 7911 {
					fmt.Println("b..", b[bix:cbix+1])
				}
				for bix <= cbix {
					ret[rix] = rune(b[bix])
					rix++
					bix++
				}
			}
		} else {
			if rix == 7911 {
				fmt.Println("SGW2", aix, alen, bix, blen, string(ca), string(cb), ca < cb || (ca == cb && bix+1 == blen))
			}
			if ca < cb || (ca == cb && bix+1 == blen) {
				ret[rix] = rune(ca)
				aix++
			} else {
				ret[rix] = rune(cb)
				bix++
			}
			rix++
			// fmt.Println("Picked", string(ret[rix]))
		}

	}
	// fmt.Println("loop done", aix, alen, bix, blen, rix, totalLen)
	if rix == 7911 {
		fmt.Println("SGW3", aix, alen, bix, blen)
	}

	if aix < alen {
		for ; aix < alen; aix++ {
			ret[rix] = rune(a[aix])
			rix++
		}
	} else if bix < blen {
		for ; bix < blen; bix++ {
			ret[rix] = rune(b[bix])
			rix++
		}
	}
	return string(ret)
}

// Complete the palindromeIndex function below.
func palindromeIndex(s string) int32 {
	// generally, start at left and right, ignore middle
	// if characters don't match, move forward and mark spot
	// else go again
	// if more than one character is wrong, fail
	/*
		ex: abcbca
		a = a check
		b = c no (acbca)
		c = c yes
		b = middle
		del 1
		or
		a = a check
		b = c no (abcba)
		b = b check
		c = middle
		del 4
	*/
	middle := int(len(s) / 2)
	right := len(s)
	del := -1
	for l := 0; l < middle; l, right = l+1, right-1 {
		if s[l] != s[right-1] {
			if del != -1 {
				fmt.Println("2nd error found", string(s[l]), string(s[right-1]))
				// can't do it
				del = -1
				break
			}
			// peek ahead, not just this char but the next one as well
			if s[l+1] == s[right-1] && s[l+2] == s[right-2] {
				fmt.Println("moving left ahead", string(s[l]), string(s[right-1]))
				del = l
				l++
			} else {
				fmt.Println("moving right behind", string(s[l]), string(s[right-1]))
				del = right - 1
				right--
			}
		}
	}

	return int32(del)

}

func steadyGene(gene string) int32 {
	genesMap := map[rune]int{'G': 0, 'A': 0, 'T': 0, 'C': 0}
	numEa := len(gene) / 4
	optLength := 0
	subArr := make([]int, 26)
	for _, c := range gene {
		genesMap[c]++
		if genesMap[c] > numEa {
			subArr[c-'A'] = genesMap[c] - numEa
			optLength++
		}
	}
	//fmt.Println(subArr[0], subArr[2], subArr[6], subArr[19]) //, 'A'-'A', 'C'-'A', 'G'-'A', 'T'-'A')
	if optLength == 0 {
		return 0
	}
	minLength := len(gene)
	left, right := 0, 0
	for left < len(gene) && right < len(gene) {
		if subArr[0] > 0 || subArr[2] > 0 || subArr[6] > 0 || subArr[19] > 0 {
			subArr[gene[right]-'A']--
			right++
		} else {
			minLength = mymin(minLength, right-left)
			subArr[gene[left]-'A']++
			left++
		}
	}
	return int32(minLength)
}

func mymin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Complete the steadyGene function below.
func steadyGeneOLD(gene string) int32 {
	genesMap := map[rune]int{'G': 0, 'A': 0, 'T': 0, 'C': 0}
	numEa := len(gene) / 4
	optLength := 0
	subMap := make(map[rune]int, 0)
	subArr := make([]int, 26)
	for _, c := range gene {
		genesMap[c]++
		if genesMap[c] > numEa {
			subMap[c] = genesMap[c] - numEa
			subArr[c-'A'] = genesMap[c] - numEa
			optLength++
		}
	}
	fmt.Println(subMap)
	if len(subMap) == 0 {
		return 0
	}
	// scan string for matches
	minString := 0
	for ix, c := range gene {
		if subMap[c] != 0 {
			str := findMinString(gene, subMap, subArr, ix, optLength)
			if minString == 0 || str < minString {
				minString = str
				if minString == optLength {
					break
				}
			}
		}
	}
	return int32(minString)
}

func findMinString(gene string, subs map[rune]int, subarr []int, six int, totchanges int) int {
	ret := 0
	subarrc := make([]int, 26)
	copy(subarrc, subarr)
	totSubs := 0
	if len(gene)-six >= totchanges {
		for i := six; i < len(gene); i++ {
			g := rune(gene[i])
			ret++
			if subarrc[g-'A'] != 0 {
				totSubs++
				subarrc[g-'A']--
				if totSubs == totchanges {
					break
				}
			}
		}
	}
	// if len(subsc) == 0 {
	if totSubs == totchanges {
		fmt.Println("found string starting at", six, "with length", ret)
		return ret
	}

	return len(gene)
}

func LCS(short string, long string) int {
	/*
			for i := 1..m
		        for j := 1..n
		            if X[i] = Y[j] //i-1 and j-1 if reading X & Y from zero
		                C[i,j] := C[i-1,j-1] + 1
		            else
		                C[i,j] := max(C[i,j-1], C[i-1,j])
			return C[m,n]
	*/
	counter := make([][]int, len(short)+1)
	for i := range counter {
		counter[i] = make([]int, len(long)+1)
		counter[i][0] = 0
	}
	for i := 1; i <= len(short); i++ {
		for j := 1; j <= len(long); j++ {
			if short[i-1] == long[j-1] {
				counter[i][j] = counter[i-1][j-1] + 1
			} else {
				counter[i][j] = int(math.Max(float64(counter[i][j-1]), float64(counter[i-1][j])))
			}
		}
	}
	return counter[len(short)][len(long)]
}

// Complete the commonChild function below.
// largest string with common characters that maintains order
// eg harry & sally = ay
func commonChild(s1 string, s2 string) int32 {
	s1chars := make([]rune, 0)
	s2chars := make([]rune, 0)
	for _, c := range s1 {
		if strings.Contains(s2, string(c)) {
			s1chars = append(s1chars, c)
		}
	}
	for _, c := range s2 {
		if strings.Contains(s1, string(c)) {
			s2chars = append(s2chars, c)
		}
	}
	s1c, s2c := string(s1chars), string(s2chars)
	fmt.Println(s1c, s2c)
	if s1c == s2c {
		return int32(len(s1chars))
	}
	if len(s1chars) < len(s2chars) {
		return int32(LCS(s1c, s2c))
	} else {
		return int32(LCS(s2c, s1c))
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// reader := bufio.NewReader(os.Stdin)
	/*var probFile string
	var answerFile string
	flag.StringVar(&probFile, "test_dat", "test_file", "location of csv file containing problems, default problems.csv")
	flag.StringVar(&answerFile, "answer_date", "answer_file", "location of csv file containing problems, default problems.csv")
	testFd, err := os.Open(probFile)
	if err != nil {
		log.Fatal(err)
		defer testFd.Close()
	}
	reader := bufio.NewReader(testFd)
	l, _ := reader.ReadString('\n')
	numItems, _ := strconv.Atoi(strings.TrimSpace(l))
	tests := make([][]string, numItems)
	for i := 0; i < numItems; i++ {
		ab := make([]string, 2)
		l, _ = reader.ReadString('\n')
		ab[0] = strings.TrimSpace(l)
		l, _ = reader.ReadString('\n')
		ab[1] = strings.TrimSpace(l)
		tests[i] = ab
	}
	// fmt.Println(tests)

	answerFd, err := os.Open(answerFile)
	if err != nil {
		log.Fatal(err)
		defer testFd.Close()
	}

	reader = bufio.NewReader(answerFd)
	ans := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		l, _ = reader.ReadString('\n')
		ans[i] = strings.TrimSpace(l)
	}
	// fmt.Println(ans)
	resultFd, err := os.Create("test_results")
	if err != nil {
		log.Fatal(err)
		defer resultFd.Close()
	}
	writer := bufio.NewWriter(resultFd)
	tstart := 0
	tstart, numItems = 4, 5
	for i := tstart; i < numItems; i++ {
		answer := morganAndString(tests[i][0], tests[i][1])
		if answer == ans[i] {
			fmt.Println("test", i, "passed")
		} else {
			i2 := 0
			for ; i2 < len(answer); i2++ {
				if answer[i2] != ans[i][i2] {
					break
				}
			}
			fmt.Println("test", i, "failed at", i2, "see file for details")
			writer.WriteString("ans len:" + strconv.Itoa(len(answer)))
			writer.WriteString("\n")
			writer.WriteString("exp ans len:" + strconv.Itoa(len(ans[i])))
			writer.WriteString("\n")
			writer.WriteString("\n")
			writer.WriteString("diff starts at col: " + strconv.Itoa(i2))
			writer.WriteString("\n")
			writer.WriteString("got")
			writer.WriteString("\n")
			writer.WriteString(answer)
			writer.WriteString("\n")
			writer.WriteString("expected")
			writer.WriteString("\n")
			writer.WriteString(ans[i])
			writer.WriteString("\n")
			writer.Flush()
		}
	}*/
	// fmt.Println(ans)
	resultFd, err := os.Create("test_out")
	if err != nil {
		log.Fatal(err)
		defer resultFd.Close()
	}
	writer := bufio.NewWriter(resultFd)
	writer.WriteString("[")
	for i := 0; i < 1000000; i++ {
		writer.WriteString(strconv.Itoa(i + 1))
		writer.WriteString(",")
	}
	writer.WriteString("]")
}
