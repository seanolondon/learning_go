package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type results struct {
	input string
}

func main() {
	//open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//use a csv parser then read the whole file
	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//for loop through each row

	//change below to a function which adds 1 each reponse.
	//
	for probs := range reader {
		fmt.Printf("Problem #%v: %v =\n", probs+1, reader[probs][0])

		//use below for input but causes huge loop of 0s
		//input := bufio.NewReader(os.Stdin)
		//	fmt.Printf("%v\n", input)

	}

	//second column is the correct answer

	// fmt.Printf("Check: %s\n", string(csvread))
	//print the first line
	//fmt.Println(reader[0])

	//io.Copy(os.Stdout, reader)

	//	fmt.Println("Quiz Game")
}
