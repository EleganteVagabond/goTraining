package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the minimumPasses function below.
func minimumPasses(m int64, w int64, p int64, n int64) int64 {
	fmt.Println(m, w, p, n)
	un := uint64(n)
	up := uint64(p)
	um := uint64(m)
	uw := uint64(w)
	totalCandies := uint64(0)
	totalPasses := int64(0)
	need := uint64(0)
	passCandies := uint64(0)
	lastNeed := uint64(0)
	for totalCandies < un {
		totalPasses++
		lastNeed = need
		need = un - totalCandies
		// overflow
		if lastNeed > need {
			fmt.Println("last > need", lastNeed, need)
			break
		}
		passCandies = um * uw
		// we did it! (check for multiplication overflowed)
		if passCandies >= need {
			break
		}
		// in this case the next pass will do it, just increment and break
		if passCandies >= need/2 {
			totalPasses++
			break
		}
		totalCandies += passCandies
		// maximum the equation M * W
		// / operator on ints returns a floor in golang
		canBuy := totalCandies / up
		// fmt.Println("got",totalCandies,"can buy", canBuy)
		if canBuy > 0 {
			// TODO optimize the if with abs value call
			if um > uw {
				// buy as many as we can to try to make them match
				buyW := min(canBuy, um-uw)
				// increase the workers
				uw += buyW
				// decrement how many we can buy
				canBuy -= buyW
			} else {
				// buy as many as we can to try to make them match
				buyM := min(canBuy, uw-um)
				// increase the machines
				um += buyM
				// decrement how many we can buy
				canBuy -= buyM
			}
			// if we still have candies, divide them as evenly as possible
			if canBuy > 0 {
				// give m any extras
				um += canBuy / 2
				um += canBuy % 2
				// give w just half
				uw += canBuy / 2
			}
			totalCandies %= up
			// fmt.Println("new m,w", m, w)
		} else {
			// since we can't buy any now, calculate how long it will take until we can buy it
			necessaryPassesToBuy := (up - totalCandies) / passCandies
			// we might blow past our goal amount so find the minimum bound
			if necessaryPassesToBuy > need/passCandies {
				necessaryPassesToBuy = need / passCandies
			}
			// fast forward
			totalPasses += int64(necessaryPassesToBuy)
			totalCandies += passCandies * necessaryPassesToBuy
			// overflow
			// fmt.Println("skipping ahead...ludicrous speed, now!", um, uw, canBuy,
			// 	necessaryPassesToBuy, totalPasses, totalCandies, p)
		}
		if totalPasses >= 617737754 {
			fmt.Println(totalPasses, un-totalCandies, passCandies, um, uw, un)
		}

		// fmt.Println(totalPasses)
	}

	return totalPasses
}

func min(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
func max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func abs(v int64) int64 {
	if v < 0 {
		return v * -1
	}
	return v
}

// Complete the unboundedKnapsack function below.
func unboundedKnapsack(k int32, arr []int32) int32 {

	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })

	for _, v := range arr {
		if k%v == 0 {
			return k
		}
	}

	for v := k; v > 0; v-- {
		// see if we can make the "change"
		// if not decrement

		return v
	}
	return 0
}

// Complete the substrings function below.
func substrings(n string) int32 {
	// 1) pull out all single characters
	// 2) find all combinations
	// 3) sum
	// 4) modulus
	MOD := 10 ^ 9 + 7
	totalSum := 0
	// the solution indicates that instead of finding all of the permutations we just
	// find the contribution
	// ex: 16
	// subs = 1, 6, 16
	//  1 adds 1 and 10 to the total = 11
	//  6 adds 6 and 6 to the total = 12
	//  sum is 23
	factor := 1
	for ix := len(n) - 1; ix >= 0; ix-- {
		totalSum = (totalSum + int(n[ix]-'0')*factor*(ix+1)) % MOD
		factor = (factor*10 + 1) % MOD
	}
	return int32(totalSum % MOD)
}

// Complete the redJohn function below.
func redJohn(n int32) int32 {
	// dynamic programming problem
	// using configurations of blocks find total number of ways
	// space is 4 x {n} (always 4 rows)
	// blocks are in size 4 x 1 , either vertical or horizontal
	// for 4x1 to 4x3 there is only 1 way to make each (lining up vertical tiles).
	// for 4x4 there are 2 ways (oriented both horizontally and vertically)
	// so create a recurrence relation
	// F(N) = F(N-1) + F(N-4)
	// Here, F(N) = # of ways to tile the 4xN rectangle with 4x1 and 1x4 tiles.
	// F(0) = 1 (no blocks), F(1)=F(2)=F(3)=1 because they are unique
	// F(N-1) describes what's left when we remove a vertical block of 4 (since it's 4x1)
	// F(N-4) describes how many ways if we take away a horizontal block of (made using 1x4 blocks)
	// ex: N=5 f(5) = f(4)+ f(0), f(4) = f(3)+f(0) = 2 so 2 + 1 = 3
	// ex: N=7 f(7) = f(6)+f(3), f(6) = f(5)+f(2) = 3 + 1, so 4 + 1 = 5
	return 0
}

// used for redjohn above. needs to be converted to the sieve of erasthones
// where we only consider numbers less than the square root of the # of solutions
// see http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func calcPrimes(solutions int) int {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	n := 0
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	} else {
		primes := 0
		nint := int(n) - 2
		for _, p := range primeNumbers {
			if p <= nint {
				primes++
			}
			if p > nint {
				break
			}
		}
		return primes
	}
}

// Complete the bigSorting function below.
func bigSorting(unsorted []string) []string {
	// convert strings to some kind of numeric value (int64?)
	// put these into a slice of the same numbers
	// sort using the built in sort method
	// convert back to strings and return slice with them
	sort.Slice(unsorted, func(i, j int) bool {
		// f1, e1 := strconv.ParseFloat(unsorted[i], 64)
		// if e1 != nil {
		// 	panic(e1)
		// }
		// f2, e2 := strconv.ParseFloat(unsorted[j], 64)
		// if e2 != nil {
		// 	panic(e2)
		// }
		// return f1 < f2
		if len(unsorted[i]) == len(unsorted[j]) {
			ret := false
			for ix := 0; ix < len(unsorted[i])-1; ix++ {
				if unsorted[i][ix] != unsorted[j][ix] {
					ret = unsorted[i][ix] < unsorted[j][ix]
					break
				}
			}
			return ret
		} else {
			return len(unsorted[i]) < len(unsorted[j])
		}
	})
	return unsorted
}

func stockmax(prices []int32) int64 {
	// Write your code here
	// 1. buy? if any future price is greater than today
	// 2. when to sell? maximization function
	// start at end then work back because we have to try to peek ahead
	endIx := len(prices) - 1
	currMax := prices[endIx]
	profit := int64(0)
	for i := endIx - 1; i > -1; i-- {
		// buy and sell
		if prices[i] < currMax {
			profit += int64(currMax - prices[i])
		}
		// new high score
		if prices[i] > currMax {
			currMax = prices[i]
		}
	}

	return profit
}

/*
 * Complete the 'getWays' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. LONG_INTEGER_ARRAY c
 */

func getWays(n int32, c []int64) int64 {
	numways := make([]int64, n+1) // numways[x] means # ways to get sum x
	numways[0] = 1                // init base case n=0

	// go thru coins 1-by-1 to build up numways[] dynamically
	// just need to consider cases where sum j>=c[i]
	// because if j < then the coin is too big
	for _, coin := range c {
		fmt.Println("checking coin", coin)
		for cCheck := coin; cCheck <= int64(n); cCheck++ {
			// find numways to get sum j given value c[i]
			// it consists of those found earlier plus
			// new ones.
			// E.g. if c[]=1,2,3... and c[i]=3,j=5,
			//      new ones will now include '3' with
			//      numways[2] = 2, that is:
			//      '3' with '2', '3' with '1'+'1'
			fmt.Println(coin, cCheck, numways[cCheck-coin], numways)
			numways[cCheck] += numways[cCheck-coin]
		}
		// fmt.Println("num ways with", coin, "=", numways[coin])
	}
	fmt.Println(numways)
	return numways[n] - 1
}

func getWaysR(change int32, coins []int64, solutions []solution) {
	getNextWay([]int64{}, change, change, coins, solutions)
}

type solution struct {
	coins  []int64
	change int32
	ways   [][]int64
}

// recursively calculate for each coin type
func getNextWay(currentCoins []int64, startChange int32, currentChange int32, coins []int64, knownWays []solution) {
	if currentChange > 0 {
		for _, coin := range coins {
			remainder := int64(currentChange) - coin
			if remainder >= 0 {
				getNextWay(append(currentCoins, coin), startChange, int32(remainder), coins, knownWays)
			}
		}
	} else if currentChange == 0 {
		fmt.Println("Coinset is ", currentCoins)
		// knownWays[currentChange] = solution{coins, startChange, [][]int64{currentCoins}}
		knownWays = append(knownWays, solution{coins, startChange, [][]int64{currentCoins}})
	}
}

// Complete the countSort function below.
func countSort(arr [][]string) {
	iMap := make(map[int]*strings.Builder)

	for ix, v := range arr {
		vix, _ := strconv.Atoi(v[0])
		if ix < len(arr)/2 {
			// replace with -
			iMap[vix].WriteString(" -")
		} else {
			iMap[vix].WriteString(" ")
			iMap[vix].WriteString(v[1])
			// iMap[vix] = strings.TrimSpace(iMap[vix] + " " + v[1])
		}
	}

	var keys []int
	for k := range iMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// val := ""
	for _, v := range keys {
		//val = val + iMap[v] + " "
		fmt.Print(strings.TrimSpace(iMap[v].String()))
		fmt.Print(" ")
	}
}

// Complete the closestNumbers function below.
func closestNumbers(arr []int32) []int32 {
	sorted := make([]int, len(arr))
	for ix, v := range arr {
		sorted[ix] = int(v)
	}
	sort.Ints(sorted)
	minDiff := 0
	pairsMap := make(map[int][]int32)
	for ix := 0; ix < len(sorted)-1; ix++ {
		diff := sorted[ix+1] - sorted[ix]
		pairsMap[diff] = append(pairsMap[diff], int32(sorted[ix]), int32(sorted[ix+1]))

		if diff < minDiff || minDiff == 0 {
			minDiff = diff
		}
	}
	return pairsMap[minDiff]
}

// Complete the maxSubarray function below.
func maxSubarray(arr []int32) []int32 {
	maxSubArray, maxSubSequence, maxNegValue := int32(0), int32(0), int32(0)
	// from kardane's algorithm
	currSum, bestSum := int32(0), int32(0)
	// process all of the values
	for _, v := range arr {
		// edge case, all values negative so select the hightest
		if v < 0 && (v > maxNegValue || maxNegValue == 0) {
			maxNegValue = v
		}
		// collect all non-negative values for subsequence
		if v > 0 {
			maxSubSequence += v
		}
		// from kardane's algorithm
		currSum = int32(math.Max(float64(0), float64(currSum+v)))
		bestSum = int32(math.Max(float64(bestSum), float64(currSum)))
	}
	maxSubArray = bestSum
	if maxSubSequence == 0 && maxNegValue != 0 {
		maxSubSequence = maxNegValue
	}
	if maxSubArray <= 0 && maxNegValue != 0 {
		maxSubArray = maxNegValue
	}
	return []int32{int32(maxSubArray), int32(maxSubSequence)}
}

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
			asumsIx, bsumsIx := aix, bix
			acont, bcont := true, true
			caix, cbix := aix, bix
			ca2, cb2 := a[asumsIx], b[bsumsIx]
			for ca2 == cb2 && asumsIx+1 < alen && bsumsIx+1 < blen {
				asumsIx++
				bsumsIx++
				ca2, cb2 = a[asumsIx], b[bsumsIx]
				// if the same char as the match and all the chars have matched
				if ca2 == ca && acont {
					caix = asumsIx
				} else {
					acont = false
				}
				if cb2 == cb && bcont {
					cbix = bsumsIx
				} else {
					bcont = false
				}
			}
			// fmt.Println("Inner loop done", aix, bix, asumsIx, bsumsIx, caix, cbix)
			// copy all of the duplicate values at once
			if rix == 7911 {
				fmt.Println("SGW1", aix, alen, bix, blen, string(ca), string(cb), ca < cb || (ca == cb && bix+1 == blen))
				fmt.Println(a[aix:])
				fmt.Println(b[bix:])
			}
			if ca2 < cb2 || (ca2 == cb2 && bsumsIx+1 == blen) {
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
