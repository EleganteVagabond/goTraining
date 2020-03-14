package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	probFile  string
	timeout   int
	randomize bool

	problems        [][]string
	probIndex       int
	questionIndeces []int
	correctAnswers  int
)

func endGame() {
	fmt.Println("The quiz has ended")
	fmt.Println("You answered ", correctAnswers, " correctly out of ", len(problems), "questions")
	os.Exit(0)
}

func main() {
	// process the command line args using flag
	flag.StringVar(&probFile, "path", "problems.csv", "location of csv file containing problems, default problems.csv")
	flag.IntVar(&timeout, "timeout", 30, "time alotted for answering questions in seconds, default 30 seconds")
	flag.BoolVar(&randomize, "randomize", false, "randomize the order of the questions, default false")
	flag.Parse()

	fmt.Println("cmd line values file:", probFile, ", timeout:", timeout)
	// open and read the file, ensuring we exit cleanly on errors
	// using := okay here because all variables are scoped to this function
	probFd, err := os.Open(probFile)
	if err != nil {
		log.Fatal(err)
		defer probFd.Close()
	}
	probReader := csv.NewReader(probFd)
	// don't use := because this will overwrite the global scope variable problems var and we already initialized err above
	problems, err = probReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		defer probFd.Close()
	}
	fmt.Println(problems)

	// setup for scoring the quiz
	probIndex, correctAnswers = 0, 0
	// start the end of game timer
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		fmt.Println("times up!")
		endGame()
	}()
	// using scanln
	answer := ""
	if randomize {
		// do some random stuff
		questionIndeces = rand.Perm(len(problems))
	} else {
		questionIndeces = make([]int, len(problems))
		for i := 0; i < len(questionIndeces); i++ {
			questionIndeces[i] = i
		}
	}
	for i := 0; i < len(problems); i++ {
		probIndex = questionIndeces[i]
		question := problems[probIndex][0]
		fmt.Println("Question #", i+1, ":", question)
		fmt.Scanln(&answer)
		if problems[probIndex][1] == answer {
			correctAnswers++
		}
	}

	endGame()

}
