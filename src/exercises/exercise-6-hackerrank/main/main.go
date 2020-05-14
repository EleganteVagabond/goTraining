package main

import (
	"exercises/exercise-6-hackerrank/camelcase"
	"fmt"
)

func main() {
	// run camel case test
	var camelstr string
	fmt.Scanf("%s", &camelstr)
	cnt := camelcase.CamelWordCount(camelstr)
	fmt.Println(cnt)

	

}
